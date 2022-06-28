package mx

import "fehu/util/stringify"

type Mix interface {
	GetQuery() string
	GetArgs() []interface{}
	With(With)
}

type Mixs []Mix

func (f Mixs) With(w With) {
	for _, f := range f {
		f.With(w)
	}
}

func (f Mixs) GetQuery() string {
	if len(f) == 0 {
		return ""
	}
	list := []string{}
	for _, f := range f {
		list = append(list, f.GetQuery())
	}

	return stringify.ToString(list, "")
}

func (f Mixs) GetArgs() []interface{} {
	a := []interface{}{}
	for _, f := range f {
		a = append(a, f.GetArgs()...)
	}
	if len(a) == 0 {
		return nil
	}
	return a
}

type ConditionMix []Mix

func (f ConditionMix) GetQuery() string {
	if len(f) == 0 {
		return ""
	}
	list := []string{}
	for _, f := range f {
		list = append(list, f.GetQuery())
	}

	return stringify.ToString(list, " AND ")
}

func (f ConditionMix) GetArgs() []interface{} {
	return Mixs(f).GetArgs()
}

func (f ConditionMix) With(w With) {
	Mixs(f).With(w)
}

type SliceMix []Mix

func (f SliceMix) GetQuery() string {
	if len(f) == 0 {
		return ""
	}
	list := []string{}
	for _, f := range f {
		list = append(list, f.GetQuery())
	}

	return stringify.ToString(list, ", ")
}

func (f SliceMix) GetArgs() []interface{} {
	return Mixs(f).GetArgs()
}

func (f SliceMix) With(w With) {
	Mixs(f).With(w)
}
