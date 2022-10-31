package object

type String struct {
	Value string
}


func (s *String) Type() Type{
	return T_STRING
}

func (s *String) String() string {
	return s.Value
}
