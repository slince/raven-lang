package ir

import (
	"github.com/slince/php-plus/ir/insts"
	"github.com/slince/php-plus/ir/value"
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

func (b *BasicBlock) NewLea(variable value.Operand, target value.Operand) *insts.Lea {
	var inst = insts.NewLea(variable, target)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewPtr(variable value.Operand, target value.Operand) *insts.Ptr {
	var inst = insts.NewPtr(variable, target)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAdd(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Add {
	var inst = insts.NewAdd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewSub(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Sub {
	var inst = insts.NewSub(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}
func (b *BasicBlock) NewMul(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Mul {
	var inst = insts.NewMul(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewDiv(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Div {
	var inst = insts.NewDiv(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewMod(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Mod {
	var inst = insts.NewMod(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGt(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Gt {
	var inst = insts.NewGt(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGeq(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Geq {
	var inst = insts.NewGeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLt(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Lt {
	var inst = insts.NewLt(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLeq(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Leq {
	var inst = insts.NewLeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewEq(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Eq {
	var inst = insts.NewEq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeq(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Neq {
	var inst = insts.NewNeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalAnd(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.LogicalAnd {
	var inst = insts.NewLogicalAnd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalOr(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.LogicalOr {
	var inst = insts.NewLogicalOr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalNot(result value.Operand, ope value.Operand) *insts.LogicalNot {
	var inst = insts.NewLogicalNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitAnd(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.And {
	var inst = insts.NewBitAnd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitOr(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Or {
	var inst = insts.NewBitOr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitXor(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Xor {
	var inst = insts.NewBitXor(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitNot(result value.Operand, ope value.Operand) *insts.Not {
	var inst = insts.NewBitNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShl(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Shl {
	var inst = insts.NewBitShl(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShr(result value.Operand, lhs value.Operand, rhs value.Operand) *insts.Shr {
	var inst = insts.NewBitShr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeg(result value.Operand, ope value.Operand) *insts.Neg {
	var inst = insts.NewNeg(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGlobal(variable value.Operand, value value.Operand) *insts.Global {
	var inst = insts.NewGlobal(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetGlobal(variable value.Operand) *insts.GetGlobal {
	var inst = insts.NewGetGlobal(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewConst(variable value.Operand, value value.Operand) *insts.Const {
	var inst = insts.NewConst(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetConst(variable value.Operand) *insts.GetConst {
	var inst = insts.NewGetConst(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLocal(variable value.Operand, value value.Operand) *insts.Local {
	var inst = insts.NewLocal(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetLocal(variable value.Operand) *insts.GetLocal {
	var inst = insts.NewGetLocal(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAssign(variable value.Operand, value value.Operand) *insts.Assign {
	var inst = insts.NewAssign(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewArg(value value.Operand) *insts.Arg {
	var inst = insts.NewArg(value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewCall(callee *Function, argNum uint64) *insts.Call {
	var inst = insts.NewCall(callee, argNum)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewRet(ope value.Operand) *insts.Ret {
	var inst = insts.NewRet(ope)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewJmp(target Block) *insts.Jmp {
	var inst = insts.NewJmp(target)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewCondJmp(cond value.Operand, trueTarget Block, falseTarget Block) *insts.CondJmp {
	var inst = insts.NewCondJmp(cond, trueTarget, falseTarget)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewSetArray(variable value.Variable, index value.Operand, value value.Operand) *insts.SetArray {
	var inst = insts.NewSetArray(variable, index, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetArray(result value.Operand, variable value.Variable, index value.Operand) *insts.GetArray {
	var inst = insts.NewGetArray(result, variable, index)
	b.AddInstruction(inst)
	return inst
}

func NewBlock(name string) *BasicBlock {
	return &BasicBlock{Name: name, Instructions: []insts.Instruction{}}
}
