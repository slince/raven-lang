package ir

import "github.com/slince/php-plus/ir/value"

type Bitwise interface {
	bitwise()
}

type bitwise struct {
	instruction
}

func (a bitwise) bitwise() {}

type And struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	bitwise
}

type Or struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	bitwise
}

type Xor struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	bitwise
}

type Not struct {
	Result value.Value
	Ope    value.Value
	bitwise
}

type Shl struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	bitwise
}

type Shr struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	bitwise
}

func NewBitAnd(result value.Value, lhs value.Value, rhs value.Value) *And {
	return &And{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitOr(result value.Value, lhs value.Value, rhs value.Value) *Or {
	return &Or{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitXor(result value.Value, lhs value.Value, rhs value.Value) *Xor {
	return &Xor{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitNot(result value.Value, ope value.Value) *Not {
	return &Not{Result: result, Ope: ope}
}

func NewBitShl(result value.Value, lhs value.Value, rhs value.Value) *Shl {
	return &Shl{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewBitShr(result value.Value, lhs value.Value, rhs value.Value) *Shr {
	return &Shr{Result: result, Lhs: lhs, Rhs: rhs}
}
