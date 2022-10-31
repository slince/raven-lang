package lexer

import (
	"fmt"
	"github.com/slince/php-plus/assembler/token"
	"github.com/slince/php-plus/utils"
)

type Lexer struct {
	source []rune
	end    int
	offset int
	line   int
	column int
}

func (l *Lexer) Lex() *token.Stream {
	var stream = token.NewStream()
	for {
		tok := l.NextToken()
		stream.Add(tok)
		if tok.Test(token.EOF) {
			break
		}
	}
	return stream
}

func (l *Lexer) NextToken() *token.Token {
	// Skip blank character
	l.skipWhitespace()
	if l.eof() {
		return token.NewToken(token.EOF, token.ValueOf(token.EOF), l.position())
	}

	var ch = l.current()

	var tok *token.Token
	var kind token.Type
	// parse source string.
	switch {
	case utils.IsDigit(ch):
		value, isFloat := l.readNumber()
		kind = token.INT
		if isFloat {
			kind = token.FLOAT
		}
		tok = token.NewToken(kind, value, l.position())
	// inst label
	case ch == '.' && utils.IsLetter(l.peek()):
		label := l.readLabel()
		kind = token.Lookup(label)
		tok = token.NewToken(kind, label, l.position())
	case utils.IsLetter(ch):
		id := l.readIdentifier()
		kind = token.Lookup(id)
		tok = token.NewToken(kind, id, l.position())
	case ch == '\'' || ch == '"' || ch == '`':
		tok = token.NewToken(token.STR, l.readString(ch), l.position())
	case ch == ';':
		tok = token.NewToken(token.COMMENT, l.readLine(), l.position())
	default:
		tok = l.lexPunctuation()
	}
	return tok
}

func (l *Lexer) skipWhitespace() {
	for !l.eof() && utils.IsWhitespace(l.current()) {
		l.next()
	}
}

func (l *Lexer) lexPunctuation() *token.Token {
	var kind token.Type
	var ch = l.current()
	switch ch {
	case '.':
		kind = token.DOT
	case '$':
		kind = token.DOLLAR
	case '%':
		kind = token.PERCENT
	case '#':
		kind = token.HASH
	case ':':
		kind = token.COLON
	case ';':
		kind = token.SEMICOLON
	default:
		l.unexpect(ch, l.position())
	}
	l.next()
	return token.NewToken(kind, token.ValueOf(kind), l.position())
}

func (l *Lexer) readNumber() (string, bool) {
	var isFloat = false
	var buf = l.readIf(func(ch rune) bool {
		if ch == '.' {
			if isFloat {
				return false
			}
			isFloat = true
			return true
		}
		return utils.IsDigit(ch)
	})
	return buf, isFloat
}

func (l *Lexer) readString(beginChar rune) string {
	l.next() // skip first ' or "
	var buf = l.readIf(func(ch rune) bool {
		return ch != beginChar
	})
	l.next() // skip end ' or "
	return buf
}

func (l *Lexer) readLabel() string {
	return l.readIf(isLabel)
}

func (l *Lexer) readIdentifier() string {
	return l.readIf(utils.IsIdentifier)
}

func (l *Lexer) readLine() string {
	return l.readIf(func(ch rune) bool {
		return ch != '\n'
	})
}

func (l *Lexer) readIf(predicate func(ch rune) bool) string {
	var buffer = make([]rune, 0)
	for !l.eof() {
		var ch = l.current()
		if !predicate(ch) {
			break
		}
		buffer = append(buffer, ch)
		l.next()
	}
	return string(buffer)
}

func (l *Lexer) next() rune {
	var ch = l.source[l.offset]
	l.offset++
	if ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
	return ch
}

func (l *Lexer) peek() rune {
	return l.source[l.offset+1]
}

func (l *Lexer) current() rune {
	return l.source[l.offset]
}

func (l *Lexer) position() *token.Position {
	return token.NewPosition(l.offset, l.line, l.column)
}

func (l *Lexer) eof() bool {
	return l.offset > l.end
}

func (l *Lexer) unexpect(ch rune, position *token.Position) {
	panic(token.NewSyntaxError(fmt.Sprintf("Unrecognized punctuation %s", string(ch)), position))
}

// isLabel whether the char is a valid label character
// Valid characters in labels are letters, numbers, _, $, #, @, ~, ., and ?.
// The only characters which may be used as the first character of an identifier are letters and "."
func isLabel(ch rune) bool {
	return ch == '$' || ch == '#' || ch == '@' || ch == '~' || ch == '.' || utils.IsIdentifier(ch)
}

func NewLexer(source string) *Lexer {
	var buf = []rune(source)

	return &Lexer{
		source: buf,
		end:    len(buf) - 1,
		offset: 0,
		line:   0,
		column: 0,
	}
}
