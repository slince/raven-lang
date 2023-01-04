package ir

import (
	"fmt"
	"github.com/slince/php-plus/ir/value"
)

type Add struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Add) String() string {
	return fmt.Sprintf("add %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

type Sub struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Sub) String() string {
	return fmt.Sprintf("sub %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

type Mul struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Mul) String() string {
	return fmt.Sprintf("mul %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

type Div struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Div) String() string {
	return fmt.Sprintf("div %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

type Mod struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

func (inst *Mod) String() string {
	return fmt.Sprintf("mod %s %s %s", inst.Variable.String(), inst.Lhs.String(), inst.Rhs.String())
}

func NewAdd(lhs value.Value, rhs value.Value) *Add {
	return &Add{Lhs: lhs, Rhs: rhs}
}

func NewSub(lhs value.Value, rhs value.Value) *Sub {
	return &Sub{Lhs: lhs, Rhs: rhs}
}

func NewMul(lhs value.Value, rhs value.Value) *Mul {
	return &Mul{Lhs: lhs, Rhs: rhs}
}

func NewDiv(lhs value.Value, rhs value.Value) *Div {
	return &Div{Lhs: lhs, Rhs: rhs}
}

func NewMod(lhs value.Value, rhs value.Value) *Mod {
	return &Mod{Lhs: lhs, Rhs: rhs}
}
