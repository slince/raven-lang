package token

type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}

type SyntaxError struct {
	message  string
	position *Position
}

func (e SyntaxError) Error() string {
	return e.message
}

func NewSyntaxError(message string, position *Position) SyntaxError {
	return SyntaxError{message, position}
}
