package ir

import (
	"fmt"
	"github.com/slince/php-plus/ir/value"
)

type And struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *And) String() string {
	return fmt.Sprintf("and %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

type Or struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Or) String() string {
	return fmt.Sprintf("or %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

type Xor struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Xor) String() string {
	return fmt.Sprintf("xor %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

type Not struct {
	value.Variable
	Ope value.Value
	instruction
}

func (inst *Not) String() string {
	return fmt.Sprintf("not %s %s", inst.Variable.String(), inst.Ope.String())
}

type Shl struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Shl) String() string {
	return fmt.Sprintf("shl %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

type Shr struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Shr) String() string {
	return fmt.Sprintf("shr %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

func NewBitAnd(lhs value.Value, rhs value.Value) *And {
	return &And{Lhs: lhs, Rhs: rhs}
}

func NewBitOr(lhs value.Value, rhs value.Value) *Or {
	return &Or{Lhs: lhs, Rhs: rhs}
}

func NewBitXor(lhs value.Value, rhs value.Value) *Xor {
	return &Xor{Lhs: lhs, Rhs: rhs}
}

func NewBitNot(ope value.Value) *Not {
	return &Not{Ope: ope}
}

func NewBitShl(lhs value.Value, rhs value.Value) *Shl {
	return &Shl{Lhs: lhs, Rhs: rhs}
}

func NewBitShr(lhs value.Value, rhs value.Value) *Shr {
	return &Shr{Lhs: lhs, Rhs: rhs}
}
