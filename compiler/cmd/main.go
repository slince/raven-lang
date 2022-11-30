package main

import (
	"fmt"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/lexer"
	"github.com/slince/php-plus/compiler/parser"
	"io"
	"os"
)

func parse(path string) *ast.Program {
	var file, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	var code, _ = io.ReadAll(file)
	lex := lexer.NewLexer(string(code))
	tokens := lex.Lex()

	p := parser.NewParser(tokens)
	return p.Parse()
}

func main() {

	var ast = parse("./examples/lang/simple.x")
	fmt.Println(ast)

	//var a  int32 = 10
	//var b int64 = 20

	var a = 10

	if true {
		const a = 12
		fmt.Println(a)
	} else {

	}

	fmt.Println(a)
}
