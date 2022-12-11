package traversal

import "github.com/slince/php-plus/ir"

type MirToLir struct {
}

func (m *MirToLir) EnterNode(inst ir.Instruction) {
	if inst, ok := inst.(*ir.If); ok {

	}
}

func (m *MirToLir) convertIfInst(inst *ir.If) {
	ir.NewCondJmp(inst.Test)
}
