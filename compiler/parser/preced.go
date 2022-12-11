package parser

type OperatorAssoc uint8

const (
	ASSOC_LEFT  OperatorAssoc = 1
	ASSOC_RIGHT OperatorAssoc = 2
)

type OperatorPrecedence struct {
	precedence int
	assoc      OperatorAssoc
}

var defaultOperatorPrecedence = OperatorPrecedence{
	precedence: -1,
}

var binaryOperators = map[string]OperatorPrecedence{
	"or":  {precedence: 10, assoc: ASSOC_LEFT},
	"||":  {precedence: 10, assoc: ASSOC_LEFT},
	"and": {precedence: 15, assoc: ASSOC_LEFT},
	"&&":  {precedence: 15, assoc: ASSOC_LEFT},
	"|":   {precedence: 16, assoc: ASSOC_LEFT},
	"^":   {precedence: 17, assoc: ASSOC_LEFT},
	"&":   {precedence: 18, assoc: ASSOC_LEFT},
	"==":  {precedence: 20, assoc: ASSOC_LEFT},
	"===": {precedence: 20, assoc: ASSOC_LEFT},
	"is":  {precedence: 20, assoc: ASSOC_LEFT},
	"!=":  {precedence: 20, assoc: ASSOC_LEFT},
	"!==": {precedence: 20, assoc: ASSOC_LEFT},
	"<":   {precedence: 20, assoc: ASSOC_LEFT},
	">":   {precedence: 20, assoc: ASSOC_LEFT},
	">=":  {precedence: 20, assoc: ASSOC_LEFT},
	"<=":  {precedence: 20, assoc: ASSOC_LEFT},
	"in":  {precedence: 20, assoc: ASSOC_LEFT},
	"<<":  {precedence: 25, assoc: ASSOC_LEFT},
	">>":  {precedence: 25, assoc: ASSOC_LEFT},
	"+":   {precedence: 30, assoc: ASSOC_LEFT},
	"-":   {precedence: 30, assoc: ASSOC_LEFT},
	"~":   {precedence: 40, assoc: ASSOC_LEFT},
	"*":   {precedence: 60, assoc: ASSOC_LEFT},
	"/":   {precedence: 60, assoc: ASSOC_LEFT},
	"%":   {precedence: 60, assoc: ASSOC_LEFT},
}

var unaryOperators = map[string]OperatorPrecedence{
	"!":  {precedence: 50},
	"~":  {precedence: 50},
	"-":  {precedence: 500},
	"+":  {precedence: 500},
	"--": {precedence: 500},
	"++": {precedence: 500},
}

func isBinaryOperator(value string) bool {
	_, ok := binaryOperators[value]
	return ok
}

func isUnaryOperator(value string) bool {
	_, ok := unaryOperators[value]
	return ok
}

func getBinaryPrecedence(value string) OperatorPrecedence {
	if precedence, ok := binaryOperators[value]; ok {
		return precedence
	}
	return defaultOperatorPrecedence
}

func getUnaryPrecedence(value string) OperatorPrecedence {
	if precedence, ok := unaryOperators[value]; ok {
		return precedence
	}
	return defaultOperatorPrecedence
}
