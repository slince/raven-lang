package ir

import (
	"github.com/slince/php-plus/ir/types"
	"github.com/slince/php-plus/ir/value"
)

type Global struct {
	value.Variable
	Init *value.Const
	instruction
}

type Const struct {
	value.Variable
	Init *value.Const
	instruction
}

type Local struct {
	value.Variable
	Init value.Value
	instruction
}

type Assign struct {
	Variable *value.Variable
	Value    value.Value
	instruction
}

type Lea struct {
	value.Variable
	Target value.Value
	instruction
}

type Ptr struct {
	value.Variable
	Target value.Value
	instruction
}

type Load struct {
	value.Variable
	Addr value.Value // PointType variable
	instruction
}

type Store struct {
	Addr  value.Value // PointType variable
	Value value.Value
	instruction
}

type PtrStride struct {
	Addr   value.Value
	Stride int64
	instruction
}

type Label struct {
	instruction
	Name string
}

func NewGlobal(name string, kind types.Type, init *value.Const) *Global {
	return &Global{Variable: *value.NewVariable(name, kind), Init: init}
}

func NewConst(name string, kind types.Type, init *value.Const) *Const {
	return &Const{Variable: *value.NewVariable(name, kind), Init: init}
}

func NewLocal(name string, kind types.Type, init value.Value) *Local {
	return &Local{Variable: *value.NewVariable(name, kind), Init: init}
}

func NewAssign(variable *value.Variable, value value.Value) *Assign {
	return &Assign{Variable: variable, Value: value}
}

func NewLea(target value.Value) *Lea {
	return &Lea{Target: target}
}

func NewPtr(target value.Value) *Ptr {
	return &Ptr{Target: target}
}
