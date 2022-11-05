package token

type Kind uint8

const (
	ILLEGAL Kind = iota

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

	AND     // &
	OR      // |
	XOR     // ^
	NOT     // ~
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

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

	T_INT // int
	T_LONG
	T_FLOAT  // float
	T_STRING // string
	T_BOOL   // bool

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

var tokens = map[Kind]string{
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
var keywords map[string]Kind

func init() {
	keywords = make(map[string]Kind)
	for i := keyword_begin; i <= keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// assignment
var Assigns []Kind

func init() {
	Assigns = make([]Kind, 0)
	for i := assign_begin; i <= assign_end; i++ {
		Assigns = append(Assigns, i)
	}
}

type Token struct {
	Kind     Kind
	Literal  string
	Position *Position
}

func (t Token) Test(kind ...Kind) bool {
	return IndexOf(kind, t.Kind) > -1
}

func (t Token) IsBinaryOperator() bool {
	return IsOperator(t.Kind)
}

func ValueOf(kind Kind) string {
	return tokens[kind]
}

func IsKeyword(kind Kind) bool {
	return keyword_begin < kind && kind < keyword_end
}

func IsOperator(kind Kind) bool {
	return operator_begin < kind && kind < operator_end
}

func IsLiteral(kind Kind) bool {
	return literal_begin < kind && kind < literal_end
}

func IsAssign(kind Kind) bool {
	return assign_begin < kind && kind < assign_end
}

func NewToken(kind Kind, literal string, position *Position) *Token {
	return &Token{
		Kind:     kind,
		Literal:  literal,
		Position: position,
	}
}

func Lookup(name string) Kind {
	if kind, ok := keywords[name]; ok {
		return kind
	}
	return ID
}

func IndexOf(haystack []Kind, needle Kind) int {
	for key, value := range haystack {
		if value == needle {
			return key
		}
	}
	return -1
}
