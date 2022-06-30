package shardingConfigBuilder

import (
	"fehu/conf"
	"gorm.io/sharding"
)

// GetShardingConfig 获取sharding配置参数
func GetShardingConfig(shardingKey string) sharding.Config {
	return sharding.Config{
		ShardingKey:    shardingKey,
		NumberOfShards: uint(conf.Mysql.Split),
	}
}
