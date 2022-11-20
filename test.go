package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/samber/lo"
)

//int g = 2;
//
//int add(int x, int y) {
//return x + y;
//}
//int main() {
//return add(1, g);
//}

var a = 10 + 56

func main() {

	fmt.Println(5 >> 2)
	var a = 10

	if a > 5 {
		var b = 10
	} else {
		var b = 5
	}

	fmt.Println(b)

	var d uint8 = 123
	var e float32 = 123

	fmt.Println(d / e)

	m := ir.NewModule()

	globalG := m.NewGlobalDef("g", constant.NewInt(types.I32, 2))

	lo.Find()
	funcAdd := m.NewFunc("add", types.I32,
		ir.NewParam("x", types.I32),
		ir.NewParam("y", types.I32),
	)
	ab := funcAdd.NewBlock("")
	ab.NewRet(ab.NewAdd(funcAdd.Params[0], funcAdd.Params[1]))
	funcMain := m.NewFunc(
		"main",
		types.I32,
	) // omit parameters
	mb := funcMain.NewBlock("") // llir/llvm would give correct default name for block without name
	mb.NewRet(mb.NewCall(funcAdd, constant.NewInt(types.I32, 1), mb.NewLoad(types.I32, globalG)))
	println(m.String())
}
