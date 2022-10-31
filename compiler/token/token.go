package token

type Type uint8

const (
	ILLEGAL Type = iota

	literal_begin
	ID    // foo
	NULL  // null
	TRUE  // true
	FALSE // false
	STR
	INT   // 123
	FLOAT // 123.56
	literal_end

	// punctuation
	operator_begin
	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %

	AND // &
	OR  // |
	XOR // ^
	NOT // ~
	SHL // <<
	SHR // >>

	assign_begin
	ASSIGN     // =
	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	DIV_ASSIGN // /=
	MOD_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=
	assign_end

	INC // ++
	DEC // --

	LOGIC_AND // &&
	LOGIC_OR  // ||
	LOGIC_NOT // !
	EQ        // ==
	LT        // <
	GT        // >
	NEQ       // !=
	GEQ       // >=
	LEQ       // <=
	operator_end

	LPAREN   // (
	LBRACKET // [
	LBRACE   // {
	RPAREN   // )
	RBRACKET // ]
	RBRACE   // }

	COMMA     // ,
	COLON     // :
	SEMICOLON // ;
	DOT       // .
	QUESTION  // ?

	ARROW        // ->
	DOUBLE_ARROW // =>

	// keyword
	keyword_begin
	IF     // if
	ELSEIF // elseif
	ELSE   // else

	SWITCH  // switch
	CASE    // case
	DEFAULT // default

	BREAK    // break
	CONTINUE // continue
	RETURN   // return

	FOR     // for
	WHILE   // while
	DO      // do
	FOREACH // foreach
	AS      // as
	IN      // in

	LET   // let
	CONST // const

	THROW   // throw
	TRY     // try
	CATCH   // catch
	FINALLY // finally

	//INT     // int
	//FLOAT   // float
	//STRING  // string
	//ARRAY   // array
	//OBJECT  // object
	//BOOL    // bool
	//BOOLEAN // boolean

	NAMESPACE // namespace
	USE       // use
	PACKAGE   // package
	IMPORT    // import

	FUNCTION   // function
	CLASS      // class
	EXTENDS    // extends
	IMPLEMENTS // implements
	PUBLIC     // public
	PROTECTED  // protected
	PRIVATE    // private
	FINAL      // final
	STATIC     // static
	ABSTRACT   // abstract

	ECHO    // echo
	PRINT   // print
	DECLARE // declare
	keyword_end

	EOF // eof
)

var tokens = map[Type]string{
	ID:    "identifier",
	NULL:  "null",
	TRUE:  "true",
	FALSE: "false",
	STR:   "string",
	INT:   "int",
	FLOAT: "float",

	//punctuation
	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	MOD: "%",

	AND: "&",
	OR:  "|",
	XOR: "^",
	NOT: "~",
	SHL: "<<",
	SHR: ">>",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	DIV_ASSIGN: "/=",
	MOD_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	INC: "++",
	DEC: "--",

	LOGIC_AND: "&&",
	LOGIC_OR:  "||",
	EQ:        "==",
	LT:        "<",
	GT:        ">",
	ASSIGN:    "=",
	LOGIC_NOT: "!",

	NEQ: "!=",
	GEQ: ">=",
	LEQ: "<=",

	LPAREN:   "(",
	LBRACKET: "[",
	LBRACE:   "{",
	RPAREN:   ")",
	RBRACKET: "]",
	RBRACE:   "}",

	COMMA:     ",",
	COLON:     ":",
	SEMICOLON: ";",
	DOT:       ".",
	QUESTION:  "?",

	ARROW:        "->",
	DOUBLE_ARROW: "=>",

	// "keywords",
	// choice
	IF:     "if",
	ELSEIF: "elseif",
	ELSE:   "else",

	SWITCH:  "switch",
	CASE:    "case",
	DEFAULT: "default",

	// control
	BREAK:    "break",
	CONTINUE: "continue",
	RETURN:   "return",

	// loop.x
	FOR:     "for",
	WHILE:   "while",
	DO:      "do",
	FOREACH: "foreach",
	AS:      "as",
	IN:      "in",
	// variable declaration
	LET:   "let",
	CONST: "const",

	// exceptions
	THROW:   "throw",
	TRY:     "try",
	CATCH:   "catch",
	FINALLY: "finally",

	//INT:     "int",
	//FLOAT:   "float",
	//STRING:  "string",
	//ARRAY:   "array",
	//OBJECT:  "object",
	//BOOL:    "bool",
	//BOOLEAN: "boolean",

	NAMESPACE: "namespace",
	USE:       "use",
	PACKAGE:   "package",
	IMPORT:    "import",

	FUNCTION:   "function",
	CLASS:      "class",
	EXTENDS:    "extends",
	IMPLEMENTS: "implements",
	PUBLIC:     "public",
	PROTECTED:  "protected",
	PRIVATE:    "private",
	FINAL:      "final",
	STATIC:     "static",
	ABSTRACT:   "abstract",

	ECHO:    "echo",
	PRINT:   "print",
	DECLARE: "declare",

	EOF: "eof",
}

// keywords
var keywords map[string]Type

func init() {
	keywords = make(map[string]Type)
	for i := keyword_begin; i <= keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// assignment
var Assigns []Type

func init() {
	Assigns = make([]Type, 0)
	for i := assign_begin; i <= assign_end; i++ {
		Assigns = append(Assigns, i)
	}
}

type Token struct {
	Type     Type
	Value    string
	Position *Position
}

func (t Token) Test(kind ...Type) bool {
	return IndexOf(kind, t.Type) > -1
}

func (t Token) IsBinaryOperator() bool {
	return IsOperator(t.Type)
}

func ValueOf(kind Type) string {
	return tokens[kind]
}

func IsKeyword(kind Type) bool {
	return keyword_begin < kind && kind < keyword_end
}

func IsOperator(kind Type) bool {
	return operator_begin < kind && kind < operator_end
}

func IsLiteral(kind Type) bool {
	return literal_begin < kind && kind < literal_end
}

func IsAssign(kind Type) bool {
	return assign_begin < kind && kind < assign_end
}

func NewToken(kind Type, value string, position *Position) *Token {
	return &Token{
		Type:     kind,
		Value:    value,
		Position: position,
	}
}

func Lookup(name string) Type {
	if kind, ok := keywords[name]; ok {
		return kind
	}
	return ID
}

func IndexOf(haystack []Type, needle Type) int {
	for key, value := range haystack {
		if value == needle {
			return key
		}
	}
	return -1
}
