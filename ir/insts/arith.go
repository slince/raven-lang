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
	Ope1   ir.Operand
	Ope2   ir.Operand
	arith
}

type Sub struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	arith
}

type Mul struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	arith
}

type Div struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	arith
}

type Mod struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	arith
}

func NewAdd(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Add {
	return &Add{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewSub(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Sub {
	return &Sub{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewMul(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Mul {
	return &Mul{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewDiv(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Div {
	return &Div{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewMod(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Mod {
	return &Mod{Result: result, Ope1: ope1, Ope2: ope2}
}
