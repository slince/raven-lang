package ast

import "github.com/slince/php-plus/compiler/token"

type Decl interface {
	Stmt
	Decl()
}

type decl struct{ stmt }

func (dec decl) Decl() {}

type VariableDeclarator struct {
	Id   *Identifier
	Kind *Identifier
	Init Expr
	node
}

type VariableDeclaration struct {
	Kind        string
	Declarators []*VariableDeclarator
	decl
}

func NewVariableDeclarator(id *Identifier, kind *Identifier, init Expr, pos *token.Position) *VariableDeclarator {
	var dec = &VariableDeclarator{
		Id:   id,
		Kind: kind,
		Init: init,
	}
	dec.pos = pos
	return dec
}

func NewVariableDeclaration(kind string, declarators []*VariableDeclarator, pos *token.Position) *VariableDeclaration {
	var dec = &VariableDeclaration{
		Kind:        kind,
		Declarators: declarators,
	}
	dec.pos = pos
	return dec
}

type FunctionArgument struct {
	Id   *Identifier
	Kind *Identifier
	node
}

type Function struct {
	Id   *Identifier // Identifier or nil
	Args []*FunctionArgument
	Kind *Identifier // Identifier or nil
	Body *BlockStmt
	node
}

type FunctionDeclaration struct {
	Function *Function
	decl
}

func (dec *FunctionDeclaration) Position() *token.Position {
	return dec.Function.Position()
}

func NewFunctionArgument(id *Identifier, kind *Identifier, pos *token.Position) *FunctionArgument {
	var arg = &FunctionArgument{
		Id:   id,
		Kind: kind,
	}
	arg.pos = pos
	return arg
}

func NewFunction(id *Identifier, args []*FunctionArgument, kind *Identifier, body *BlockStmt, pos *token.Position) *Function {
	var fn = &Function{
		Id:   id,
		Args: args,
		Kind: kind,
		Body: body,
	}
	fn.pos = pos
	return fn
}

func NewFunctionDeclaration(function *Function) *FunctionDeclaration {
	return &FunctionDeclaration{
		Function: function,
	}
}

// Class ast collection
type MemberModifier struct {
	Value string
	node
}

type PropertyDefinition struct {
	Visibility *MemberModifier
	Static     *MemberModifier
	Kind       string // const or generic
	Value      *VariableDeclarator
	node
}

type MethodDefinition struct {
	Final      *MemberModifier
	Abstract   *MemberModifier
	Visibility *MemberModifier
	Static     *MemberModifier
	Value      *Function
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

type ClassDeclaration struct {
	Class *Class
	decl
}

func (dec *ClassDeclaration) Position() *token.Position {
	return dec.Class.Position()
}

func NewModifier(value string, pos *token.Position) *MemberModifier {
	var mod = &MemberModifier{
		Value: value,
	}
	mod.pos = pos
	return mod
}

func NewPropertyDefinition(visibility *MemberModifier, static *MemberModifier, kind string, value *VariableDeclarator, pos *token.Position) *PropertyDefinition {
	var def = &PropertyDefinition{
		Visibility: visibility,
		Static:     static,
		Kind:       kind,
		Value:      value,
	}
	def.pos = pos
	return def
}

func NewMethodDefinition(final *MemberModifier, abstract *MemberModifier, visibility *MemberModifier, static *MemberModifier, value *Function, pos *token.Position) *MethodDefinition {
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

func NewClassDeclaration(class *Class) *ClassDeclaration {
	return &ClassDeclaration{
		Class: class,
	}
}
