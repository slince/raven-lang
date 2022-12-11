package traversal

import "github.com/slince/php-plus/ir"

type CfgBuilder struct {
}

func (c *CfgBuilder) EnterNode(inst ir.Instruction) {
	if inst, ok := inst.(*ir.Label); ok {

	}
}

func (c *CfgBuilder) convertIfInst(inst *ir.If) {
	ir.NewCondJmp(inst.Test)
}
