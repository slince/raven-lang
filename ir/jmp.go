package ir

import (
	"fmt"
	"github.com/slince/php-plus/ir/value"
)

type Jmp struct {
	Target Block
	instruction
}

func (inst *Jmp) String() string {
	return fmt.Sprintf("jmp %s", inst.Target.String())
}

func NewJmp(target Block) *Jmp {
	return &Jmp{Target: target}
}

type CondJmp struct {
	Cond        value.Value
	TrueTarget  Block
	FalseTarget Block
	instruction
}

func NewCondJmp(cond value.Value, trueTarget Block, falseTarget Block) *CondJmp {
	return &CondJmp{Cond: cond, TrueTarget: trueTarget, FalseTarget: falseTarget}
}
