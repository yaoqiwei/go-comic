package model

type Bytebool byte

func (v *Bytebool) Val() byte {
	return byte(*v)
}

func (v *Bytebool) Set(b byte) *Bytebool {
	*v = Bytebool(b)
	return v
}

func (v *Bytebool) SetBool(b bool) *Bytebool {
	if b {
		*v = 1
	} else {
		*v = 0
	}
	return v
}

func (v *Bytebool) Bool() bool {
	if *v == 0 {
		return false
	}
	return true
}

func NewByteBool(b bool) *Bytebool {
	var i Bytebool
	i.SetBool(b)
	return &i
}

func NewByte(b byte) *Bytebool {
	var i Bytebool
	i.Set(b)
	return &i
}
