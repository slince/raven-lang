package ir

import "github.com/slince/php-plus/ir/value"

type Logical interface {
	logical()
}

type logical struct {
	instruction
}

func (a logical) logical() {}

type Gt struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	logical
}

type Geq struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	logical
}

type Lt struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	logical
}

type Leq struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	logical
}

type Eq struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	logical
}

type Neq struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	logical
}

type LogicalAnd struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	logical
}

type LogicalOr struct {
	Result value.Value
	Lhs    value.Value
	Rhs    value.Value
	logical
}

type LogicalNot struct {
	Result value.Value
	Ope    value.Value
	logical
}

func NewGt(result value.Value, lhs value.Value, rhs value.Value) *Gt {
	return &Gt{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewGeq(result value.Value, lhs value.Value, rhs value.Value) *Geq {
	return &Geq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLt(result value.Value, lhs value.Value, rhs value.Value) *Lt {
	return &Lt{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLeq(result value.Value, lhs value.Value, rhs value.Value) *Leq {
	return &Leq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewEq(result value.Value, lhs value.Value, rhs value.Value) *Eq {
	return &Eq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewNeq(result value.Value, lhs value.Value, rhs value.Value) *Neq {
	return &Neq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalAnd(result value.Value, lhs value.Value, rhs value.Value) *LogicalAnd {
	return &LogicalAnd{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalOr(result value.Value, lhs value.Value, rhs value.Value) *LogicalOr {
	return &LogicalOr{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalNot(result value.Value, ope value.Value) *LogicalNot {
	return &LogicalNot{Result: result, Ope: ope}
}
