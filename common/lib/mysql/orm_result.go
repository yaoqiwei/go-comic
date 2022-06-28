package mysql

import (
	"time"

	"fehu/common/lib/mysql/field"

	"fehu/util/stringify"
)

func (v *Orm) startQuery(t sqlType) {

	var sql string
	if len(v.table) > 0 {
		switch t {
		case SQL_SELECT:
			sql = v.transformSelectSql()
		case SQL_UPDATE:
			sql = v.transformUpdateSql()
		case SQL_INSERT:
			sql = v.transformInsertSql()
		case SQL_DELETE:
			sql = v.transformDeleteSql()
		case SQL_REPLACE:
			sql = v.transformReplaceSql()
		}
	} else {
		sql = v.transformQuery()
	}

	v.StartQueryTime = time.Now()
	v.Sql = sql
}

func (v *Orm) afterQuery() {
	if v.db.afterQueryHandler != nil {
		v.db.afterQueryHandler(v)
	}
}

// 获取多条数据
func (v *Orm) Select() error {

	v.startQuery(SQL_SELECT)

	if len(v.orms) > 0 {
		for _, o := range v.orms {
			o.Select()
			v.Sql += " UNION "
			if o.unionAll {
				v.Sql += "ALL "
			}
			o.Select()
			v.Sql += "(" + o.Sql + ")"
		}
	}

	if v.b {
		return nil
	}

	rows, err := v.db.q.Query(v.Sql, v.GetArgs()...)
	if err != nil {
		v.setErr(err)
		return err
	}

	defer rows.Close()
	v.afterQuery()

	err = scanSlice(v.dest, rows)
	if err != nil {
		v.setErr(err)
		return err
	}

	return nil
}

// 获取单条数据
func (v *Orm) FetchOne() error {

	v.Limit(1)

	v.startQuery(SQL_SELECT)

	if v.b {
		return nil
	}

	rows, err := v.db.q.Query(v.Sql, v.GetArgs()...)
	if err != nil {
		v.setErr(err)
		return err
	}

	defer rows.Close()
	v.afterQuery()

	err = scanOne(v.dest, rows)
	if err != nil {
		v.setErr(err)
		return err
	}

	return nil
}

func (v *Orm) exec(s sqlType, r resultType) (int64, error) {

	v.startQuery(s)

	if v.b {
		return 0, nil
	}

	result, err := v.db.q.Exec(v.Sql, v.GetArgs()...)
	if err != nil {
		v.setErr(err)
		return 0, err
	}

	v.afterQuery()
	if r == RESULT_LAST_INSERT_ID {
		return result.LastInsertId()
	}

	return result.RowsAffected()
}

// 更新
func (v *Orm) Update() (int64, error) {
	return v.exec(SQL_UPDATE, RESULT_ROWS_AFFECTED)
}

// 插入
func (v *Orm) Insert() (int64, error) {
	return v.exec(SQL_INSERT, RESULT_LAST_INSERT_ID)
}

func (v *Orm) Replace() (int64, error) {
	return v.exec(SQL_REPLACE, RESULT_LAST_INSERT_ID)
}

// 删除
func (v *Orm) Delete() (int64, error) {
	return v.exec(SQL_DELETE, RESULT_ROWS_AFFECTED)
}

// 获取单个字段的值
func (v *Orm) GetField(name interface{}) error {

	v.Field(name)
	v.Limit(1)

	v.startQuery(SQL_SELECT)

	if v.b {
		return nil
	}

	rows, err := v.db.q.Query(v.Sql, v.GetArgs()...)
	if err != nil {
		v.setErr(err)
		return err
	}
	defer rows.Close()
	v.afterQuery()

	err = scanField(v.dest, rows)
	if err != nil {
		v.setErr(err)
	}

	return err
}

// 获取单个字段的值的slice
func (v *Orm) GetFields(name string) error {
	v.Field(name)

	v.startQuery(SQL_SELECT)

	if v.b {
		return nil
	}

	rows, err := v.db.q.Query(v.Sql, v.GetArgs()...)
	if err != nil {
		v.setErr(err)
		return err
	}
	defer rows.Close()
	v.afterQuery()

	err = scanFields(v.dest, rows)
	if err != nil {
		v.setErr(err)
	}

	return err
}

func (v *Orm) GetFieldString(f interface{}) string {
	var data *string
	v.Dest(&data).GetField(f)
	if data == nil {
		return ""
	}
	return *data
}

func (v *Orm) GetFieldInt(f interface{}) int64 {
	var data *int64
	v.Dest(&data).GetField(f)
	if data == nil {
		return 0
	}
	return *data
}

func (v *Orm) Count(f ...string) int64 {
	if len(f) > 0 {
		return v.GetFieldInt(field.NewMutiField("COUNT(%t)", Field(f[0])))
	}
	return v.GetFieldInt(RawField("COUNT(1)"))
}

func (v *Orm) Sum(f string) int64 {
	return v.GetFieldInt(field.NewMutiField("SUM(%t)", Field(f)))
}

func (v *Orm) SumFloat(f string) float64 {
	return stringify.ToFloat(v.GetFieldString(field.NewMutiField("SUM(%t)", Field(f))))
}

func (v *Orm) Exist() bool {
	return v.GetFieldInt(RawField("1")) != 0
}

func (v *Orm) GetFieldsString(name string) []string {
	data := []string{}
	v.Dest(&data).GetFields(name)
	return data
}

func (v *Orm) GetFieldsInt(name string) []int64 {
	data := []int64{}
	v.Dest(&data).GetFields(name)
	return data
}
