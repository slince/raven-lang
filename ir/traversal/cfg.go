package traversal

import "github.com/slince/php-plus/ir"

type CfgBuilder struct {
	block *ir.BasicBlock
}

func (c *CfgBuilder) EnterNode(inst ir.Instruction) {
	if inst, ok := inst.(*ir.Label); ok {
		c.block = ir.NewBlock(inst.Name)
	}
}

func (c *CfgBuilder) convertIfInst(inst *ir.If) {
	ir.NewCondJmp(inst.Test)
}
