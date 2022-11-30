package insts

import "github.com/slince/php-plus/ir/value"

type SetArrayElement struct {
	Array value.Operand
	Value value.Operand
	Index value.Operand
}

func NewSetArrayElement(array value.Operand, value value.Operand, index value.Operand) *SetArrayElement {
	return &SetArrayElement{Array: array, Value: value, Index: index}
}

type GetArrayElement struct {
	Variable value.Operand
	Array    value.Operand
	Index    value.Operand
}

func NewGetArrayElement(variable value.Operand, array value.Operand, index value.Operand) *GetArrayElement {
	return &GetArrayElement{Variable: variable, Array: array, Index: index}
}

type SetSliceElement struct {
	Slice value.Operand
	Value value.Operand
	Index value.Operand
}

func NewSetSliceElement(slice value.Operand, value value.Operand, index value.Operand) *SetSliceElement {
	return &SetSliceElement{Slice: slice, Value: value, Index: index}
}

type GetSliceElement struct {
	Variable value.Operand
	Slice    value.Operand
	Index    value.Operand
}

func NewGetSliceElement(variable value.Operand, slice value.Operand, index value.Operand) *GetSliceElement {
	return &GetSliceElement{Variable: variable, Slice: slice, Index: index}
}

type SetMapElement struct {
	Map   value.Operand
	Value value.Operand
	Key   value.Operand
}

func NewSetMapElement(_map value.Operand, value value.Operand, key value.Operand) *SetMapElement {
	return &SetMapElement{Map: _map, Value: value, Key: key}
}

type GetMapElement struct {
	Variable value.Operand
	Map      value.Operand
	Key      value.Operand
}

func NewGetMapElement(variable value.Operand, _map value.Operand, key value.Operand) *GetMapElement {
	return &GetMapElement{Variable: variable, Map: _map, Key: key}
}

type Len struct {
	Variable value.Operand
}

func NewLen(variable value.Operand) *Len {
	return &Len{Variable: variable}
}

type Cap struct {
	Variable value.Operand
}

func NewCap(variable value.Operand) *Cap {
	return &Cap{Variable: variable}
}
