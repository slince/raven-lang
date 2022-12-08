package ir

import (
	"github.com/slince/php-plus/ir/types"
	"github.com/slince/php-plus/ir/value"
)

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
	Globals   []*Global
	Consts    []*Const
	Functions []*Function
}

func (m *Module) NewFunction(name string, retType types.Type, params ...*FunctionArgument) *Function {
	var fun = NewFunction(name, retType, params...)
	m.Functions = append(m.Functions, fun)
	return fun
}

func (m *Module) NewGlobal(name string, kind types.Type, init *value.Const) *Global {
	var global = NewGlobal(name, kind, init)
	m.Globals = append(m.Globals, global)
	return global
}

func (m *Module) NewConst(name string, kind types.Type, init *value.Const) *Const {
	var constant = NewConst(name, kind, init)
	m.Consts = append(m.Consts, constant)
	return constant
}

func NewModule(name string) *Module {
	return &Module{
		Name:      name,
		Types:     []types.Type{},
		Globals:   []*Global{},
		Consts:    []*Const{},
		Functions: []*Function{},
	}
}
