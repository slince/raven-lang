package compiler

import (
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir/types"
)

func (c *Compiler) compileType(node *ast.Identifier) types.Type {
	var _type types.Type
	switch node.Value {
	case "int4":
		_type = types.I4
	case "int8":
		_type = types.I8
	case "int16":
		_type = types.I16
	case "int32":
		_type = types.I32
	case "int64":
		_type = types.I64
	case "int128":
		_type = types.I128
	case "int256":
		_type = types.I256
	case "int512":
		_type = types.I512
	case "int1024":
		_type = types.I1024
	case "uint4":
		_type = types.U4
	case "uint8":
		_type = types.U8
	case "uint16":
		_type = types.U16
	case "uint32":
		_type = types.U32
	case "uint64":
		_type = types.U64
	case "uint128":
		_type = types.U128
	case "uint256":
		_type = types.U256
	case "uint512":
		_type = types.U512
	case "uint1024":
		_type = types.U1024
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
