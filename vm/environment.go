package vm

type Environment struct {
	functions map[string]BuiltinFunction
}

func (e *Environment) GetFunction(name string) (BuiltinFunction, bool) {
	fun, ok := e.functions[name]
	return fun, ok
}
