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
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	arith
}

type Sub struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	arith
}

type Mul struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	arith
}

type Div struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	arith
}

type Mod struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	arith
}

func NewAdd(result value.Operand, lhs value.Operand, rhs value.Operand) *Add {
	return &Add{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewSub(result value.Operand, lhs value.Operand, rhs value.Operand) *Sub {
	return &Sub{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewMul(result value.Operand, lhs value.Operand, rhs value.Operand) *Mul {
	return &Mul{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewDiv(result value.Operand, lhs value.Operand, rhs value.Operand) *Div {
	return &Div{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewMod(result value.Operand, lhs value.Operand, rhs value.Operand) *Mod {
	return &Mod{Result: result, Lhs: lhs, Rhs: rhs}
}
