package insts

import "github.com/slince/php-plus/ir"

type Arith interface {
	arith()
}

type arith struct {
	instruction
}

func (a arith) arith() {}

type Add struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	arith
}

type Sub struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	arith
}

type Mul struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	arith
}

type Div struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	arith
}

type Mod struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	arith
}

func NewAdd(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Add {
	return &Add{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewSub(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Sub {
	return &Sub{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewMul(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Mul {
	return &Mul{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewDiv(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Div {
	return &Div{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewMod(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Mod {
	return &Mod{Result: result, Lhs: lhs, Rhs: rhs}
}
