package lib

import (
	"database/sql"
	"errors"
	"fehu/conf"
	"runtime/debug"
	"strings"
	"time"

	"fehu/common/lib/mysql"
	"fehu/util/convert"
	"fehu/util/stringify"

	"github.com/sirupsen/logrus"
)

var DBMapPool map[string]*mysql.DB
var DBDefaultPool *mysql.DB

func InitDBPool() error {
	if len(conf.Mysql.List) == 0 {
		logrus.Errorf("empty mysql config.")
	}
	DBMapPool = map[string]*mysql.DB{}
	for confName, DbConf := range conf.Mysql.List {
		dbPool, err := mysql.Open("mysql", DbConf.DataSourceName)
		if err != nil {
			return err
		}
		dbPool.WithPrefix(DbConf.Prefix)
		dbPool.SetMaxOpenConns(DbConf.MaxOpenConn)
		dbPool.SetMaxIdleConns(DbConf.MaxIdleConn)
		dbPool.SetConnMaxLifetime(time.Duration(DbConf.MaxConnLifeTime) * time.Second)
		dbPool.WithSuffix(subTable)
		dbPool.WithErrHandler(errorHandle)
		dbPool.WithAfterQueryHandler(afterQueryHandle)

		err = dbPool.Ping()
		if err != nil {
			return err
		}

		DBMapPool[confName] = dbPool
		logrus.Infof("%s db pool start", confName)
	}

	//手动配置连接
	if dbPool, err := GetDBPool("default"); err == nil {
		DBDefaultPool = dbPool
	}

	return nil
}

func GetDBPool(name string) (*mysql.DB, error) {
	if dbPool, ok := DBMapPool[name]; ok {
		return dbPool, nil
	}
	return nil, errors.New("get pool error")
}

func CloseDB() error {
	for _, dbPool := range DBMapPool {
		dbPool.Close()
	}
	return nil
}

func subTable(i interface{}) string {
	if conf.Mysql.Split == 0 {
		return ""
	}
	v := stringify.ToInt(i)
	return "_" + stringify.ToString(v%int64(conf.Mysql.Split))
}

func errorHandle(err error, o *mysql.Orm) {
	if err == sql.ErrNoRows {
		return
	}
	logrus.Warnf("mysql err: %s", err.Error())
	logrus.Warnf("sql: %s", o.Sql)
	logrus.Debug(string(debug.Stack()))
}

func afterQueryHandle(m *mysql.Orm) {
	t := time.Since(m.StartQueryTime)
	if t > time.Millisecond*100 {
		logrus.Warnf("sql exec time too long: %s,%s,%v", m.Sql, convert.ToJson(m.GetArgs()), t)
	}

	if strings.Contains(m.Sql, " LIKE ") || strings.Contains(m.Sql, "RAND()") {
		logrus.Warnf("sql exec use LIKE query: %s,%s,%v", m.Sql, convert.ToJson(m.GetArgs()), t)
	} else {
		// logrus.Warnf("sql query log: %s,%s,%v", m.Sql, convert.ToJson(m.GetArgs()), t)
	}
}
