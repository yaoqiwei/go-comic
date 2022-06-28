package mx

type With int

const (
	WithAlias With = 1 << iota
	WithTable
	WithBackquote
)

type WithTrait struct {
	Status With
	query  bool
}

func (wt *WithTrait) With(w With) {
	wt.Status |= w
}

func (wt *WithTrait) SetQuery() {
	wt.query = true
}

func (wt *WithTrait) Reset() {
	wt.Status = 0
	wt.query = false
}

func (wt *WithTrait) IsWithAlias() bool {
	return wt.Status&WithAlias > 0
}

func (wt *WithTrait) IsWithBackquote() bool {
	return wt.Status&WithBackquote > 0
}

func (wt *WithTrait) IsWithTable() bool {
	return wt.Status&WithTable > 0
}

func (wt *WithTrait) IsQuery() bool {
	return wt.query
}
