package ir

import (
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
	Instructions []Instruction
	Phis         []Instruction
	Leader       Instruction
	Terminator   Instruction
	Predecessors []*BasicBlock
	Successors   []*BasicBlock
	Dominators   []*BasicBlock
}

func (b *BasicBlock) block() {}

func (b *BasicBlock) HasTerminator() bool {
	return b.Terminator != nil
}

func (b *BasicBlock) AddInstruction(instruction Instruction) {
	b.Instructions = append(b.Instructions, instruction)
}

func (b *BasicBlock) NewLea(variable value.Value, target value.Value) *Lea {
	var inst = NewLea(variable, target)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewPtr(variable value.Value, target value.Value) *Ptr {
	var inst = NewPtr(variable, target)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAdd(lhs value.Value, rhs value.Value) *Add {
	var inst = NewAdd(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewSub(lhs value.Value, rhs value.Value) *Sub {
	var inst = NewSub(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}
func (b *BasicBlock) NewMul(lhs value.Value, rhs value.Value) *Mul {
	var inst = NewMul(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewDiv(lhs value.Value, rhs value.Value) *Div {
	var inst = NewDiv(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewMod(lhs value.Value, rhs value.Value) *Mod {
	var inst = NewMod(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGt(lhs value.Value, rhs value.Value) *Gt {
	var inst = NewGt(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGeq(lhs value.Value, rhs value.Value) *Geq {
	var inst = NewGeq(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLt(lhs value.Value, rhs value.Value) *Lt {
	var inst = NewLt(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLeq(lhs value.Value, rhs value.Value) *Leq {
	var inst = NewLeq(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewEq(lhs value.Value, rhs value.Value) *Eq {
	var inst = NewEq(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeq(lhs value.Value, rhs value.Value) *Neq {
	var inst = NewNeq(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalAnd(lhs value.Value, rhs value.Value) *LogicalAnd {
	var inst = NewLogicalAnd(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalOr(lhs value.Value, rhs value.Value) *LogicalOr {
	var inst = NewLogicalOr(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalNot(ope value.Value) *LogicalNot {
	var inst = NewLogicalNot(ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitAnd(lhs value.Value, rhs value.Value) *And {
	var inst = NewBitAnd(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitOr(lhs value.Value, rhs value.Value) *Or {
	var inst = NewBitOr(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitXor(lhs value.Value, rhs value.Value) *Xor {
	var inst = NewBitXor(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitNot(ope value.Value) *Not {
	var inst = NewBitNot(ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShl(lhs value.Value, rhs value.Value) *Shl {
	var inst = NewBitShl(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShr(lhs value.Value, rhs value.Value) *Shr {
	var inst = NewBitShr(lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeg(ope value.Value) *Neg {
	var inst = NewNeg(ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGlobal(variable value.Value, value value.Value) *Global {
	var inst = NewGlobal(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetGlobal(variable value.Value) *GetGlobal {
	var inst = NewGetGlobal(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewConst(variable value.Value, value value.Value) *Const {
	var inst = NewConst(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetConst(variable value.Value) *GetConst {
	var inst = NewGetConst(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLocal(variable value.Value, value value.Value) *Local {
	var inst = NewLocal(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetLocal(variable value.Value) *GetLocal {
	var inst = NewGetLocal(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAssign(variable value.Value, value value.Value) *Assign {
	var inst = NewAssign(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewArg(value value.Value) *Arg {
	var inst = NewArg(value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewCall(callee *Function, argNum uint64) *Call {
	var inst = NewCall(callee, argNum)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewRet(ope value.Value) *Ret {
	var inst = NewRet(ope)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewJmp(target Block) *Jmp {
	var inst = NewJmp(target)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewCondJmp(cond value.Value, trueTarget Block, falseTarget Block) *CondJmp {
	var inst = NewCondJmp(cond, trueTarget, falseTarget)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewSetArray(variable value.Variable, index value.Value, value value.Value) *SetArray {
	var inst = NewSetArray(variable, index, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetArray(variable value.Variable, index value.Value) *GetArray {
	var inst = NewGetArray(variable, index)
	b.AddInstruction(inst)
	return inst
}

func NewBlock(name string) *BasicBlock {
	return &BasicBlock{Name: name, Instructions: []Instruction{}}
}
