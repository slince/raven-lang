package insts

import "github.com/slince/php-plus/ir"

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
