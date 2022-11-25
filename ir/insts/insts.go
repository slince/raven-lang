package insts

type Instruction interface {
	inst()
}

type instruction struct {
}

func (i instruction) inst() {}
