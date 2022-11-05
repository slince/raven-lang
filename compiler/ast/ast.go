package ast

import (
	"github.com/slince/php-plus/compiler/token"
)

type Node interface {
	Position() *token.Position
	node()
}

type Expr interface {
	Node
	expr()
}

type Stmt interface {
	Node
	stmt()
}

type Decl interface {
	Stmt
	decl()
}

type node struct{ pos *token.Position }

func (n node) Position() *token.Position { return n.pos }
func (n node) node()                     {}

type expr struct{ node }

func (exp expr) expr() {}

type stmt struct{ node }

func (smt stmt) stmt() {}

type decl struct{ stmt }

func (dec decl) decl() {}

// Identifier represents an identifier.
type Identifier struct {
	Value string
	node
}

func NewIdentifier(value string, pos *token.Position) *Identifier {
	var ident = &Identifier{Value: value}
	ident.pos = pos
	return ident
}
