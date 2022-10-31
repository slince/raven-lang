package main

import (
	"fmt"
	"github.com/slince/php-plus/assembler/lexer"
	"github.com/slince/php-plus/assembler/parser"
)

const ASM = `

section .data
   hello string "helloword"; 字符串常量初始化
   world long 1234

.code
 a:  global main; 程序入口

main:
   load reg0 10 ;这是嵌入行注释
   load reg1 20
   add reg2 reg0  reg1

.loop.x:
   add reg2 reg0  reg1

.loop2x:
   add reg2 reg0  reg1
;这是独立行注释
`

func main() {

	//var str = "re"
	//fmt.Println(strings.Index(str, "reg") == 0)
	//
	//fmt.Println(strconv.ParseUint(str, 10, 8))

	var lex = lexer.NewLexer(ASM)
	var tokens = lex.Lex()

	var p = parser.NewParser(tokens)
	ast := p.Parse()
	fmt.Println(tokens, ast)
}
