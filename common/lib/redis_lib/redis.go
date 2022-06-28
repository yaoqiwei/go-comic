package redis_lib

import (
	"encoding/json"
	"fehu/common/lib"
	"fmt"
	"runtime/debug"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

// Do 执行redis命令并返回结果。执行时从连接池获取连接并在执行完命令后关闭连接。
func Do(clusterName string, commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := lib.GetPool(clusterName)
	defer conn.Close()

	reply, err = conn.Do(commandName, args...)
	if err != nil {
		errorHandle(err)
	}

	return reply, err
}

// Get 获取键值。一般不直接使用该值，而是配合下面的工具类方法获取具体类型的值，或者直接使用github.com/gomodule/redigo/redis包的工具方法。
func Get(key string, clusterName string) (interface{}, error) {
	return Do(clusterName, "GET", key)
}

// GetString 获取string类型的键值
func GetString(key string, clusterName string) (string, error) {
	return String(Get(key, clusterName))
}

// GetInt 获取int类型的键值
func GetInt(key string, clusterName string) (int, error) {
	return Int(Get(key, clusterName))
}

// GetInt64 获取int64类型的键值
func GetInt64(key string, clusterName string) (int64, error) {
	return Int64(Get(key, clusterName))
}

// GetBool 获取bool类型的键值
func GetBool(key string, clusterName string) (bool, error) {
	return Bool(Get(key, clusterName))
}

// GetObject 获取非基本类型stuct的键值。在实现上，使用json的Marshal和Unmarshal做序列化存取。
func GetObject(key string, val interface{}, clusterName string) error {
	reply, err := Get(key, clusterName)
	if err != nil {
		errorHandle(err)
	}
	return decode(reply, err, val)
}

// Set 存并设置有效时长。时长的单位为秒。
// 基础类型直接保存，其他用json.Marshal后转成string保存。
func Set(key string, val interface{}, expire int64, clusterName string) error {
	value, err := encode(val)
	if err != nil {
		return err
	}
	if expire > 0 {
		_, err := Do(clusterName, "SETEX", key, expire, value)
		return err
	}
	_, err = Do(clusterName, "SET", key, value)
	return err
}

// SetNx: 添加key值，如果存在添加失败，不会修改原来的值
func SetNx(key string, val interface{}, expire int64, clusterName string) (bool, error) {
	value, err := encode(val)
	if err != nil {
		return false, err
	}

	e, err := Int64(Do(clusterName, "SETNX", key, value))
	if err != nil || e == 0 {
		return false, err
	}

	if expire > 0 {
		_, err = Do(clusterName, "EXPIRE", key, expire)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// Exists 检查键是否存在
func Exists(key string, clusterName string) (bool, error) {
	return Bool(Do(clusterName, "EXISTS", key))
}

//Del 删除键
func Del(key string, clusterName string) error {
	_, err := Do(clusterName, "DEL", key)
	return err
}

// Flush 清空当前数据库中的所有 key，慎用！
func Flush(clusterName string) error {
	//TODO 这里应该删除所有
	_, err := Do(clusterName, "db", "FLUSHDB")
	return err
}

// TTL 以秒为单位。当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1
func TTL(key string, clusterName string) (ttl int64, err error) {
	return Int64(Do(clusterName, "TTL", key))
}

// Expire 设置键过期时间，expire的单位为秒
func Expire(key string, expire int64, clusterName string) error {
	_, err := Bool(Do(clusterName, "EXPIRE", key, expire))
	return err
}

// Incr 将 key 中储存的数字值增一
func Incr(key string, clusterName string) (val int64, err error) {
	return Int64(Do(clusterName, "INCR", key))
}

// IncrBy 将 key 所储存的值加上给定的增量值（increment）。
func IncrBy(key string, amount int64, clusterName string) (val int64, err error) {
	return Int64(Do(clusterName, "INCRBY", key, amount))
}

// Decr 将 key 中储存的数字值减一。
func Decr(key string, clusterName string) (val int64, err error) {
	return Int64(Do(clusterName, "DECR", key))
}

// DecrBy key 所储存的值减去给定的减量值（decrement）。
func DecrBy(key string, amount int64, clusterName string) (val int64, err error) {
	return Int64(Do(clusterName, "DECRBY", key, amount))
}

// HMSet 将一个map存到Redis hash，同时设置有效期，单位：秒
// Example:
//
// ```golang
// m := make(map[string]interface{})
// m["name"] = "corel"
// m["age"] = 23
// err := HMSet("user", m, 10)
// ```
func HMSet(key string, val interface{}, expire int, clusterName string) (err error) {
	conn := lib.GetPool(clusterName)
	defer conn.Close()
	err = conn.Send("HMSET", redis.Args{}.Add(key).AddFlat(val)...)
	if err != nil {
		errorHandle(err)
		return
	}
	if expire > 0 {
		err = conn.Send("EXPIRE", key, int64(expire))
	}
	if err != nil {
		errorHandle(err)
		return
	}
	conn.Flush()
	_, err = conn.Receive()
	if err != nil {
		errorHandle(err)
	}
	return
}

/** Redis hash 是一个string类型的field和value的映射表，hash特别适合用于存储对象。 **/

// HSet 将哈希表 key 中的字段 field 的值设为 val
// Example:
//
// ```golang
// _, err := HSet("user", "age", 23)
// ```
func HSet(key, field string, val interface{}, clusterName string) (interface{}, error) {
	value, err := encode(val)
	if err != nil {
		return nil, err
	}
	return Do(clusterName, "HSET", key, field, value)
}

// HSetExpire 将哈希表 key 中的字段 field 的值设为 val，并设置key的过期时间
// Example:
//
// ```golang
// _, err := HSet("user", "age", 23)
// ```
func HSetExpire(key, field string, val interface{}, expire int64, clusterName string) (err error) {
	value, err := encode(val)
	if err != nil {
		return err
	}
	conn := lib.GetPool(clusterName)
	defer conn.Close()
	err = conn.Send("HSET", key, field, value)
	if err != nil {
		errorHandle(err)
		return err
	}

	if expire > 0 {
		err = conn.Send("EXPIRE", key, expire)
	}
	if err != nil {
		errorHandle(err)
		return
	}
	conn.Flush()
	_, err = conn.Receive()
	if err != nil {
		errorHandle(err)
	}
	return
}

func HDel(key, field interface{}, clusterName string) (interface{}, error) {
	return Do(clusterName, "HDEL", key, field)
}

// HGet 获取存储在哈希表中指定字段的值
// Example:
//
// ```golang
// val, err := HGet("user", "age")
// ```
func HGet(key, field string, clusterName string) (reply interface{}, err error) {
	reply, err = Do(clusterName, "HGET", key, field)
	return
}

// HGetString HGet的工具方法，当字段值为字符串类型时使用
func HGetString(key, field string, clusterName string) (reply string, err error) {
	reply, err = String(HGet(clusterName, key, field))
	return
}

// HGetInt HGet的工具方法，当字段值为int类型时使用
func HGetInt(key, field string, clusterName string) (reply int, err error) {
	reply, err = Int(HGet(clusterName, key, field))
	return
}

// HGetInt64 HGet的工具方法，当字段值为int64类型时使用
func HGetInt64(key, field string, clusterName string) (reply int64, err error) {
	reply, err = Int64(HGet(key, field, clusterName))
	return
}

// HGetBool HGet的工具方法，当字段值为bool类型时使用
func HGetBool(key, field string, clusterName string) (reply bool, err error) {
	reply, err = Bool(HGet(clusterName, key, field))
	return
}

// HGetObject HGet的工具方法，当字段值为非基本类型的stuct时使用
func HGetObject(key, field string, val interface{}, clusterName string) error {
	reply, err := HGet(clusterName, key, field)
	return decode(reply, err, val)
}

// HGetAll HGetAll("key", &val)
func HGetAll(key string, val interface{}, clusterName string) error {
	v, err := redis.Values(Do(clusterName, "HGETALL", key))
	if err != nil {
		errorHandle(err)
		return err
	}

	if err := redis.ScanStruct(v, val); err != nil {
		errorHandle(err)
	}
	return err
}

func SIsMember(key, val interface{}, clusterName string) (bool, error) {
	reply, err := Do(clusterName, "SISMEMBER", key, val)
	return Bool(reply, err)
}

func SCard(key, clusterName string) (int, error) {
	return Int(Do(clusterName, "SCARD", key))
}

// Redis Sadd 命令将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略。
func SAdd(key, val interface{}, expire int64, clusterName string) (bool, error) {
	reply, err := Do(clusterName, "SADD", key, val)
	if expire > 0 {
		Do(clusterName, "EXPIRE", key, expire)
	}
	return Bool(reply, err)
}

/**
Redis列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）
**/

// BLPop 它是 LPOP 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 BLPOP 命令阻塞，直到等待超时或发现可弹出元素为止。
// 超时参数 timeout 接受一个以秒为单位的数字作为值。超时参数设为 0 表示阻塞时间可以无限期延长(block indefinitely) 。
func BLPop(clusterName string, key string, timeout int) (interface{}, error) {
	values, err := redis.Values(Do(clusterName, "BLPOP", key, timeout))
	if err != nil {
		return nil, err
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("redisgo: unexpected number of values, got %d", len(values))
	}
	return values[1], err
}

// BLPopInt BLPop的工具方法，元素类型为int时
func BLPopInt(key string, timeout int, clusterName string) (int, error) {
	return Int(BLPop(clusterName, key, timeout))
}

// BLPopInt64 BLPop的工具方法，元素类型为int64时
func BLPopInt64(key string, timeout int, clusterName string) (int64, error) {
	return Int64(BLPop(clusterName, key, timeout))
}

// BLPopString BLPop的工具方法，元素类型为string时
func BLPopString(key string, timeout int, clusterName string) (string, error) {
	return String(BLPop(clusterName, key, timeout))
}

// BLPopBool BLPop的工具方法，元素类型为bool时
func BLPopBool(key string, timeout int, clusterName string) (bool, error) {
	return Bool(BLPop(clusterName, key, timeout))
}

// BLPopObject BLPop的工具方法，元素类型为object时
func BLPopObject(key string, timeout int, val interface{}, clusterName string) error {
	reply, err := BLPop(clusterName, key, timeout)
	return decode(reply, err, val)
}

// BRPop 它是 RPOP 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 BRPOP 命令阻塞，直到等待超时或发现可弹出元素为止。
// 超时参数 timeout 接受一个以秒为单位的数字作为值。超时参数设为 0 表示阻塞时间可以无限期延长(block indefinitely) 。
func BRPop(clusterName string, key string, timeout int) (interface{}, error) {
	values, err := redis.Values(Do(clusterName, "BRPOP", key, timeout))
	if err != nil {
		return nil, err
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("redisgo: unexpected number of values, got %d", len(values))
	}
	return values[1], err
}

// BRPopInt BRPop的工具方法，元素类型为int时
func BRPopInt(key string, timeout int, clusterName string) (int, error) {
	return Int(BRPop(clusterName, key, timeout))
}

// BRPopInt64 BRPop的工具方法，元素类型为int64时
func BRPopInt64(key string, timeout int, clusterName string) (int64, error) {
	return Int64(BRPop(clusterName, key, timeout))
}

// BRPopString BRPop的工具方法，元素类型为string时
func BRPopString(key string, timeout int, clusterName string) (string, error) {
	return String(BRPop(clusterName, key, timeout))
}

// BRPopBool BRPop的工具方法，元素类型为bool时
func BRPopBool(key string, timeout int, clusterName string) (bool, error) {
	return Bool(BRPop(clusterName, key, timeout))
}

// BRPopObject BRPop的工具方法，元素类型为object时
func BRPopObject(key string, timeout int, val interface{}, clusterName string) error {
	reply, err := BRPop(clusterName, key, timeout)
	return decode(reply, err, val)
}

func LLen(key string, clusterName string) (int, error) {
	return Int(Do(clusterName, "LLen", key))
}

// LPop 移出并获取列表中的第一个元素（表头，左边）
func LPop(key string, clusterName string) (interface{}, error) {
	return Do(clusterName, "LPOP", key)
}

// LPopInt 移出并获取列表中的第一个元素（表头，左边），元素类型为int
func LPopInt(key string, clusterName string) (int, error) {
	return Int(LPop(key, clusterName))
}

// LPopInt64 移出并获取列表中的第一个元素（表头，左边），元素类型为int64
func LPopInt64(key string, clusterName string) (int64, error) {
	return Int64(LPop(key, clusterName))
}

// LPopString 移出并获取列表中的第一个元素（表头，左边），元素类型为string
func LPopString(key string, clusterName string) (string, error) {
	return String(LPop(key, clusterName))
}

// LPopBool 移出并获取列表中的第一个元素（表头，左边），元素类型为bool
func LPopBool(key string, clusterName string) (bool, error) {
	return Bool(LPop(key, clusterName))
}

// LPopObject 移出并获取列表中的第一个元素（表头，左边），元素类型为非基本类型的struct
func LPopObject(key string, val interface{}, clusterName string) error {
	reply, err := LPop(key, clusterName)
	return decode(reply, err, val)
}

// RPop 移出并获取列表中的最后一个元素（表尾，右边）
func RPop(key string, clusterName string) (interface{}, error) {
	return Do(clusterName, "RPOP", key)
}

// RPopInt 移出并获取列表中的最后一个元素（表尾，右边），元素类型为int
func RPopInt(key string, clusterName string) (int, error) {
	return Int(RPop(clusterName, key))
}

// RPopInt64 移出并获取列表中的最后一个元素（表尾，右边），元素类型为int64
func RPopInt64(key string, clusterName string) (int64, error) {
	return Int64(RPop(clusterName, key))
}

// RPopString 移出并获取列表中的最后一个元素（表尾，右边），元素类型为string
func RPopString(key string, clusterName string) (string, error) {
	return String(RPop(clusterName, key))
}

// RPopBool 移出并获取列表中的最后一个元素（表尾，右边），元素类型为bool
func RPopBool(key string, clusterName string) (bool, error) {
	return Bool(RPop(clusterName, key))
}

// RPopObject 移出并获取列表中的最后一个元素（表尾，右边），元素类型为非基本类型的struct
func RPopObject(key string, val interface{}, clusterName string) error {
	reply, err := RPop(clusterName, key)
	return decode(reply, err, val)
}

// LPush 将一个值插入到列表头部
func LPush(key string, member interface{}, clusterName string) error {
	value, err := encode(member)
	if err != nil {
		return err
	}
	_, err = Do(clusterName, "LPUSH", key, value)
	return err
}

// RPush 将一个值插入到列表尾部
func RPush(key string, member interface{}, clusterName string) error {
	value, err := encode(member)
	if err != nil {
		return err
	}
	_, err = Do(clusterName, "RPUSH", key, value)
	return err
}

// ZCard 获取有序集合的成员数
func Zcard(key string, clusterName string) (int, error) {
	return Int(Do(clusterName, "ZCARD", key))
}

// ZScore 命令返回有序集中，成员的分数值。
func ZScore(key string, field interface{}, clusterName string) (int64, error) {
	return Int64(Do(clusterName, "ZSCORE", key, field))
}

// Zrevrange 命令返回有序集中，指定区间内的成员
func ZRevRangeWithScore(key string, start interface{}, end interface{}, clusterName string) (map[string]string, error) {

	m := map[string]string{}
	reply, err := Do(clusterName, "ZREVRANGE", key, start, end, "WITHSCORES")
	if err != nil {
		return m, err
	}

	kvs := reply.([]interface{})

	if len(kvs) >= 2 {
		for i := 0; i < len(kvs); i += 2 {
			m[string(kvs[i].([]byte))] = string(kvs[i+1].([]byte))
		}
	}
	return m, nil
}

func ZRevRange(key string, start interface{}, end interface{}, clusterName string) ([]string, error) {

	m := []string{}
	reply, err := Do(clusterName, "ZREVRANGE", key, start, end, "WITHSCORES")
	if err != nil {
		return m, err
	}

	kvs := reply.([]interface{})

	if len(kvs) > 0 {
		for i := 0; i < len(kvs); i += 2 {
			m = append(m, string(kvs[i].([]byte)))
		}
	}
	return m, nil
}

// ZAdd 将一个或多个 member 元素及其 score 值加入到有序集 key 当中。
func ZAdd(key string, score interface{}, field interface{}, clusterName string) (interface{}, error) {
	return Do(clusterName, "ZADD", key, score, field)
}

// ZIncrBy 命令对有序集合中指定成员的分数加上增量
func ZIncrBy(key string, val int64, field string, clusterName string) (interface{}, error) {
	return Do(clusterName, "ZINCRBY", key, val, field)
}

// HIncrBy 命令用于为哈希表中的字段值加上指定增量值
func HIncrBy(key string, field string, val int64, clusterName string) (interface{}, error) {
	return Do(clusterName, "HINCRBY", key, field, val)
}

// encode 序列化要保存的值
func encode(val interface{}) (interface{}, error) {
	var value interface{}
	switch v := val.(type) {
	case string, int, uint, int8, int16, int32, int64, float32, float64, bool:
		value = v
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		value = string(b)
	}
	return value, nil
}

// decode 反序列化保存的struct对象
func decode(reply interface{}, err error, val interface{}) error {
	str, err := String(reply, err)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), val)
}

func errorHandle(err error) {
	if err == redis.ErrNil {
		return
	}
	logrus.Debug("", map[string]interface{}{
		"error": fmt.Sprint(err),
		"stack": string(debug.Stack()),
	})
}
