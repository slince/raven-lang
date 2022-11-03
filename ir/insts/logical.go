package insts

import "github.com/slince/php-plus/ir"

type Logical interface {
	logical()
}

type logical struct {
	instruction
}

func (a logical) logical() {}

type Gt struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	logical
}

type Geq struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	logical
}

type Lt struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	logical
}

type Leq struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	logical
}

type Eq struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	logical
}

type Neq struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	logical
}

type LogicalAnd struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	logical
}

type LogicalOr struct {
	Result ir.Operand
	Lhs    ir.Operand
	Rhs    ir.Operand
	logical
}

type LogicalNot struct {
	Result ir.Operand
	Ope    ir.Operand
	logical
}

func NewGt(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Gt {
	return &Gt{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewGeq(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Geq {
	return &Geq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLt(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Lt {
	return &Lt{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLeq(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Leq {
	return &Leq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewEq(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Eq {
	return &Eq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewNeq(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *Neq {
	return &Neq{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalAnd(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *LogicalAnd {
	return &LogicalAnd{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalOr(result ir.Operand, lhs ir.Operand, rhs ir.Operand) *LogicalOr {
	return &LogicalOr{Result: result, Lhs: lhs, Rhs: rhs}
}

func NewLogicalNot(result ir.Operand, ope ir.Operand) *LogicalNot {
	return &LogicalNot{Result: result, Ope: ope}
}
