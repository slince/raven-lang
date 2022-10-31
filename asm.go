package main

import (
	"github.com/slince/php-plus/assembler/lexer"
)

const ASM = `

section .data
   hello .string "helloword"
   world .long 1234

section .code
   global main
main:
   load reg0 10 // 这是嵌入行注释
   load reg1 20
   add reg2 reg0  reg1
   print reg2
.loop.x:
   add reg2 reg0  reg1

// 这是独立行注释
`

func main() {
	var lex = lexer.NewLexer(ASM)
	var tokens = lex.Lex()
	println(tokens)
}
