package shardingConfigBuilder

import (
	"fehu/conf"
	"gorm.io/sharding"
)

func GetShardingConfig(shardingKey string) sharding.Config {
	return sharding.Config{
		ShardingKey:    shardingKey,
		NumberOfShards: uint(conf.Mysql.Split),
	}
}
