package ast

import "github.com/slince/php-plus/assembler/token"

// buf string "helloworld"
// num long 123
// len equ 123
type Directive struct {
	Label   *Label
	Kind    *Identifier
	Value   *Literal
	Comment *Comment
}

func (d *Directive) Position() *token.Position {
	if d.Label != nil {
		return d.Label.position
	}
	return d.Kind.position
}

func NewDirective(label *Label, kind *Identifier, value *Literal, comment *Comment) *Directive {
	return &Directive{Label: label, Kind: kind, Value: value, Comment: comment}
}
