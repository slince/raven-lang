package object

type Closure struct {
	Argument []Object
	Return   Object
	Func     *emulator.Function
}

func (f *Closure) Type() Type {
	return T_FUNCTION
}

func (f *Closure) String() string {
	return "closure"
}
