package insts

import "github.com/slince/php-plus/ir/value"

type Bitwise interface {
	bitwise()
}

type bitwise struct {
	instruction
}

func (a bitwise) bitwise() {}

type And struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	bitwise
}

type Or struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	bitwise
}

type Xor struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	bitwise
}

type Not struct {
	Result value.Operand
	Ope    value.Operand
	bitwise
}

type Shl struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	bitwise
}

type Shr struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	bitwise
}

func NewBitAnd(result value.Operand, lhs value.Operand, rhs value.Operand) *And {
	return &And{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitOr(result value.Operand, lhs value.Operand, rhs value.Operand) *Or {
	return &Or{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitXor(result value.Operand, lhs value.Operand, rhs value.Operand) *Xor {
	return &Xor{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitNot(result value.Operand, ope value.Operand) *Not {
	return &Not{Result: result, Ope: ope}
}

func NewBitShl(result value.Operand, lhs value.Operand, rhs value.Operand) *Shl {
	return &Shl{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitShr(result value.Operand, lhs value.Operand, rhs value.Operand) *Shr {
	return &Shr{Result: result, Lhs: lhs, Rhs: rhs}
}
