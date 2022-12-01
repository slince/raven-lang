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

func (b *BasicBlock) NewLea(variable value.Operand, target value.Operand) *Lea {
	var inst = NewLea(variable, target)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewPtr(variable value.Operand, target value.Operand) *Ptr {
	var inst = NewPtr(variable, target)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAdd(result value.Operand, lhs value.Operand, rhs value.Operand) *Add {
	var inst = NewAdd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewSub(result value.Operand, lhs value.Operand, rhs value.Operand) *Sub {
	var inst = NewSub(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}
func (b *BasicBlock) NewMul(result value.Operand, lhs value.Operand, rhs value.Operand) *Mul {
	var inst = NewMul(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewDiv(result value.Operand, lhs value.Operand, rhs value.Operand) *Div {
	var inst = NewDiv(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewMod(result value.Operand, lhs value.Operand, rhs value.Operand) *Mod {
	var inst = NewMod(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGt(result value.Operand, lhs value.Operand, rhs value.Operand) *Gt {
	var inst = NewGt(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGeq(result value.Operand, lhs value.Operand, rhs value.Operand) *Geq {
	var inst = NewGeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLt(result value.Operand, lhs value.Operand, rhs value.Operand) *Lt {
	var inst = NewLt(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLeq(result value.Operand, lhs value.Operand, rhs value.Operand) *Leq {
	var inst = NewLeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewEq(result value.Operand, lhs value.Operand, rhs value.Operand) *Eq {
	var inst = NewEq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeq(result value.Operand, lhs value.Operand, rhs value.Operand) *Neq {
	var inst = NewNeq(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalAnd(result value.Operand, lhs value.Operand, rhs value.Operand) *LogicalAnd {
	var inst = NewLogicalAnd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalOr(result value.Operand, lhs value.Operand, rhs value.Operand) *LogicalOr {
	var inst = NewLogicalOr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLogicalNot(result value.Operand, ope value.Operand) *LogicalNot {
	var inst = NewLogicalNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitAnd(result value.Operand, lhs value.Operand, rhs value.Operand) *And {
	var inst = NewBitAnd(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitOr(result value.Operand, lhs value.Operand, rhs value.Operand) *Or {
	var inst = NewBitOr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitXor(result value.Operand, lhs value.Operand, rhs value.Operand) *Xor {
	var inst = NewBitXor(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitNot(result value.Operand, ope value.Operand) *Not {
	var inst = NewBitNot(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShl(result value.Operand, lhs value.Operand, rhs value.Operand) *Shl {
	var inst = NewBitShl(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewBitShr(result value.Operand, lhs value.Operand, rhs value.Operand) *Shr {
	var inst = NewBitShr(result, lhs, rhs)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewNeg(result value.Operand, ope value.Operand) *Neg {
	var inst = NewNeg(result, ope)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGlobal(variable value.Operand, value value.Operand) *Global {
	var inst = NewGlobal(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetGlobal(variable value.Operand) *GetGlobal {
	var inst = NewGetGlobal(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewConst(variable value.Operand, value value.Operand) *Const {
	var inst = NewConst(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetConst(variable value.Operand) *GetConst {
	var inst = NewGetConst(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewLocal(variable value.Operand, value value.Operand) *Local {
	var inst = NewLocal(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetLocal(variable value.Operand) *GetLocal {
	var inst = NewGetLocal(variable)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewAssign(variable value.Operand, value value.Operand) *Assign {
	var inst = NewAssign(variable, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewArg(value value.Operand) *Arg {
	var inst = NewArg(value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewCall(callee *Function, argNum uint64) *Call {
	var inst = NewCall(callee, argNum)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewRet(ope value.Operand) *Ret {
	var inst = NewRet(ope)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewJmp(target Block) *Jmp {
	var inst = NewJmp(target)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewCondJmp(cond value.Operand, trueTarget Block, falseTarget Block) *CondJmp {
	var inst = NewCondJmp(cond, trueTarget, falseTarget)
	b.Terminator = inst
	return inst
}

func (b *BasicBlock) NewSetArray(variable value.Variable, index value.Operand, value value.Operand) *SetArray {
	var inst = NewSetArray(variable, index, value)
	b.AddInstruction(inst)
	return inst
}

func (b *BasicBlock) NewGetArray(result value.Operand, variable value.Variable, index value.Operand) *GetArray {
	var inst = NewGetArray(result, variable, index)
	b.AddInstruction(inst)
	return inst
}

func NewBlock(name string) *BasicBlock {
	return &BasicBlock{Name: name, Instructions: []Instruction{}}
}
