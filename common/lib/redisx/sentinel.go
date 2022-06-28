package redisx

import (
	"math/rand"
	"time"

	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
)

type SentinelConf struct {
	Master   *ProxyConf
	Slave    *ProxyConf
	UseSlave bool
	ProxyConf
}

func initSentinel(conf *SentinelConf) (*redis.Pool, *redis.Pool) {
	sentinel := getSentinel(conf)
	master := getSentinelMasterPool(sentinel, conf.Master)

	if conf.UseSlave {
		return master, getSentinelSlavePool(sentinel, conf.Slave)
	}
	return master, nil
}

func getSentinel(opts *SentinelConf) *sentinel.Sentinel {
	setDefaultOpts(&opts.ProxyConf)
	if opts.Name == "" {
		opts.Name = "mymaster"
	}
	return &sentinel.Sentinel{
		Addrs:      opts.AddrList,
		MasterName: opts.Name,
		Pool:       pool(&opts.ProxyConf),
	}
}

func getSentinelMasterPool(sntnl *sentinel.Sentinel, opts *ProxyConf) *redis.Pool {
	setDefaultOpts(opts)
	return poolWithDial(opts, func() (redis.Conn, error) {
		masterAddr, err := sntnl.MasterAddr()
		if err != nil {
			return nil, err
		}
		return dial(masterAddr, opts)()
	})
}

func getSentinelSlavePool(sntnl *sentinel.Sentinel, opts *ProxyConf) *redis.Pool {
	setDefaultOpts(opts)
	return poolWithDial(opts, func() (redis.Conn, error) {
		slaveAddrs, err := sntnl.SlaveAddrs()
		if err != nil {
			return nil, err
		}
		rand.Seed(time.Now().UnixNano())
		slaveAddr := slaveAddrs[rand.Intn(len(slaveAddrs))]
		return dial(slaveAddr, opts)()
	})
}
