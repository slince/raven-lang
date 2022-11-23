package insts

import "github.com/slince/php-plus/ir"

type Array struct {
	instruction
}

type PushArray struct {
	instruction
	Variable ir.Variable
}

func NewPushArray(variable ir.Variable) *PushArray {
	return &PushArray{
		Variable: variable,
	}
}

type SetArray struct {
	instruction
	Variable ir.Variable
	Index    ir.Operand
	Value    ir.Operand
}

func NewSetArray(variable ir.Variable, index ir.Operand, value ir.Operand) *SetArray {
	return &SetArray{
		Variable: variable,
		Index:    index,
		Value:    value,
	}
}

type GetArray struct {
	instruction
	Result   ir.Operand
	Variable ir.Variable
	Index    ir.Operand
}

func NewGetArray(result ir.Operand, variable ir.Variable, index ir.Operand) *GetArray {
	return &GetArray{
		Result:   result,
		Variable: variable,
		Index:    index,
	}
}
