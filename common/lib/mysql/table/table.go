package table

import (
	"fehu/common/lib/mysql/mx"
)

type Table struct {
	DBName  string
	Name    string
	Alias   string
	RawName string
	suffix  func(interface{}) string
	join    joins

	with mx.WithTrait
}

func (t *Table) With(w mx.With) {
	t.with.With(w)

	if len(t.join) > 0 {
		for _, j := range t.join {
			j.table.With(w)
			j.joinCondition.With(w)
		}
	}

}

func (t *Table) GetName() string {
	return mx.GetName(t.Alias, t.RawName, &t.with)
}

func (t *Table) Suffix(s interface{}) {
	t.RawName += t.suffix(s)
}

func (t *Table) GetQuery() string {

	t.with.SetQuery()
	query := t.RawName

	if t.with.IsWithBackquote() {
		query = "`" + query + "`"
	}

	if t.DBName != "" {
		dbName := t.DBName
		if t.with.IsWithBackquote() {
			dbName = "`" + dbName + "`"
		}
		query = dbName + "." + query
	}

	if t.with.IsWithAlias() && t.Alias != "" {
		query += " " + t.GetName()
	}

	if t.join != nil {
		for _, j := range t.join {
			if t.with.IsWithAlias() {
				j.table.With(mx.WithAlias)
			}
			if t.with.IsWithBackquote() {
				j.table.With(mx.WithBackquote)
			}

			query += " " + j.joinType.String() + " " + j.table.GetQuery() + " " + j.joinCondition.GetQuery()
		}
	}
	t.with.Reset()

	return query
}

func (t *Table) GetArgs() []interface{} {

	t.with.SetQuery()
	args := []interface{}{}
	if t.join != nil {
		for _, j := range t.join {
			args = append(args, j.joinCondition.GetArgs()...)
		}
	}
	if len(args) == 0 {
		return nil
	}
	t.with.Reset()
	return args
}

func (t *Table) SetSuffix(f func(interface{}) string) *Table {
	t.suffix = f
	return t
}

func (t *Table) SetAlias(a string) *Table {
	t.Alias = a
	return t
}

func (t *Table) SetDBName(n string) *Table {
	t.DBName = n
	return t
}

func (t *Table) IsMuti() bool {
	return len(t.join) > 0
}

func (t *Table) Join(table mx.Container, typ mx.JoinType, condition mx.ConditionMix) mx.Container {

	if t.join == nil {
		t.join = make(joins, 0)
	}

	j := &join{
		table:         table,
		joinType:      typ,
		joinCondition: condition,
	}

	t.join = append(t.join, j)
	return t
}

func NewTable(name string, rawNames ...string) *Table {
	rawName := name
	if len(rawNames) > 0 {
		rawName = rawNames[0]
	}
	return &Table{Name: name, RawName: rawName}
}
