package ir

import "github.com/slince/php-plus/ir/value"

type Unary interface {
	unary()
}

type unary struct {
	instruction
}

type Neg struct {
	value.Variable
	Ope value.Value
	unary
}

func NewNeg(ope value.Value) *Neg {
	return &Neg{Ope: ope}
}
