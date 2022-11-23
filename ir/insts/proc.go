package insts

import "github.com/slince/php-plus/ir"

type Arg struct {
	Value ir.Operand
	instruction
}

func NewArg(value ir.Operand) *Arg {
	return &Arg{Value: value}
}

type Call struct {
	Callee *ir.Function
	ArgNum uint64
	instruction
}

func NewCall(callee *ir.Function, argNum uint64) *Call {
	return &Call{Callee: callee, ArgNum: argNum}
}

type Ret struct {
	Ope ir.Operand
	instruction
}

func NewRet(ope ir.Operand) *Ret {
	return &Ret{Ope: ope}
}
