package ast

import "github.com/slince/php-plus/assembler/token"

type Label struct {
	Value    string
	position *token.Position
}

func (l *Label) Position() *token.Position {
	return l.position
}

func NewLabel(value string, position *token.Position) *Label {
	return &Label{Value: value, position: position}
}
