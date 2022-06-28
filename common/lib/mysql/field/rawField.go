package field

import (
	"fehu/common/lib/mysql/mx"
)

type RawField struct {
	q     string
	Alias string
	with  mx.WithTrait
}

func (t *RawField) With(w mx.With) {
	t.with.With(w)
}

func (f *RawField) SetAlias(n string) mx.Field {
	f.Alias = n
	return f
}

func (f *RawField) GetQuery() string {

	f.with.SetQuery()
	query := f.q
	if f.Alias != "" {
		if f.with.IsWithBackquote() {
			query += " `" + f.Alias + "`"
		} else {
			query += " " + f.Alias
		}
	}

	f.with.Reset()
	return query
}

func (f *RawField) GetArgs() []interface{} {
	return nil
}

func NewRawField(q string) *RawField {
	return &RawField{q: q}
}
