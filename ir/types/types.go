package types

import (
	"fmt"
)

var (
	I4    = NewInt(4, false)    // i4
	I8    = NewInt(8, false)    // i8
	I16   = NewInt(16, false)   // i16
	I32   = NewInt(32, false)   // i32
	I64   = NewInt(64, false)   // i64
	I128  = NewInt(128, false)  // i128
	I256  = NewInt(256, false)  // i256
	I512  = NewInt(512, false)  // i512
	I1024 = NewInt(1024, false) // i1024

	U4    = NewInt(4, true)    // i4
	U8    = NewInt(8, true)    // i8
	U16   = NewInt(16, true)   // i16
	U32   = NewInt(32, true)   // i32
	U64   = NewInt(64, true)   // i64
	U128  = NewInt(128, true)  // i128
	U256  = NewInt(256, true)  // i256
	U512  = NewInt(512, true)  // i512
	U1024 = NewInt(1024, true) // i1024

	F32 = NewFloat(32) // f32
	F64 = NewFloat(64) // f64

	Bool = &BoolType{}

	String = &StringType{}

	Void = &VoidType{}
	
	Nop = &NopType{}
)

type Type interface {
	fmt.Stringer
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

type NopType struct {
}

func (n *NopType) Equal(u Type) bool {
	_, ok := u.(*NopType)
	return ok
}

func (n *NopType) String() string {
	return "nop"
}

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
