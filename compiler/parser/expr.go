package parser

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
	"strconv"
)

func (p *Parser) parseExpr(precedence int) ast.Expr {
	var expr = p.parsePrimaryExpr()
	var cur = p.tokens.Current()
	for token.IsOperator(cur.Kind) && isBinaryOperator(cur.Literal) && getBinaryPrecedence(cur.Literal).precedence > precedence {
		var op = getBinaryPrecedence(p.tokens.Current().Literal)
		var expr2 = p.parseExpr(op.precedence)
		expr = ast.NewBinaryExpr(expr, cur.Literal, expr2, expr.Position())
		cur = p.tokens.Current()
	}
	return expr
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	var tok = p.tokens.Current()
	var exp ast.Expr

	switch tok.Kind {
	// constant
	case token.INT:
		num, _ := strconv.ParseInt(tok.Literal, 10, 64)
		exp = ast.NewLiteral(num, tok.Literal, tok.Position)
		p.tokens.Next()
	case token.FLOAT:
		num, _ := strconv.ParseFloat(tok.Literal, 64)
		exp = ast.NewLiteral(num, tok.Literal, tok.Position)
		p.tokens.Next()
	case token.STR:
		exp = ast.NewLiteral(tok.Literal, tok.Literal, tok.Position)
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
	return p.parsePostfixExpr(exp)
}

func (p *Parser) parsePostfixExpr(exp ast.Expr) ast.Expr {
	for !p.tokens.Eof() {
		var tok = p.tokens.Current()
		var end = false
		switch tok.Kind {
		case token.INC, token.DEC: // a ++; b--
			exp = p.parseUpdateExpr(false, exp)
		case token.LPAREN:
			exp = ast.NewCallExpr(exp, p.parseArguments(), exp.Position())
		case token.DOT:
			exp = p.parseObjectAccessExpr(exp)
		case token.LBRACKET: // array[1] map['property']
			exp = p.parseAccessExpr(exp)
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
	var rhs = p.parseExpr(0)
	return ast.NewAssignmentExpr(converted, tok.Literal, rhs, lhs.Position())
}

func (p *Parser) parseIdentifierExpr() ast.Expr {
	var tok = p.tokens.Current()
	var exp ast.Expr
	switch tok.Literal {
	case "true":
	case "TRUE":
		exp = ast.NewLiteral(true, tok.Literal, tok.Position)
	case "false":
	case "FALSE":
		exp = ast.NewLiteral(false, tok.Literal, tok.Position)
	case "null":
	case "NULL":
		exp = ast.NewLiteral(nil, tok.Literal, tok.Position)
	default:
		exp = ast.NewIdentifier(tok.Literal, tok.Position)
	}
	p.tokens.Next()
	return exp
}

func (p *Parser) parseObjectAccessExpr(object ast.Expr) ast.Expr {
	p.tokens.Expect(token.DOT)
	var tok = p.tokens.Expect(token.ID)
	var property = ast.NewIdentifier(tok.Literal, tok.Position)
	var exp ast.Expr = ast.NewMemberExpr(object, property, object.Position())
	if p.tokens.Test(token.LPAREN) {
		exp = ast.NewCallExpr(exp, p.parseArguments(), exp.Position())
	}
	return exp
}

func (p *Parser) parseAccessExpr(object ast.Expr) *ast.MemberExpr {
	p.tokens.Expect(token.LBRACKET)
	var property = p.parseExpr(0)
	p.tokens.Expect(token.RBRACKET)
	return ast.NewMemberExpr(object, property, object.Position())
}

func (p *Parser) parseArrayExpr() *ast.ArrayExpr {
	var tok = p.tokens.Expect(token.LBRACKET)
	var exp = ast.NewArrayExpr(make([]ast.Expr, 0), tok.Position)
	for !p.tokens.Test(token.RBRACKET) {
		if !exp.IsEmpty() {
			p.tokens.Expect(token.COMMA)
		}
		exp.AddElement(p.parseExpr(0))
	}
	p.tokens.Expect(token.RBRACKET)
	return exp
}

func (p *Parser) parseMapExpr() *ast.MapExpr {
	var tok = p.tokens.Expect(token.LBRACE)
	var exp = ast.NewMapExpr(make(map[ast.Expr]ast.Expr), tok.Position)
	for p.tokens.Test(token.RBRACE) {
		if !exp.IsEmpty() {
			p.tokens.Expect(token.COMMA)
		}
		var key = p.parseExpr(0)
		p.tokens.Expect(token.COLON)
		var value = p.parseExpr(0)
		exp.AddElement(key, value)
	}
	return exp
}

func (p *Parser) parseParenExpr() ast.Expr {
	p.tokens.Expect(token.LPAREN)
	var exp = p.parseExpr(0)
	p.tokens.Expect(token.RPAREN)
	return exp
}

func (p *Parser) parseUnaryExpr() *ast.UnaryExpr {
	var tok = p.tokens.Current()
	var operator = tok.Literal
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
	return ast.NewUpdateExpr(tok.Literal, argument, prefix, tok.Position)
}

func (p *Parser) parseBinaryExpr(exp ast.Expr) *ast.BinaryExpr {
	// a + b * c / d
	// a * b + c
	for isBinaryOperator(p.tokens.Current().Literal) {
		exp = p.doParseBinary(exp, defaultOperatorPrecedence)
	}
	return exp.(*ast.BinaryExpr)
}

func (p *Parser) doParseBinary(lhs ast.Expr, prevPrecedence OperatorPrecedence) *ast.BinaryExpr {
	for isBinaryOperator(p.tokens.Current().Literal) {
		var tok = p.tokens.Current()
		var operator = tok.Literal
		var currentPrecedence = getBinaryPrecedence(operator)

		// if the current less than prev, don't consume token.
		if currentPrecedence.precedence < prevPrecedence.precedence {
			break
		}
		// rhs
		p.tokens.Next()
		var rhs = p.parsePrimaryExpr()
		var nextPrecedence = getBinaryPrecedence(p.tokens.Current().Literal)
		if currentPrecedence.precedence < nextPrecedence.precedence {
			rhs = p.doParseBinary(rhs, currentPrecedence)
		}
		prevPrecedence = currentPrecedence
		lhs = ast.NewBinaryExpr(lhs, operator, rhs, lhs.Position())
	}
	return lhs.(*ast.BinaryExpr)
}

func (p *Parser) parseFunctionExpr() *ast.FunctionExpr {
	return ast.NewFunctionExpr(p.parseFunction())
}

func (p *Parser) parseClassExpr() *ast.ClassExpr {
	return ast.NewClassExpr(p.parseClass())
}
