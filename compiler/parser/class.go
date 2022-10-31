package parser

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
)

type ClassPropertyContext struct {
	final      *ast.MemberModifier
	abstract   *ast.MemberModifier
	visibility *ast.MemberModifier
	static     *ast.MemberModifier

	finalDefined      bool
	abstractDefined   bool
	visibilityDefined bool
	staticDefined     bool
}

func (ctx *ClassPropertyContext) setFinal(final *ast.MemberModifier) error {
	if ctx.finalDefined {
		return token.NewSyntaxError("illegal combination of MemberModifiers", final.Position())
	}
	ctx.final = final
	ctx.finalDefined = true
	return nil
}

func (ctx *ClassPropertyContext) setAbstract(abstract *ast.MemberModifier) error {
	if ctx.abstractDefined {
		return token.NewSyntaxError("illegal combination of MemberModifiers", abstract.Position())
	}
	ctx.abstract = abstract
	ctx.abstractDefined = true
	return nil
}

func (ctx *ClassPropertyContext) setVisibility(visibility *ast.MemberModifier) error {
	if ctx.visibilityDefined {
		return token.NewSyntaxError("illegal combination of MemberModifiers", visibility.Position())
	}
	ctx.visibility = visibility
	ctx.visibilityDefined = true
	return nil
}

func (ctx *ClassPropertyContext) setStatic(static *ast.MemberModifier) error {
	if ctx.staticDefined {
		return token.NewSyntaxError("illegal combination of MemberModifiers", static.Position())
	}
	ctx.static = static
	ctx.staticDefined = true
	return nil
}

func NewClassPropertyContext() *ClassPropertyContext {
	return &ClassPropertyContext{}
}
