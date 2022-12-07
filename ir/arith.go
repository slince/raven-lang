package ir

import "github.com/slince/php-plus/ir/value"

type Add struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Sub struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Mul struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Div struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
}

type Mod struct {
	value.Variable
	Lhs value.Value
	Rhs value.Value
	instruction
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
