package ir

import (
	"github.com/slince/php-plus/ir/value"
)

type Arg struct {
	Value value.Operand
	instruction
}

func NewArg(value value.Operand) *Arg {
	return &Arg{Value: value}
}

type Call struct {
	Callee *Function
	ArgNum uint64
	instruction
}

func NewCall(callee *Function, argNum uint64) *Call {
	return &Call{Callee: callee, ArgNum: argNum}
}

type Ret struct {
	Ope value.Operand
	instruction
}

func NewRet(ope value.Operand) *Ret {
	return &Ret{Ope: ope}
}