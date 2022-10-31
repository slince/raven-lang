package ast

import (
	"github.com/slince/php-plus/assembler/token"
)

type Comment struct {
	Text     *Literal
	position *token.Position
}

func (c *Comment) Position() *token.Position {
	return c.position
}

func NewComment(Value *Literal, position *token.Position) *Comment {
	return &Comment{
		Text:     Value,
		position: position,
	}
}
