package lib

import (
	"fehu/conf"
	"time"

	"fehu/common/lib/redisx"

	"github.com/gomodule/redigo/redis"
)

var Pools redisx.Pools

type GlobalRedisUitl int

func InitRedis() (err error) {

	confList := []*redisx.Conf{}

	for k, v := range conf.Redis {
		conf := &redisx.ProxyConf{}
		conf.Name = k
		conf.AddrList = v.ProxyList
		conf.MaxActive = v.MaxActive
		conf.MaxIdle = v.MaxIdle
		conf.Downgrade = v.Downgrade
		conf.Network = v.Network
		conf.Password = v.Password
		conf.Db = v.Db
		conf.IdleTimeout = time.Duration(v.IdleTimeout) * time.Second
		conf.ConnectTimeout = time.Duration(v.ConnectTimeout) * time.Second
		conf.ReadTimeout = time.Duration(v.ReadTimeout) * time.Second
		conf.WriteTimeout = time.Duration(v.WriteTimeout) * time.Second
		conf.Wait = v.Wait

		confList = append(confList, &redisx.Conf{
			Mode:       "single",
			SingleConf: conf,
		})

	}

	Pools = redisx.InitRedis(confList)

	return nil
}

func GetPool(clusterList ...string) redis.Conn {
	return Pools.GetPool(clusterList...)
}
