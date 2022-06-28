package redisx

import (
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

func initNormal(conf *ProxyConf) *redis.Pool {
	setDefaultOpts(conf)
	return getNormalPool(conf)
}

func getNormalPool(opts *ProxyConf) *redis.Pool {
	return poolWithDial(opts, func() (redis.Conn, error) {
		addrs := opts.AddrList
		var addr string
		if len(addrs) == 0 {
			addr = "127.0.0.1:6379"
		} else {
			rand.Seed(time.Now().UnixNano())
			addr = addrs[rand.Intn(len(addrs))]
		}
		return dial(addr, opts)()
	})
}
