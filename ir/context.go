package ir

type BlockContext struct {
	*BasicBlock
	LeaveBlock Block
	Prev       *BlockContext
}

func NewBlockContext(block *BasicBlock, prev *BlockContext) *BlockContext {
	return &BlockContext{
		BasicBlock: block,
		Prev:       prev,
	}
}
