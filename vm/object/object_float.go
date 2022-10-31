package object

import "fmt"

type Float struct {
	Value float64
}

func (f *Float) Type() Type{
	return T_FLOAT
}

func (f *Float) String() string {
	return fmt.Sprintf("%f", f.Value)
}
