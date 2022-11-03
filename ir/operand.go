package ir

import "github.com/slince/php-plus/ir/types"

var (
	Zero = NewLiteral(0, types.I1)
	One  = NewLiteral(1, types.I1)
)

type Operand interface {
	operand()
}

type operand struct {
}

func (o operand) operand() {}

type Const struct {
	Value interface{}
	Kind  types.Type
	operand
}

func NewLiteral(value interface{}, kind types.Type) *Const {
	return &Const{
		Value: value,
		Kind:  kind,
	}
}

type Variable struct {
	// Parameter name.
	Name string
	// Parameter type.
	Kind types.Type
	Init bool
	operand
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

func NewTemporary(variable *Variable) *Temporary {
	return &Temporary{Original: variable}
}
