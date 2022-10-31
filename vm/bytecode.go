package vm

import (
	"github.com/slince/php-plus/vm/object"
)

// Function 用户自定义函数，此函数不是符号表的函数表达
type Function struct {
	Instructions Instructions
}

type Bytecode struct {
	Entry     string
	Constants []object.Object
	Globals   []object.Object
	Functions map[string]*Function
}
