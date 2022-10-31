package native

type register = uint8

const (
	rax register = iota
	rbx
	rcx
	rdx
	rsp
	rbp
	rsi
	rdi
	r8
	r9
	r10
	r11
	r12
	r13
	r14
	r15
	eax
	edi
	edx
)

var fnArgRegs = []register{rdi, rsi, rdx, rcx, r8, r9}
