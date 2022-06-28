package mysql

import (
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"strings"

	"fehu/common/lib/mysql/field"
	"fehu/common/lib/mysql/mix"
	"fehu/common/lib/mysql/mx"
	"fehu/common/lib/mysql/table"

	"fehu/util/stringify"
)

var (
	ErrNotPointer            = errors.New("not pass a pointer")
	ErrNotSlice              = errors.New("not pass a slice")
	ErrNoDest                = errors.New("not pass a dest")
	ErrNotMap                = errors.New("not pass a map")
	ErrNotStringMapKey       = errors.New("not pass a string key of map")
	ErrNotStruct             = errors.New("not pass a struct")
	ErrNotStructOrMap        = errors.New("not pass a struct or map")
	ErrNotStructInSlice      = errors.New("not pass a struct in slice")
	ErrNotStructOrMapInSlice = errors.New("not pass a struct or map in slice")
	ErrNilPointer            = errors.New("pass a nil pointer")
	ErrOddNumberOfParams     = errors.New("odd number of parameters")
	ErrNoContainer           = errors.New("no container")
	ErrNoTable               = errors.New("no table")
	ErrType                  = errors.New("error type")
	ErrNoRows                = sql.ErrNoRows
)

type sqlType byte
type resultType byte

const (
	SQL_SELECT sqlType = iota
	SQL_DELETE
	SQL_UPDATE
	SQL_INSERT
	SQL_REPLACE
)

const (
	RESULT_LAST_INSERT_ID resultType = iota
	RESULT_ROWS_AFFECTED
)

type column struct {
	Dest interface{}
}

type query interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
	QueryRow(string, ...interface{}) *sql.Row
}

type scan interface {
	Scan(...interface{}) error
}

type rows interface {
	scan
	Columns() ([]string, error)
	Next() bool
}

func scanSlice(dest interface{}, rows rows) error {

	value, err := getReflectSliceValue(dest)
	if err != nil {
		return err
	}

	columns, _ := rows.Columns()

	base := value.Type().Elem()
	var isPtr bool
	if base.Kind() == reflect.Ptr {
		isPtr = true
		base = base.Elem()
	}

	for rows.Next() {
		rv := reflect.New(base).Elem()
		if err := scanRow(rv, rows, columns); err != nil {
			return err
		}
		if isPtr {
			rv = rv.Addr()
		}
		value.Set(reflect.Append(value, rv))
	}

	return nil
}

func scanField(dest interface{}, rows rows) error {

	if dest == nil {
		return ErrNoDest
	}

	if !rows.Next() {
		return ErrNoRows
	}

	return rows.Scan(dest)
}

func scanFields(dest interface{}, rows rows) error {

	if dest == nil {
		return ErrNoDest
	}

	value, err := getReflectSliceValue(dest)
	if err != nil {
		return err
	}

	base, isPtr := getSliceBase(value)

	for rows.Next() {
		rv := reflect.New(base)
		if err := rows.Scan(rv.Interface()); err != nil {
			return err
		}
		if !isPtr {
			rv = rv.Elem()
		}
		value.Set(reflect.Append(value, rv))
	}

	return nil
}

func checkMap(rv reflect.Value) error {

	rvt := rv.Type()
	if rvt.Kind() != reflect.Map {
		return ErrNotStructOrMap
	}

	rv.Set(reflect.MakeMap(rvt))
	if rvt.Key().Kind() != reflect.String {
		return ErrNotStringMapKey
	}

	return nil
}

func scanRow(rv reflect.Value, rs scan, columns []string) error {

	if rv.Kind() == reflect.Struct {
		rs.Scan(generateScanData(rv, columns)...)
		return nil
	}

	if err := checkMap(rv); err != nil {
		return err
	}

	data := make([]interface{}, 0)
	for i := 0; i < len(columns); i++ {
		data = append(data, reflect.New(rv.Type().Elem()).Interface())
	}

	if err := rs.Scan(data...); err != nil {
		return err
	}

	for k := range columns {
		key := reflect.ValueOf(&columns[k]).Elem()
		rv.SetMapIndex(key, reflect.ValueOf(data[k]).Elem())
	}

	return nil

}

func scanOne(dest interface{}, rows rows) error {

	if dest == nil {
		return ErrNoDest
	}

	if !rows.Next() {
		return ErrNoRows
	}

	value := reflect.ValueOf(dest).Elem()
	columns, _ := rows.Columns()

	return scanRow(value, rows, columns)
}

func generateScanData(rv reflect.Value, columns []string) []interface{} {

	s := []interface{}{}

	columnMap := map[string]*column{}
	for _, v := range columns {
		columnMap[v] = &column{}
	}

	loopStruct(rv, func(v reflect.Value, s reflect.StructField) bool {
		if name := s.Tag.Get("db"); name != "" {
			if column, ok := columnMap[name]; ok {
				if v.CanAddr() && v.CanInterface() {
					column.Dest = v.Addr().Interface()
				}
			}
			return true
		}
		return false
	})

	for _, v := range columns {
		s = append(s, columnMap[v].Dest)
	}

	return s
}

func getReflectSliceValue(dest interface{}) (reflect.Value, error) {

	if dest == nil {
		return reflect.Value{}, ErrNoDest
	}

	value := reflect.ValueOf(dest).Elem()

	if value.Kind() != reflect.Slice {
		return reflect.Value{}, ErrNotSlice
	}

	if value.IsNil() {
		value.Set(reflect.MakeSlice(value.Type(), 0, 0))
	}

	return value, nil
}

func getSliceBase(value reflect.Value) (base reflect.Type, isPtr bool) {
	base = value.Type().Elem()
	if base.Kind() == reflect.Ptr {
		isPtr = true
		base = base.Elem()
	}
	return
}

