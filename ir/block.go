package ir

import (
	"github.com/slince/php-plus/ir/insts"
)

type Block interface {
	block()
}

// Reference references to other block
type Reference struct {
	Name string
}

type Incoming struct {
	Name string
}

func NewReference(name string) *Reference {
	return &Reference{Name: name}
}

func NewIncoming(name string) *Incoming {
	return &Incoming{Name: name}
}

func (r *Reference) block() {}

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

func (b *BasicBlock) NewLea(variable Operand, target Operand) *insts.Lea {
	var inst = insts.NewLea(variable, target)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewPtr(variable Operand, target Operand) *insts.Ptr {
	var inst = insts.NewPtr(variable, target)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAdd(result Operand, lhs Operand, rhs Operand) *insts.Add {
	var inst = insts.NewAdd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewSub(result Operand, lhs Operand, rhs Operand) *insts.Sub {
	var inst = insts.NewSub(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}
func (b *BasicBlock) NewMul(result Operand, lhs Operand, rhs Operand) *insts.Mul {
	var inst = insts.NewMul(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewDiv(result Operand, lhs Operand, rhs Operand) *insts.Div {
	var inst = insts.NewDiv(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewMod(result Operand, lhs Operand, rhs Operand) *insts.Mod {
	var inst = insts.NewMod(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGt(result Operand, lhs Operand, rhs Operand) *insts.Gt {
	var inst = insts.NewGt(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGeq(result Operand, lhs Operand, rhs Operand) *insts.Geq {
	var inst = insts.NewGeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLt(result Operand, lhs Operand, rhs Operand) *insts.Lt {
	var inst = insts.NewLt(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLeq(result Operand, lhs Operand, rhs Operand) *insts.Leq {
	var inst = insts.NewLeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewEq(result Operand, lhs Operand, rhs Operand) *insts.Eq {
	var inst = insts.NewEq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeq(result Operand, lhs Operand, rhs Operand) *insts.Neq {
	var inst = insts.NewNeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalAnd(result Operand, lhs Operand, rhs Operand) *insts.LogicalAnd {
	var inst = insts.NewLogicalAnd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalOr(result Operand, lhs Operand, rhs Operand) *insts.LogicalOr {
	var inst = insts.NewLogicalOr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalNot(result Operand, ope Operand) *insts.LogicalNot {
	var inst = insts.NewLogicalNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitAnd(result Operand, lhs Operand, rhs Operand) *insts.And {
	var inst = insts.NewBitAnd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitOr(result Operand, lhs Operand, rhs Operand) *insts.Or {
	var inst = insts.NewBitOr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitXor(result Operand, lhs Operand, rhs Operand) *insts.Xor {
	var inst = insts.NewBitXor(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitNot(result Operand, ope Operand) *insts.Not {
	var inst = insts.NewBitNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShl(result Operand, lhs Operand, rhs Operand) *insts.Shl {
	var inst = insts.NewBitShl(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShr(result Operand, lhs Operand, rhs Operand) *insts.Shr {
	var inst = insts.NewBitShr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeg(result Operand, ope Operand) *insts.Neg {
	var inst = insts.NewNeg(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGlobal(variable Operand, value Operand) *insts.Global {
	var inst = insts.NewGlobal(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetGlobal(variable Operand) *insts.GetGlobal {
	var inst = insts.NewGetGlobal(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewConst(variable Operand, value Operand) *insts.Const {
	var inst = insts.NewConst(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetConst(variable Operand) *insts.GetConst {
	var inst = insts.NewGetConst(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLocal(variable Operand, value Operand) *insts.Local {
	var inst = insts.NewLocal(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetLocal(variable Operand) *insts.GetLocal {
	var inst = insts.NewGetLocal(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAssign(variable Operand, value Operand) *insts.Assign {
	var inst = insts.NewAssign(variable, value)
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

func (b *BasicBlock) NewSetArray(variable Variable, index Operand, value Operand) *insts.SetArray {
	var inst = insts.NewSetArray(variable, index, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetArray(result Operand, variable Variable, index Operand) *insts.GetArray {
	var inst = insts.NewGetArray(result, variable, index)
	b.AddInstruction(inst)
	return inst
}

func NewBlock(name string) *BasicBlock {
	return &BasicBlock{Name: name, Instructions: []insts.Instruction{}}
}
