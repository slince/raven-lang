package ast

import "github.com/slince/php-plus/assembler/token"

type Program struct {
	Body     []Node
	position *token.Position
}

func (p *Program) Position() *token.Position {
	return p.position
}

func NewProgram(body []Node, position *token.Position) *Program {
	return &Program{
		Body:     body,
		position: position,
	}
}

type Section struct {
	Name     *Label
	Body     []Node
	position *token.Position
}

func (b *Section) Position() *token.Position {
	return b.position
}

func NewSection(Name *Label, body []Node, position *token.Position) *Section {
	return &Section{
		Name:     Name,
		Body:     body,
		position: position,
	}
}

type Block struct {
	Label    *Identifier
	Body     []Node
	position *token.Position
}

func (b *Block) Position() *token.Position {
	return b.Label.position
}

func NewBlock(label *Identifier, body []Node) *Block {
	return &Block{
		Label: label,
		Body:  body,
	}
}
