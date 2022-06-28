package redisx

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)

type ClusterConf struct {
	ProxyConf
}

func initCluster(options *ClusterConf) *redisc.Cluster {

	setDefaultOpts(&options.ProxyConf)
	cluster := &redisc.Cluster{
		StartupNodes: options.AddrList,
		CreatePool: func(addr string, opts ...redis.DialOption) (*redis.Pool, error) {
			return poolWithDial(&options.ProxyConf, dial(addr, &options.ProxyConf)), nil
		},
	}

	if err := cluster.Refresh(); err != nil {
		panic(err)
	}

	return cluster

}
