package ir

import (
	"fmt"
	"github.com/slince/php-plus/ir/types"
	"github.com/slince/php-plus/ir/value"
	"strconv"
)

type Global struct {
	value.Variable
	Init *value.Const
	instruction
}

func (inst *Global) String() string {
	var init = ""
	if inst.Init != nil {
		init = inst.Init.String()
	}
	return fmt.Sprintf("global %s %s", inst.Variable.String(), init)
}

type Const struct {
	value.Variable
	Init *value.Const
	instruction
}

func (inst *Const) String() string {
	var init = inst.Init.String()
	return fmt.Sprintf("const %s %s", inst.Variable.String(), init)
}

type Local struct {
	value.Variable
	Init value.Value
	instruction
}

func (inst *Local) String() string {
	var init = ""
	if inst.Init != nil {
		init = inst.Init.String()
	}
	return fmt.Sprintf("local %s %s", inst.Variable.String(), init)
}

type Assign struct {
	Variable *value.Variable
	Value    value.Value
	instruction
}

func (inst *Assign) String() string {
	return fmt.Sprintf("assign %s %s", inst.Variable.String(), inst.Value.String())
}

type Lea struct {
	value.Variable
	Value value.Value
	instruction
}

func (inst *Lea) String() string {
	return fmt.Sprintf("lea %s %s", inst.Variable.String(), inst.Value.String())
}

type Ptr struct {
	value.Variable
	Value value.Value
	instruction
}

func (inst *Ptr) String() string {
	return fmt.Sprintf("ptr %s %s", inst.Variable.String(), inst.Value.String())
}

type Load struct {
	value.Variable
	Addr value.Value // PointType variable
	instruction
}

func (inst *Load) String() string {
	return fmt.Sprintf("load %s %s", inst.Variable.String(), inst.Addr.String())
}

type Store struct {
	Addr  value.Value // PointType variable
	Value value.Value
	instruction
}

func (inst *Store) String() string {
	return fmt.Sprintf("store %s %s", inst.Addr.String(), inst.Value.String())
}

type PtrStride struct {
	Addr   value.Value
	Stride int64
	instruction
}

func (inst *PtrStride) String() string {
	return fmt.Sprintf("stride %s %s", inst.Addr.String(), strconv.FormatInt(inst.Stride, 10))
}

type Label struct {
	instruction
	Name string
}

func (inst *Label) String() string {
	return fmt.Sprintf("%s:", inst.Name)
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

func NewLea(value value.Value) *Lea {
	return &Lea{Value: value}
}

func NewPtr(target value.Value) *Ptr {
	return &Ptr{Value: target}
}

func NewLabel(name string) *Label {
	return &Label{Name: name}
}
