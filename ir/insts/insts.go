package insts

import "github.com/slince/php-plus/ir"

type Instruction interface {
	inst()
}

type instruction struct {
}

func (i instruction) inst() {}

type Global struct {
	Variable ir.Operand
	Value    ir.Operand
	Init     bool
	instruction
}

func NewGlobal(variable ir.Operand, value ir.Operand) *Global {
	return &Global{Variable: variable, Value: value}
}

type GetGlobal struct {
	Variable ir.Operand
	instruction
}

func NewGetGlobal(variable ir.Operand) *GetGlobal {
	return &GetGlobal{Variable: variable}
}

type Const struct {
	Variable ir.Operand
	Value    ir.Operand
	instruction
}

func NewConst(variable ir.Operand, value ir.Operand) *Const {
	return &Const{Variable: variable, Value: value}
}

type GetConst struct {
	Variable ir.Operand
	instruction
}

func NewGetConst(variable ir.Operand) *GetConst {
	return &GetConst{Variable: variable}
}

type Local struct {
	Variable ir.Operand
	Value    ir.Operand
	instruction
}

func NewLocal(variable ir.Operand, value ir.Operand) *Local {
	return &Local{Variable: variable, Value: value}
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

type Lea struct {
	instruction
	Variable ir.Operand
	Target   ir.Operand
}

func NewLea(variable ir.Operand, target ir.Operand) *Lea {
	return &Lea{Variable: variable, Target: target}
}

type Ptr struct {
	instruction
	Variable ir.Operand
	Target   ir.Operand
}

func NewPtr(variable ir.Operand, target ir.Operand) *Ptr {
	return &Ptr{Variable: variable, Target: target}
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

type Entry struct {
}

func NewEntry() *Entry {
	return &Entry{}
}

type Exit struct {
}

func NewExit() *Exit {
	return &Exit{}
}
