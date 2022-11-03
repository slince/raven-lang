package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir"
)

func (c *Compiler) compileExpr(node ast.Expr) ir.Operand {
	switch expr := node.(type) {
	case *ast.BinaryExpr:
		return c.compileBinaryExpr(expr)
	case *ast.UnaryExpr:
		return c.compileUnaryExpr(expr)
	case *ast.UpdateExpr:
		return c.compileUnaryExpr(expr)
	}
}

func (c *Compiler) compileUpdateExpr(expr *ast.UpdateExpr) ir.Operand {
	var target = c.compileExpr(expr.Target)
	var result = ir.NewTemporary(nil)
	switch expr.Operator {
	case "++":
		c.ctx.NewAdd(result, target, ir.NewLiteral(1, c.compileType(expr.Target)))
	case "--":
		c.ctx.NewBitNot(result, target)
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
