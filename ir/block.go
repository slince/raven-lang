package ir

import (
	"github.com/slince/php-plus/ir/insts"
)

type Block interface {
	block()
}

type BasicBlock struct {
	Name         string
	Instructions []insts.Instruction
	Phis         []insts.Instruction
	Leader       insts.Instruction
	Terminator   insts.Instruction
	Predecessors []*BasicBlock
	Successors   []*BasicBlock
	Dominators   []*BasicBlock
}

func (b *BasicBlock) block() {}

func (b *BasicBlock) HasTerminator() bool {
	return b.Terminator != nil
}

func (b *BasicBlock) AddInstruction(instruction insts.Instruction) {
	b.Instructions = append(b.Instructions, instruction)
}

func (b *BasicBlock) NewAdd(result Operand, ope1 Operand, ope2 Operand) *insts.Add {
	var inst = insts.NewAdd(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewSub(result Operand, ope1 Operand, ope2 Operand) *insts.Sub {
	var inst = insts.NewSub(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}
func (b *BasicBlock) NewMul(result Operand, ope1 Operand, ope2 Operand) *insts.Mul {
	var inst = insts.NewMul(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewDiv(result Operand, ope1 Operand, ope2 Operand) *insts.Div {
	var inst = insts.NewDiv(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewMod(result Operand, ope1 Operand, ope2 Operand) *insts.Mod {
	var inst = insts.NewMod(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGt(result Operand, ope1 Operand, ope2 Operand) *insts.Gt {
	var inst = insts.NewGt(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGeq(result Operand, ope1 Operand, ope2 Operand) *insts.Geq {
	var inst = insts.NewGeq(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLt(result Operand, ope1 Operand, ope2 Operand) *insts.Lt {
	var inst = insts.NewLt(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLeq(result Operand, ope1 Operand, ope2 Operand) *insts.Leq {
	var inst = insts.NewLeq(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewEq(result Operand, ope1 Operand, ope2 Operand) *insts.Eq {
	var inst = insts.NewEq(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeq(result Operand, ope1 Operand, ope2 Operand) *insts.Neq {
	var inst = insts.NewNeq(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalAnd(result Operand, ope1 Operand, ope2 Operand) *insts.LogicalAnd {
	var inst = insts.NewLogicalAnd(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalOr(result Operand, ope1 Operand, ope2 Operand) *insts.LogicalOr {
	var inst = insts.NewLogicalOr(result, ope1, ope2)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalNot(result Operand, ope Operand) *insts.LogicalNot {
	var inst = insts.NewLogicalNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAssign(variable Operand, value Operand) *insts.Assign {
	var inst = insts.NewAssign(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeg(result Operand, ope Operand) *insts.Neg {
	var inst = insts.NewNeg(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewArg(value Operand) *insts.Arg {
	var inst = insts.NewArg(value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewCall(callee *Function, argNum uint64) *insts.Call {
	var inst = insts.NewCall(callee, argNum)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewRet(ope Operand) *insts.Ret {
	var inst = insts.NewRet(ope)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewJmp(target Block) *insts.Jmp {
	var inst = insts.NewJmp(target)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewCondJmp(cond Operand, trueTarget Block, falseTarget Block) *insts.CondJmp {
	var inst = insts.NewCondJmp(cond, trueTarget, falseTarget)
	b.Terminator = inst
	return inst
}

func NewBlock(name string) *BasicBlock {
	return &BasicBlock{Name: name, Instructions: []insts.Instruction{}}
}

// Reference references to other block
type Reference struct {
	Name string
}

func (r *Reference) block() {}

func NewReference(name string) *Reference {
	return &Reference{Name: name}
}
