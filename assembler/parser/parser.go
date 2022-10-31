package parser

import (
	"fmt"
	"github.com/slince/php-plus/assembler/ast"
	"github.com/slince/php-plus/assembler/token"
	"strconv"
)

type Parser struct {
	tokens *token.Stream
}

func (p *Parser) Parse() *ast.Program {
	var tok = p.tokens.Current()
	var body = make([]ast.Node, 0)
	for !p.tokens.Eof() {
		body = append(body, p.parsePrimary())
	}
	return ast.NewProgram(body, tok.Position)
}

func (p *Parser) parsePrimary() ast.Node {
	var tok = p.tokens.Current()
	var node ast.Node
	switch tok.Type {
	case token.COMMENT:
		node = p.parseComment()
	case token.DIR_SECTION, token.DIR_DATA, token.DIR_CODE:
		node = p.parseSection()
	default:
		p.unexpect(tok)
	}
	return node
}

func (p *Parser) parseSection() *ast.Section {
	var section = p.tokens.SkipIfTest(token.DIR_SECTION)
	var name = p.tokens.Expect(token.DIR_DATA, token.DIR_CODE)
	var body = make([]ast.Node, 0)
	for !p.tokens.Current().Test(token.DIR_SECTION, token.DIR_DATA, token.DIR_CODE, token.COMMENT, token.EOF) {
		body = append(body, p.parseNode())
	}
	var pos *token.Position
	if section != nil {
		pos = section.Position
	} else {
		pos = name.Position
	}
	return ast.NewSection(createLabel(name), body, pos)
}

func (p *Parser) parseNode() ast.Node {
	var tok = p.tokens.Current()
	var node ast.Node
	switch {
	case tok.Test(token.COMMENT):
		node = p.parseComment()
	case tok.Test(token.IDENT) && p.nextIsInstruction():
		node = p.parseBlock()
	case tok.IsInstruction() || tok.Test(token.LABEL) && p.nextIsInstruction():
		node = p.parseInstruction()
	case tok.IsDirective() || tok.Test(token.LABEL, token.IDENT) && p.nextIsDirective():
		node = p.parseDirective()
	default:
		p.unexpect(tok)
	}
	return node
}

func (p *Parser) nextIsInstruction() bool {
	return p.tokens.Peek().IsInstruction() || p.tokens.Peek().Test(token.COLON) && p.tokens.Look(2).IsInstruction()
}

func (p *Parser) nextIsDirective() bool {
	return p.tokens.Peek().IsDirective() || p.tokens.Peek().Test(token.COLON) && p.tokens.Look(2).IsDirective()
}

func (p *Parser) parseBlock() *ast.Block {
	// block cannot start with local label
	var label = p.tokens.Expect(token.IDENT)
	p.tokens.SkipIfTest(token.COLON)
	var body = make([]ast.Node, 0)
	for !p.tokens.Current().Test(token.IDENT, token.EOF) {
		body = append(body, p.parseInstruction())
	}
	return ast.NewBlock(createIdentifier(label), body)
}

func (p *Parser) parseCommentOptional() *ast.Comment {
	if p.tokens.Current().Test(token.COMMENT) {
		return p.parseComment()
	}
	return nil
}

func (p *Parser) parseComment() *ast.Comment {
	var tok = p.tokens.Expect(token.COMMENT)
	return ast.NewComment(createLiteral(tok), tok.Position)
}

func (p *Parser) parseInstruction() *ast.Instruction {
	var label = p.parseLabelOptional()
	var kind = p.tokens.ExpectInstruction()
	var def = InstDefinitionOf(kind.Value)
	var ope1 = p.parseOperand(def.First)
	var ope2 = p.parseOperand(def.Second)
	var ope3 = p.parseOperand(def.Third)
	return ast.NewInstruction(label, createIdentifier(kind), ope1, ope2, ope3, p.parseCommentOptional())
}

func (p *Parser) parseOperand(kind OperandType) *ast.Operand {
	var ope *ast.Operand
	switch kind {
	case OPE_REG:
		ope = p.parseRegOperand()
	case OPE_IMM:
		ope = p.parseImmOperand()
	case OPE_CONST:
		ope = p.parseConstOperand()
	case OPE_NONE:
		ope = nil
	default:
		ope = p.parseDynOperand()
	}
	return ope
}

