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
	Variable value.Value
	instruction
}

type Const struct {
	Variable value.Value
	Value    value.Value
	instruction
}

type GetConst struct {
	Variable value.Value
	instruction
}

type Local struct {
	Variable value.Value
	Value    value.Value
	instruction
}

type GetLocal struct {
	Variable value.Value
	instruction
}

type Assign struct {
	Variable value.Value
	Value    value.Value
	instruction
}

type Lea struct {
	instruction
	Variable value.Value
	Target   value.Value
}

type Ptr struct {
	instruction
	Variable value.Value
	Target   value.Value
}

type Load struct {
	instruction
	value.Variable
	Addr value.Value // PointType variable
}

type Store struct {
	instruction
	Addr  value.Value
	Value value.Value // PointType variable
}

type PtrStride struct {
	instruction
	Addr   value.Value
	Stride int64
}

type Label struct {
	instruction
	Name string
}

func NewGlobal(name string, kind types.Type, init *value.Const) *Global {
	return &Global{Name: name, Kind: kind, Init: init}
}

func NewGetGlobal(variable value.Value) *GetGlobal {
	return &GetGlobal{Variable: variable}
}

func NewConst(variable value.Value, value value.Value) *Const {
	return &Const{Variable: variable, Value: value}
}

func NewGetConst(variable value.Value) *GetConst {
	return &GetConst{Variable: variable}
}

func NewGetLocal(variable value.Value) *GetLocal {
	return &GetLocal{Variable: variable}
}

func NewLocal(variable value.Value, value value.Value) *Local {
	return &Local{Variable: variable, Value: value}
}

func NewAssign(variable value.Value, value value.Value) *Assign {
	return &Assign{Variable: variable, Value: value}
}

func NewLea(variable value.Value, target value.Value) *Lea {
	return &Lea{Variable: variable, Target: target}
}

func NewPtr(variable value.Value, target value.Value) *Ptr {
	return &Ptr{Variable: variable, Target: target}
}
