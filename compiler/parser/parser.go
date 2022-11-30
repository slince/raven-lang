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
		smt = p.parseVariableDeclaration()
	case token.FUNCTION:
		smt = p.parseFunctionDeclaration()
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
		smt = p.parseClassDeclaration()
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

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	var tok = p.tokens.Expect(token.LET, token.CONST)
	var declarators = make([]*ast.VariableDeclarator, 0)

	for {
		declarators = append(declarators, p.parseVariableDeclarator())
		if p.tokens.Test(token.COMMA) { // if comma and go on
			p.tokens.Next()
			continue
		}
		break
	}
	return ast.NewVariableDeclaration(tok.Literal, declarators, tok.Position)
}

func (p *Parser) parseVariableDeclarator() *ast.VariableDeclarator {
	var id = p.parseIdentifier()
	var kind *ast.Identifier
	var init ast.Expr

	// parse variable type
	if p.tokens.Test(token.COLON) { // let a: string
		p.tokens.Next()
		kind = p.parseIdentifier()
		// parse variable init
		if p.tokens.Test(token.ASSIGN) { // variable init
			p.tokens.Next()
			init = p.parseExpr(0)
		}
	} else {
		p.tokens.Expect(token.ASSIGN)
		init = p.parseExpr(0)
	}
	return ast.NewVariableDeclarator(id, kind, init, id.Position())
}

func (p *Parser) parseFunctionDeclaration() *ast.FunctionDeclaration {
	return ast.NewFunctionDeclaration(p.parseFunction())
}

func (p *Parser) parseFunction() *ast.Function {
	var tok = p.tokens.Expect(token.FUNCTION)
	var id *ast.Identifier // optional
	if p.tokens.Test(token.ID) {
		id = p.parseIdentifier()
	}
	var args = make([]*ast.FunctionArgument, 0)
	p.tokens.Expect(token.LPAREN)
	for !p.tokens.Test(token.RPAREN) {
		if len(args) > 0 {
			p.tokens.Expect(token.COMMA)
		}
		args = append(args, p.parseFunctionArgument())
	}
	p.tokens.Expect(token.RPAREN)
	// function return type is optional.
	var kind *ast.Identifier // optional
	if p.tokens.Test(token.COLON) {
		p.tokens.Next()
		kind = p.parseIdentifier()
	}
	var body = p.parseBlockStmt()
	return ast.NewFunction(id, args, kind, body, tok.Position)
}

func (p *Parser) parseFunctionArgument() *ast.FunctionArgument {
	var id = p.parseIdentifier()
	p.tokens.Expect(token.COLON)
	var kind = p.parseIdentifier()
	return ast.NewFunctionArgument(id, kind, id.Position())
}

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
		init = p.parseVariableDeclaration()
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

func (p *Parser) parseClassDeclaration() *ast.ClassDeclaration {
	return ast.NewClassDeclaration(p.parseClass())
}

func (p *Parser) parseClass() *ast.Class {
	var tok = p.tokens.Expect(token.CLASS)
	var id *ast.Identifier
	var extends *ast.Identifier
	var impls = make([]*ast.Identifier, 0)
	if p.tokens.Test(token.ID) {
		id = p.parseIdentifier()
	}

	if p.tokens.Test(token.EXTENDS) {
		p.tokens.Next()
		extends = p.parseIdentifier()
	}

	if p.tokens.Test(token.IMPLEMENTS) {
		p.tokens.Next()
		for !p.tokens.Test(token.LBRACE) {
			if len(impls) > 0 {
				p.tokens.Expect(token.COMMA)
			}
			var impl = p.parseIdentifier()
			impls = append(impls, impl)
		}
	}
	return ast.NewClass(id, extends, impls, p.parseClassBody(), tok.Position)
}

func (p *Parser) parseClassBody() *ast.ClassBody {
	var tok = p.tokens.Expect(token.LBRACE)
	var props = make([]*ast.PropertyDefinition, 0)
	var methods = make([]*ast.MethodDefinition, 0)

	for !p.tokens.Test(token.RBRACE) {
		var tok = p.tokens.Current()
		var context = NewClassPropertyContext()
		var prop *ast.PropertyDefinition
		var method *ast.MethodDefinition
		for {
			var err error
			var end = false
			var cur = p.tokens.Current()
			switch cur.Kind {
			case token.PUBLIC, token.PROTECTED, token.PRIVATE:
				err = context.setVisibility(p.parseClassMemberModifier())
			case token.STATIC:
				err = context.setStatic(p.parseClassMemberModifier())
			case token.FINAL:
				err = context.setFinal(p.parseClassMemberModifier())
			case token.ABSTRACT:
				err = context.setAbstract(p.parseClassMemberModifier())
			case token.CONST:
				p.tokens.Next()
				var variable = p.parseVariableDeclarator()
				prop = ast.NewPropertyDefinition(context.visibility, context.static, cur.Literal, variable, tok.Position)
				end = true
			case token.ID:
				var variable = p.parseVariableDeclarator()
				prop = ast.NewPropertyDefinition(context.visibility, context.static, "", variable, tok.Position)
				end = true
			case token.FUNCTION:
				var function = p.parseFunction()
				method = ast.NewMethodDefinition(context.final, context.abstract, context.visibility, context.static, function, tok.Position)
				end = true
			default:
				p.unexpect(tok)
			}
			if err != nil {
				p.error(err.(token.SyntaxError))
			}
			if end {
				break
			}
		}
		if prop != nil {
			props = append(props, prop)
		} else {
			methods = append(methods, method)
		}
	}
	p.tokens.Expect(token.RBRACE)
	return ast.NewClassBody(props, methods, tok.Position)
}

func (p *Parser) parseClassMemberModifier() *ast.MemberModifier {
	var tok = p.tokens.Current()
	p.tokens.Next()
	return ast.NewModifier(tok.Literal, tok.Position)
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	var tok = p.tokens.Expect(token.ID)
	return ast.NewIdentifier(tok.Literal, tok.Position)
}

//	func (p *Parser) parseLiteral() *ast.Literal {
//		var tok = p.tokens.Expect(token.ID)
//		return ast.NewIdentifier(tok.Literal, tok.Position)
//	}
func (p *Parser) parseArguments() []ast.Expr {
	// the_foo_func(1, "foo")
	var args = make([]ast.Expr, 0)
	p.tokens.Expect(token.LPAREN)
	for !p.tokens.Test(token.RPAREN) {
		if len(args) > 0 {
			p.tokens.Expect(token.COMMA)
		}
		args = append(args, p.parseExpr(0))
	}
	p.tokens.Expect(token.RPAREN)
	return args
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
