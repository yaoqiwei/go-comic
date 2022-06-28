package field

import (
	"fehu/common/lib/mysql/mx"
)

type Field struct {
	Table string
	Name  string
	Alias string

	with mx.WithTrait
}

func NewField(name string) *Field {
	return &Field{Name: name}
}

func (t *Field) With(w mx.With) {
	t.with.With(w)
}

func (f *Field) SetAlias(n string) mx.Field {
	f.Alias = n
	return f
}

func (f *Field) SetTable(n string) *Field {
	f.Table = n
	return f
}

func (t *Field) GetName() string {
	return mx.GetName(t.Alias, t.Name, &t.with)
}

func (f *Field) GetQuery() string {

	f.with.SetQuery()
	query := f.Name

	if f.with.IsWithBackquote() {
		query = "`" + query + "`"
	}

	if f.with.IsWithAlias() && f.Alias != "" {
		query += " " + f.GetName()
	}

	if f.with.IsWithTable() && f.Table != "" {
		tableName := f.Table
		if f.with.IsWithBackquote() {
			tableName = "`" + tableName + "`"
		}
		query = tableName + "." + query
	}
	f.with.Reset()
	return query
}

func (f *Field) GetArgs() []interface{} {
	return nil
}
