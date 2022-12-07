package compiler

import (
	"fmt"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/compiler/token"
	"github.com/slince/php-plus/ir/types"
)

func (c *Compiler) compileType(node *ast.Identifier) (types.Type, error) {
	var kind types.Type
	var err error
	switch node.Value {
	case "int4":
		kind = types.I4
	case "int8":
		kind = types.I8
	case "int16":
		kind = types.I16
	case "int32":
		kind = types.I32
	case "int64":
		kind = types.I64
	case "int128":
		kind = types.I128
	case "int256":
		kind = types.I256
	case "int512":
		kind = types.I512
	case "int1024":
		kind = types.I1024
	case "uint4":
		kind = types.U4
	case "uint8":
		kind = types.U8
	case "uint16":
		kind = types.U16
	case "uint32":
		kind = types.U32
	case "uint64":
		kind = types.U64
	case "uint128":
		kind = types.U128
	case "uint256":
		kind = types.U256
	case "uint512":
		kind = types.U512
	case "uint1024":
		kind = types.U1024
	case "float32":
		kind = types.F32
	case "float64":
		kind = types.F64
	case "string":
		kind = types.String
	case "bool":
		kind = types.Bool
	case "void":
		kind = types.Void
	default:
		err = token.NewSyntaxError(fmt.Sprintf("unknown identifier %s", node.Value), node.Position())
	}
	return kind, err
}
