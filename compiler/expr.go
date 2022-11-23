package compiler

import (
	"fmt"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
	"github.com/slince/php-plus/ir"
	"github.com/slince/php-plus/ir/types"
	"math"
)

func (c *Compiler) compileLiteral(node *ast.Literal) (*ir.Const, error) {
	var kind types.Type
	var err error
	switch node.Value.(type) {
	case int64:
		kind = types.I64
	case bool:
		kind = types.Bool
	case float64:
		var num = node.Value.(float64)
		if num > math.MaxFloat32 {
			kind = types.F32
		} else {
			kind = types.F64
		}
	case string:
		kind = types.String
	default:
		err = token.NewSyntaxError(fmt.Sprintf("unknown identifier %s", node.Value), node.Position())
	}
	if err != nil {
		return nil, err
	}
	return ir.NewConst(node.Value, kind), err
}

func (c *Compiler) compileExpr(node ast.Expr) ir.Operand {
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

	}
}

func (c *Compiler) compileUpdateExpr(expr *ast.UpdateExpr) ir.Operand {
	var target = c.compileExpr(expr.Target)
	var result = ir.NewTemporary(nil)
	switch expr.Operator {
	case "++":
		c.ctx.NewAdd(result, target, ir.NewConst(1, target.Type()))
	case "--":
		c.ctx.NewSub(result, target, ir.NewConst(1, target.Type()))
	}
	return result
}

func (c *Compiler) compileBinaryExpr(expr *ast.BinaryExpr) ir.Operand {
	var l, r = c.compileExpr(expr.Left), c.compileExpr(expr.Right)
	var result = ir.NewTemporary(nil)
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
	return result
}

func (c *Compiler) compileUnaryExpr(expr *ast.UnaryExpr) ir.Operand {
	var target = c.compileExpr(expr.Target)
	var result = ir.NewTemporary(nil)
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

func (c *Compiler) compileVariableDecl(expr *ast.VariableDeclarator, declType string) (ir.Operand, error) {
	var name = c.compileIdentifier(expr.Id)
	var kind types.Type
	var err error
	if expr.Kind != nil {
		kind, err = c.compileType(expr.Kind)
		if err != nil {
			return nil, err
		}
	}
	var init ir.Operand
	if expr.Init != nil {
		init = c.compileExpr(expr.Init)
	}
	var variable = ir.NewVariable(name, kind)
	c.ctx.NewLocal(variable, init)
	return variable, err
}

func (c *Compiler) compileAssignmentExpr(expr *ast.AssignmentExpr) ir.Operand {
	var target = c.compileIdentifier(expr.Left)
	var result = ir.NewTemporary(nil)
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
