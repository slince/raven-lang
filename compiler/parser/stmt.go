package parser

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
)

func (p *Parser) parseBlockStmt() *ast.BlockStmt {
	var tok = p.tokens.Expect(token.LBRACE)
	var stmts = make([]ast.Stmt, 0)
	for !p.tokens.Test(token.RBRACE) {
		stmts = append(stmts, p.parseStmt())
	}
	p.tokens.Expect(token.RBRACE)
	return ast.NewBlockStmt(stmts, tok.Position)
}

func (p *Parser) parseIfStmt() *ast.IfStmt {
	var tok = p.tokens.Expect(token.IF, token.ELSEIF)
	p.tokens.Expect(token.LPAREN)
	var test = p.parseExpr(0)
	p.tokens.Expect(token.RPAREN)
	var consequent = p.parseBlockStmt()
	var alternate ast.Stmt
	if cur := p.tokens.Current(); cur.Test(token.ELSE) {
		p.tokens.Next()
		alternate = p.parseStmt()
	} else if cur.Test(token.ELSEIF) {
		alternate = p.parseIfStmt()
	}
	return ast.NewIfStmt(test, consequent, alternate, tok.Position)
}

func (p *Parser) parseSwitchStmt() *ast.SwitchStmt {
	var tok = p.tokens.Expect(token.SWITCH)
	p.tokens.Expect(token.LPAREN)
	var discriminant = p.parseExpr(0)
	p.tokens.Expect(token.RPAREN)
	// parse cases
	var cases = make([]*ast.SwitchCase, 0)
	var hasDefault = false
	p.tokens.Expect(token.LBRACE)
	for p.tokens.Test(token.CASE, token.DEFAULT) {
		var tok = p.tokens.Current()
		p.tokens.Next()
		var isCase = tok.Kind == token.CASE
		if !isCase {
			if hasDefault {
				p.error(token.NewSyntaxError("Multiple default clauses", tok.Position))
			} else {
				hasDefault = true
			}
		}
		// parse case item
		var test ast.Expr
		if isCase { // "case" requires test expr
			test = p.parseExpr(0)
		}
		p.tokens.Expect(token.COLON)
		var consequent = make([]ast.Stmt, 0)
		for !p.tokens.Test(token.CASE, token.DEFAULT, token.RBRACE) {
			consequent = append(consequent, p.parseStmt())
		}
		var _case = ast.NewSwitchCase(test, consequent, !isCase, tok.Position)
		cases = append(cases, _case)
	}
	p.tokens.Expect(token.RBRACE)
	return ast.NewSwitchStmt(discriminant, cases, tok.Position)
}

func (p *Parser) parseForStmt() *ast.ForStmt {
	var tok = p.tokens.Expect(token.FOR)
	p.tokens.Expect(token.LPAREN)
	// parse init
	var init ast.Node
	if cur := p.tokens.Current(); cur.Test(token.LET, token.CONST) { // variable declaration
		init = p.parseVarDecl()
	} else if !cur.Test(token.SEMICOLON) {
		init = p.parseExpr(0)
	}
	p.tokens.Expect(token.SEMICOLON)
	// parse test
	var test ast.Expr
	if !p.tokens.Test(token.SEMICOLON) {
		test = p.parseExpr(0)
	}
	p.tokens.Expect(token.SEMICOLON)
	// parse update
	var update ast.Expr
	if !p.tokens.Test(token.RPAREN) {
		update = p.parseExpr(0)
	}
	p.tokens.Expect(token.RPAREN)
	var body = p.parseBlockStmt()
	return ast.NewForStmt(init, test, update, body, tok.Position)
}

func (p *Parser) parseWhileStmt() *ast.WhileStmt {
	var tok = p.tokens.Expect(token.WHILE)
	p.tokens.Expect(token.LPAREN)
	var test = p.parseExpr(0)
	p.tokens.Expect(token.RPAREN)
	var body = p.parseBlockStmt()
	return ast.NewWhileStmt(test, body, tok.Position)
}

func (p *Parser) parseDoWhileStmt() *ast.DoWhileStmt {
	var tok = p.tokens.Expect(token.DO)
	var body = p.parseBlockStmt()
	p.tokens.Expect(token.WHILE)
	p.tokens.Expect(token.LPAREN)
	var test = p.parseExpr(0)
	p.tokens.Expect(token.RPAREN)

	return ast.NewDoWhileStmt(test, body, tok.Position)
}

func (p *Parser) parseForeachStmt() *ast.ForeachStmt {
	var tok = p.tokens.Expect(token.FOREACH)
	p.tokens.Expect(token.LPAREN)
	var source = p.parseExpr(0)
	p.tokens.Expect(token.AS)
	var key *ast.Identifier
	var cur = p.tokens.Expect(token.ID)
	var value = ast.NewIdentifier(cur.Literal, cur.Position)
	if p.tokens.Test(token.DOUBLE_ARROW) {
		key = value
		p.tokens.Next()
		cur = p.tokens.Expect(token.ID)
		value = ast.NewIdentifier(cur.Literal, cur.Position)
	}
	p.tokens.Expect(token.RPAREN)
	return ast.NewForeachStmt(source, key, value, p.parseBlockStmt(), tok.Position)
}

func (p *Parser) parseTryStmt() *ast.TryStmt {
	var tok = p.tokens.Expect(token.TRY)
	var body = p.parseBlockStmt()
	var catches = make([]*ast.CatchClause, 0)
	for p.tokens.Test(token.CATCH) {
		var tok = p.tokens.Current()
		p.tokens.Next()
		p.tokens.Expect(token.LPAREN)
		var variable = p.parseIdentifier()
		p.tokens.Expect(token.COLON)
		var kind = p.parseIdentifier()
		p.tokens.Expect(token.RPAREN)
		var body = p.parseBlockStmt()
		catches = append(catches, ast.NewCatchClause(variable, kind, body, tok.Position))
	}
	var finalizer *ast.BlockStmt
	if p.tokens.Test(token.FINALLY) {
		p.tokens.Next()
		finalizer = p.parseBlockStmt()
	}
	return ast.NewTryStmt(body, catches, finalizer, tok.Position)
}
