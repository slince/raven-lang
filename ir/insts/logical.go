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
	Ope1   ir.Operand
	Ope2   ir.Operand
	logical
}

type Geq struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	logical
}

type Lt struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	logical
}

type Leq struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	logical
}

type Eq struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	logical
}

type Neq struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	logical
}

type LogicalAnd struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	logical
}

type LogicalOr struct {
	Result ir.Operand
	Ope1   ir.Operand
	Ope2   ir.Operand
	logical
}

type LogicalNot struct {
	Result ir.Operand
	Ope    ir.Operand
	logical
}

func NewGt(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Gt {
	return &Gt{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewGeq(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Geq {
	return &Geq{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewLt(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Lt {
	return &Lt{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewLeq(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Leq {
	return &Leq{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewEq(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Eq {
	return &Eq{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewNeq(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *Neq {
	return &Neq{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewAnd(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *LogicalAnd {
	return &LogicalAnd{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewOr(result ir.Operand, ope1 ir.Operand, ope2 ir.Operand) *LogicalOr {
	return &LogicalOr{Result: result, Ope1: ope1, Ope2: ope2}
}

func NewNot(result ir.Operand, ope ir.Operand) *LogicalNot {
	return &LogicalNot{Result: result, Ope: ope}
}
