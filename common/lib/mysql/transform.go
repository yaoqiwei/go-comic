package mysql

import (
	"reflect"
	"strings"

	"fehu/common/lib/mysql/mx"

	"fehu/util/stringify"
)

func (v *Orm) transformTable() string {
	if v.table.IsMuti() {
		v.table.With(mx.WithAlias)
		v.table.With(mx.WithTable)
	}
	v.table.With(mx.WithBackquote)
	return v.table.GetQuery()
}
func (v *Orm) transformQuery() string {
	if v.table.IsMuti() {
		v.mix.With(mx.WithTable)
	}
	v.mix.With(mx.WithBackquote)

	query := v.mix.GetQuery()
	if len(v.table) > 0 && query != "" {
		return query
	}

	return strings.Trim(query, " ")
}

func (v *Orm) transformDestToField() mx.Fields {
	val := stringify.GetReflectValue(v.dest).Type()
	if val.Kind() == reflect.Slice {
		val = val.Elem()
		for val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
	}

	keys := mx.Fields{}
	if val.Kind() == reflect.Struct {
		fields := []string{}
		loopStructType(val, func(s reflect.StructField) bool {
			name := s.Tag.Get("db")
			if name != "" {
				if name != "-" {
					fields = append(fields, name)
				}
				return true
			}
			return false
		})
		fields = removeRep(fields)
		if len(fields) > 0 {
			for _, f := range fields {
				keys = append(keys, Field(f))
			}
		} else {
			keys = append(keys, RawField("1"))
		}
	} else {
		keys = append(keys, RawField("*"))
	}

	return keys
}

func (v *Orm) transformFields() string {
	if len(v.fields) == 0 {
		if v.dest != nil {
			fields := v.transformDestToField()
			fields.With(mx.WithBackquote | mx.WithAlias)
			if v.table.IsMuti() {
				fields.With(mx.WithTable)
			}
			return fields.GetQuery()
		}
		return "*"
	}

	if v.table.IsMuti() {
		v.fields.With(mx.WithTable)
	}
	v.fields.With(mx.WithBackquote | mx.WithAlias)
	return v.fields.GetQuery()
}

func (v *Orm) transformSelectSql() string {
	return "SELECT " + v.transformFields() + " FROM " + v.transformTable() + v.transformQuery()
}

func (v *Orm) transformUpdateSql() string {
	return "UPDATE " + v.transformTable() + v.transformQuery()
}

func (v *Orm) transformDeleteSql() string {
	return "DELETE FROM " + v.transformTable() + v.transformQuery()
}

func (v *Orm) transformInsertSql() string {
	return "INSERT INTO " + v.transformTable() + v.transformQuery()
}

func (v *Orm) transformReplaceSql() string {
	return "REPLACE INTO " + v.transformTable() + v.transformQuery()
}
