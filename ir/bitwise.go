package ir

import "github.com/slince/php-plus/ir/value"

type And struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Or struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Xor struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Not struct {
	value.Variable
	Ope value.Value
	instruction
}

type Shl struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Shr struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func NewBitAnd(lhs value.Value, rhs value.Value) *And {
	return &And{Lhs: lhs, Rhs: rhs}
}

func NewBitOr(lhs value.Value, rhs value.Value) *Or {
	return &Or{Lhs: lhs, Rhs: rhs}
}

func NewBitXor(lhs value.Value, rhs value.Value) *Xor {
	return &Xor{Lhs: lhs, Rhs: rhs}
}

func NewBitNot(ope value.Value) *Not {
	return &Not{Ope: ope}
}

func NewBitShl(lhs value.Value, rhs value.Value) *Shl {
	return &Shl{Lhs: lhs, Rhs: rhs}
}

func NewBitShr(lhs value.Value, rhs value.Value) *Shr {
	return &Shr{Lhs: lhs, Rhs: rhs}
}
