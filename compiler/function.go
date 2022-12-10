package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir"
	"github.com/slince/php-plus/ir/types"
)

func (c *Compiler) compileFunctionDecl(node *ast.FunctionDeclaration) error {
	var retType types.Type
	var err error
	if node.Function.Kind != nil {
		retType, err = c.compileType(node.Function.Kind)
		if err != nil {
			return err
		}
	} else {
		retType = types.Void
	}
	var name = c.compileIdentifier(node.Function.Id)
	// function arguments
	var args = make([]*ir.FunctionArgument, 0)
	for _, arg := range node.Function.Args {
		var compiled, err = c.compileFunctionArgument(arg)
		if err != nil {
			return err
		}
		args = append(args, compiled)
	}
	c.function = c.module.NewFunction(name, retType, args...)
	//c.enterBlock(c.function.NewBlock("entry"), c.function.NewBlock("leave"))
	_, err = c.compileBlockStmt(node.Function.Body, "entry")
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
