package ast

import (
	"github.com/slince/php-plus/compiler/token"
)

type Node interface {
	Position() *token.Position
	Node()
}

type node struct{ pos *token.Position }

func (n node) Position() *token.Position { return n.pos }
func (n node) Node()                     {}
