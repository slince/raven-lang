package vm

type Opcode byte

const (
	OP_NOP Opcode = iota
	OP_HALT
	// 数据加载指令
	OP_MOVE
	OP_LOAD
	OP_LOAD_NULL
	OP_LOAD_BOOL
	OP_LOAD_CONST
	OP_STORE_GLOBAL
	OP_LOAD_GLOBAL
	// 跳转指令
	OP_JMP
	OP_JZ
	OP_JNZ
	// 函数调用指令
	OP_ARG
	OP_CALL
	OP_ENTER
	OP_LEAVE
	OP_RET
	// 四则运算
	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_MOD
	// 位运算
	OP_AND // &
	OP_OR  // |
	OP_XOR // ^
	OP_NOT // ~
	OP_SHL
	OP_SHR
	// 逻辑运算
	OP_LOGIC_AND // &&
	OP_LOGIC_OR  // ||
	OP_LOGIC_NOT // !
	OP_EQ        // ==
	OP_LT        // <
	OP_GT        // >
	OP_NEQ       // !=
	OP_LEQ       // <=
	OP_GEQ       // >=
	// 自增/自减
	OP_INC // ++
	OP_DEC // --
	// 数据构造
	OP_ARRAY
	OP_MAP
	OP_FUNC
	OP_LEN
	OP_CAP
)

// OperandType 立即数类型
type OperandType = uint8

const (
	OPE_NONE    OperandType = iota // 不使用
	OPE_IMM                        // 立即数 32位
	OPE_SMA_IMM                    // 小立即数 8位
	OPE_CONST                      // 常量地址
	OPE_GLOBAL                     // 全局变量
	OPE_REG                        // 寄存器
	OPE_DYN                        // 混合类型，最高位表示标志位
)

type Operand struct {
	Type  OperandType
	Value uint32
}

type InstructionDefinition struct {
	// 操作码名称
	Name string
	// 操作数位宽
	First  OperandType
	Second OperandType
	Third  OperandType
}

