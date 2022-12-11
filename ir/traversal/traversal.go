package traversal

import "github.com/slince/php-plus/ir"

type Traversal interface {
	EnterNode(inst ir.Instruction)
	LeaveNode(inst ir.Instruction)
	GetPriority() int
}
