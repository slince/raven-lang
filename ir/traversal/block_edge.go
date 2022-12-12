package traversal

import "github.com/slince/php-plus/ir"

type BlockEdgeBuilder struct {
	blocks map[string]*ir.BasicBlock
}

func (b *BlockEdgeBuilder) EnterBlock(bb *ir.BasicBlock) {
	if bb.Terminator == nil {
		return
	}
	if inst, ok := bb.Terminator.(*ir.Jmp); ok {

	}
}

func (b *BlockEdgeBuilder) findBasicBlock(bb *ir.BasicBlock) {
}
