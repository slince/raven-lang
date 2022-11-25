package types

import "fmt"

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

// IntType is an LLVM IR integer type.
type IntType struct {
	// Integer size in number of bits.
	BitWidth uint64
	Unsigned bool
}

// NewInt returns a new integer type based on the given integer bit size.
func NewInt(bitWidth uint64, unsigned bool) *IntType {
	return &IntType{
		BitWidth: bitWidth,
		Unsigned: unsigned,
	}
}

// Equal reports whether t and u are of equal type.
func (i *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return i.BitWidth == u.BitWidth && i.Unsigned == u.Unsigned
	}
	return false
}

// String returns the string representation of the integer type.
func (i *IntType) String() string {
	var unsigned byte
	if i.Unsigned {
		unsigned = 'c'
	}
	return fmt.Sprintf("%cint%d", unsigned, i.BitWidth)
}

// FloatType is an LLVM IR floating-point type.
type FloatType struct {
	// Integer size in number of bits.
	BitWidth uint64
}

// Equal reports whether t and u are of equal type.
func (f *FloatType) Equal(u Type) bool {
	if u, ok := u.(*FloatType); ok {
		return f.BitWidth == u.BitWidth
	}
	return false
}

// String returns the string representation of the floating-point type.
func (f *FloatType) String() string {
	return fmt.Sprintf("float%d", f.BitWidth)
}

// NewFloat returns a new float type based on the given integer bit size.
func NewFloat(bitWidth uint64) *FloatType {
	return &FloatType{
		BitWidth: bitWidth,
	}
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

type PointerType struct {
	BaseType Type
}

// Equal reports whether t and u are of equal type.
func (t *PointerType) Equal(u Type) bool {
	if u, ok := u.(*PointerType); ok {
		return t.BaseType.Equal(u.BaseType)
	}
	return false
}

// String returns the string representation of the vector type.
func (t *PointerType) String() string {
	return fmt.Sprintf("* %s", t.BaseType.String())
}

func NewPoint(base Type) *PointerType {
	return &PointerType{
		BaseType: base,
	}
}
