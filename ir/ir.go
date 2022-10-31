package ir

import "github.com/slince/php-plus/ir/types"

type Ir struct {
	Global string
	Blocks map[string]*BasicBlock
}

type Program struct {
	Modules []*Module
}

func (p *Program) NewModule(name string) *Module {
	var mod = NewModule(name)
	p.Modules = append(p.Modules, mod)
	return mod
}

func NewProgram() *Program {
	return &Program{Modules: []*Module{}}
}

type Module struct {
	Name      string
	Types     []types.Type
	Globals   []*Variable
	Consts    []*Const
	Functions []*Function
}

func (m *Module) NewFunction(name string, retType types.Type, params ...*FunctionArgument) *Function {
	var fun = NewFunction(name, retType, params...)
	m.Functions = append(m.Functions, fun)
	return fun
}

func NewModule(name string) *Module {
	return &Module{Name: name, Functions: []*Function{}}
}
