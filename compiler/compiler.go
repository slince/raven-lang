package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/lexer"
	"github.com/slince/php-plus/compiler/parser"
	"github.com/slince/php-plus/ir"
	"github.com/slince/php-plus/ir/value"
)

type Compiler struct {
	ctx         *ir.BlockContext
	symbolTable *ir.SymbolTable
	module      *ir.Module
	function    *ir.Function
	program     *ir.Program
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
	if node == nil {
		return ""
	}
	return node.Value
}

func (c *Compiler) compileVariable(node *ast.Identifier) (*value.Variable, error) {
	var name = c.compileIdentifier(node)
	return c.symbolTable.GetVariable(name)
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
	var name = c.compileIdentifier(node.Name)
	c.module = c.program.NewModule(name)
	var err error
	for _, stmt := range node.Body.Body {
		err = c.compileStmt(stmt)
	}
	return err
}

func (c *Compiler) compileStmt(node ast.Stmt) error {
	var err error
	switch node.(type) {
	case *ast.FunctionDeclaration:
		err = c.compileFunctionDecl(node.(*ast.FunctionDeclaration))
	case *ast.VariableDeclaration:
		err = c.compileVariableDecl(node.(*ast.VariableDeclaration))
	case *ast.BlockStmt:
		_, err = c.compileBlockStmt(node.(*ast.BlockStmt), "")
	case *ast.ExpressionStmt:
		err = c.compileExprStmt(node.(*ast.ExpressionStmt))
	case *ast.WhileStmt:
		c.compileWhileStmt(node.(*ast.WhileStmt))
	case *ast.DoWhileStmt:
		c.compileDoWhileStmt(node.(*ast.DoWhileStmt))
	case *ast.SwitchStmt:
		c.compileSwitchStmt(node.(*ast.SwitchStmt))
	case *ast.ReturnStmt:
		err = c.compileReturnStmt(node.(*ast.ReturnStmt))
	case *ast.BreakStmt:
		c.compileBreakStmt(node.(*ast.BreakStmt))
	}
	c.enterBlock(c.ctx.LeaveBlock, nil)
	return err
}

func (c *Compiler) compileBlockStmt(node *ast.BlockStmt, label string) (*ir.BasicBlock, error) {
	c.enterScope()
	var block = c.function.NewBlock(label)
	var err = c.compileBlock(block, func() error {
		for _, stmt := range node.Body {
			var err = c.compileStmt(stmt)
			if err != nil {
				return err
			}
		}
		return nil
	})
	c.leaveScope()
	return block, err
}

func (c *Compiler) compileExprStmt(node *ast.ExpressionStmt) error {
	_, err := c.compileExpr(node.Expr)
	return err
}

func (c *Compiler) compileBlock(block *ir.BasicBlock, executor func() error) error {
	c.enterBlock(block, c.ctx.LeaveBlock)
	var err = executor()
	c.leaveBlock()
	return err
}

func (c *Compiler) createBlock(label string, executor func() error) (ir.Block, error) {
	var block = c.function.NewBlock(label)
	var err = c.compileBlock(block, executor)
	return block, err
}

func NewCompiler() *Compiler {
	return &Compiler{}
}
