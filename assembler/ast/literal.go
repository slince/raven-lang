package ast

import "github.com/slince/php-plus/assembler/token"

type Literal struct {
	raw      string
	value    interface{}
	position *token.Position
}

func (l *Literal) Position() *token.Position {
	return l.position
}

func NewLiteral(value interface{}, raw string, position *token.Position) *Literal {
	return &Literal{
		raw:      raw,
		value:    value,
		position: position,
	}
}
