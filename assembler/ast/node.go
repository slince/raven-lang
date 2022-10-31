package ast

import "github.com/slince/php-plus/assembler/token"

type Node interface {
	Position() *token.Position
}
