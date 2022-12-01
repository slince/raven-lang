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
	Index    value.Operand
	Value    value.Operand
}

func NewSetArray(variable value.Variable, index value.Operand, value value.Operand) *SetArray {
	return &SetArray{
		Variable: variable,
		Index:    index,
		Value:    value,
	}
}

type GetArray struct {
	instruction
	Result   value.Operand
	Variable value.Variable
	Index    value.Operand
}

func NewGetArray(result value.Operand, variable value.Variable, index value.Operand) *GetArray {
	return &GetArray{
		Result:   result,
		Variable: variable,
		Index:    index,
	}
}
