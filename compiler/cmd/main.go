package main

import (
	"fmt"
	"github.com/slince/php-plus/compiler"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/lexer"
	"github.com/slince/php-plus/compiler/parser"
	"github.com/slince/php-plus/ir/graph"
	"io"
	"os"
)

func source(path string) string {
	var file, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	var code, _ = io.ReadAll(file)
	return string(code)
}

func parse(path string) *ast.Program {
	lex := lexer.NewLexer(source(path))
	tokens := lex.Lex()
	p := parser.NewParser(tokens)
	return p.Parse()
}

func main() {
	//
	//var ast = parse("./examples/lang/simple.x")
	//fmt.Println(ast)

	var c = compiler.NewCompiler()
	var program, err = c.Compile(source("./examples/lang/simple.x"))

	var canvas = graph.NewCanvas(program)
	err = canvas.Draw()
	err = canvas.SaveTo("./graph.png")
	fmt.Println(program)
	if err != nil {
		panic(err)
	}
}
