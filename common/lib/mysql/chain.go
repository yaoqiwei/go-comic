package mysql

import (
	"reflect"

	"fehu/common/lib/mysql/field"
	"fehu/common/lib/mysql/mx"
	"fehu/common/lib/mysql/table"

	"fehu/util/stringify"
)

func (v *Orm) Fields(fields []interface{}) *Orm {
	for _, f := range fields {

		if k, ok := f.(mx.Field); ok {
			v.addField(k)
			continue
		}
		if k, ok := f.(string); ok {
			v.addField(Field(k))
		}
	}
	return v
}

func (v *Orm) Field(fields ...interface{}) *Orm {
	return v.Fields(fields)
}

func (v *Orm) using(tagName, typ string, s ...interface{}) *Orm {
	if len(s) == 0 {
		return v
	}
	mixs, err := transformToMixs(tagName, s...)
	if err != nil {
		return v.setErr(err)
	}

	if typ == "set" {
		return v.addMix(mx.SliceMix(mixs), typ)
	}

	return v.addMix(mx.ConditionMix(mixs), typ)
}

func (v *Orm) Where(s ...interface{}) *Orm {
	return v.using("dbwhere", "where", s...)
}

func (v *Orm) Having(s ...interface{}) *Orm {
	return v.using("dbwhere", "having", s...)
}

func (v *Orm) Set(s ...interface{}) *Orm {
	return v.using("dbset", "set", s...)
}

// 添加表
func (v *Orm) Table(s ...interface{}) *Orm {
	for _, t := range s {
		if t, ok := t.(mx.Table); ok {
			v.addTable(t)
			continue
		}
		if t, ok := t.(string); ok {
			for _, t := range stringify.ToStringSlice(t, ",") {
				v.addTable(Table(t, v.db.prefix).SetSuffix(v.db.suffix))
			}
		}
	}
	return v
}

func (v *Orm) Limit(l ...uint) *Orm {
	if len(l) == 1 {
		v.addMix(Mix("?", l[0]), "limit")
	} else if len(l) == 2 {
		v.addMix(Mix("?,?", l[0], l[1]), "limit")
	}
	return v
}

func (v *Orm) Page(p, l uint) *Orm {
	if p == 0 {
		p = 1
	}
	offset := l * (p - 1)
	return v.Limit(offset, l)
}

func (v *Orm) Order(s string) *Orm {
	keys := transformToKeyList(s)
	p := mx.SliceMix{}
	for _, k := range keys {
		// Alias 为ASC/DESC
		f := field.NewField(k.Name).SetTable(k.Parent)
		if k.Alias == "DESC" || k.Alias == "desc" {
			p = append(p, Mix("%t DESC", f))
		} else {
			p = append(p, f)
		}

	}
	return v.addMix(p, "order")
}

func (v *Orm) Group(s string) *Orm {
	keys := transformToKeyList(s)
	p := mx.SliceMix{}
	for _, k := range keys {
		f := field.NewField(k.Name).SetTable(k.Parent)
		p = append(p, f)
	}
	return v.addMix(p, "group")
}

func (v *Orm) Join(s interface{}, c ...interface{}) *Orm {
	return v.addJoin(mx.INNER_JOIN, s, c...)
}

func (v *Orm) LeftJoin(s interface{}, c ...interface{}) *Orm {
	return v.addJoin(mx.LEFT_JOIN, s, c...)
}

func (v *Orm) RightJoin(s interface{}, c ...interface{}) *Orm {
	return v.addJoin(mx.RIGHT_JOIN, s, c...)
}

func (v *Orm) Alias(n string) *Orm {
	if len(v.table) > 0 {
		if t, ok := v.table[0].(*table.Table); ok {
			t.SetAlias(n)
		}
	}
	return v
}

func (v *Orm) Suffix(i interface{}) *Orm {
	if len(v.table) > 0 {
		if t, ok := v.table[0].(*table.Table); ok {
			t.Suffix(i)
		}
	}
	return v
}

func (v *Orm) Union(o ...*Orm) *Orm {
	for _, o := range o {
		v.addUnion(o.Exec(false))
	}
	return v
}

func (v *Orm) UnionAll(o ...*Orm) *Orm {
	for _, o := range o {
		o.unionAll = true
		v.addUnion(o.Exec(false))
	}
	return v
}

func (v *Orm) Exec(e bool) *Orm {
	v.b = !e
	return v
}

func (v *Orm) Dest(dest interface{}) *Orm {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return v.setErr(ErrNotPointer)
	}
	v.dest = dest
	return v
}
