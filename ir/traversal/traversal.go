package traversal

import "github.com/slince/php-plus/ir"

type Traversal interface {
	EnterBlock(bb *ir.BasicBlock)
	LeaveBlock(bb *ir.BasicBlock)

	EnterNode(inst ir.Instruction)
	LeaveNode(inst ir.Instruction)
	GetPriority() int
}
