package insts

type Unary interface {
	unary()
}

type unary struct {
	instruction
}

type Neg struct {
	Result ir.Operand
	Ope    ir.Operand
	unary
}

func NewNeg(result ir.Operand, ope ir.Operand) *Neg {
	return &Neg{Result: result, Ope: ope}
}
