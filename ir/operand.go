package ir

import "github.com/slince/php-plus/ir/types"

var (
	Zero = NewConst(0, types.U4)
	One  = NewConst(1, types.U4)
)

type Operand interface {
	operand()
	Type() types.Type
}

type operand struct {
}

func (o operand) operand() {}
func (o operand) Type()    {}

type Const struct {
	Value interface{}
	Kind  types.Type
	operand
}

func (c *Const) Type() types.Type {
	return c.Kind
}

func NewConst(value interface{}, kind types.Type) *Const {
	return &Const{
		Value: value,
		Kind:  kind,
	}
}

type Variable struct {
	Name string
	Kind types.Type
	Init bool
	operand
}

func (v *Variable) Type() types.Type {
	return v.Kind
}

func NewVariable(name string, kind types.Type) *Variable {
	return &Variable{
		Name: name,
		Kind: kind,
		Init: false,
	}
}

type Temporary struct {
	Original *Variable
	operand
}

func (v *Temporary) Type() types.Type {
	if v.Original != nil {
		return v.Original.Type()
	}
	return nil
}

func NewTemporary(variable *Variable) *Temporary {
	return &Temporary{Original: variable}
}

type Label struct {
	Name string
	operand
}

func (v *Label) Type() types.Type {
	return nil
}

func NewLabel(name string) *Label {
	return &Label{Name: name}
}
