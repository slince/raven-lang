package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir"
)

func (c *Compiler) compileFunctionDecl(node *ast.FunctionDeclaration) error {
	var name = c.compileIdentifier(node.Function.Id)
	var retType, err = c.compileType(node.Function.Kind)
	if err != nil {
		return err
	}
	// function arguments
	var args = make([]*ir.FunctionArgument, 0)
	for _, arg := range node.Function.Args {
		var compiled, err = c.compileFunctionArgument(arg)
		if err != nil {
			return err
		}
		args = append(args, compiled)
	}
	var fun = c.module.NewFunction(name, retType, args...)
	c.function = fun
	_, err = c.compileBlockStmt(node.Function.Body, "")
	return err
}

func (c *Compiler) compileFunctionArgument(node *ast.FunctionArgument) (*ir.FunctionArgument, error) {
	var kind, err = c.compileType(node.Kind)
	if err != nil {
		return nil, err
	}
	return ir.NewFuncParam(c.compileIdentifier(node.Id), kind), err
}

func (c *Compiler) compileReturnStmt(node *ast.ReturnStmt) error {
	var ret, err = c.compileExpr(node.Argument)
	if err != nil {
		return err
	}
	c.ctx.NewRet(ret)
	return nil
}
