package mix

import (
	"strings"

	"fehu/common/lib/mysql/mx"
)

type Mix struct {
	q string
	m mx.Mixs
	a []interface{}

	with mx.WithTrait
}

func (t *Mix) With(w mx.With) {
	t.with.With(w)
	t.m.With(w)
}

func (m *Mix) GetQuery() string {
	m.with.SetQuery()
	query := m.q
	for _, v := range m.m {
		query = strings.Replace(query, "%t", v.GetQuery(), 1)
	}
	m.with.Reset()
	return query
}

func (m *Mix) GetArgs() []interface{} {
	return m.a
}

func NewMix(q string, mixs mx.Mixs, args []interface{}) *Mix {
	mix := &Mix{q: q}
	if len(mixs) > 0 {
		mix.m = mixs
	}
	if len(args) > 0 {
		mix.a = args
	}
	return mix
}
