package ir

import (
	"fmt"
	"github.com/slince/php-plus/ir/types"
)

// ___ [ function argument ] __________________________________________________

// FuncArg is a function argument.
type FuncArg struct {
	// Parameter name.
	Name string
	// Parameter type.
	Kind types.Type
}

// String returns the LLVM syntax representation of the function argument as a
// type-value pair.
func (p *FuncArg) String() string {
	return fmt.Sprintf("%s %s", p.Kind, p.Name)
}

// NewFuncArg returns a new function argument based on the given name and type.
func NewFuncArg(name string, kind types.Type) *FuncArg {
	return &FuncArg{
		Name: name,
		Kind: kind,
	}
}

type Func struct {
	Name string
	// Func signature.
	Signature *types.FuncType
	// Func arguments.
	Arguments []*FuncArg
	Blocks    []*BasicBlock
	Alias     map[string]string
}

func (f *Func) NewBlock(name string) *BasicBlock {
	var block = NewBlock(name)
	f.Blocks = append(f.Blocks, block)
	return block
}

func (f *Func) SetBlockAlias(source string, alias string) {
	f.Alias[alias] = source
}

// NewFunc returns a new function based on the given function name, return type
// and function arguments.
func NewFunc(name string, retType types.Type, arguments ...*FuncArg) *Func {
	paramTypes := make([]types.Type, len(arguments))
	for i, param := range arguments {
		paramTypes[i] = param.Kind
	}
	var sig = types.NewFunc(retType, paramTypes...)
	return &Func{
		Name:      name,
		Signature: sig,
		Arguments: arguments,
		Blocks:    []*BasicBlock{},
	}
}