// Definitions 指令集接受的操作数是可变的,可以接收0-3个操作数；不同类型操作数字宽如下
//   - 寄存器索引类型采用8 bit存储，索引取值范围0-255
//   - 常量表 32 bit存储，取值范围0-4294967295；
//   - 全局变量表，同上
//   - 立即数，无符号扩展则同上；若立即数需要符号扩展，则取值范围变更为-2147483648-2147483647
//   - 小立即数 8bit存储，索引取值范围0-255
//   - 混合类型操作数（即操作数可能表示常量表，全局变量表，立即数，寄存器）采用32 bit存储，最高位为标志位，剩下31位表示操作数
var Definitions = map[Opcode]InstructionDefinition{
	OP_NOP:          {"OP_NOP", OPE_NONE, OPE_NONE, OPE_NONE},            // 空操作
	OP_HALT:         {"OP_HALT", OPE_NONE, OPE_NONE, OPE_NONE},           // 终止虚拟机执行
	OP_MOVE:         {"OP_MOVE", OPE_REG, OPE_REG, OPE_NONE},             // 寄存器间拷贝
	OP_LOAD:         {"OP_LOAD", OPE_REG, OPE_IMM, OPE_NONE},             //将有符号立即数拷贝到寄存器
	OP_LOAD_NULL:    {"OP_LOAD_NULL", OPE_REG, OPE_SMA_IMM, OPE_NONE},    // 给连续的寄存器设置null
	OP_LOAD_BOOL:    {"OP_LOAD_BOOL", OPE_REG, OPE_SMA_IMM, OPE_SMA_IMM}, // 给寄存器设置布尔值
	OP_LOAD_CONST:   {"OP_LOAD_CONST", OPE_REG, OPE_CONST, OPE_NONE},     //从常量表拷贝到寄存器
	OP_LOAD_GLOBAL:  {"OP_LOAD_GLOBAL", OPE_REG, OPE_GLOBAL, OPE_NONE},   // 从全局变量拷贝到寄存器
	OP_STORE_GLOBAL: {"OP_SET_GLOBAL", OPE_REG, OPE_GLOBAL, OPE_NONE},    // 将寄存器的值拷贝到全局变量
	OP_JMP:          {"OP_JMP", OPE_IMM, OPE_NONE, OPE_NONE},             // 无条件跳转
	OP_JZ:           {"OP_JZ", OPE_IMM, OPE_NONE, OPE_NONE},              // 标志寄存器的值为1则跳转
	OP_JNZ:          {"OP_JNZ", OPE_IMM, OPE_NONE, OPE_NONE},             // 标志寄存器的值不为1则跳转

	OP_CALL: {"OP_CALL", OPE_REG, OPE_NONE, OPE_NONE}, // 标志寄存器的值不为1则跳转

	// 二元运算，将两个常量表/全局变量/立即数/寄存器的值运算后存入寄存器
	OP_ADD: {"OP_ADD", OPE_REG, OPE_DYN, OPE_DYN},
	OP_SUB: {"OP_SUB", OPE_REG, OPE_DYN, OPE_DYN},
	OP_MUL: {"OP_MUL", OPE_REG, OPE_DYN, OPE_DYN},
	OP_DIV: {"OP_DIV", OPE_REG, OPE_DYN, OPE_DYN},
	OP_MOD: {"OP_MOD", OPE_REG, OPE_DYN, OPE_DYN},

	OP_AND: {"OP_AND", OPE_REG, OPE_DYN, OPE_DYN},
	OP_OR:  {"OP_OR", OPE_REG, OPE_DYN, OPE_DYN},
	OP_XOR: {"OP_XOR", OPE_REG, OPE_DYN, OPE_DYN},
	OP_NOT: {"OP_NOT", OPE_REG, OPE_DYN, OPE_DYN},
	OP_SHL: {"OP_SHL", OPE_REG, OPE_DYN, OPE_DYN},
	OP_SHR: {"OP_SHR", OPE_REG, OPE_DYN, OPE_DYN},

	// 逻辑运算，将两个常量表/全局变量/立即数/寄存器的值运算后存入标志寄存器
	OP_LOGIC_AND: {"OP_LOGIC_AND", OPE_DYN, OPE_DYN, OPE_NONE},
	OP_LOGIC_OR:  {"OP_LOGIC_OR", OPE_DYN, OPE_DYN, OPE_NONE},
	OP_LOGIC_NOT: {"OP_LOGIC_NOT", OPE_DYN, OPE_NONE, OPE_NONE},
	OP_EQ:        {"OP_EQ", OPE_DYN, OPE_DYN, OPE_NONE},
	OP_LT:        {"OP_LT", OPE_DYN, OPE_DYN, OPE_NONE},
	OP_GT:        {"OP_GT", OPE_DYN, OPE_DYN, OPE_NONE},
	OP_NEQ:       {"OP_NEQ", OPE_DYN, OPE_DYN, OPE_NONE},
	OP_LEQ:       {"OP_GT", OPE_DYN, OPE_DYN, OPE_NONE},
	OP_GEQ:       {"OP_GT", OPE_DYN, OPE_DYN, OPE_NONE},

	OP_INC: {"OP_INC", OPE_DYN, OPE_NONE, OPE_NONE},
	OP_DEC: {"OP_GT", OPE_DYN, OPE_NONE, OPE_NONE},

	OP_ARRAY: {"OP_ARRAY", OPE_DYN, OPE_NONE, OPE_NONE}, // 构建数组，根据常量表/全局变量/立即数/寄存器的值构建数组
	OP_MAP:   {"OP_MAP", OPE_DYN, OPE_DYN, OPE_NONE},    // 构建map，根据常量表/全局变量/立即数/寄存器的值构建map
	OP_LEN:   {"OP_LEN", OPE_REG, OPE_NONE, OPE_NONE},   // 计算数组或者map的长度
	OP_CAP:   {"OP_CAP", OPE_REG, OPE_NONE, OPE_NONE},   // 计算map的容量
}

type Instructions = []byte
