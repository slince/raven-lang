package ir

import (
	"github.com/pkg/errors"
	"github.com/slince/php-plus/ir/value"
)

type SymbolTable struct {
	Outer *SymbolTable
	vars  map[string]*value.Variable
}

func (s *SymbolTable) GetVariable(name string) (*value.Variable, error) {
	if v, ok := s.vars[name]; ok {
		return v, nil
	} else if s.Outer != nil {
		return s.Outer.GetVariable(name)
	}
	return nil, errors.Errorf("unresolved reference '%s'", name)
}

func (s *SymbolTable) AddVariable(variable *value.Variable) {
	s.vars[variable.Name] = variable
}

func NewSymbolTable(outer *SymbolTable) *SymbolTable {
	return &SymbolTable{
		Outer: outer,
		vars:  make(map[string]*value.Variable),
	}
}
