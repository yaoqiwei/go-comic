package snowflake

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

const (
	epoch             = int64(1656578574000)                                              //起始时间（时间戳/毫秒）：2020-01-01 00：00：00
	timestampBits     = uint(41)                                                          //时间戳占用位数
	datacenterIdBits  = uint(2)                                                           //数据中心id所占位数
	workerIdBits      = uint(4)                                                           //机器id所占位数
	sequenceBits      = uint(12)                                                          //序列所占的位数
	randomNumberBits  = uint(4)                                                           //随机数所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))                                 //时间戳最大值
	datacenterIdMax   = int64(-1 ^ (-1 << datacenterIdBits))                              //支持的最大数据中心id数量
	workerIdMax       = int64(-1 ^ (-1 << workerIdBits))                                  //支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))                                  //支持的最大序列id数量
	randomNumberMax   = int64(-1 ^ (-1 << randomNumberBits))                              //支持的最大随机数数量
	sequenceShift     = randomNumberBits                                                  //序列号左移位数
	workerIdShift     = sequenceBits + randomNumberBits                                   //机器id左移位数
	datacenterIdShift = sequenceBits + workerIdBits + randomNumberBits                    //数据中心id左移位数
	timestampShift    = sequenceBits + workerIdBits + datacenterIdBits + randomNumberBits //时间戳左移位数
)

type Snowflake struct {
	sync.Mutex
	Timestamp    int64 //时间戳，毫秒
	WorkerId     int64 //工作节点
	DatacenterId int64 //数据中心机房id
	Sequence     int64 //序列号
}

func (s *Snowflake) GetId() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1000000 //转毫秒
	if s.Timestamp == now {
		//当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.Sequence = (s.Sequence + 1) & sequenceMask
		if s.Sequence == 0 {
			//如果当前序列超出12bit长度，则需要等待下一毫秒
			//下一毫秒将使用sequence：0
			for now <= s.Timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.Sequence = 0
	}
	t := now - epoch
	if t > timestampMax || t < 0 {
		logrus.Errorf("epoch must be between 0 and %d", timestampMax)
		return 0
	}
	s.Timestamp = now
	rand.Seed(time.Now().UnixNano())
	random := int64(rand.Intn(16))
	r := (t)<<timestampShift | (s.DatacenterId << datacenterIdShift) | (s.WorkerId << workerIdShift) | (s.Sequence << sequenceShift) | (random)
	fmt.Println(r)
	return r
}

func (s *Snowflake) GetRecombinationId(id int64) int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1000000 //转毫秒
	if s.Timestamp == now {
		//当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.Sequence = (s.Sequence + 1) & sequenceMask
		if s.Sequence == 0 {
			//如果当前序列超出12bit长度，则需要等待下一毫秒
			//下一毫秒将使用sequence：0
			for now <= s.Timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.Sequence = 0
	}
	t := now - epoch
	if t > timestampMax || t < 0 {
		logrus.Errorf("epoch must be between 0 and %d", timestampMax)
		return 0
	}
	s.Timestamp = now
	random := s.GetRandomNumber(id)
	r := (t)<<timestampShift | (s.DatacenterId << datacenterIdShift) | (s.WorkerId << workerIdShift) | (s.Sequence << sequenceShift) | (random)
	fmt.Println(r)
	return r
}

func (s *Snowflake) GetRandomNumber(id int64) int64 {
	id = id & randomNumberMax
	return id
}

/*func main() {
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		random := int64(rand.Intn(16))
		fmt.Println(random)
	}
}*/
