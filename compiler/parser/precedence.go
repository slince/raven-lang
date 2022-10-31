package parser

type OperatorAssoc uint8

const (
	ASSOC_LEFT  OperatorAssoc = 1
	ASSOC_RIGHT OperatorAssoc = 2
)

type OperatorPrecedence struct {
	Precedence int
	Assoc      OperatorAssoc
}

var defaultOperatorPrecedence = OperatorPrecedence{
	Precedence: -1,
}

var binaryOperators = map[string]OperatorPrecedence{
	"or":  {Precedence: 10, Assoc: ASSOC_LEFT},
	"||":  {Precedence: 10, Assoc: ASSOC_LEFT},
	"and": {Precedence: 15, Assoc: ASSOC_LEFT},
	"&&":  {Precedence: 15, Assoc: ASSOC_LEFT},
	"|":   {Precedence: 16, Assoc: ASSOC_LEFT},
	"^":   {Precedence: 17, Assoc: ASSOC_LEFT},
	"&":   {Precedence: 18, Assoc: ASSOC_LEFT},
	"==":  {Precedence: 20, Assoc: ASSOC_LEFT},
	"===": {Precedence: 20, Assoc: ASSOC_LEFT},
	"is":  {Precedence: 20, Assoc: ASSOC_LEFT},
	"!=":  {Precedence: 20, Assoc: ASSOC_LEFT},
	"!==": {Precedence: 20, Assoc: ASSOC_LEFT},
	"<":   {Precedence: 20, Assoc: ASSOC_LEFT},
	">":   {Precedence: 20, Assoc: ASSOC_LEFT},
	">=":  {Precedence: 20, Assoc: ASSOC_LEFT},
	"<=":  {Precedence: 20, Assoc: ASSOC_LEFT},
	"in":  {Precedence: 20, Assoc: ASSOC_LEFT},
	"<<":  {Precedence: 25, Assoc: ASSOC_LEFT},
	">>":  {Precedence: 25, Assoc: ASSOC_LEFT},
	"+":   {Precedence: 30, Assoc: ASSOC_LEFT},
	"-":   {Precedence: 30, Assoc: ASSOC_LEFT},
	"~":   {Precedence: 40, Assoc: ASSOC_LEFT},
	"*":   {Precedence: 60, Assoc: ASSOC_LEFT},
	"/":   {Precedence: 60, Assoc: ASSOC_LEFT},
	"%":   {Precedence: 60, Assoc: ASSOC_LEFT},
}

var unaryOperators = map[string]OperatorPrecedence{
	"not": {Precedence: 50},
	"!":   {Precedence: 50},
	"-":   {Precedence: 500},
	"+":   {Precedence: 500},
	"--":  {Precedence: 500},
	"++":  {Precedence: 500},
}

func IsBinaryOperator(value string) bool {
	_, ok := binaryOperators[value]
	return ok
}

func IsUnaryOperator(value string) bool {
	_, ok := unaryOperators[value]
	return ok
}

func BinaryPrecedence(value string) OperatorPrecedence {
	if precedence, ok := binaryOperators[value]; ok {
		return precedence
	}
	return defaultOperatorPrecedence
}

func UnaryPrecedence(value string) OperatorPrecedence {
	if precedence, ok := unaryOperators[value]; ok {
		return precedence
	}
	return defaultOperatorPrecedence
}
