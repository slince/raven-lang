package insts

import "github.com/slince/php-plus/ir/value"

type Phi struct {
	Variable value.Operand
	First    *value.Temporary
	Second   *value.Temporary
	instruction
}
