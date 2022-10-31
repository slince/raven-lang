package main

import (
	"fmt"
	"github.com/slince/php-plus/compiler/lexer"
	"github.com/slince/php-plus/compiler/parser"
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
	//var a []int = []int{1, 2, 3, 4, 5}
	//
	//var a: int = 10 + 20

	//fmt.Println(a)
	var code = `
let a : string = "abc", b : int = 10 + 20;

if (c > d) {
   f = 123
   const sub = a
} elseif (c) {
    f = 12
} else if (d) {
    f = 10
} else {
   v = 10
}

switch (a){
	case 1:
		break;
	case 2:
        a = 3;
		break
    default:
		a = 4;
}

for (let a = 10; a <= 100; a++) {
  
}

foreach (1 + 2 as key => value) {

	let a = 1 + 2
}

while (a > b) {
	a = 4;
}

do {
	a = 4;
} while (a > b);

try {
	throw a+1;
} catch (e: Exception) {
} catch (e : Exception1) {
} finally {

}

function hello(a: string, b: int64): string{
	return a + b
}

let func = function hello2(a: string, b: int64): string{
	return a + b
}

class FooClass extends ParentClass implements InterfaceA, InterfaceB{
	public const a: string = "123"
	public static b: bool = false
	
	abstract public function hello(): string{
		return "hello" + "world";
    }

	final public static function world(): void{
        return 1 + 2
    }
}

let cls = class FooClass extends ParentClass implements InterfaceA, InterfaceB{
	public const a: string = "123"
	public static b: bool = false
	
	public function hello(): string{
		return "hello" + "world";
    }

	final public static function world(): void{
        return 1 + 2
    }
}
`
	lex := lexer.NewLexer(code)
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
