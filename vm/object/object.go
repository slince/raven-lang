package object

import "fmt"

type Type uint8

const (
	primitiveStart  Type = iota
	T_VOID
	T_INT
	T_FLOAT
	T_BOOL
	T_NULL
	T_CHAR
	primitiveEnd
	T_STRING
	T_ARRAY
	T_MAP
	T_TUPLE
	T_FUNCTION
)

// TypeSizes 类型宽度
var TypeSizes = map[Type]uint8{
	T_VOID: 0,
	T_INT: 8,
	T_FLOAT: 8,
	T_BOOL: 1,
	T_NULL: 0,
	T_CHAR: 1,
}

var TypeNames = map[Type]string{
	T_VOID: "void",
	T_INT: "int",
	T_FLOAT: "float",
	T_BOOL: "bool",
	T_NULL: "null",
	T_CHAR: "char",
	T_STRING: "string",
	T_ARRAY: "array",
	T_MAP: "map",
	T_TUPLE: "tuple",
	T_FUNCTION: "function",
}

func (t Type) IsPrimitive() bool{
	return primitiveStart < t && t < primitiveEnd
}

func (t Type) Size() (uint8, error){
	if t.IsPrimitive() {
		return TypeSizes[t], nil
	}
	return 0, fmt.Errorf("the type is not primitive")
}

func (t Type) Name() string{
	return TypeNames[t]
}

type Object interface {
	// Type return its variable type
	Type() Type

	// String returns a string-representation of the given value.
	String() string
}

func SizeOf(val Object) uint8{
	var size, err = val.Type().Size()
	if err != nil {
		return 0
	}
	return size
}

var (
	TrueValue = &Bool{true}
	FalseValue = &Bool{false}
	NullValue = &Null{}
	VoidValue = &Void{}
)

type Computable interface{

}

