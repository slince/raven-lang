package assembler

import (
	"github.com/slince/php-plus/ir"
)

// 指令（instruction）、伪指令（directive）、标号（label）及注释（comment）。
// https://zhuanlan.zhihu.com/p/443522525
// .byte（1字节）、.word（2字节）、.long（4字节）及.quad（8字节）。

type Assembler struct {
}

func (asm *Assembler) Assemble(source []byte) *ir.Ir {
	return nil
}
