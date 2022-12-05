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

func (b *BasicBlock) NewAdd(result value.Value, lhs value.Value, rhs value.Value) *Add {
	var inst = NewAdd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewSub(result value.Value, lhs value.Value, rhs value.Value) *Sub {
	var inst = NewSub(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}
func (b *BasicBlock) NewMul(result value.Value, lhs value.Value, rhs value.Value) *Mul {
	var inst = NewMul(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewDiv(result value.Value, lhs value.Value, rhs value.Value) *Div {
	var inst = NewDiv(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewMod(result value.Value, lhs value.Value, rhs value.Value) *Mod {
	var inst = NewMod(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGt(result value.Value, lhs value.Value, rhs value.Value) *Gt {
	var inst = NewGt(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGeq(result value.Value, lhs value.Value, rhs value.Value) *Geq {
	var inst = NewGeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLt(result value.Value, lhs value.Value, rhs value.Value) *Lt {
	var inst = NewLt(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLeq(result value.Value, lhs value.Value, rhs value.Value) *Leq {
	var inst = NewLeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewEq(result value.Value, lhs value.Value, rhs value.Value) *Eq {
	var inst = NewEq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeq(result value.Value, lhs value.Value, rhs value.Value) *Neq {
	var inst = NewNeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalAnd(result value.Value, lhs value.Value, rhs value.Value) *LogicalAnd {
	var inst = NewLogicalAnd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalOr(result value.Value, lhs value.Value, rhs value.Value) *LogicalOr {
	var inst = NewLogicalOr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalNot(result value.Value, ope value.Value) *LogicalNot {
	var inst = NewLogicalNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitAnd(result value.Value, lhs value.Value, rhs value.Value) *And {
	var inst = NewBitAnd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitOr(result value.Value, lhs value.Value, rhs value.Value) *Or {
	var inst = NewBitOr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitXor(result value.Value, lhs value.Value, rhs value.Value) *Xor {
	var inst = NewBitXor(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitNot(result value.Value, ope value.Value) *Not {
	var inst = NewBitNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShl(result value.Value, lhs value.Value, rhs value.Value) *Shl {
	var inst = NewBitShl(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShr(result value.Value, lhs value.Value, rhs value.Value) *Shr {
	var inst = NewBitShr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeg(result value.Value, ope value.Value) *Neg {
	var inst = NewNeg(result, ope)
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

func (b *BasicBlock) NewGetArray(result value.Value, variable value.Variable, index value.Value) *GetArray {
	var inst = NewGetArray(result, variable, index)
	b.AddInstruction(inst)
	return inst
}

func NewBlock(name string) *BasicBlock {
	return &BasicBlock{Name: name, Instructions: []Instruction{}}
}
