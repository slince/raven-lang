package insts

import "github.com/slince/php-plus/ir"

type Bitwise interface {
	bitwise()
}

type bitwise struct {
}

func (a bitwise) bitwise() {}

type And struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	bitwise
}

type Or struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	bitwise
}

type Xor struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	bitwise
}

type Not struct {
	Result ir.Operand
	Ope    ir.Operand
	bitwise
}

type Shl struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	bitwise
}

type Shr struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	bitwise
}

func NewBitwiseAnd(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *And {
	return &And{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewBitwiseOr(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Or {
	return &Or{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewBitwiseXor(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Xor {
	return &Xor{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewBitwiseNot(result ir.Operand, ope ir.Operand) *Not {
	return &Not{Result: result, Ope: ope}
}

func NewBitwiseShl(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Shl {
	return &Shl{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewBitwiseShr(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Shr {
	return &Shr{Result: result, Ope1: ope1, Ope2: ope2}
}
