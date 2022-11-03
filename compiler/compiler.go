package compiler

import (
	"github.com/samber/lo"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/lexer"
	"github.com/slince/php-plus/compiler/parser"
	"github.com/slince/php-plus/ir"
	"github.com/slince/php-plus/ir/types"
	"math"
	"strconv"
)

type Compiler struct {
	ctx         *ir.BlockContext
	symbolTable *ir.SymbolTable
	Module      *ir.Module
	Function    *ir.Function
	Program     *ir.Program
}

func (c *Compiler) Compile(source string) *ir.Program {
	var lex = lexer.NewLexer(source)
	var tokens = lex.Lex()
	p := parser.NewParser(tokens)
	var program = p.Parse()
	return c.compileProgram(program)
}

func (c *Compiler) compileProgram(node *ast.Program) *ir.Program {
	var program = ir.NewProgram()
	for _, module := range node.Modules {
		c.compileModule(program, module)
	}
	return program
}

func (c *Compiler) compileModule(program *ir.Program, node *ast.Module) *ir.Module {
	var module = program.NewModule(c.compileLiteral(node.Name).Value.(string))
	c.Module = module
	c.compileBlockStmt(node.Body, "")
	c.Module = nil
	return module
}

func (c *Compiler) compileStmt(node ast.Stmt) {
	switch node.(type) {
	case *ast.FunctionDeclaration:
		c.compileFunctionDecl(node.(*ast.FunctionDeclaration))
	case *ast.BlockStmt:
		c.compileBlockStmt(node.(*ast.BlockStmt), "")
	case *ast.ExpressionStmt:
		c.compileExprStmt(node.(*ast.ExpressionStmt))
	case *ast.WhileStmt:
		c.compileWhileStmt(node.(*ast.WhileStmt))
	case *ast.DoWhileStmt:
		c.compileDoWhileStmt(node.(*ast.DoWhileStmt))
	case *ast.SwitchStmt:
		c.compileSwitchStmt(node.(*ast.SwitchStmt))
	case *ast.ReturnStmt:
		c.compileReturnStmt(node.(*ast.ReturnStmt))
	case *ast.BreakStmt:
		c.compileBreakStmt(node.(*ast.BreakStmt))
	}
	c.enterBlock(c.ctx.LeaveBlock, nil)
}

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

func (c *Compiler) compileBlockStmt(node *ast.BlockStmt, label string) *ir.BasicBlock {
	c.enterScope()
	var block = c.Function.NewBlock(label)
	c.compileBlock(block, func() {
		for _, stmt := range node.Body {
			c.compileStmt(stmt)
		}
	})
	c.leaveScope()
	return block
}

func (c *Compiler) compileExpr(node ast.Expr) ir.Operand {
	switch expr := node.(type) {
	case *ast.BinaryExpr:
		return c.ctx.NewAdd(l, r)
	case *ast.Identifier:
		return c.compileIdentifier(node.(*ast.Identifier))
	}
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
		c.ctx.NewAdd(result, l, r)
	case "|":
		c.ctx.NewAdd(result, l, r)
	case "^":
		c.ctx.NewAdd(result, l, r)
	case "~":
		c.ctx.NewAdd(result, l, r)
	case "<<":
		c.ctx.NewAdd(result, l, r)
	case ">>":
		c.ctx.NewAdd(result, l, r)
	}
	return result
}

func (c *Compiler) compileIdentifier(node *ast.Identifier) string {
	return node.Value
}

func (c *Compiler) compileLiteral(node *ast.Literal) *ir.Const {
	var kind types.Type
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
	}
	return ir.NewLiteral(node.Value, kind)
}

func (c *Compiler) compileExprStmt(node *ast.ExpressionStmt) {
	c.compileExpr(node.Expr)
}

func (c *Compiler) compileReturnStmt(node *ast.ReturnStmt) {
	c.ctx.NewRet(c.compileExpr(node.Argument))
}

func (c *Compiler) compileBreakStmt(node *ast.BreakStmt) {
	c.ctx.NewJmp(c.ctx.LeaveBlock)
}

func (c *Compiler) compileIfStmt(node *ast.IfStmt) *ir.BasicBlock {
	c.ctx.LeaveBlock = c.Function.NewBlock("if.done")
	var consequent = c.compileBlockStmt(node.Consequent, "if.then")
	if consequent.Terminator == nil {
		consequent.NewJmp(c.ctx.LeaveBlock)
	}
	var ifElse ir.Block = c.ctx.LeaveBlock
	if node.Alternate != nil {
		if alternate, ok := node.Alternate.(*ast.BlockStmt); ok {
			ifElse = c.compileBlockStmt(alternate, "if.else")
		} else if alternate, ok := node.Alternate.(*ast.IfStmt); ok {
			ifElse = c.compileIfStmt(alternate)
		}
	}
	var test = c.Function.NewBlock("if.test")
	c.compileBlock(test, func() {
		c.ctx.NewCondJmp(c.compileExpr(node.Test), consequent, ifElse)
	})
	return test
}

