package parser

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
)

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
