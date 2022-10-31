package insts

import "github.com/slince/php-plus/ir"

type Instruction interface {
	inst()
}

type instruction struct {
}

func (i instruction) inst() {}

type SetGlobal struct {
	Variable ir.Operand
	Value    ir.Operand
	Init     bool
	instruction
}

func NewSetGlobal(variable ir.Operand, value ir.Operand) *SetGlobal {
	return &SetGlobal{Variable: variable, Value: value}
}

type GetGlobal struct {
	Variable ir.Operand
	instruction
}

func NewGetGlobal(variable ir.Operand) *GetGlobal {
	return &GetGlobal{Variable: variable}
}

type SetConst struct {
	Variable ir.Operand
	Value    ir.Operand
	instruction
}

func NewSetConst(variable ir.Operand, value ir.Operand) *SetConst {
	return &SetConst{Variable: variable, Value: value}
}

type GetConst struct {
	Variable ir.Operand
	instruction
}

func NewGetConst(variable ir.Operand) *GetConst {
	return &GetConst{Variable: variable}
}

type SetLocal struct {
	Variable ir.Operand
	Value    ir.Operand
	instruction
}

func NewSetLocal(variable ir.Operand, value ir.Operand) *SetLocal {
	return &SetLocal{Variable: variable, Value: value}
}

type GetLocal struct {
	Variable ir.Operand
	instruction
}

func NewGetLocal(variable ir.Operand) *GetLocal {
	return &GetLocal{Variable: variable}
}

type Assign struct {
	Variable ir.Operand
	Value    ir.Operand
	instruction
}

func NewAssign(variable ir.Operand, value ir.Operand) *Assign {
	return &Assign{Variable: variable, Value: value}
}

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

type Jmp struct {
	Target ir.Block
	instruction
}

func NewJmp(target ir.Block) *Jmp {
	return &Jmp{Target: target}
}

type CondJmp struct {
	Cond        ir.Operand
	TrueTarget  ir.Block
	FalseTarget ir.Block
	instruction
}

func NewCondJmp(cond ir.Operand, trueTarget ir.Block, falseTarget ir.Block) *CondJmp {
	return &CondJmp{Cond: cond, TrueTarget: trueTarget, FalseTarget: falseTarget}
}
