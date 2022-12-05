package ir

import "github.com/slince/php-plus/ir/value"

type Unary interface {
	unary()
}

type unary struct {
	instruction
}

type Neg struct {
	Result value.Value
	Ope    value.Value
	unary
}

func NewNeg(result value.Value, ope value.Value) *Neg {
	return &Neg{Result: result, Ope: ope}
}
