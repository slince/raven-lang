package ir

type BlockContext struct {
	*BasicBlock
	LeaveBlock *BasicBlock
	Prev       *BlockContext
}

func NewBlockContext(block *BasicBlock, prev *BlockContext) *BlockContext {
	return &BlockContext{
		BasicBlock: block,
		Prev:       prev,
	}
}
