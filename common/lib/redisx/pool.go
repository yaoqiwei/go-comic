package redisx

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Pool interface {
	Get() redis.Conn
	Close() error
}

type ProxyConf struct {
	Name           string
	AddrList       []string
	MaxActive      int
	MaxIdle        int
	Downgrade      bool
	Network        string
	Password       string
	Db             int
	IdleTimeout    time.Duration
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	Wait           bool
}

type Conf struct {
	Mode         string        // 模式，支持 single/sentinel/cluster
	SingleConf   *ProxyConf    // 地址配置
	SentinelConf *SentinelConf // 哨兵配置
	ClusterConf  *ClusterConf  // 集群配置
}

type Pools struct {
	pools map[string]Pool
}

var NoDefaultPool = errors.New("redisx:no default pool")

func (v *Pools) GetPool(l ...string) redis.Conn {
	name := "default"
	if len(l) > 0 {
		name = l[0]
	}
	pool, ok := v.pools[name]
	if !ok {
		pool, ok = v.pools["default"]
		if !ok {
			panic(NoDefaultPool)
		}
	}
	return pool.Get()
}

func (v *Pools) closeIfDown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("redis close")
	for _, v := range v.pools {
		v.Close()
	}
}

func InitRedis(confs []*Conf) Pools {
	p := Pools{
		pools: map[string]Pool{},
	}
	go p.closeIfDown()

	for _, conf := range confs {

		if conf.Mode == "sentinel" {
			if conf.SentinelConf.Master.Name == "" {
				conf.SentinelConf.Master.Name = "master"
			}
			if conf.SentinelConf.Slave.Name == "" {
				conf.SentinelConf.Slave.Name = "slave"
			}
			p.pools[conf.SentinelConf.Master.Name], p.pools[conf.SentinelConf.Slave.Name] = initSentinel(conf.SentinelConf)
		} else if conf.Mode == "cluster" {
			if conf.ClusterConf.Name == "" {
				conf.ClusterConf.Name = "default"
			}
			p.pools[conf.ClusterConf.Name] = initCluster(conf.ClusterConf)
		} else {
			if conf.SingleConf.Name == "" {
				conf.SingleConf.Name = "default"
			}
			p.pools[conf.SingleConf.Name] = initNormal(conf.SingleConf)
		}

	}
	return p
}

func setDefaultOpts(opts *ProxyConf) {

	if opts.Network == "" {
		opts.Network = "tcp"
	}

	if opts.MaxIdle == 0 {
		opts.MaxIdle = 3
	}

	if opts.MaxActive == 0 {
		opts.MaxIdle = 8
	}

	if opts.IdleTimeout == 0 {
		opts.IdleTimeout = 10 * time.Second
	}

	if opts.ConnectTimeout == 0 {
		opts.ConnectTimeout = 10 * time.Second
	}

	if opts.ReadTimeout == 0 {
		opts.ReadTimeout = 5 * time.Second
	}

	if opts.WriteTimeout == 0 {
		opts.WriteTimeout = 5 * time.Second
	}
}

var NoAddr = errors.New("redisx:no addr")

func pool(opts *ProxyConf) func(addr string) *redis.Pool {
	return func(addr string) *redis.Pool {
		return poolWithDial(opts, dial(addr, opts))
	}
}

func poolWithDial(opts *ProxyConf, dial func() (redis.Conn, error)) *redis.Pool {
	return &redis.Pool{
		MaxIdle:      opts.MaxIdle,
		MaxActive:    opts.MaxActive,
		Wait:         opts.Wait,
		IdleTimeout:  opts.IdleTimeout,
		Dial:         dial,
		TestOnBorrow: testPing(),
	}
}

func dial(addr string, opts *ProxyConf) func() (redis.Conn, error) {
	return func() (redis.Conn, error) {

		conn, err := redis.Dial(opts.Network, addr,
			redis.DialConnectTimeout(opts.ConnectTimeout),
			redis.DialReadTimeout(opts.ReadTimeout),
			redis.DialWriteTimeout(opts.WriteTimeout),
		)
		if err != nil {
			return nil, err
		}
		if opts.Password != "" {
			if _, err := conn.Do("AUTH", opts.Password); err != nil {
				conn.Close()
				return nil, err
			}
		}
		if opts.Db != 0 {
			if _, err := conn.Do("SELECT", opts.Db); err != nil {
				conn.Close()
				return nil, err
			}
		}
		return conn, err
	}
}

func testPing() func(conn redis.Conn, t time.Time) error {
	return func(conn redis.Conn, t time.Time) error {
		_, err := conn.Do("PING")
		return err
	}
}
