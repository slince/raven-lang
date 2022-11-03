package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir/types"
)

func (c *Compiler) compileType(node *ast.Identifier) types.Type {
	var _type types.Type
	switch node.Value {
	case "int64":
		_type = types.I64
	case "int32":
		_type = types.I32
	case "float32":
		_type = types.F32
	case "float64":
		_type = types.F64
	case "string":
		_type = types.String
	case "bool":
		_type = types.Bool
	case "void":
		_type = types.Void
	}
	return _type
}
