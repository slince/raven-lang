package insts

import "github.com/slince/php-plus/ir"

type SetArrayElement struct {
	Array ir.Operand
	Value ir.Operand
	Index ir.Operand
}

func NewSetArrayElement(array ir.Operand, value ir.Operand, index ir.Operand) *SetArrayElement {
	return &SetArrayElement{Array: array, Value: value, Index: index}
}

type GetArrayElement struct {
	Variable ir.Operand
	Array    ir.Operand
	Index    ir.Operand
}

func NewGetArrayElement(variable ir.Operand, array ir.Operand, index ir.Operand) *GetArrayElement {
	return &GetArrayElement{Variable: variable, Array: array, Index: index}
}

type SetSliceElement struct {
	Slice ir.Operand
	Value ir.Operand
	Index ir.Operand
}

func NewSetSliceElement(slice ir.Operand, value ir.Operand, index ir.Operand) *SetSliceElement {
	return &SetSliceElement{Slice: slice, Value: value, Index: index}
}

type GetSliceElement struct {
	Variable ir.Operand
	Slice    ir.Operand
	Index    ir.Operand
}

func NewGetSliceElement(variable ir.Operand, slice ir.Operand, index ir.Operand) *GetSliceElement {
	return &GetSliceElement{Variable: variable, Slice: slice, Index: index}
}

type SetMapElement struct {
	Map   ir.Operand
	Value ir.Operand
	Key   ir.Operand
}

func NewSetMapElement(_map ir.Operand, value ir.Operand, key ir.Operand) *SetMapElement {
	return &SetMapElement{Map: _map, Value: value, Key: key}
}

type GetMapElement struct {
	Variable ir.Operand
	Map      ir.Operand
	Key      ir.Operand
}

func NewGetMapElement(variable ir.Operand, _map ir.Operand, key ir.Operand) *GetMapElement {
	return &GetMapElement{Variable: variable, Map: _map, Key: key}
}

type Len struct {
	Variable ir.Operand
}

func NewLen(variable ir.Operand) *Len {
	return &Len{Variable: variable}
}

type Cap struct {
	Variable ir.Operand
}

func NewCap(variable ir.Operand) *Cap {
	return &Cap{Variable: variable}
}
