package ast

import (
	"github.com/slince/php-plus/compiler/token"
)

type Node interface {
	Position() *token.Position
	Node()
}

type Expr interface {
	Node
	Expr()
}

type Stmt interface {
	Node
	Stmt()
}

type Decl interface {
	Stmt
	Decl()
}

type node struct{ pos *token.Position }

func (n node) Position() *token.Position { return n.pos }
func (n node) Node()                     {}

type expr struct{ node }

func (exp expr) Expr() {}

type stmt struct{ node }

func (smt stmt) Stmt() {}

type decl struct{ stmt }

func (dec decl) Decl() {}
