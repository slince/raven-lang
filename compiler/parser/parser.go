package parser

import (
	"fmt"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
)

type Parser struct {
	tokens *token.Stream
}

func (p *Parser) Parse() *ast.Program {
	var tok = p.tokens.Current()
	var modules = make([]*ast.Module, 0)
	for !p.tokens.Eof() {
		modules = append(modules, p.parseModule())
	}
	return ast.NewProgram(modules, tok.Position)
}

func (p *Parser) parseModule() *ast.Module {
	var tok = p.tokens.Current()
	var stmts = make([]ast.Stmt, 0)
	for !p.tokens.Eof() {
		stmts = append(stmts, p.parseStmt())
	}
	var body = ast.NewBlockStmt(stmts, tok.Position)
	return ast.NewModule(body, tok.Position)
}

func (p *Parser) parseStmt() ast.Stmt {
	var tok = p.tokens.Current()
	var smt ast.Stmt
	switch tok.Kind {
	// variable declaration
	case token.LET, token.CONST:
		smt = p.parseVarDecl()
	case token.FUNCTION:
		smt = p.parseFuncDecl()
	// choice statement
	case token.IF:
		smt = p.parseIfStmt()
	case token.SWITCH:
		smt = p.parseSwitchStmt()
	// loop.x statement
	case token.FOR:
		smt = p.parseForStmt()
	case token.WHILE:
		smt = p.parseWhileStmt()
	case token.DO:
		smt = p.parseDoWhileStmt()
	case token.FOREACH:
		smt = p.parseForeachStmt()
	// control statement
	case token.BREAK:
		smt = ast.NewBreakStmt(tok.Position)
		p.tokens.Next()
	case token.CONTINUE:
		smt = ast.NewContinueStmt(tok.Position)
		p.tokens.Next()
	case token.RETURN:
		p.tokens.Next()
		smt = ast.NewReturnStmt(p.parseExpr(0), tok.Position)
	// exceptions statement
	case token.THROW:
		p.tokens.Next()
		smt = ast.NewThrowStmt(p.parseExpr(0), tok.Position)
	case token.TRY:
		smt = p.parseTryStmt()
	// class
	case token.CLASS:
		smt = p.parseClassDecl()
		break
	// block statement
	case token.LBRACE:
		smt = p.parseBlockStmt()
	// expression statement
	default:
		exp := p.parseExpr(0)
		smt = ast.NewExpressionStmt(exp, tok.Position)
	}
	//p.tokens.Expect(token.SEMICOLON)
	// ignore semicolon
	for p.tokens.Test(token.SEMICOLON) {
		p.tokens.Next()
	}
	return smt
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	var tok = p.tokens.Expect(token.ID)
	return ast.NewIdentifier(tok.Literal, tok.Position)
}

func (p *Parser) unexpect(tok *token.Token) {
	p.error(token.NewSyntaxError(fmt.Sprintf("Unexpected token \"%d\" of value \"%s\"", tok.Kind, tok.Literal), tok.Position))
}

func (p *Parser) error(err token.SyntaxError) {
	panic(err)
}

func NewParser(tokens *token.Stream) *Parser {
	return &Parser{
		tokens: tokens,
	}
}
