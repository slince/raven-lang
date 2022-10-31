package types

import (
	"fmt"
	"strings"
)

var (
	// Integer types.
	I1    = &IntType{BitWidth: 1}    // i1
	I2    = &IntType{BitWidth: 2}    // i2
	I3    = &IntType{BitWidth: 3}    // i3
	I4    = &IntType{BitWidth: 4}    // i4
	I5    = &IntType{BitWidth: 5}    // i5
	I6    = &IntType{BitWidth: 6}    // i6
	I7    = &IntType{BitWidth: 7}    // i7
	I8    = &IntType{BitWidth: 8}    // i8
	I16   = &IntType{BitWidth: 16}   // i16
	I32   = &IntType{BitWidth: 32}   // i32
	I64   = &IntType{BitWidth: 64}   // i64
	I128  = &IntType{BitWidth: 128}  // i128
	I256  = &IntType{BitWidth: 256}  // i256
	I512  = &IntType{BitWidth: 512}  // i512
	I1024 = &IntType{BitWidth: 1024} // i1024

	// Float types.
	F32 = &FloatType{BitWidth: 32} // f32
	F64 = &FloatType{BitWidth: 64} // f64

	// Bool
	Bool = &BoolType{}

	Void = &VoidType{}

	String = &StringType{}
)

// IsVoid reports whether the given type is a void type.
func IsVoid(t Type) bool {
	_, ok := t.(*VoidType)
	return ok
}

// IsFunc reports whether the given type is a function type.
func IsFunc(t Type) bool {
	_, ok := t.(*FuncType)
	return ok
}

// IsInt reports whether the given type is an integer type.
func IsInt(t Type) bool {
	_, ok := t.(*IntType)
	return ok
}

// IsFloat reports whether the given type is a floating-point type.
func IsFloat(t Type) bool {
	_, ok := t.(*FloatType)
	return ok
}

// IsArray reports whether the given type is an array type.
func IsArray(t Type) bool {
	_, ok := t.(*ArrayType)
	return ok
}

// Equal reports whether t and u are of equal type.
func Equal(t, u Type) bool {
	return t.Equal(u)
}

// Type is an LLVM IR type.
//
// A Type has one of the following underlying types.
//
//   - [*types.VoidType]
//   - [*types.FuncType]
//   - [*types.IntType]
//   - [*types.FloatType]
//   - [*types.MMXType]
//   - [*types.PointerType]
//   - [*types.VectorType]
//   - [*types.LabelType]
//   - [*types.TokenType]
//   - [*types.MetadataType]
//   - [*types.ArrayType]
//   - [*types.StructType]
type Type interface {
	fmt.Stringer
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

// IntType is an LLVM IR integer type.
type IntType struct {
	// Integer size in number of bits.
	BitWidth uint64
}

// NewInt returns a new integer type based on the given integer bit size.
func NewInt(bitWidth uint64) *IntType {
	return &IntType{
		BitWidth: bitWidth,
	}
}

// Equal reports whether t and u are of equal type.
func (t *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.BitWidth == u.BitWidth
	}
	return false
}

// String returns the string representation of the integer type.
func (t *IntType) String() string {
	return fmt.Sprintf("int%d", t.BitWidth)
}

// FloatType is an LLVM IR floating-point type.
type FloatType struct {
	// Integer size in number of bits.
	BitWidth uint64
}

// Equal reports whether t and u are of equal type.
func (t *FloatType) Equal(u Type) bool {
	if u, ok := u.(*FloatType); ok {
		return t.BitWidth == u.BitWidth
	}
	return false
}

// String returns the string representation of the floating-point type.
func (t *FloatType) String() string {
	return fmt.Sprintf("float%d", t.BitWidth)
}

type BoolType struct {
}

// Equal reports whether t and u are of equal type.
func (b *BoolType) Equal(u Type) bool {
	u, ok := u.(*BoolType)
	return ok
}

// String returns the string representation of the floating-point type.
func (b *BoolType) String() string {
	return "bool"
}

type StringType struct {
}

// Equal reports whether t and u are of equal type.
func (s *StringType) Equal(u Type) bool {
	u, ok := u.(*StringType)
	return ok
}

// String returns the string representation of the floating-point type.
func (s *StringType) String() string {
	return "string"
}

// --- [ Void types ] ----------------------------------------------------------

// VoidType is an LLVM IR void type.
type VoidType struct {
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	_, ok := u.(*VoidType)
	return ok
}

// String returns the string representation of the void type.
func (t *VoidType) String() string {
	// 'void'
	return "void"
}

// --- [ Function types ] ------------------------------------------------------

// FuncType is an LLVM IR function type.
type FuncType struct {
	// Type name; or empty if not present.
	TypeName string
	// Return type.
	RetType Type
	// Function parameters.
	Params []Type
	// Variable number of function arguments.
	Variadic bool
}

// NewFunc returns a new function type based on the given return type and
// function parameter types.
func NewFunc(retType Type, params ...Type) *FuncType {
	return &FuncType{
		RetType: retType,
		Params:  params,
	}
}

