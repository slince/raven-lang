package main

import (
	"fmt"
	"github.com/slince/php-plus/compiler/lexer"
	"github.com/slince/php-plus/compiler/parser"
	"io"
	"os"
)

type STT struct {
	Name string
}

var a = STT{
	Name: "asa",
}

func read() string {
	return "asas"
}

var b = read()

func main() {
	var file, err = os.Open("./examples/lang/simple.x")
	if err != nil {
		panic(err)
	}
	var code, _ = io.ReadAll(file)
	lex := lexer.NewLexer(string(code))
	tokens := lex.Lex()

	p := parser.NewParser(tokens)
	ast := p.Parse()
	fmt.Println(tokens, ast, b)

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
