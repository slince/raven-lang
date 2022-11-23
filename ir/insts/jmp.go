package insts

import "github.com/slince/php-plus/ir"

type Jmp struct {
	Target ir.Block
	instruction
}

func NewJmp(target ir.Block) *Jmp {
	return &Jmp{Target: target}
}

type CondJmp struct {
	Cond        ir.Operand
	TrueTarget  ir.Block
	FalseTarget ir.Block
	instruction
}

func NewCondJmp(cond ir.Operand, trueTarget ir.Block, falseTarget ir.Block) *CondJmp {
	return &CondJmp{Cond: cond, TrueTarget: trueTarget, FalseTarget: falseTarget}
}
