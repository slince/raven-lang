package ast

import "github.com/slince/php-plus/compiler/token"

type VarSpec struct {
	Id   *Identifier
	Kind *Identifier
	Init Expr
	node
}

type VarDecl struct {
	Kind  string
	Specs []*VarSpec
	decl
}

func NewVarSpec(id *Identifier, kind *Identifier, init Expr, pos *token.Position) *VarSpec {
	var dec = &VarSpec{
		Id:   id,
		Kind: kind,
		Init: init,
	}
	dec.pos = pos
	return dec
}

func NewVarDecl(kind string, specs []*VarSpec, pos *token.Position) *VarDecl {
	var dec = &VarDecl{
		Kind:  kind,
		Specs: specs,
	}
	dec.pos = pos
	return dec
}

type FuncArg struct {
	Id   *Identifier
	Kind *Identifier
	node
}

type Func struct {
	Id   *Identifier // Identifier or nil
	Args []*FuncArg
	Kind *Identifier // Identifier or nil
	Body *BlockStmt
	node
}

type FuncDecl struct {
	Func *Func
	decl
}

func (dec *FuncDecl) Position() *token.Position {
	return dec.Func.Position()
}

func NewFuncArg(id *Identifier, kind *Identifier, pos *token.Position) *FuncArg {
	var arg = &FuncArg{
		Id:   id,
		Kind: kind,
	}
	arg.pos = pos
	return arg
}

func NewFunc(id *Identifier, args []*FuncArg, kind *Identifier, body *BlockStmt, pos *token.Position) *Func {
	var fn = &Func{
		Id:   id,
		Args: args,
		Kind: kind,
		Body: body,
	}
	fn.pos = pos
	return fn
}

func NewFuncDecl(function *Func) *FuncDecl {
	return &FuncDecl{
		Func: function,
	}
}

type MemberModifier struct {
	Value string
	node
}

type PropertyDefinition struct {
	Visibility *MemberModifier
	Static     *MemberModifier
	Kind       string // const or generic
	Value      *VarSpec
	node
}

type MethodDefinition struct {
	Final      *MemberModifier
	Abstract   *MemberModifier
	Visibility *MemberModifier
	Static     *MemberModifier
	Value      *Func
	node
}

type ClassBody struct {
	Props   []*PropertyDefinition
	Methods []*MethodDefinition
	node
}

type Class struct {
	Id         *Identifier //Identifier
	Extends    *Identifier // Identifier
	Implements []*Identifier
	Body       *ClassBody
	node
}

type ClassDecl struct {
	Class *Class
	decl
}

func (dec *ClassDecl) Position() *token.Position {
	return dec.Class.Position()
}

func NewModifier(value string, pos *token.Position) *MemberModifier {
	var mod = &MemberModifier{
		Value: value,
	}
	mod.pos = pos
	return mod
}

func NewPropertyDefinition(visibility *MemberModifier, static *MemberModifier, kind string, value *VarSpec, pos *token.Position) *PropertyDefinition {
	var def = &PropertyDefinition{
		Visibility: visibility,
		Static:     static,
		Kind:       kind,
		Value:      value,
	}
	def.pos = pos
	return def
}

func NewMethodDefinition(final *MemberModifier, abstract *MemberModifier, visibility *MemberModifier, static *MemberModifier, value *Func, pos *token.Position) *MethodDefinition {
	var def = &MethodDefinition{
		Final:      final,
		Abstract:   abstract,
		Visibility: visibility,
		Static:     static,
		Value:      value,
	}
	def.pos = pos
	return def
}

func NewClassBody(props []*PropertyDefinition, methods []*MethodDefinition, pos *token.Position) *ClassBody {
	var body = &ClassBody{
		Props:   props,
		Methods: methods,
	}
	body.pos = pos
	return body
}

func NewClass(id *Identifier, extends *Identifier, implements []*Identifier, body *ClassBody, pos *token.Position) *Class {
	var cls = &Class{
		Id:         id,
		Extends:    extends,
		Implements: implements,
		Body:       body,
	}
	cls.pos = pos
	return cls
}

func NewClassDecl(class *Class) *ClassDecl {
	return &ClassDecl{
		Class: class,
	}
}
