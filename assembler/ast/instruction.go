package ast

import "github.com/slince/php-plus/assembler/token"

type OperandType uint8

const (
	REG OperandType = iota
	IMM
	CONST
)

type Operand struct {
	Kind     OperandType
	Value    *Literal
	position *token.Position
}

func (o *Operand) Position() *token.Position {
	return o.position
}

func NewOperand(kind OperandType, value *Literal, position *token.Position) *Operand {
	return &Operand{Kind: kind, Value: value, position: position}
}

type Instruction struct {
	Label    *Label
	Kind     *Identifier
	Operand1 *Operand
	Operand2 *Operand
	Operand3 *Operand
	Comment  *Comment
}

func (d *Instruction) Position() *token.Position {
	if d.Label != nil {
		return d.Label.position
	}
	return d.Kind.position
}

func NewInstruction(label *Label, kind *Identifier, operand1 *Operand, operand2 *Operand, operand3 *Operand, comment *Comment) *Instruction {
	return &Instruction{
		Label:    label,
		Kind:     kind,
		Operand1: operand1,
		Operand2: operand2,
		Operand3: operand3,
		Comment:  comment,
	}
}
