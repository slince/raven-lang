package insts

import (
	"github.com/slince/php-plus/ir"
)

type If struct {
	instruction
	Cond        ir.Block
	Body        ir.Block
	Alternative ir.Block
}

type Loop struct {
	instruction
	Cond ir.Block
	Step ir.Block
	Body ir.Block
}

type SwitchCase struct {
	Cond    ir.Block
	Body    ir.Block
	Default bool
}

type Switch struct {
	instruction
	Cond  ir.Block
	Cases []*SwitchCase
}

func NewIf(cond ir.Block, body ir.Block, alternative ir.Block) *If {
	return &If{
		Cond:        cond,
		Body:        body,
		Alternative: alternative,
	}
}

func NewLoop(cond ir.Block, step ir.Block, body ir.Block) *Loop {
	return &Loop{
		Cond: cond, Step: step, Body: body,
	}
}

func NewSwitchCase(cond ir.Block, body ir.Block, defaults bool) *SwitchCase {
	return &SwitchCase{
		Cond:    cond,
		Body:    body,
		Default: defaults,
	}
}

func NewSwitch(cond ir.Block, cases []*SwitchCase) *Switch {
	return &Switch{
		Cond:  cond,
		Cases: cases,
	}
}
