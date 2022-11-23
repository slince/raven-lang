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
	module   *ir.Module
	function *ir.Function
	program  *ir.Program
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

func (c *Compiler) Compile(source string) (*ir.Program, error) {
	var lex = lexer.NewLexer(source)
	var tokens = lex.Lex()
	var p = parser.NewParser(tokens)
	var program = p.Parse()
	var err = c.compileProgram(program)
	if err != nil {
		return nil, err
	}
	return c.program, nil
}

func (c *Compiler) compileProgram(node *ast.Program) error {
	c.program = ir.NewProgram()
	for _, module := range node.Modules {
		var err = c.compileModule(module)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) compileModule(node *ast.Module) error {
	var name, err = c.compileLiteral(node.Name)
	if err != nil {
		return err
	}
	c.module = c.program.NewModule(name.Value.(string))
	c.compileBlockStmt(node.Body, "")
	return nil
}

func (c *Compiler) compileStmt(node ast.Stmt) error {
	var err error
	switch node.(type) {
	case *ast.FunctionDeclaration:
		err = c.compileFunctionDecl(node.(*ast.FunctionDeclaration))
	case *ast.BlockStmt:
		c.compileBlockStmt(node.(*ast.BlockStmt), "")
	case *ast.ExpressionStmt:
		err = c.compileExprStmt(node.(*ast.ExpressionStmt))
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
	return err
}

func (c *Compiler) compileBlockStmt(node *ast.BlockStmt, label string) *ir.BasicBlock {
	c.enterScope()
	var block = c.function.NewBlock(label)
	c.compileBlock(block, func() {
		for _, stmt := range node.Body {
			c.compileStmt(stmt)
		}
	})
	c.leaveScope()
	return block
}

func (c *Compiler) compileExprStmt(node *ast.ExpressionStmt) error {
	_, err := c.compileExpr(node.Expr)
	return err
}

func (c *Compiler) compileBlock(block *ir.BasicBlock, executor func()) {
	c.enterBlock(block, c.ctx.LeaveBlock)
	executor()
	c.leaveBlock()
}
