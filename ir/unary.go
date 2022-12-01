package ir

import "github.com/slince/php-plus/ir/value"

type Unary interface {
	unary()
}

type unary struct {
	instruction
}

type Neg struct {
	Result value.Operand
	Ope    value.Operand
	unary
}

func NewNeg(result value.Operand, ope value.Operand) *Neg {
	return &Neg{Result: result, Ope: ope}
}
