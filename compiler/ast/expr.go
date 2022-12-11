package ast

import (
	"github.com/slince/php-plus/compiler/token"
)

// Identifier represents an identifier.
type Identifier struct {
	Value string
	expr
}

type Literal struct {
	Kind  string
	Raw   string
	Value interface{}
	expr
}

// array list expression.
type ArrayExpr struct {
	Elements []Expr
	expr
}

func (exp *ArrayExpr) AddElement(elements ...Expr) {
	exp.Elements = append(exp.Elements, elements...)
}

func (exp ArrayExpr) IsEmpty() bool {
	return len(exp.Elements) == 0
}

type AssignmentExpr struct {
	Left     *Identifier
	Operator string
	Right    Expr
	expr
}

type BinaryExpr struct {
	Left     Expr
	Operator string
	Right    Expr
	expr
}

type CallExpr struct {
	Callee    Expr
	Arguments []Expr
	expr
}

type ClassExpr struct {
	*Class
}

func (exp ClassExpr) Position() *token.Position {
	return exp.Class.Position()
}
func (exp ClassExpr) Expr() {}

type FunctionExpr struct {
	*Function
}

func (exp FunctionExpr) Position() *token.Position {
	return exp.Function.Position()
}
func (exp FunctionExpr) Expr() {}

type MapExpr struct {
	Elements map[Expr]Expr
	expr
}

func (exp *MapExpr) AddElement(key Expr, element Expr) {
	exp.Elements[key] = element
}

func (exp MapExpr) IsEmpty() bool {
	return len(exp.Elements) == 0
}

type MemberExpr struct {
	Object   Expr
	Property Expr
	expr
}

type UnaryExpr struct {
	Operator string
	Target   Expr
	expr
}

type UpdateExpr struct {
	Operator string
	Target   Expr
	Prefix   bool
	expr
}

type VariableExpr struct {
	Value string
	expr
}

func NewIdentifier(value string, pos *token.Position) *Identifier {
	var ident = &Identifier{Value: value}
	ident.pos = pos
	return ident
}

func NewLiteral(kind string, value interface{}, raw string, pos *token.Position) *Literal {
	var exp = &Literal{
		Kind:  kind,
		Raw:   raw,
		Value: value,
	}
	exp.pos = pos
	return exp
}

func NewArrayExpr(elements []Expr, pos *token.Position) *ArrayExpr {
	var exp = &ArrayExpr{Elements: elements}
	exp.pos = pos
	return exp
}

func NewAssignmentExpr(left *Identifier, operator string, right Expr, pos *token.Position) *AssignmentExpr {
	var exp = &AssignmentExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
	exp.pos = pos
	return exp
}

func NewBinaryExpr(operator string, left Expr, right Expr, pos *token.Position) *BinaryExpr {
	var exp = &BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
	exp.pos = pos
	return exp
}

func NewCallExpr(callee Expr, arguments []Expr, pos *token.Position) *CallExpr {
	var exp = &CallExpr{
		Callee:    callee,
		Arguments: arguments,
	}
	exp.pos = pos
	return exp
}

func NewClassExpr(class *Class) *ClassExpr {
	return &ClassExpr{
		Class: class,
	}
}

func NewFunctionExpr(function *Function) *FunctionExpr {
	return &FunctionExpr{
		Function: function,
	}
}

func NewMapExpr(elements map[Expr]Expr, pos *token.Position) *MapExpr {
	var exp = &MapExpr{
		Elements: elements,
	}
	exp.pos = pos
	return exp
}

func NewMemberExpr(object Expr, property Expr, pos *token.Position) *MemberExpr {
	var exp = &MemberExpr{
		Object:   object,
		Property: property,
	}
	exp.pos = pos
	return exp
}

func NewUnaryExpr(operator string, target Expr, pos *token.Position) *UnaryExpr {
	var exp = &UnaryExpr{
		Operator: operator,
		Target:   target,
	}
	exp.pos = pos
	return exp
}

func NewUpdateExpr(operator string, target Expr, prefix bool, pos *token.Position) *UpdateExpr {
	var exp = &UpdateExpr{
		Operator: operator,
		Target:   target,
		Prefix:   prefix,
	}
	exp.pos = pos
	return exp
}
func NewVariableExpr(value string, pos *token.Position) *VariableExpr {
	var exp = &VariableExpr{
		Value: value,
	}
	exp.pos = pos
	return exp
}
