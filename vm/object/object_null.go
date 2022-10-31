package object

// Null wraps nothing and implements our Object interface.
type Null struct{}

func (n *Null) Type() Type{
	return T_NULL
}

// Inspect returns a string-representation of the given object.
func (n *Null) String() string {
	return "null"
}