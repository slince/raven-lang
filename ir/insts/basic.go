package insts

import (
	"github.com/slince/php-plus/ir"
)

type Global struct {
	Variable ir.Operand
	Value    ir.Operand
	Init     bool
	instruction
}

type GetGlobal struct {
	Variable ir.Operand
	instruction
}

type Const struct {
	Variable ir.Operand
	Value    ir.Operand
	instruction
}

type GetConst struct {
	Variable ir.Operand
	instruction
}

type Local struct {
	Variable ir.Operand
	Value    ir.Operand
	instruction
}

type GetLocal struct {
	Variable ir.Operand
	instruction
}

type Assign struct {
	Variable ir.Operand
	Value    ir.Operand
	instruction
}

type Lea struct {
	instruction
	Variable ir.Operand
	Target   ir.Operand
}

type Ptr struct {
	instruction
	Variable ir.Operand
	Target   ir.Operand
}

type Load struct {
	Result ir.Operand
	Addr   ir.Operand // PointType variable
}

type Store struct {
	Addr  ir.Operand
	Value ir.Operand // PointType variable
}

type PtrStride struct {
	Addr   ir.Operand
	Stride int64
}

func NewGlobal(variable ir.Operand, value ir.Operand) *Global {
	return &Global{Variable: variable, Value: value}
}

func NewGetGlobal(variable ir.Operand) *GetGlobal {
	return &GetGlobal{Variable: variable}
}

func NewConst(variable ir.Operand, value ir.Operand) *Const {
	return &Const{Variable: variable, Value: value}
}

func NewGetConst(variable ir.Operand) *GetConst {
	return &GetConst{Variable: variable}
}

func NewGetLocal(variable ir.Operand) *GetLocal {
	return &GetLocal{Variable: variable}
}

func NewLocal(variable ir.Operand, value ir.Operand) *Local {
	return &Local{Variable: variable, Value: value}
}

func NewAssign(variable ir.Operand, value ir.Operand) *Assign {
	return &Assign{Variable: variable, Value: value}
}

func NewLea(variable ir.Operand, target ir.Operand) *Lea {
	return &Lea{Variable: variable, Target: target}
}

func NewPtr(variable ir.Operand, target ir.Operand) *Ptr {
	return &Ptr{Variable: variable, Target: target}
}
