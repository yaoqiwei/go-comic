package field

import (
	"strings"

	"fehu/common/lib/mysql/mx"
)

type MutiField struct {
	q     string
	f     mx.Fields
	Alias string
	with  mx.WithTrait
}

func (t *MutiField) With(w mx.With) {
	t.with.With(w)
	for _, f := range t.f {
		f.With(w)
	}
}

func (f *MutiField) SetAlias(n string) mx.Field {
	f.Alias = n
	return f
}

func (m *MutiField) GetQuery() string {
	m.with.SetQuery()
	query := m.q
	for _, f := range m.f {
		query = strings.Replace(query, "%t", f.GetQuery(), 1)
	}

	if m.Alias != "" {
		if m.with.IsWithBackquote() {
			query += " `" + m.Alias + "`"
		} else {
			query += " " + m.Alias
		}
	}

	m.with.Reset()
	return query
}

func NewMutiField(q string, f ...mx.Field) *MutiField {
	return &MutiField{q: q, f: f}
}

func (f *MutiField) GetArgs() []interface{} {
	return nil
}
