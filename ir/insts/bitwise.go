package insts

import "github.com/slince/php-plus/ir"

type Bitwise interface {
	bitwise()
}

type bitwise struct {
	instruction
}

func (a bitwise) bitwise() {}

type And struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	bitwise
}

type Or struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	bitwise
}

type Xor struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	bitwise
}

type Not struct {
	Result ir.Operand
	Ope    ir.Operand
	bitwise
}

type Shl struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	bitwise
}

type Shr struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	bitwise
}

func NewBitAnd(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *And {
	return &And{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitOr(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Or {
	return &Or{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitXor(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Xor {
	return &Xor{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitNot(result ir.Operand, ope ir.Operand) *Not {
	return &Not{Result: result, Ope: ope}
}

func NewBitShl(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Shl {
	return &Shl{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitShr(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Shr {
	return &Shr{Result: result, Lhs: lhs, Rhs: rhs}
}
