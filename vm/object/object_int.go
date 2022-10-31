package object

import "fmt"

type Int struct {
	Value int64
}

func (i *Int) Type() Type{
	return T_INT
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.Value)
}