// Equal reports whether t and u are of equal type.
func (t *FuncType) Equal(u Type) bool {
	if u, ok := u.(*FuncType); ok {
		if !t.RetType.Equal(u.RetType) {
			return false
		}
		if len(t.Params) != len(u.Params) {
			return false
		}
		for i := range t.Params {
			if !t.Params[i].Equal(u.Params[i]) {
				return false
			}
		}
		return t.Variadic == u.Variadic
	}
	return false
}

// String returns the string representation of the function type.
func (t *FuncType) String() string {
	// RetType=Type '(' Arguments ')'
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s (", t.RetType)
	for i, param := range t.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.String())
	}
	if t.Variadic {
		if len(t.Params) > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("...")
	}
	buf.WriteString(")")
	return buf.String()
}

// --- [ Array types ] ---------------------------------------------------------

// ArrayType is an LLVM IR array type.
type ArrayType struct {
	// Type name; or empty if not present.
	TypeName string
	// Array length.
	Len uint64
	// Element type.
	ElemType Type
}

// NewArray returns a new array type based on the given array length and element
// type.
func NewArray(len uint64, elemType Type) *ArrayType {
	return &ArrayType{
		Len:      len,
		ElemType: elemType,
	}
}

// Equal reports whether t and u are of equal type.
func (t *ArrayType) Equal(u Type) bool {
	if u, ok := u.(*ArrayType); ok {
		if t.Len != u.Len {
			return false
		}
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

// String returns the string representation of the array type.
func (t *ArrayType) String() string {
	// '[' Len=UintLit 'x' Elem=Type ']'
	return fmt.Sprintf("[%d x %s]", t.Len, t.ElemType)
}

// --- [ Vector types ] --------------------------------------------------------

// VectorType is an LLVM IR vector type.
type VectorType struct {
	// Type name; or empty if not present.
	TypeName string
	// Scalable vector type.
	Scalable bool
	// Vector length.
	Len uint64
	// Element type.
	ElemType Type
}

// NewVector returns a new vector type based on the given vector length and
// element type.
func NewVector(len uint64, elemType Type) *VectorType {
	return &VectorType{
		Len:      len,
		ElemType: elemType,
	}
}

// Equal reports whether t and u are of equal type.
func (t *VectorType) Equal(u Type) bool {
	if u, ok := u.(*VectorType); ok {
		if t.Scalable != u.Scalable {
			return false
		}
		if t.Len != u.Len {
			return false
		}
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

// String returns the string representation of the vector type.
func (t *VectorType) String() string {
	// scalable: '<' 'vscale' 'x' Len=UintLit 'x' Elem=Type '>' non-scalable: '<'
	// Len=UintLit 'x' Elem=Type '>'
	if t.Scalable {
		// '<' 'vscale' 'x' Len=UintLit 'x' Elem=Type '>'
		return fmt.Sprintf("<vscale x %d x %s>", t.Len, t.ElemType)
	}
	// '<' Len=UintLit 'x' Elem=Type '>'
	return fmt.Sprintf("<%d x %s>", t.Len, t.ElemType)
}

// --- [ Structure types ] -----------------------------------------------------

// StructType is an LLVM IR structure type. Identified (named) struct types are
// uniqued by type names, not by structural identity.
type StructType struct {
	// Type name; or empty if not present.
	TypeName string
	// Packed memory layout.
	Packed bool
	// Struct fields.
	Fields []Type
	// Opaque struct type.
	Opaque bool
}

// NewStruct returns a new struct type based on the given field types.
func NewStruct(fields ...Type) *StructType {
	return &StructType{
		Fields: fields,
	}
}

// Equal reports whether t and u are of equal type.
func (t *StructType) Equal(u Type) bool {
	if u, ok := u.(*StructType); ok {
		if len(t.TypeName) > 0 || len(u.TypeName) > 0 {
			// Identified struct types are uniqued by type names, not by structural
			// identity.
			//
			// t or u is an identified struct type.
			return t.TypeName == u.TypeName
		}
		// Literal struct types are uniqued by structural identity.
		if t.Packed != u.Packed {
			return false
		}
		if len(t.Fields) != len(u.Fields) {
			return false
		}
		for i := range t.Fields {
			if !t.Fields[i].Equal(u.Fields[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// String returns the string representation of the structure type.
func (t *StructType) String() string {
	// Opaque struct type.
	//
	//	'opaque'
	//
	// Struct type.
	//
	//	'{' Fields=(Type separator ',')+? '}'
	//
	// Packed struct type.
	//
	//	'<' '{' Fields=(Type separator ',')+? '}' '>'   -> PackedStructType
	if t.Opaque {
		return "opaque"
	}
	if len(t.Fields) == 0 {
		if t.Packed {
			return "<{}>"
		}
		return "{}"
	}
	buf := &strings.Builder{}
	if t.Packed {
		buf.WriteString("<")
	}
	buf.WriteString("{ ")
	for i, field := range t.Fields {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	buf.WriteString(" }")
	if t.Packed {
		buf.WriteString(">")
	}
	return buf.String()
}
