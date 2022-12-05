package ir

import (
	"github.com/slince/php-plus/ir/value"
)

type Arg struct {
	Value value.Value
	instruction
}

func NewArg(value value.Value) *Arg {
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
	Ope value.Value
	instruction
}

func NewRet(ope value.Value) *Ret {
	return &Ret{Ope: ope}
}
