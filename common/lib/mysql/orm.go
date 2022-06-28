package mysql

import (
	"time"

	"fehu/common/lib/mysql/mx"
)

type Orm struct {
	db             *DB
	table          mx.Tables
	mix            mx.Mixs
	fields         mx.Fields
	mixType        string
	dest           interface{}
	err            error
	StartQueryTime time.Time
	Sql            string
	b              bool // true 不执行sql
	orms           []*Orm
	unionAll       bool
}

func (v *Orm) Query(query string, args ...interface{}) *Orm {

	v.addMix(Mix(" "+query, args...))
	return v
}

func (v *Orm) addMix(m mx.Mix, typs ...string) *Orm {

	if v.mix == nil {
		v.mix = make(mx.Mixs, 0)
	}

	if len(typs) == 0 {
		v.mixType = ""
		v.mix = append(v.mix, m)
		return v
	}

	typ := typs[0]
	var ms mx.Mix

	switch typ {
	case "group":
		ms = Raw(" GROUP BY ")
	case "limit":
		ms = Raw(" LIMIT ")
	case "order":
		ms = Raw(" ORDER BY ")
	case "where":
		ms = Raw(" WHERE ")
	case "having":
		ms = Raw(" HAVING ")
	case "set":
		ms = Raw(" SET ")
	}

	if v.mixType == typ {
		switch typ {
		case "where", "having":
			ms = Raw(" AND ")
		case "set":
			ms = Raw(", ")
		}
	}

	if ms != nil {
		v.mixType = typ
		v.mix = append(v.mix, ms)
	}

	v.mix = append(v.mix, m)
	return v
}

func (v *Orm) addField(field mx.Field) *Orm {
	if v.fields == nil {
		v.fields = make(mx.Fields, 0)
	}
	v.fields = append(v.fields, field)
	return v
}

func (v *Orm) addTable(table mx.Table) *Orm {
	if v.table == nil {
		v.table = make(mx.Tables, 0)
	}
	v.table = append(v.table, table)
	return v
}

func (v *Orm) addJoin(typ mx.JoinType, s interface{}, c ...interface{}) *Orm {
	if len(v.table) == 0 {
		return v.setErr(ErrNoTable)
	}

	var container mx.Container
	if c, ok := s.(mx.Container); ok {
		container = c
	} else if s, ok := s.(string); ok {
		container = Table(s, v.db.prefix)
	} else {
		return v.setErr(ErrNoContainer)
	}

	mixs, err := transformToMixs("dbwhere", c...)
	if err != nil {
		return v.setErr(err)
	}

	v.table[0].Join(container, typ, mx.ConditionMix(mixs))

	return v
}

func (v *Orm) addUnion(o *Orm) *Orm {
	if v.orms == nil {
		v.orms = make([]*Orm, 0)
	}
	v.orms = append(v.orms, o)
	return v
}

func (v *Orm) Err() error {
	return v.err
}

func (v *Orm) setErr(e error) *Orm {

	if v.err != nil {
		return v
	}

	v.err = e
	if v.db.errHandler != nil {
		v.db.errHandler(e, v)
	}

	return v
}

func (v *Orm) GetArgs() []interface{} {
	return v.mix.GetArgs()
}
