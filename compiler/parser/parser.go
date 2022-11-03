package parser

import (
	"fmt"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
	"strconv"
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
	for !p.tokens.Test(token.EOF, token.PACKAGE) {
		stmts = append(stmts, p.parseStmt())
	}
	var body = ast.NewBlockStmt(stmts, tok.Position)
	return ast.NewModule(body, tok.Position)
}

func (p *Parser) parseStmt() ast.Stmt {
	var tok = p.tokens.Current()
	var smt ast.Stmt
	switch tok.Type {
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
		smt = ast.NewReturnStmt(p.parseExpr(), tok.Position)
	// exceptions statement
	case token.THROW:
		p.tokens.Next()
		smt = ast.NewThrowStmt(p.parseExpr(), tok.Position)
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
		exp := p.parseExpr()
		smt = ast.NewExpressionStmt(exp, tok.Position)
	}
	//p.tokens.Expect(token.SEMICOLON)
	// ignore semicolon
	for p.tokens.Current().Test(token.SEMICOLON) {
		p.tokens.Next()
	}
	return smt
}

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	var tok = p.tokens.Expect(token.LET, token.CONST)
	var declarators = make([]*ast.VariableDeclarator, 0)

	for {
		declarators = append(declarators, p.parseVariableDeclarator())
		if p.tokens.Current().Test(token.COMMA) { // if comma and go on
			p.tokens.Next()
			continue
		}
		break
	}
	return ast.NewVariableDeclaration(tok.Value, declarators, tok.Position)
}

func (p *Parser) parseVariableDeclarator() *ast.VariableDeclarator {
	var id = p.parseIdentifier()
	var kind *ast.Identifier
	var init ast.Expr

	// parse variable type
	if p.tokens.Current().Test(token.COLON) { // let a: string
		p.tokens.Next()
		kind = p.parseIdentifier()
		// parse variable init
		if p.tokens.Current().Test(token.ASSIGN) { // variable init
			p.tokens.Next()
			init = p.parseExpr()
		}
	} else {
		p.tokens.Expect(token.ASSIGN)
		init = p.parseExpr()
	}
	return ast.NewVariableDeclarator(id, kind, init, id.Position())
}

func (p *Parser) parseFunctionDeclaration() *ast.FunctionDeclaration {
	return ast.NewFunctionDeclaration(p.parseFunction())
}

