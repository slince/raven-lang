package utils

import (
	"unicode"
	"unicode/utf8"
)

func IsDecimal(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func IsDigit(ch rune) bool {
	return IsDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func IsLetter(ch rune) bool {
	return 'A' <= ch && ch <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func IsIdentifier(ch rune) bool {
	return IsDigit(ch) || IsLetter(ch)
}

// IsWhitespace is white space
func IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
