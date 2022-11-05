package lexer

import (
	"fmt"
	"github.com/slince/php-plus/compiler/token"
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
	// skip blank char
	l.skipWhitespace()
	if l.eof() {
		return token.NewToken(token.EOF, token.ValueOf(token.EOF), l.position())
	}

	var ch = l.current()

	var tok *token.Token
	var kind token.Kind

	// parse source string.
	switch {
	case utils.IsDecimal(ch):
		value, isFloat := l.readNumber()
		kind = token.INT
		if isFloat {
			kind = token.FLOAT
		}
		tok = token.NewToken(kind, value, l.position())
	case utils.IsLetter(ch):
		identifier := l.readIdentifier()
		kind = token.Lookup(identifier)
		tok = token.NewToken(kind, identifier, l.position())
	case ch == '\'' || ch == '"':
		tok = token.NewToken(token.STR, l.readString(ch), l.position())
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
	var kind token.Kind
	var ch = l.current()
	var next = l.look()
	switch ch {
	case '=':
		kind = token.ASSIGN
		if next == '=' {
			kind = token.EQ
			l.next()
		} else if next == '>' {
			kind = token.DOUBLE_ARROW
			l.next()
		}
	case '!':
		kind = token.LOGIC_NOT
		if next == '=' {
			kind = token.NEQ
			l.next()
		}
	case '<':
		kind = token.LT
		if next == '=' {
			kind = token.LEQ
			l.next()
		}
	case '>':
		kind = token.GT
		if next == '=' {
			kind = token.GEQ
			l.next()
		}
	case '&':
		kind = token.AND
		if next == '&' {
			kind = token.LOGIC_AND
			l.next()
		}
	case '|':
		kind = token.OR
		if next == '|' {
			kind = token.LOGIC_OR
			l.next()
		}
	case '+':
		kind = token.ADD
		if next == '+' {
			kind = token.INC
			l.next()
		}
	case '-':
		kind = token.SUB
		if next == '-' {
			kind = token.DEC
			l.next()
		}
	case '*':
		kind = token.MUL
	case '/':
		kind = token.DIV
	case '%':
		kind = token.MOD
	case '(':
		kind = token.LPAREN
	case '[':
		kind = token.LBRACKET
	case '{':
		kind = token.LBRACE
	case ')':
		kind = token.RPAREN
	case ']':
		kind = token.RBRACKET
	case '}':
		kind = token.RBRACE
	case ',':
		kind = token.COMMA
	case ':':
		kind = token.COLON
	case ';':
		kind = token.SEMICOLON
	case '.':
		kind = token.DOT
	case '?':
		kind = token.QUESTION
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

func (l *Lexer) readIdentifier() string {
	return l.readIf(func(ch rune) bool {
		return utils.IsIdentifier(ch)
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

func (l *Lexer) look() rune {
	return l.source[l.offset+1]
}

func (l *Lexer) current() rune {
	return l.source[l.offset]
}

func (l *Lexer) position() *token.Position {
	return token.NewPosition(l.offset, l.line, l.column)
}

func (l *Lexer) eof() bool {
	return l.offset >= l.end
}

func (l *Lexer) unexpect(ch rune, position *token.Position) {
	panic(token.NewSyntaxError(fmt.Sprintf("Unrecognized punctuation %s", string(ch)), position))
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
