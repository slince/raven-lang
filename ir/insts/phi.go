package insts

import "github.com/slince/php-plus/ir"

type Phi struct {
	Variable ir.Operand
	First    *ir.Temporary
	Second   *ir.Temporary
	instruction
}
