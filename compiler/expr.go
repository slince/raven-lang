package compiler

import (
	"fmt"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
	"github.com/slince/php-plus/ir/insts"
	"github.com/slince/php-plus/ir/types"
	"github.com/slince/php-plus/ir/value"
	"math"
)

func (c *Compiler) compileLiteral(node *ast.Literal) (*insts.Const, error) {
	var kind types.Type
	var err error
	switch node.Kind {
	case "int":
		kind = types.I64
	case "float":
		if node.Value.(float64) > math.MaxFloat32 {
			kind = types.F64
		} else {
			kind = types.F32
		}
	case "string":
		kind = types.String
	case "bool":
		kind = types.Bool
	case "null":
		kind = types.Nop
	default:
		err = token.NewSyntaxError(fmt.Sprintf("unknown identifier %s", node.Value), node.Position())
	}
	if err != nil {
		return nil, err
	}
	return insts.NewConst(node.Value, kind), err
}

func (c *Compiler) compileExpr(node ast.Expr) (value.Operand, error) {
	switch expr := node.(type) {
	case *ast.Literal:
		return c.compileLiteral(expr)
	case *ast.BinaryExpr:
		return c.compileBinaryExpr(expr)
	case *ast.UnaryExpr:
		return c.compileUnaryExpr(expr)
	case *ast.UpdateExpr:
		return c.compileUpdateExpr(expr)
	case *ast.ArrayExpr:
		return c.compileArrayExpr(expr)
	}
}

func (c *Compiler) compileArrayExpr(expr *ast.ArrayExpr) (value.Operand, error) {

}

func (c *Compiler) compileUpdateExpr(expr *ast.UpdateExpr) (value.Operand, error) {
	var target, err = c.compileExpr(expr.Target)
	if err != nil {
		return nil, err
	}
	var result = value.NewTemporary(nil)
	switch expr.Operator {
	case "++":
		c.ctx.NewAdd(result, target, insts.NewConst(1, target.Type()))
	case "--":
		c.ctx.NewSub(result, target, insts.NewConst(1, target.Type()))
	}
	return result, nil
}

func (c *Compiler) compileBinaryExpr(expr *ast.BinaryExpr) (value.Operand, error) {
	var l, err = c.compileExpr(expr.Left)
	if err != nil {
		return nil, err
	}
	r, err := c.compileExpr(expr.Right)
	if err != nil {
		return nil, err
	}
	var result = value.NewTemporary(nil)
	switch expr.Operator {
	case "+":
		c.ctx.NewAdd(result, l, r)
	case "-":
		c.ctx.NewAdd(result, l, r)
	case "*":
		c.ctx.NewAdd(result, l, r)
	case "/":
		c.ctx.NewAdd(result, l, r)
	case "%":
		c.ctx.NewMod(result, l, r)

	case "&":
		c.ctx.NewBitAnd(result, l, r)
	case "|":
		c.ctx.NewBitOr(result, l, r)
	case "^":
		c.ctx.NewBitXor(result, l, r)
	case "<<":
		c.ctx.NewBitShr(result, l, r)
	case ">>":
		c.ctx.NewBitShr(result, l, r)

	case "&&":
		c.ctx.NewLogicalAnd(result, l, r)
	case "||":
		c.ctx.NewLogicalOr(result, l, r)
	case "==":
		c.ctx.NewEq(result, l, r)
	case "!=":
		c.ctx.NewNeq(result, l, r)
	case "<":
		c.ctx.NewLt(result, l, r)
	case "<=":
		c.ctx.NewLeq(result, l, r)
	case ">":
		c.ctx.NewGt(result, l, r)
	case ">=":
		c.ctx.NewGeq(result, l, r)
	}
	return result, nil
}

func (c *Compiler) compileUnaryExpr(expr *ast.UnaryExpr) (value.Operand, error) {
	var target, err = c.compileExpr(expr.Target)
	if err != nil {
		return nil, err
	}
	var result = value.NewTemporary(nil)
	switch expr.Operator {
	case "!":
		c.ctx.NewLogicalNot(result, target)
	case "~":
		c.ctx.NewBitNot(result, target)
	case "+":
	case "-":
		c.ctx.NewNeg(result, target)
	}
	return result, nil
}

func (c *Compiler) compileVariableDecl(expr *ast.VariableDeclarator, declType string) (value.Operand, error) {
	var name = c.compileIdentifier(expr.Id)
	var kind types.Type
	var err error
	if expr.Kind != nil {
		kind, err = c.compileType(expr.Kind)
		if err != nil {
			return nil, err
		}
	}
	var init value.Operand
	if expr.Init != nil {
		init, err = c.compileExpr(expr.Init)
		if err != nil {
			return nil, err
		}
	}
	var variable = value.NewVariable(name, kind)
	c.ctx.NewLocal(variable, init)
	return variable, err
}

func (c *Compiler) compileAssignmentExpr(expr *ast.AssignmentExpr) value.Operand {
	var target = c.compileIdentifier(expr.Left)
	var result = value.NewTemporary(nil)
	switch expr.Operator {
	case "!":
		c.ctx.NewLogicalNot(result, target)
	case "~":
		c.ctx.NewBitNot(result, target)
	case "+":
	case "-":
		c.ctx.NewNeg(result, target)
	}
	return result
}
