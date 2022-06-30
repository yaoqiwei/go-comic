package sharding

import (
	"fehu/conf"
	"gorm.io/sharding"
)

type shardingConfigBuilder struct {
	sharding.Config
}

func (*shardingConfigBuilder) GetShardingConfig(shardingKey string) sharding.Config {
	return sharding.Config{
		ShardingKey:    shardingKey,
		NumberOfShards: uint(conf.Mysql.Split),
	}
}
