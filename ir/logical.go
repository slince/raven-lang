package ir

import "github.com/slince/php-plus/ir/value"

type Gt struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Geq struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Lt struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Leq struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Eq struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Neq struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type LogicalAnd struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type LogicalOr struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type LogicalNot struct {
	value.Variable
	Ope value.Value
	instruction
}

func NewGt(lhs value.Value, rhs value.Value) *Gt {
	return &Gt{Lhs: lhs, Rhs: rhs}
}

func NewGeq(lhs value.Value, rhs value.Value) *Geq {
	return &Geq{Lhs: lhs, Rhs: rhs}
}

func NewLt(lhs value.Value, rhs value.Value) *Lt {
	return &Lt{Lhs: lhs, Rhs: rhs}
}

func NewLeq(lhs value.Value, rhs value.Value) *Leq {
	return &Leq{Lhs: lhs, Rhs: rhs}
}

func NewEq(lhs value.Value, rhs value.Value) *Eq {
	return &Eq{Lhs: lhs, Rhs: rhs}
}

func NewNeq(lhs value.Value, rhs value.Value) *Neq {
	return &Neq{Lhs: lhs, Rhs: rhs}
}

func NewLogicalAnd(lhs value.Value, rhs value.Value) *LogicalAnd {
	return &LogicalAnd{Lhs: lhs, Rhs: rhs}
}

func NewLogicalOr(lhs value.Value, rhs value.Value) *LogicalOr {
	return &LogicalOr{Lhs: lhs, Rhs: rhs}
}

func NewLogicalNot(ope value.Value) *LogicalNot {
	return &LogicalNot{Ope: ope}
}
