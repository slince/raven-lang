package token

type Position struct {
	Offset int // offset, starting at 0
	Line   int // line number, starting at 1
	Column int // column number, starting at 1 (byte count)
}

func NewPosition(offset int, line int, column int) *Position {
	return &Position{offset, line, column}
}
