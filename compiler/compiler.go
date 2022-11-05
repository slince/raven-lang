package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/lexer"
	"github.com/slince/php-plus/compiler/parser"
	"github.com/slince/php-plus/ir"
)

type Compiler struct {
	ctx         *ir.BlockContext
	symbolTable *ir.SymbolTable
	Module      *ir.Module
	Function    *ir.Function
	Program     *ir.Program
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

func (c *Compiler) compileIdentifier(node *ast.Identifier) string {
	return node.Value
}

func (c *Compiler) Compile(source string) *ir.Program {
	var lex = lexer.NewLexer(source)
	var tokens = lex.Lex()
	var p = parser.NewParser(tokens)
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

func (c *Compiler) compileExprStmt(node *ast.ExpressionStmt) {
	c.compileExpr(node.Expr)
}

func (c *Compiler) compileBlock(block *ir.BasicBlock, executor func()) {
	c.enterBlock(block, c.ctx.LeaveBlock)
	executor()
	c.leaveBlock()
}
