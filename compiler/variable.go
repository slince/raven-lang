package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir/types"
	"github.com/slince/php-plus/ir/value"
)

func (c *Compiler) compileVarDecl(node *ast.VariableDeclaration) error {
	var err error
	for _, decl := range node.Declarators {
		err = c.compileVarDeclarator(decl, node.Kind == "const")
	}
	return err
}

func (c *Compiler) compileVarDeclarator(node *ast.VariableDeclarator, immutable bool) error {
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
			c.module.NewConst(name, kind, init.(*value.Const))
		} else {
			c.module.NewGlobal(name, kind, init.(*value.Const))
		}
	} else {
		var variable = value.NewVariable(name, kind)
		variable.Immutable = immutable
		c.ctx.NewLocal(variable, init)
	}
	return nil
}
