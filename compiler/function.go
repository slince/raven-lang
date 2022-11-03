package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir"
)

func (c *Compiler) compileFunctionDecl(node *ast.FunctionDeclaration) {
	var name = c.compileIdentifier(node.Function.Id)
	var retType = c.compileType(node.Function.Kind)
	// function arguments
	var args = make([]*ir.FunctionArgument, 0)
	for _, arg := range node.Function.Args {
		args = append(args, c.compileFunctionArgument(arg))
	}
	var fun = c.Module.NewFunction(name, retType, args...)
	c.Function = fun
	c.compileBlockStmt(node.Function.Body, "")
}

func (c *Compiler) compileFunctionArgument(node *ast.FunctionArgument) *ir.FunctionArgument {
	return ir.NewFuncParam(c.compileIdentifier(node.Id), c.compileType(node.Kind))
}

func (c *Compiler) compileReturnStmt(node *ast.ReturnStmt) {
	c.ctx.NewRet(c.compileExpr(node.Argument))
}
