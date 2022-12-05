package ir

import (
	"github.com/slince/php-plus/ir/types"
	"github.com/slince/php-plus/ir/value"
)

type Global struct {
	Name string
	Kind types.Type
	Init *value.Const
	instruction
}

type GetGlobal struct {
	Variable value.Operand
	instruction
}

type Const struct {
	Variable value.Operand
	Value    value.Operand
	instruction
}

type GetConst struct {
	Variable value.Operand
	instruction
}

type Local struct {
	Variable value.Operand
	Value    value.Operand
	instruction
}

type GetLocal struct {
	Variable value.Operand
	instruction
}

type Assign struct {
	Variable value.Operand
	Value    value.Operand
	instruction
}

type Lea struct {
	instruction
	Variable value.Operand
	Target   value.Operand
}

type Ptr struct {
	instruction
	Variable value.Operand
	Target   value.Operand
}

type Load struct {
	instruction
	Result value.Operand
	Addr   value.Operand // PointType variable
}

type Store struct {
	instruction
	Addr  value.Operand
	Value value.Operand // PointType variable
}

type PtrStride struct {
	instruction
	Addr   value.Operand
	Stride int64
}

type Label struct {
	instruction
	Name string
}

func NewGlobal(name string, kind types.Type, init *value.Const) *Global {
	return &Global{Name: name, Kind: kind, Init: init}
}

func NewGetGlobal(variable value.Operand) *GetGlobal {
	return &GetGlobal{Variable: variable}
}

func NewConst(variable value.Operand, value value.Operand) *Const {
	return &Const{Variable: variable, Value: value}
}

func NewGetConst(variable value.Operand) *GetConst {
	return &GetConst{Variable: variable}
}

func NewGetLocal(variable value.Operand) *GetLocal {
	return &GetLocal{Variable: variable}
}

func NewLocal(variable value.Operand, value value.Operand) *Local {
	return &Local{Variable: variable, Value: value}
}

func NewAssign(variable value.Operand, value value.Operand) *Assign {
	return &Assign{Variable: variable, Value: value}
}

func NewLea(variable value.Operand, target value.Operand) *Lea {
	return &Lea{Variable: variable, Target: target}
}

func NewPtr(variable value.Operand, target value.Operand) *Ptr {
	return &Ptr{Variable: variable, Target: target}
}
