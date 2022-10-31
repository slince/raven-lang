package object

import "fmt"

type Bool struct {
	Value bool
}

func (b *Bool) Type() Type{
	return T_BOOL
}

func (b *Bool) String() string {
	return fmt.Sprintf("%t", b.Value)
}
