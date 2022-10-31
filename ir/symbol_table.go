package ir

type SymbolTable struct {
	Outer *SymbolTable
	vars  map[string]*Variable
}

func (s *SymbolTable) GetVariable(name string) *Variable {
	if v, ok := s.vars[name]; ok {
		return v
	} else if s.Outer != nil {
		return s.Outer.GetVariable(name)
	}
	return nil
}

func NewSymbolTable(outer *SymbolTable) *SymbolTable {
	return &SymbolTable{
		Outer: outer,
		vars:  make(map[string]*Variable),
	}
}
