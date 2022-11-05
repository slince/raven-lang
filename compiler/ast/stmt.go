package ast

import (
	"github.com/slince/php-plus/compiler/token"
)

type Stmt interface {
	Node
	stmt()
}

type stmt struct{ node }

func (smt stmt) stmt() {}

type Program struct {
	Modules []*Module
	stmt
}

func (p *Program) AddModule(module *Module) {
	p.Modules = append(p.Modules, module)
}

type Module struct {
	Name *Literal
	Body *BlockStmt
	stmt
}

type BlockStmt struct {
	Body []Stmt
	stmt
}

type BreakStmt struct {
	stmt
}

type ContinueStmt struct {
	stmt
}

type ReturnStmt struct {
	Argument Expr
	stmt
}

type ExpressionStmt struct {
	Expr Expr
	stmt
}

type ForStmt struct {
	Init   Node // VariableDeclaration | Expr
	Test   Expr
	Update Expr
	Body   Stmt
	stmt
}

type ForeachStmt struct {
	Source Expr // VariableDeclaration | Expr
	Key    Expr
	Value  Expr
	Body   Stmt
	stmt
}

type IfStmt struct {
	Test       Expr
	Consequent *BlockStmt
	Alternate  Stmt
	stmt
}

type SwitchCase struct {
	Test       Expr
	Consequent []Stmt
	Default    bool
	node
}

type SwitchStmt struct {
	Discriminant Expr
	Cases        []*SwitchCase
	stmt
}

type WhileStmt struct {
	Test Expr
	Body *BlockStmt
	stmt
}

type DoWhileStmt struct {
	Test Expr
	Body *BlockStmt
	stmt
}

type ThrowStmt struct {
	Argument Expr
	stmt
}

type CatchClause struct {
	Variable *Identifier
	Kind     *Identifier
	Body     *BlockStmt
	node
}

type TryStmt struct {
	Body      *BlockStmt
	Catches   []*CatchClause
	Finalizer *BlockStmt
	stmt
}

func NewProgram(modules []*Module, pos *token.Position) *Program {
	var smt = &Program{Modules: modules}
	smt.pos = pos
	return smt
}

func NewModule(body *BlockStmt, pos *token.Position) *Module {
	var module = &Module{Body: body}
	module.pos = pos
	return module
}

func NewExpressionStmt(expr Expr, pos *token.Position) *ExpressionStmt {
	var smt = &ExpressionStmt{
		Expr: expr,
	}
	smt.pos = pos
	return smt
}

func NewForStmt(init Node, test Expr, update Expr, body Stmt, pos *token.Position) *ForStmt {
	var smt = &ForStmt{
		Init:   init,
		Test:   test,
		Update: update,
		Body:   body,
	}
	smt.pos = pos
	return smt
}

func NewForeachStmt(source Expr, key Expr, value Expr, body Stmt, pos *token.Position) *ForeachStmt {
	var smt = &ForeachStmt{
		Source: source,
		Key:    key,
		Value:  value,
		Body:   body,
	}
	smt.pos = pos
	return smt
}

func NewIfStmt(test Expr, consequent *BlockStmt, alternate Stmt, pos *token.Position) *IfStmt {
	var smt = &IfStmt{
		Test:       test,
		Consequent: consequent,
		Alternate:  alternate,
	}
	smt.pos = pos
	return smt
}

func NewSwitchCase(test Expr, consequent []Stmt, defaults bool, pos *token.Position) *SwitchCase {
	var smt = &SwitchCase{
		Test:       test,
		Consequent: consequent,
		Default:    defaults,
	}
	smt.pos = pos
	return smt
}

func NewSwitchStmt(discriminant Expr, cases []*SwitchCase, pos *token.Position) *SwitchStmt {
	var smt = &SwitchStmt{
		Discriminant: discriminant,
		Cases:        cases,
	}
	smt.pos = pos
	return smt
}

func NewBlockStmt(body []Stmt, pos *token.Position) *BlockStmt {
	var smt = &BlockStmt{Body: body}
	smt.pos = pos
	return smt
}

func NewReturnStmt(argument Expr, pos *token.Position) *ReturnStmt {
	var smt = &ReturnStmt{
		Argument: argument,
	}
	smt.pos = pos
	return smt
}

func NewBreakStmt(pos *token.Position) *BreakStmt {
	var smt = &BreakStmt{}
	smt.pos = pos
	return smt
}

func NewContinueStmt(pos *token.Position) *ContinueStmt {
	var smt = &ContinueStmt{}
	smt.pos = pos
	return smt
}

func NewWhileStmt(test Expr, body *BlockStmt, pos *token.Position) *WhileStmt {
	var smt = &WhileStmt{
		Test: test,
		Body: body,
	}
	smt.pos = pos
	return smt
}

func NewDoWhileStmt(test Expr, body *BlockStmt, pos *token.Position) *DoWhileStmt {
	var smt = &DoWhileStmt{
		Test: test,
		Body: body,
	}
	smt.pos = pos
	return smt
}

func NewThrowStmt(argument Expr, pos *token.Position) *ThrowStmt {
	var smt = &ThrowStmt{
		Argument: argument,
	}
	smt.pos = pos
	return smt
}

func NewCatchClause(variable *Identifier, kind *Identifier, body *BlockStmt, pos *token.Position) *CatchClause {
	var clause = &CatchClause{
		Variable: variable,
		Kind:     kind,
		Body:     body,
	}
	clause.pos = pos
	return clause
}

func NewTryStmt(body *BlockStmt, catches []*CatchClause, finalizer *BlockStmt, pos *token.Position) *TryStmt {
	var smt = &TryStmt{
		Body:      body,
		Catches:   catches,
		Finalizer: finalizer,
	}
	smt.pos = pos
	return smt
}
