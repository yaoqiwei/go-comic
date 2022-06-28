package mysql_lib

import (
	"fehu/common/lib"
	mx "fehu/common/lib/mysql"
)

func DB(name ...interface{}) *mx.Orm {
	Db, err := lib.GetDBPool("default")
	if err != nil {
		return nil
	}
	return Db.Table(name...)
}

func Select(dest interface{}, query string, args ...interface{}) error {

	Db, err := lib.GetDBPool("read")
	if err != nil {
		return err
	}

	err = Db.Table().Query(query, args...).Dest(dest).Select()

	if err != nil {
		return err
	}

	return nil
}

func FetchOne(dest interface{}, sql string, params ...interface{}) error {

	Db, err := lib.GetDBPool("read")
	if err != nil {
		return err
	}

	err = Db.Table().Query(sql, params...).Dest(dest).FetchOne()
	if err != nil {
		return err
	}

	return nil
}

func Count(sql string, params ...interface{}) int64 {
	var data MysqlCountData
	FetchOne(&data, sql, params...)
	return data.Count
}

type MysqlCountData struct {
	Count int64 `db:"c"`
}
