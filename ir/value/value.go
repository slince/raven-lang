package value

import (
	"fmt"
	"github.com/slince/php-plus/ir/types"
	"strconv"
)

type Value interface {
	String() string
	Type() types.Type
}

type Const struct {
	Value any
	Kind  types.Type
}

func (c *Const) Type() types.Type {
	return c.Kind
}

func (c *Const) String() string {
	var value string
	switch c.Kind {
	case types.I4, types.I8, types.I32, types.I64:
		value = strconv.FormatInt(c.Value.(int64), 10)
	case types.U4, types.U8, types.U32, types.U64:
		value = strconv.FormatUint(c.Value.(uint64), 10)
	case types.F32, types.F64:
		value = strconv.FormatFloat(c.Value.(float64), 'g', -1, 64)
	case types.Bool:
		value = strconv.FormatBool(c.Value.(bool))
	case types.String:
		value = c.Value.(string)
	}
	return value
}

func NewConst(value interface{}, kind types.Type) *Const {
	return &Const{
		Value: value,
		Kind:  kind,
	}
}

type Variable struct {
	Name      string
	Kind      types.Type
	Immutable bool
	Init      bool
}

func (v *Variable) Type() types.Type {
	return v.Kind
}

func (v *Variable) String() string {
	return fmt.Sprintf("%s: %s", v.Name, v.Kind.String())
}

func NewVariable(name string, kind types.Type) *Variable {
	return &Variable{
		Name:      name,
		Kind:      kind,
		Immutable: false,
		Init:      false,
	}
}