func (c *Compiler) compileDoWhileStmt(node *ast.DoWhileStmt) {
	c.ctx.LeaveBlock = c.Function.NewBlock("do.while.done")
	var body = c.compileBlockStmt(node.Body, "do.while.body")
	var test = c.Function.NewBlock("do.while.test")
	c.compileBlock(test, func() {
		c.ctx.NewCondJmp(c.compileExpr(node.Test), body, c.ctx.LeaveBlock)
	})
	if body.Terminator == nil {
		body.NewJmp(test)
	}
	c.ctx.NewJmp(body)
}

func (c *Compiler) compileWhileStmt(node *ast.WhileStmt) {
	c.ctx.LeaveBlock = c.Function.NewBlock("while.done")
	var body = c.compileBlockStmt(node.Body, "while.body")
	var test = c.Function.NewBlock("while.test")
	c.compileBlock(test, func() {
		c.ctx.NewCondJmp(c.compileExpr(node.Test), body, c.ctx.LeaveBlock)
	})
	if body.Terminator == nil {
		body.NewJmp(test)
	}
	c.ctx.NewJmp(test)
}

func (c *Compiler) compileSwitchStmt(node *ast.SwitchStmt) {
	// compile switch cases
	var disc = c.compileExpr(node.Discriminant)
	c.ctx.LeaveBlock = c.Function.NewBlock("switch.done")
	c.enterScope()
	var caseNum = len(node.Cases)
	var _, defaultIdx, _ = lo.FindIndexOf(node.Cases, func(clause *ast.SwitchCase) bool {
		return clause.Default
	})
	for idx, clause := range node.Cases {
		var caseBody = c.compileSwitchCaseBody(clause, idx, idx == caseNum-1)
		c.compileSwitchCaseDisc(disc, caseBody, clause, idx, idx == caseNum-1, defaultIdx)
	}
	c.leaveScope()
	// jmp first case discriminant
	c.ctx.NewJmp(ir.NewReference("switch.case.disc.0"))
}

func (c *Compiler) compileSwitchCaseDisc(disc ir.Operand, caseBody *ir.BasicBlock, node *ast.SwitchCase, idx int, last bool, defaultIdx int) *ir.BasicBlock {
	var discBlock = c.Function.NewBlock("switch.case.disc." + strconv.Itoa(idx))
	c.compileBlock(discBlock, func() {
		if node.Default {
			c.ctx.NewJmp(caseBody)
			return
		}
		var cond = ir.NewTemporary(nil)
		c.ctx.NewLogicalEq(cond, disc, c.compileExpr(node.Test))
		var leaveTarget ir.Block
		if last {
			// jump to default case when not match the case, if the default case is present.
			if defaultIdx > -1 {
				leaveTarget = ir.NewReference("switch.case.disc." + strconv.Itoa(defaultIdx))
			} else {
				leaveTarget = c.ctx.LeaveBlock
			}
		} else {
			// Skip default branch
			if defaultIdx == idx+1 {
				leaveTarget = ir.NewReference("switch.case.disc." + strconv.Itoa(idx+1))
			} else {
				leaveTarget = ir.NewReference("switch.case.disc." + strconv.Itoa(idx+2))
			}
		}
		c.ctx.NewCondJmp(cond, caseBody, leaveTarget)
	})
	return discBlock
}

func (c *Compiler) compileSwitchCaseBody(node *ast.SwitchCase, idx int, last bool) *ir.BasicBlock {
	var caseBlock = c.Function.NewBlock("switch.case." + strconv.Itoa(idx))
	c.compileBlock(caseBlock, func() {
		c.compileSwitchCaseConsequent(node)
		if c.ctx.Terminator == nil {
			var leaveTarget ir.Block
			if last {
				leaveTarget = c.ctx.LeaveBlock
			} else {
				leaveTarget = ir.NewReference("switch.case." + strconv.Itoa(idx+1))
			}
			c.ctx.NewJmp(leaveTarget)
		}
	})
	return caseBlock
}

func (c *Compiler) compileSwitchCaseConsequent(node *ast.SwitchCase) {
	for _, consequent := range node.Consequent {
		c.compileStmt(consequent)
	}
}

func (c *Compiler) compileVarDecl(node *ast.VariableDeclarator) {

}
func (c *Compiler) compileType(node *ast.Identifier) types.Type {
	var _type types.Type
	switch node.Value {
	case "int64":
		_type = types.I64
	case "int32":
		_type = types.I32
	case "float32":
		_type = types.F32
	case "float64":
		_type = types.F64
	case "string":
		_type = types.String
	case "bool":
		_type = types.Bool
	case "void":
		_type = types.Void
	}
	return _type
}

func (c *Compiler) compileBlock(b *ir.BasicBlock, executor func()) {
	c.compileBlockWith(b, c.ctx.LeaveBlock, executor)
}

func (c *Compiler) compileBlockWith(block *ir.BasicBlock, leave *ir.BasicBlock, executor func()) {
	c.enterBlock(block, leave)
	executor()
	c.leaveBlock()
}

func (c *Compiler) enterScope() {
	c.symbolTable = ir.NewSymbolTable(c.symbolTable)
}

func (c *Compiler) leaveScope() {
	c.symbolTable = c.symbolTable.Outer
}

func (c *Compiler) enterBlock(block *ir.BasicBlock, leave *ir.BasicBlock) {
	c.ctx = ir.NewBlockContext(block, c.ctx)
	c.ctx.LeaveBlock = leave
}

func (c *Compiler) leaveBlock() {
	c.ctx = c.ctx.Prev
}
