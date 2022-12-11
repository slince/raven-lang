package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir/types"
	"github.com/slince/php-plus/ir/value"
)

func (c *Compiler) compileVarDecl(node *ast.VarDecl) error {
	var err error
	for _, decl := range node.Specs {
		err = c.compileVarSpec(decl, node.Kind == "const")
	}
	return err
}

func (c *Compiler) compileVarSpec(node *ast.VarSpec, immutable bool) error {
	var name string
	var init value.Value
	var kind types.Type
	var err error
	if node.Init != nil {
		init, err = c.compileExpr(node.Init)
		if err != nil {
			return err
		}
	}
	if node.Kind != nil {
		kind, err = c.compileType(node.Kind)
		if err != nil {
			return err
		}
	}
	name = c.compileIdentifier(node.Id)
	if c.function == nil {
		if immutable {
			var constant = c.module.NewConst(name, kind, init.(*value.Const))
			c.symbolTable.AddVariable(&constant.Variable)
		} else {
			var global = c.module.NewGlobal(name, kind, init.(*value.Const))
			c.symbolTable.AddVariable(&global.Variable)
		}
	} else {
		var local = c.ctx.NewLocal(name, kind, init)
		local.Immutable = immutable
		c.symbolTable.AddVariable(&local.Variable)
	}
	return nil
}
