package ir

import "github.com/slince/php-plus/ir/value"

type Array struct {
	instruction
}

type PushArray struct {
	instruction
	Variable value.Variable
}

func NewPushArray(variable value.Variable) *PushArray {
	return &PushArray{
		Variable: variable,
	}
}

type SetArray struct {
	instruction
	Variable value.Variable
	Index    value.Value
	Value    value.Value
}

func NewSetArray(variable value.Variable, index value.Value, value value.Value) *SetArray {
	return &SetArray{
		Variable: variable,
		Index:    index,
		Value:    value,
	}
}

type GetArray struct {
	instruction
	Result   value.Value
	Variable value.Variable
	Index    value.Value
}

func NewGetArray(result value.Value, variable value.Variable, index value.Value) *GetArray {
	return &GetArray{
		Result:   result,
		Variable: variable,
		Index:    index,
	}
}
