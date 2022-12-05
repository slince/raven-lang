package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

//int g = 2;
//
//int add(int x, int y) {
//return x + y;
//}
//int main() {
//return add(1, g);
//}

func hello(a int) int {
	return a + 1
}

var b = hello(12)
var a = b + 56

func main() {

	m := ir.NewModule()

	globalG := m.NewGlobalDef("g", constant.NewInt(types.I32, 2))

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
	var ld = mb.NewLoad(types.I32, globalG)
	fmt.Println(ld)
	ld.SetName("xxxx")
	mb.NewStore(constant.NewInt(types.I32, 100), ld)
	mb.NewRet(mb.NewCall(funcAdd, constant.NewInt(types.I32, 1), ld))

	println(m.String())
}
