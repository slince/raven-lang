package compiler

import (
	"fmt"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
	"github.com/slince/php-plus/ir/types"
	"github.com/slince/php-plus/ir/value"
	"math"
)

func (c *Compiler) compileLiteral(node *ast.Literal) (*value.Const, error) {
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
	return value.NewConst(node.Value, kind), err
}

func (c *Compiler) compileExpr(node ast.Expr) (value.Value, error) {
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
	case *ast.AssignmentExpr:
		return c.compileAssignmentExpr(expr)
	}
	return nil, nil
}

func (c *Compiler) compileArrayExpr(expr *ast.ArrayExpr) (value.Value, error) {
	return nil, nil
}

func (c *Compiler) compileUpdateExpr(expr *ast.UpdateExpr) (value.Value, error) {
	var target, err = c.compileExpr(expr.Target)
	var result value.Value
	if err == nil {
		switch expr.Operator {
		case "++":
			result = c.ctx.NewAdd(target, value.NewConst(1, target.Type()))
		case "--":
			result = c.ctx.NewSub(target, value.NewConst(1, target.Type()))
		}
	}
	return result, err
}

func (c *Compiler) compileBinaryExpr(expr *ast.BinaryExpr) (value.Value, error) {
	var l, err1 = c.compileExpr(expr.Left)
	if err1 != nil {
		return nil, err1
	}
	var r, err2 = c.compileExpr(expr.Right)
	if err2 != nil {
		return nil, err2
	}
	var result value.Value
	switch expr.Operator {
	case "+":
		result = c.ctx.NewAdd(l, r)
	case "-":
		result = c.ctx.NewSub(l, r)
	case "*":
		result = c.ctx.NewMul(l, r)
	case "/":
		result = c.ctx.NewDiv(l, r)
	case "%":
		result = c.ctx.NewMod(l, r)

	case "&":
		result = c.ctx.NewBitAnd(l, r)
	case "|":
		result = c.ctx.NewBitOr(l, r)
	case "^":
		result = c.ctx.NewBitXor(l, r)
	case "<<":
		result = c.ctx.NewBitShr(l, r)
	case ">>":
		result = c.ctx.NewBitShr(l, r)

	case "&&":
		result = c.ctx.NewLogicalAnd(l, r)
	case "||":
		result = c.ctx.NewLogicalOr(l, r)
	case "==":
		result = c.ctx.NewEq(l, r)
	case "!=":
		result = c.ctx.NewNeq(l, r)
	case "<":
		result = c.ctx.NewLt(l, r)
	case "<=":
		result = c.ctx.NewLeq(l, r)
	case ">":
		result = c.ctx.NewGt(l, r)
	case ">=":
		result = c.ctx.NewGeq(l, r)
	}
	return result, nil
}

func (c *Compiler) compileUnaryExpr(expr *ast.UnaryExpr) (value.Value, error) {
	var target, err = c.compileExpr(expr.Target)
	var result value.Value
	if err == nil {
		switch expr.Operator {
		case "!":
			result = c.ctx.NewLogicalNot(target)
		case "~":
			result = c.ctx.NewBitNot(target)
		case "+":
		case "-":
			result = c.ctx.NewNeg(target)
		}
	}
	return result, err
}

func (c *Compiler) compileAssignmentExpr(expr *ast.AssignmentExpr) (value.Value, error) {
	var target, err = c.compileVariable(expr.Left)
	var result value.Value
	if err == nil {
		switch expr.Operator {
		case "!":
			result = c.ctx.NewLogicalNot(target)
		case "~":
			result = c.ctx.NewBitNot(target)
		case "+":
		case "-":
			result = c.ctx.NewNeg(target)
		}
	}
	return result, err
}
