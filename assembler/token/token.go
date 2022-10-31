package token

import (
	"strings"
)

type Type uint8

const (
	literal_begin Type = iota
	IDENT              // foo
	LABEL              // letter _ number $ # @ .
	NULL               // null
	TRUE               // true
	FALSE              // false
	STR                // "abc"
	INT                // 123
	FLOAT              // 123.56
	literal_end

	DOT       // .
	DOLLAR    // $
	HASH      // #
	PERCENT   // %
	COLON     // :
	SEMICOLON // ;

	LPAREN   // (
	LBRACKET // [
	LBRACE   // {
	RPAREN   // )
	RBRACKET // ]
	RBRACE   // }

	COMMENT // //
	REG     // reg0
	// directive
	directive_begin
	DIR_SECTION
	DIR_DATA
	DIR_CODE
	DIR_GLOBAL
	DIR_STRING
	DIR_LONG
	DIR_DECIMAL
	directive_end

	// instruction
	instruction_begin
	INS_NOP
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
	instruction_end
	EOF // end
)

var tokens = map[Type]string{
	IDENT:     "id",
	LABEL:     "label",
	DOT:       ".",
	DOLLAR:    "$",
	HASH:      "#",
	PERCENT:   "%",
	COLON:     ":",
	SEMICOLON: ";",

	LPAREN:   "(",
	LBRACKET: "[",
	LBRACE:   "{",
	RPAREN:   ")",
	RBRACKET: "]",
	RBRACE:   "}",

	COMMENT: "comment",
	REG:     "reg",

	DIR_SECTION: "section",
	DIR_DATA:    ".data",
	DIR_CODE:    ".code",
	DIR_GLOBAL:  "global",
	DIR_STRING:  "string",
	DIR_LONG:    "long",
	DIR_DECIMAL: "decimal",

	INS_NOP:  "nop",
	INS_HALT: "halt",
	// 数据加载指令
	INS_LOAD: "load",
	// 跳转指令
	INS_JMP: "jmp",
	INS_JZ:  "jz",
	INS_JNZ: "jnz",
	// 函数调用指令
	INS_ARG:   "arg",
	INS_CALL:  "call",
	INS_ENTER: "enter",
	INS_LEAVE: "leave",
	INS_RET:   "ret",
	// 四则运算
	INS_ADD: "add",
	INS_SUB: "sub",
	INS_MUL: "mul",
	INS_DIV: "div",
	INS_MOD: "mod",
	// 位运算
	INS_AND: "biand",
	INS_OR:  "bior",
	INS_XOR: "bixor",
	INS_NOT: "binot",
	INS_SHL: "bishl",
	INS_SHR: "bishr",
	// 逻辑运算
	INS_LOGIC_NOT: "logic_not",
	INS_EQ:        "eq",
	INS_LT:        "lt",
	INS_GT:        "gt",
	INS_NEQ:       "neq",
	INS_LEQ:       "leq",
	INS_GEQ:       "geq",
	// 自增/自减
	INS_INC: "inc",
	INS_DEC: "dec",
	// 数据构造
	INS_ARRAY: "array",
	INS_MAP:   "map",
	INS_FUNC:  "func",
	INS_LEN:   "len",
	INS_CAP:   "cap",
	EOF:       "eof",
}

// directive
var directives map[string]Type
var directiveTypes []Type

// instructions
var instructions map[string]Type
var instructionTypes []Type

func init() {
	directives = make(map[string]Type)
	directiveTypes = make([]Type, directive_end-directive_begin)
	for i := directive_begin; i <= directive_end; i++ {
		directives[tokens[i]] = i
		directiveTypes = append(directiveTypes, i)
	}
	instructions = make(map[string]Type)
	instructionTypes = make([]Type, instruction_end-instruction_begin)
	for i := instruction_begin; i <= instruction_end; i++ {
		instructions[tokens[i]] = i
		instructionTypes = append(instructionTypes, i)
	}
}

func IsDirective(kind Type) bool {
	return directive_begin < kind && kind <= directive_end
}

func IsInstruction(kind Type) bool {
	return instruction_begin < kind && kind <= instruction_end
}

type Token struct {
	Type     Type
	Value    string
	Position *Position
}

func (t *Token) Test(kind ...Type) bool {
	return IndexOf(kind, t.Type) > -1
}

func (t *Token) IsDirective() bool {
	return IsDirective(t.Type)
}

func (t *Token) IsInstruction() bool {
	return IsInstruction(t.Type)
}

func NewToken(kind Type, value string, position *Position) *Token {
	return &Token{
		Type:     kind,
		Value:    value,
		Position: position,
	}
}

func Lookup(value string) Type {
	if kind, ok := directives[value]; ok {
		return kind
	}
	if kind, ok := instructions[value]; ok {
		return kind
	}
	if value[0] == '.' {
		return LABEL
	}
	if strings.Index(value, "reg") == 0 {
		return REG
	}
	return IDENT
}

func ValueOf(kind Type) string {
	return tokens[kind]
}

func IndexOf(haystack []Type, needle Type) int {
	for key, value := range haystack {
		if value == needle {
			return key
		}
	}
	return -1
}
