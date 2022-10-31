package ir

import (
	"fmt"
	"github.com/slince/php-plus/ir/types"
)

// ___ [ Function argument ] __________________________________________________

// FunctionArgument is a function argument.
type FunctionArgument struct {
	// Parameter name.
	Name string
	// Parameter type.
	Kind types.Type
}

// String returns the LLVM syntax representation of the function argument as a
// type-value pair.
func (p *FunctionArgument) String() string {
	return fmt.Sprintf("%s %s", p.Kind, p.Name)
}

// NewFuncParam returns a new function argument based on the given name and type.
func NewFuncParam(name string, kind types.Type) *FunctionArgument {
	return &FunctionArgument{
		Name: name,
		Kind: kind,
	}
}

type Function struct {
	Name string
	// Function signature.
	Signature *types.FuncType
	// Function arguments.
	Arguments []*FunctionArgument
	Blocks    []*BasicBlock
}

func (f *Function) NewBlock(name string) *BasicBlock {
	var block = NewBlock(name)
	f.Blocks = append(f.Blocks, block)
	return block
}

func (f *Function) Alias(name string) *BasicBlock {

}

// NewFunction returns a new function based on the given function name, return type
// and function arguments.
func NewFunction(name string, retType types.Type, arguments ...*FunctionArgument) *Function {
	paramTypes := make([]types.Type, len(arguments))
	for i, param := range arguments {
		paramTypes[i] = param.Kind
	}
	sig := types.NewFunc(retType, paramTypes...)
	f := &Function{Name: name, Signature: sig, Arguments: arguments}
	return f
}
