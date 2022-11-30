package insts

import "github.com/slince/php-plus/ir/value"

type Logical interface {
	logical()
}

type logical struct {
	instruction
}

func (a logical) logical() {}

type Gt struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	logical
}

type Geq struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	logical
}

type Lt struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	logical
}

type Leq struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	logical
}

type Eq struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	logical
}

type Neq struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	logical
}

type LogicalAnd struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	logical
}

type LogicalOr struct {
	Result value.Operand
	Lhs    value.Operand
	Rhs    value.Operand
	logical
}

type LogicalNot struct {
	Result value.Operand
	Ope    value.Operand
	logical
}

func NewGt(result value.Operand, lhs value.Operand, rhs value.Operand) *Gt {
	return &Gt{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewGeq(result value.Operand, lhs value.Operand, rhs value.Operand) *Geq {
	return &Geq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLt(result value.Operand, lhs value.Operand, rhs value.Operand) *Lt {
	return &Lt{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLeq(result value.Operand, lhs value.Operand, rhs value.Operand) *Leq {
	return &Leq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewEq(result value.Operand, lhs value.Operand, rhs value.Operand) *Eq {
	return &Eq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewNeq(result value.Operand, lhs value.Operand, rhs value.Operand) *Neq {
	return &Neq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalAnd(result value.Operand, lhs value.Operand, rhs value.Operand) *LogicalAnd {
	return &LogicalAnd{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalOr(result value.Operand, lhs value.Operand, rhs value.Operand) *LogicalOr {
	return &LogicalOr{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalNot(result value.Operand, ope value.Operand) *LogicalNot {
	return &LogicalNot{Result: result, Ope: ope}
}