func (p *Parser) parseDynOperand() *ast.Operand {
	var tok = p.tokens.Current()
	var ope *ast.Operand
	switch tok.Type {
	case token.REG:
		ope = p.parseRegOperand()
	case token.LBRACKET:
		ope = p.parseConstOperand()
	case token.INT, token.HASH:
		ope = p.parseImmOperand()
	default:
		p.unexpect(tok)
	}
	return ope
}

// Register operand is start with "%"
func (p *Parser) parseRegOperand() *ast.Operand {
	var tok = p.tokens.Expect(token.REG)
	return ast.NewOperand(ast.REG, createLiteral(tok), tok.Position)
}

// Const operand is start with "["
func (p *Parser) parseConstOperand() *ast.Operand {
	var tok = p.tokens.Expect(token.LBRACKET)
	var val = p.tokens.Expect(token.IDENT)
	p.tokens.Expect(token.RBRACKET)
	return ast.NewOperand(ast.CONST, createLiteral(val), tok.Position)
}

// Imm operand is start with "$" or number
func (p *Parser) parseImmOperand() *ast.Operand {
	var tok = p.tokens.SkipIfTest(token.DOLLAR)
	var val = p.tokens.Expect(token.INT)
	var pos *token.Position
	if tok != nil {
		pos = tok.Position
	} else {
		pos = val.Position
	}
	return ast.NewOperand(ast.REG, createLiteral(val), pos)
}

func (p *Parser) parseDirective() *ast.Directive {
	var label = p.parseLabelOptional()
	var kind = p.tokens.ExpectDirective()
	var tok *token.Token
	switch kind.Type {
	case token.DIR_STRING:
		tok = p.tokens.Expect(token.STR)
	case token.DIR_LONG:
		tok = p.tokens.Expect(token.INT)
	case token.DIR_DECIMAL:
		tok = p.tokens.Expect(token.FLOAT)
	case token.DIR_GLOBAL:
		tok = p.tokens.Expect(token.IDENT)
	default:
		p.unexpect(tok)
	}
	return ast.NewDirective(label, createIdentifier(kind), createLiteral(tok), p.parseCommentOptional())
}

func (p *Parser) parseLabelOptional() *ast.Label {
	if p.tokens.Current().Test(token.LABEL, token.IDENT) {
		return p.parseLabel()
	}
	return nil
}

func (p *Parser) parseLabel() *ast.Label {
	var tok = p.tokens.Expect(token.LABEL, token.IDENT)
	p.tokens.SkipIfTest(token.COLON)
	return createLabel(tok)
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	var tok = p.tokens.Expect(token.IDENT)
	return ast.NewIdentifier(tok.Value, tok.Position)
}

func (p *Parser) parseLiteral() *ast.Literal {
	var tok = p.tokens.Current()
	var literal = createLiteral(tok)
	p.tokens.Next()
	return literal
}

func (p *Parser) unexpect(tok *token.Token) {
	panic(token.NewSyntaxError(fmt.Sprintf("Unexpected token \"%d\" of value \"%s\"", tok.Type, tok.Value), tok.Position))
}

func createLabel(tok *token.Token) *ast.Label {
	return ast.NewLabel(tok.Value, tok.Position)
}

func createIdentifier(tok *token.Token) *ast.Identifier {
	return ast.NewIdentifier(tok.Value, tok.Position)
}

func createLiteral(tok *token.Token) *ast.Literal {
	var exp *ast.Literal
	switch tok.Type {
	// constant
	case token.INT:
		num, _ := strconv.ParseInt(tok.Value, 10, 64)
		exp = ast.NewLiteral(num, tok.Value, tok.Position)
	case token.FLOAT:
		num, _ := strconv.ParseFloat(tok.Value, 64)
		exp = ast.NewLiteral(num, tok.Value, tok.Position)
	case token.STR:
		fallthrough
	default:
		exp = ast.NewLiteral(tok.Value, tok.Value, tok.Position)
	}
	return exp
}

func NewParser(tokens *token.Stream) *Parser {
	return &Parser{tokens: tokens}
}
