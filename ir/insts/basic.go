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
