package ast

import "github.com/slince/php-plus/assembler/token"

type Identifier struct {
	Value    string
	position *token.Position
}

func (i *Identifier) Position() *token.Position {
	return i.position
}

func NewIdentifier(value string, position *token.Position) *Identifier {
	return &Identifier{Value: value, position: position}
}
