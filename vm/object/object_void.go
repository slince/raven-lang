package object

type Void struct {
}

func (n *Void) Type() Type{
	return T_VOID
}

func (f *Void) String() string {
	return "void"
}
