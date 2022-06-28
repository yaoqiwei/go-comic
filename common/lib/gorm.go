package lib

import (
	"errors"
	"fehu/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/sharding"
	"time"
)

var GormPool map[string]*gorm.DB
var Db *gorm.DB

func InitGormPool() error {
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
		gormDB.Use(sharding.Register(sharding.Config{
			ShardingKey:         "user_id",
			NumberOfShards:      4,
			PrimaryKeyGenerator: sharding.PKSnowflake,
		}, "cmf_order"))
		sqlDB, _ := gormDB.DB()
		sqlDB.SetMaxIdleConns(DbConf.MaxIdleConn)
		sqlDB.SetMaxOpenConns(DbConf.MaxOpenConn)
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
