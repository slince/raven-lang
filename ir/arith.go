package ir

import "github.com/slince/php-plus/ir/value"

type Arith interface {
	arith()
}

type arith struct {
	instruction
}

func (a arith) arith() {}

type Add struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	arith
}

type Sub struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	arith
}

type Mul struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	arith
}

type Div struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	arith
}

type Mod struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	arith
}

func NewAdd(result value.Value, lhs value.Value, rhs value.Value) *Add {
	return &Add{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewSub(result value.Value, lhs value.Value, rhs value.Value) *Sub {
	return &Sub{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewMul(result value.Value, lhs value.Value, rhs value.Value) *Mul {
	return &Mul{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewDiv(result value.Value, lhs value.Value, rhs value.Value) *Div {
	return &Div{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewMod(result value.Value, lhs value.Value, rhs value.Value) *Mod {
	return &Mod{Result: result, Lhs: lhs, Rhs: rhs}
}