func loopStructType(val reflect.Type, f func(s reflect.StructField) bool) {

	for k := 0; k < val.NumField(); k++ {
		ft := val.Field(k).Type
		for ft.Kind() == reflect.Ptr || ft.Kind() == reflect.Interface {
			ft = ft.Elem()
		}

		if !f(val.Field(k)) && ft.Kind() == reflect.Struct {
			loopStructType(ft, f)
		}

	}
}

func loopStruct(val reflect.Value, f func(v reflect.Value, s reflect.StructField) bool) {

	for k := 0; k < val.NumField(); k++ {
		ft := val.Field(k)
		if f(ft, val.Type().Field(k)) {
			continue
		}

		for ft.Kind() == reflect.Ptr || ft.Kind() == reflect.Interface {
			if ft.Kind() == reflect.Ptr &&
				ft.Type().Elem().Kind() == reflect.Struct &&
				ft.IsNil() && ft.CanSet() {
				ft.Set(reflect.New(ft.Type().Elem()))
			}
			ft = ft.Elem()
		}
		if ft.Kind() == reflect.Struct {
			loopStruct(ft, f)
		}
	}
}

func removeRep(s []string) []string {
	r := []string{}
	t := map[string]bool{}
	for _, e := range s {
		l := len(t)
		t[e] = false
		if len(t) != l {
			r = append(r, e)
		}
	}
	return r
}

type key struct {
	Alias  string
	Name   string
	Parent string
}

type keyList []*key

func transformToKeyList(i string) keyList {

	list := keyList{}
	slist := stringify.ToStringSlice(i, ",")

	for _, s := range slist {

		s = strings.Trim(s, " ")
		s = strings.ReplaceAll(s, "`", "")

		sli := stringify.ToStringSlice(s, ".")

		k := &key{}
		names := sli[0]
		if len(sli) > 1 {
			k.Parent = sli[0]
			names = sli[1]
		}

		sli = stringify.ToStringSlice(names, " ")
		k.Name = sli[0]
		if len(sli) > 1 {
			k.Alias = sli[len(sli)-1]
		}

		list = append(list, k)
	}

	return list
}

func transformToKey(i string) *key {
	return transformToKeyList(i)[0]
}

func transformStructToMixs(s interface{}, tagName string) mx.Mixs {

	p := mx.Mixs{}
	rv := stringify.GetReflectValue(s)

	loopStruct(rv, func(v reflect.Value, s reflect.StructField) bool {
		db := s.Tag.Get("db")
		dbset := s.Tag.Get(tagName)
		if dbset != "" {
			db = dbset
		}
		if db == "-" {
			return true
		}
		if db != "" && v.CanInterface() {
			p = append(p, Mix("%t=?", Field(db), v.Interface()))

			return true
		}
		return false
	})

	if len(p) == 0 {
		p = nil
	}
	return p
}

func transformMapToMixs(s map[string]interface{}) mx.Mixs {
	p := mx.Mixs{}
	for k, v := range s {
		p = append(p, Mix("%t=?", Field(k), v))
	}
	return p
}

func transformSliceToMixs(s ...interface{}) mx.Mixs {
	p := mx.Mixs{}
	for k := 0; k < len(s); k += 2 {
		p = append(p, Mix("%t=?", Field(s[k].(string)), s[k+1]))
	}
	return p
}

func transformToMixs(tagName string, s ...interface{}) (mx.Mixs, error) {
	var mixs mx.Mixs
	if len(s) == 1 {
		if s, ok := s[0].(mx.Mix); ok {
			return mx.Mixs{s}, nil
		}
		rv := stringify.GetReflectValue(s[0])
		if rv.Kind() == reflect.Struct {
			mixs = transformStructToMixs(s[0], tagName)
		} else if rv.Kind() == reflect.Map {
			mixs = transformMapToMixs(s[0].(map[string]interface{}))
		}
	}

	if mixs == nil {
		if len(s)%2 == 1 {
			return nil, ErrOddNumberOfParams
		}
		mixs = transformSliceToMixs(s...)
	}

	return mixs, nil
}

func Field(f string) mx.Field {
	k := transformToKey(f)
	return field.NewField(k.Name).SetTable(k.Parent).SetAlias(k.Alias)
}

func Table(f string, prefix ...string) *table.Table {
	k := transformToKey(f)
	var pre string
	if len(prefix) > 0 {
		pre = prefix[0]
	}
	return table.NewTable(k.Name, pre+k.Name).SetAlias(k.Alias).SetDBName(k.Parent)
}

func Raw(f string) *mix.Raw {
	return mix.NewRawMix(f)
}

func RawField(f string) *field.RawField {
	return field.NewRawField(f)
}

func Mix(q string, f ...interface{}) *mix.Mix {
	r := regexp.MustCompile(`(?i)%t|\?`)

	k := -1
	mixs := mx.Mixs{}
	args := []interface{}{}

	q = r.ReplaceAllStringFunc(q, func(s string) string {

		k++

		if s == "?" {
			r := stringify.GetReflectValue(f[k])
			if r.Kind() != reflect.Slice {
				args = append(args, f[k])
				return s
			}

			if r.Len() == 0 {
				return "NULL"
			}

			for i := 0; i < r.Len(); i++ {
				if i != 0 {
					s += ", ?"
				}
				args = append(args, r.Index(i).Interface())
			}

		} else if v, ok := f[k].(mx.Mix); ok {
			mixs = append(mixs, v)
			args = append(args, v.GetArgs()...)
		} else if v, ok := f[k].(string); ok {
			mixs = append(mixs, Field(v))
		} else {
			mixs = append(mixs, Raw("NULL"))
		}

		return s

	})

	return mix.NewMix(q, mixs, args)
}
