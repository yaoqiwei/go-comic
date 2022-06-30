package gorm

import (
	"errors"
	"fehu/common/lib/gorm/shardingConfigBuilder"
	"fehu/conf"
	snowflake2 "fehu/util/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/sharding"
	"sync"
	"time"
)

var GormPool map[string]*gorm.DB
var Db *gorm.DB
var Snowflake snowflake2.Snowflake

func InitGormPool() error {
	Snowflake = snowflake2.Snowflake{
		Mutex:        sync.Mutex{},
		WorkerId:     0,
		DatacenterId: 0,
	}
	GormPool = map[string]*gorm.DB{}
	for confName, DbConf := range conf.Mysql.List {
		gormDB, err := gorm.Open(mysql.Open(DbConf.DataSourceName), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   DbConf.Prefix,
				SingularTable: true, // 使用单数表名
			},
		})
		if err != nil {
			return err
		}
		//使用分表插件
		gormDB.Use(sharding.Register(shardingConfigBuilder.GetShardingConfig("user_id"), "comic_order"))
		sqlDB, err := gormDB.DB()
		if err != nil {
			return err
		}
		//最大闲置连接数
		sqlDB.SetMaxIdleConns(DbConf.MaxIdleConn)
		//最大的连接数，默认值为0表示不限制
		sqlDB.SetMaxOpenConns(DbConf.MaxOpenConn)
		//最大连接超时
		sqlDB.SetConnMaxLifetime(time.Duration(DbConf.MaxConnLifeTime) * time.Second)
		GormPool[confName] = gormDB
	}
	//手动配置连接
	if gormPool, err := GetGormPool("default"); err == nil {
		Db = gormPool
	}
	return nil
}

func GetGormPool(name string) (*gorm.DB, error) {
	if dbPool, ok := GormPool[name]; ok {
		return dbPool, nil
	}
	return nil, errors.New("get gormPool error")
}
