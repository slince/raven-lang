package ir

type If struct {
	instruction
	Cond        Block
	Body        Block
	Alternative Block
}

type Loop struct {
	instruction
	Cond Block
	Step Block
	Body Block
}

type SwitchCase struct {
	Cond    Block
	Body    Block
	Default bool
}

type Switch struct {
	instruction
	Cond  Block
	Cases []*SwitchCase
}

func NewIf(cond Block, body Block, alternative Block) *If {
	return &If{
		Cond:        cond,
		Body:        body,
		Alternative: alternative,
	}
}

func NewLoop(cond Block, step Block, body Block) *Loop {
	return &Loop{
		Cond: cond, Step: step, Body: body,
	}
}

func NewSwitchCase(cond Block, body Block, defaults bool) *SwitchCase {
	return &SwitchCase{
		Cond:    cond,
		Body:    body,
		Default: defaults,
	}
}

func NewSwitch(cond Block, cases []*SwitchCase) *Switch {
	return &Switch{
		Cond:  cond,
		Cases: cases,
	}
}
