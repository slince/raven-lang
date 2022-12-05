package ir

import "github.com/slince/php-plus/ir/value"

type SetArrayElement struct {
	Array value.Value
	Value value.Value
	Index value.Value
}

func NewSetArrayElement(array value.Value, value value.Value, index value.Value) *SetArrayElement {
	return &SetArrayElement{Array: array, Value: value, Index: index}
}

type GetArrayElement struct {
	Variable value.Value
	Array    value.Value
	Index    value.Value
}

func NewGetArrayElement(variable value.Value, array value.Value, index value.Value) *GetArrayElement {
	return &GetArrayElement{Variable: variable, Array: array, Index: index}
}

type SetSliceElement struct {
	Slice value.Value
	Value value.Value
	Index value.Value
}

func NewSetSliceElement(slice value.Value, value value.Value, index value.Value) *SetSliceElement {
	return &SetSliceElement{Slice: slice, Value: value, Index: index}
}

type GetSliceElement struct {
	Variable value.Value
	Slice    value.Value
	Index    value.Value
}

func NewGetSliceElement(variable value.Value, slice value.Value, index value.Value) *GetSliceElement {
	return &GetSliceElement{Variable: variable, Slice: slice, Index: index}
}

type SetMapElement struct {
	Map   value.Value
	Value value.Value
	Key   value.Value
}

func NewSetMapElement(_map value.Value, value value.Value, key value.Value) *SetMapElement {
	return &SetMapElement{Map: _map, Value: value, Key: key}
}

type GetMapElement struct {
	Variable value.Value
	Map      value.Value
	Key      value.Value
}

func NewGetMapElement(variable value.Value, _map value.Value, key value.Value) *GetMapElement {
	return &GetMapElement{Variable: variable, Map: _map, Key: key}
}

type Len struct {
	Variable value.Value
}

func NewLen(variable value.Value) *Len {
	return &Len{Variable: variable}
}

type Cap struct {
	Variable value.Value
}

func NewCap(variable value.Value) *Cap {
	return &Cap{Variable: variable}
}
