package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir"
	"github.com/slince/php-plus/ir/types"
)

func (c *Compiler) compileFuncDecl(node *ast.FuncDecl) error {
	var retType types.Type
	var err error
	if node.Func.Kind != nil {
		retType, err = c.compileType(node.Func.Kind)
		if err != nil {
			return err
		}
	} else {
		retType = types.Void
	}
	var name = c.compileIdentifier(node.Func.Id)
	// function arguments
	var args = make([]*ir.FuncArg, 0)
	for _, arg := range node.Func.Args {
		var compiled, err = c.compileFuncArg(arg)
		if err != nil {
			return err
		}
		args = append(args, compiled)
	}
	c.function = c.module.NewFunction(name, retType, args...)
	//c.enterBlock(c.function.NewBlock("entry"), c.function.NewBlock("leave"))
	_, err = c.compileBlockStmt(node.Func.Body, "entry")
	return err
}

func (c *Compiler) compileFuncArg(node *ast.FuncArg) (*ir.FuncArg, error) {
	var kind, err = c.compileType(node.Kind)
	if err != nil {
		return nil, err
	}
	return ir.NewFuncArg(c.compileIdentifier(node.Id), kind), err
}

func (c *Compiler) compileReturnStmt(node *ast.ReturnStmt) error {
	var ret, err = c.compileExpr(node.Argument)
	if err != nil {
		return err
	}
	c.ctx.NewRet(ret)
	return nil
}