func (p *Parser) parseFunction() *ast.Function {
	var tok = p.tokens.Expect(token.FUNCTION)
	var id *ast.Identifier // optional
	if p.tokens.Current().Test(token.ID) {
		id = p.parseIdentifier()
	}
	var args = make([]*ast.FunctionArgument, 0)
	p.tokens.Expect(token.LPAREN)
	for !p.tokens.Current().Test(token.RPAREN) {
		if len(args) > 0 {
			p.tokens.Expect(token.COMMA)
		}
		args = append(args, p.parseFunctionArgument())
	}
	p.tokens.Expect(token.RPAREN)
	// function return type is optional.
	var kind *ast.Identifier // optional
	if p.tokens.Current().Test(token.COLON) {
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
	for !p.tokens.Current().Test(token.RBRACE) {
		stmts = append(stmts, p.parseStmt())
	}
	p.tokens.Expect(token.RBRACE)
	return ast.NewBlockStmt(stmts, tok.Position)
}

func (p *Parser) parseIfStmt() *ast.IfStmt {
	var tok = p.tokens.Expect(token.IF, token.ELSEIF)
	p.tokens.Expect(token.LPAREN)
	var test = p.parseExpr()
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
	var discriminant = p.parseExpr()
	p.tokens.Expect(token.RPAREN)
	// parse cases
	var cases = make([]*ast.SwitchCase, 0)
	var hasDefault = false
	p.tokens.Expect(token.LBRACE)
	for p.tokens.Current().Test(token.CASE, token.DEFAULT) {
		var tok = p.tokens.Current()
		p.tokens.Next()
		var isCase = tok.Type == token.CASE
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
			test = p.parseExpr()
		}
		p.tokens.Expect(token.COLON)
		var consequent = make([]ast.Stmt, 0)
		for !p.tokens.Current().Test(token.CASE, token.DEFAULT, token.RBRACE) {
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
		init = p.parseExpr()
	}
	p.tokens.Expect(token.SEMICOLON)
	// parse test
	var test ast.Expr
	if !p.tokens.Current().Test(token.SEMICOLON) {
		test = p.parseExpr()
	}
	p.tokens.Expect(token.SEMICOLON)
	// parse update
	var update ast.Expr
	if !p.tokens.Current().Test(token.RPAREN) {
		update = p.parseExpr()
	}
	p.tokens.Expect(token.RPAREN)
	var body = p.parseBlockStmt()
	return ast.NewForStmt(init, test, update, body, tok.Position)
}

func (p *Parser) parseWhileStmt() *ast.WhileStmt {
	var tok = p.tokens.Expect(token.WHILE)
	p.tokens.Expect(token.LPAREN)
	var test = p.parseExpr()
	p.tokens.Expect(token.RPAREN)
	var body = p.parseBlockStmt()
	return ast.NewWhileStmt(test, body, tok.Position)
}

func (p *Parser) parseDoWhileStmt() *ast.DoWhileStmt {
	var tok = p.tokens.Expect(token.DO)
	var body = p.parseBlockStmt()
	p.tokens.Expect(token.WHILE)
	p.tokens.Expect(token.LPAREN)
	var test = p.parseExpr()
	p.tokens.Expect(token.RPAREN)

	return ast.NewDoWhileStmt(test, body, tok.Position)
}

func (p *Parser) parseForeachStmt() *ast.ForeachStmt {
	var tok = p.tokens.Expect(token.FOREACH)
	p.tokens.Expect(token.LPAREN)
	var source = p.parseExpr()
	p.tokens.Expect(token.AS)
	var key ast.Expr
	var cur = p.tokens.Expect(token.ID)
	var value = ast.NewIdentifier(cur.Value, cur.Position)
	if p.tokens.Current().Test(token.DOUBLE_ARROW) {
		key = value
		p.tokens.Next()
		cur = p.tokens.Expect(token.ID)
		value = ast.NewIdentifier(cur.Value, cur.Position)
	}
	p.tokens.Expect(token.RPAREN)
	return ast.NewForeachStmt(source, key, value, p.parseBlockStmt(), tok.Position)
}

func (p *Parser) parseTryStmt() *ast.TryStmt {
	var tok = p.tokens.Expect(token.TRY)
	var body = p.parseBlockStmt()
	var catches = make([]*ast.CatchClause, 0)
	for p.tokens.Current().Test(token.CATCH) {
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
	if p.tokens.Current().Test(token.FINALLY) {
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
	if p.tokens.Current().Test(token.ID) {
		id = p.parseIdentifier()
	}

	if p.tokens.Current().Test(token.EXTENDS) {
		p.tokens.Next()
		extends = p.parseIdentifier()
	}

	if p.tokens.Current().Test(token.IMPLEMENTS) {
		p.tokens.Next()
		for !p.tokens.Current().Test(token.LBRACE) {
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

	for !p.tokens.Current().Test(token.RBRACE) {
		var tok = p.tokens.Current()
		var context = NewClassPropertyContext()
		var prop *ast.PropertyDefinition
		var method *ast.MethodDefinition
		for {
			var err error
			var end = false
			var cur = p.tokens.Current()
			switch cur.Type {
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
				prop = ast.NewPropertyDefinition(context.visibility, context.static, cur.Value, variable, tok.Position)
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
	return ast.NewModifier(tok.Value, tok.Position)
}
func (p *Parser) parseExpr() ast.Expr {
	var exp = p.parsePrimaryExpr()
	if token.IsOperator(p.tokens.Current().Type) {
		exp = p.parseBinaryExpr(exp)
	}
	return exp
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	var tok = p.tokens.Expect(token.ID)
	return ast.NewIdentifier(tok.Value, tok.Position)
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	var tok = p.tokens.Current()
	var exp ast.Expr

	switch tok.Type {
	// constant
	case token.INT:
		num, _ := strconv.ParseInt(tok.Value, 10, 64)
		exp = ast.NewLiteral(num, tok.Value, tok.Position)
		p.tokens.Next()
	case token.FLOAT:
		num, _ := strconv.ParseFloat(tok.Value, 64)
		exp = ast.NewLiteral(num, tok.Value, tok.Position)
		p.tokens.Next()
	case token.STR:
		exp = ast.NewLiteral(tok.Value, tok.Value, tok.Position)
		p.tokens.Next()
	// identifier
	case token.ID:
		exp = p.parseIdentifierExpr()
	// function
	case token.FUNCTION:
		exp = p.parseFunctionExpr()
	case token.CLASS:
		exp = p.parseClassExpr()
	// punctuation
	case token.LBRACKET:
		exp = p.parseArrayExpr()
	case token.LBRACE:
		exp = p.parseMapExpr()
	case token.LPAREN:
		exp = p.parseParenExpr()
	// unary operator
	case token.INC, token.DEC:
		exp = p.parseUpdateExpr(true, nil)
	case token.LOGIC_NOT, token.NOT, token.ADD, token.SUB:
		exp = p.parseUnaryExpr()
	default:
		p.unexpect(tok)
	}

	return p.parsePosixExpr(exp)
}

func (p *Parser) parsePosixExpr(exp ast.Expr) ast.Expr {
	for !p.tokens.Eof() {
		var tok = p.tokens.Current()
		var end = false
		switch tok.Type {
		case token.LPAREN:
			exp = ast.NewCallExpr(exp, p.parseArguments(), exp.Position())
		case token.DOT:
			exp = p.parseObjectAccessExpr(exp)
		case token.LBRACKET: // array[1] map['property']
			exp = p.parseAccessExpr(exp)
		case token.INC, token.DEC: // a ++; b--
			exp = p.parseUpdateExpr(false, exp)
		case token.ASSIGN,
			token.ADD_ASSIGN, token.SUB_ASSIGN, token.MUL_ASSIGN, token.DIV_ASSIGN, token.MOD_ASSIGN,
			token.AND_ASSIGN, token.OR_ASSIGN, token.XOR_ASSIGN, token.SHL_ASSIGN, token.SHR_ASSIGN, token.AND_NOT_ASSIGN:
			exp = p.parseAssignmentExpr(exp)
		default:
			end = true
		}
		if end {
			break
		}
	}
	return exp
}

func (p *Parser) parseAssignmentExpr(lhs ast.Expr) *ast.AssignmentExpr {
	var converted, isIdent = interface{}(lhs).(*ast.Identifier)
	if !isIdent {
		p.error(token.NewSyntaxError("Assigning to rvalue", lhs.Position()))
	}
	var tok = p.tokens.Expect(token.Assigns...)
	var rhs = p.parseExpr()
	return ast.NewAssignmentExpr(converted, tok.Value, rhs, lhs.Position())
}

func (p *Parser) parseIdentifierExpr() ast.Expr {
	var tok = p.tokens.Current()
	var exp ast.Expr
	switch tok.Value {
	case "true":
	case "TRUE":
		exp = ast.NewLiteral(true, tok.Value, tok.Position)
	case "false":
	case "FALSE":
		exp = ast.NewLiteral(false, tok.Value, tok.Position)
	case "null":
	case "NULL":
		exp = ast.NewLiteral(nil, tok.Value, tok.Position)
	default:
		exp = ast.NewIdentifier(tok.Value, tok.Position)
	}
	p.tokens.Next()
	return exp
}

func (p *Parser) parseObjectAccessExpr(object ast.Expr) ast.Expr {
	p.tokens.Expect(token.DOT)
	var tok = p.tokens.Expect(token.ID)
	var property = ast.NewIdentifier(tok.Value, tok.Position)
	var exp ast.Expr = ast.NewMemberExpr(object, property, object.Position())
	if p.tokens.Current().Test(token.LPAREN) {
		exp = ast.NewCallExpr(exp, p.parseArguments(), exp.Position())
	}
	return exp
}

func (p *Parser) parseAccessExpr(object ast.Expr) *ast.MemberExpr {
	p.tokens.Expect(token.LBRACKET)
	var property = p.parseExpr()
	p.tokens.Expect(token.RBRACKET)
	return ast.NewMemberExpr(object, property, object.Position())
}

func (p *Parser) parseArrayExpr() *ast.ArrayExpr {
	var tok = p.tokens.Expect(token.LBRACKET)
	var exp = ast.NewArrayExpr(make([]ast.Expr, 0), tok.Position)
	for !p.tokens.Current().Test(token.RBRACKET) {
		if !exp.IsEmpty() {
			p.tokens.Expect(token.COMMA)
		}
		exp.AddElement(p.parseExpr())
	}
	p.tokens.Expect(token.RBRACKET)
	return exp
}

func (p *Parser) parseMapExpr() *ast.MapExpr {
	var tok = p.tokens.Expect(token.LBRACE)
	var exp = ast.NewMapExpr(make(map[ast.Expr]ast.Expr), tok.Position)
	for p.tokens.Current().Test(token.RBRACE) {
		if !exp.IsEmpty() {
			p.tokens.Expect(token.COMMA)
		}
		var key = p.parseExpr()
		p.tokens.Expect(token.COLON)
		var value = p.parseExpr()
		exp.AddElement(key, value)
	}
	return exp
}

func (p *Parser) parseParenExpr() ast.Expr {
	p.tokens.Expect(token.LPAREN)
	var exp = p.parseExpr()
	p.tokens.Expect(token.RPAREN)
	return exp
}

func (p *Parser) parseUnaryExpr() *ast.UnaryExpr {
	var tok = p.tokens.Current()
	var operator = tok.Value
	p.tokens.Next()
	var target = p.parsePrimaryExpr()
	return ast.NewUnaryExpr(operator, target, tok.Position)
}

func (p *Parser) parseUpdateExpr(prefix bool, argument ast.Expr) *ast.UpdateExpr {
	var tok = p.tokens.Expect(token.INC, token.DEC)
	if prefix { // ++a --b
		argument = p.parsePrimaryExpr()
	} // a++ b--
	var converted = interface{}(argument)
	_, isVariable := converted.(*ast.Identifier)
	_, isMember := converted.(*ast.MemberExpr)
	if !isVariable && !isMember {
		var msg = "Invalid left-hand side in postfix operation"
		if prefix {
			msg = "Invalid left-hand side in prefix operation"
		}
		p.error(token.NewSyntaxError(msg, argument.Position()))
	}
	return ast.NewUpdateExpr(tok.Value, argument, prefix, tok.Position)
}

func (p *Parser) parseBinaryExpr(exp ast.Expr) *ast.BinaryExpr {
	// a + b * c / d
	// a * b + c
	for IsBinaryOperator(p.tokens.Current().Value) {
		exp = p.doParseBinary(exp, defaultOperatorPrecedence)
	}
	return exp.(*ast.BinaryExpr)
}

func (p *Parser) doParseBinary(lhs ast.Expr, prevPrecedence OperatorPrecedence) *ast.BinaryExpr {
	for IsBinaryOperator(p.tokens.Current().Value) {
		var tok = p.tokens.Current()
		var operator = tok.Value
		var currentPrecedence = BinaryPrecedence(operator)

		// if the current less than prev, don't consume token.
		if currentPrecedence.Precedence < prevPrecedence.Precedence {
			break
		}
		// rhs
		p.tokens.Next()
		var rhs = p.parsePrimaryExpr()
		var nextPrecedence = BinaryPrecedence(p.tokens.Current().Value)
		if currentPrecedence.Precedence < nextPrecedence.Precedence {
			rhs = p.doParseBinary(rhs, currentPrecedence)
		}
		prevPrecedence = currentPrecedence
		lhs = ast.NewBinaryExpr(lhs, operator, rhs, lhs.Position())
	}
	return lhs.(*ast.BinaryExpr)
}

func (p *Parser) parseArguments() []ast.Expr {
	// the_foo_func(1, "foo")
	var args = make([]ast.Expr, 0)
	p.tokens.Expect(token.LPAREN)
	for !p.tokens.Current().Test(token.RPAREN) {
		if len(args) > 0 {
			p.tokens.Expect(token.COMMA)
		}
		args = append(args, p.parseExpr())
	}
	p.tokens.Expect(token.RPAREN)
	return args
}

func (p *Parser) parseFunctionExpr() *ast.FunctionExpr {
	return ast.NewFunctionExpr(p.parseFunction())
}

func (p *Parser) parseClassExpr() *ast.ClassExpr {
	return ast.NewClassExpr(p.parseClass())
}

func (p *Parser) unexpect(tok *token.Token) {
	p.error(token.NewSyntaxError(fmt.Sprintf("Unexpected token \"%d\" of value \"%s\"", tok.Type, tok.Value), tok.Position))
}

func (p *Parser) error(err token.SyntaxError) {
	panic(err)
}

func NewParser(tokens *token.Stream) *Parser {
	return &Parser{
		tokens: tokens,
	}
}
