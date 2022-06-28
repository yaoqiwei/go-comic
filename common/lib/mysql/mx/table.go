package mx

import "fehu/util/stringify"

type Container interface {
	Mix
	Join(Container, JoinType, ConditionMix) Container
	IsMuti() bool
}

type Table interface {
	GetName() string
	Container
}

type Tables []Table

func (ts Tables) With(w With) {
	for _, t := range ts {
		t.With(w)
	}
}

func (ts Tables) GetQuery() string {
	s := []string{}
	for _, t := range ts {
		s = append(s, t.GetQuery())
	}
	return stringify.ToString(s, ", ")
}

func (ts Tables) GetArgs() []interface{} {
	args := []interface{}{}
	for _, t := range ts {
		args = append(args, t.GetArgs()...)
	}
	if len(args) == 0 {
		return nil
	}
	return args
}

func (ts Tables) IsMuti() bool {
	if len(ts) == 0 {
		return false
	}

	if len(ts) == 1 {
		return ts[0].IsMuti()
	}

	return true
}
