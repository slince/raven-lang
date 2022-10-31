package parser

// OperandType 立即数类型
type OperandType = uint8

const (
	OPE_NONE  OperandType = iota // 不使用
	OPE_REG                      // 寄存器
	OPE_IMM                      // 立即数 32位
	OPE_CONST                    // 常量地址
	OPE_DYN
)

type InstructionType uint8

const (
	INS_NOP InstructionType = iota
	INS_HALT
	// 数据加载指令
	INS_LOAD
	// 跳转指令
	INS_JMP
	INS_JZ
	INS_JNZ
	// 函数调用指令
	INS_ARG
	INS_CALL
	INS_ENTER
	INS_LEAVE
	INS_RET
	// 四则运算
	INS_ADD
	INS_SUB
	INS_MUL
	INS_DIV
	INS_MOD
	// 位运算
	INS_AND // &
	INS_OR  // |
	INS_XOR // ^
	INS_NOT // ~
	INS_SHL
	INS_SHR
	// 逻辑运算
	INS_LOGIC_NOT // !
	INS_EQ        // ==
	INS_LT        // <
	INS_GT        // >
	INS_NEQ       // !=
	INS_LEQ       // <=
	INS_GEQ       // >=
	// 自增/自减
	INS_INC // ++
	INS_DEC // --
	// 数据构造
	INS_ARRAY
	INS_MAP
	INS_FUNC
	INS_LEN
	INS_CAP
)

type InstructionDefinition struct {
	// 操作码名称
	InstructionType InstructionType
	// 操作数位宽
	First  OperandType
	Second OperandType
	Third  OperandType
}

var InstructionDefinitions = map[string]*InstructionDefinition{
	"nop":  {INS_NOP, OPE_NONE, OPE_NONE, OPE_NONE},  // 空操作
	"halt": {INS_HALT, OPE_NONE, OPE_NONE, OPE_NONE}, // 终止虚拟机执行
	"load": {INS_LOAD, OPE_REG, OPE_DYN, OPE_NONE},   //将有符号立即数拷贝到寄存器
	"jmp":  {INS_JMP, OPE_IMM, OPE_NONE, OPE_NONE},   // 无条件跳转
	"jz":   {INS_JZ, OPE_IMM, OPE_NONE, OPE_NONE},    // 标志寄存器的值为1则跳转
	"jnz":  {INS_JNZ, OPE_IMM, OPE_NONE, OPE_NONE},   // 标志寄存器的值不为1则跳转

	"call": {INS_CALL, OPE_REG, OPE_NONE, OPE_NONE}, // 标志寄存器的值不为1则跳转

	// 二元运算，将两个常量表/全局变量/立即数/寄存器的值运算后存入寄存器
	"add": {INS_ADD, OPE_REG, OPE_DYN, OPE_DYN},
	"sub": {INS_SUB, OPE_REG, OPE_DYN, OPE_DYN},
	"mul": {INS_MUL, OPE_REG, OPE_DYN, OPE_DYN},
	"div": {INS_DIV, OPE_REG, OPE_DYN, OPE_DYN},
	"mod": {INS_MOD, OPE_REG, OPE_DYN, OPE_DYN},

	"and": {INS_AND, OPE_REG, OPE_DYN, OPE_DYN},
	"or":  {INS_OR, OPE_REG, OPE_DYN, OPE_DYN},
	"xor": {INS_XOR, OPE_REG, OPE_DYN, OPE_DYN},
	"not": {INS_NOT, OPE_REG, OPE_DYN, OPE_DYN},
	"shl": {INS_SHL, OPE_REG, OPE_DYN, OPE_DYN},
	"shr": {INS_SHR, OPE_REG, OPE_DYN, OPE_DYN},

	// 逻辑运算，将两个常量表/全局变量/立即数/寄存器的值运算后存入标志寄存器
	"lnot": {INS_LOGIC_NOT, OPE_DYN, OPE_NONE, OPE_NONE},
	"eq":   {INS_EQ, OPE_DYN, OPE_DYN, OPE_NONE},
	"lt":   {INS_LT, OPE_DYN, OPE_DYN, OPE_NONE},
	"gt":   {INS_GT, OPE_DYN, OPE_DYN, OPE_NONE},
	"neq":  {INS_NEQ, OPE_DYN, OPE_DYN, OPE_NONE},
	"leq":  {INS_GT, OPE_DYN, OPE_DYN, OPE_NONE},
	"geq":  {INS_GT, OPE_DYN, OPE_DYN, OPE_NONE},

	"inc": {INS_INC, OPE_DYN, OPE_NONE, OPE_NONE},
	"dec": {INS_GT, OPE_DYN, OPE_NONE, OPE_NONE},

	"array": {INS_ARRAY, OPE_DYN, OPE_NONE, OPE_NONE}, // 构建数组，根据常量表/全局变量/立即数/寄存器的值构建数组
	"map":   {INS_MAP, OPE_DYN, OPE_DYN, OPE_NONE},    // 构建map，根据常量表/全局变量/立即数/寄存器的值构建map
	"len":   {INS_LEN, OPE_REG, OPE_NONE, OPE_NONE},   // 计算数组或者map的长度
	"cap":   {INS_CAP, OPE_REG, OPE_NONE, OPE_NONE},   // 计算map的容量
}

func InstDefinitionOf(mnemonic string) *InstructionDefinition {
	return InstructionDefinitions[mnemonic]
}
