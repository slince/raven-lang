package token

import (
	"fmt"
	"strings"
)

type Stream struct {
	index  int
	Tokens []*Token
}

func (s *Stream) Add(token *Token) {
	s.Tokens = append(s.Tokens, token)
}

func (s *Stream) Next() *Token {
	token := s.Tokens[s.index]
	s.index++
	return token
}

func (s *Stream) Look(number int) *Token {
	return s.Tokens[s.index+number]
}

func (s *Stream) Current() *Token {
	return s.Tokens[s.index]
}

func (s *Stream) Count() int {
	return len(s.Tokens)
}

func (s *Stream) Eof() bool {
	return s.Current().Test(EOF)
}

func (s *Stream) Test(kinds ...Kind) bool {
	return s.Current().Test(kinds...)
}

func (s *Stream) Expect(kinds ...Kind) *Token {
	tok, err := s.expect(kinds...)
	if err != nil {
		panic(err)
	}
	return tok
}

func (s *Stream) SkipIfTest(kinds ...Kind) *Token {
	tok := s.Current()
	if tok.Test(kinds...) {
		s.Next()
		return tok
	}
	return nil
}

func (s *Stream) expect(kinds ...Kind) (tok *Token, err error) {
	tok = s.Current()
	if !tok.Test(kinds...) {
		var expected = make([]string, 0)
		for _, item := range kinds {
			expected = append(expected, Literal(item))
		}
		var msg = fmt.Sprintf("Unexpected token \"%s\" (expected \"%s\")", tok.Literal, strings.Join(expected, ","))
		err = NewSyntaxError(msg, tok.Position)
		return
	}
	s.Next()
	return
}

func (s *Stream) Dump() []string {
	s.index = 0
	var toks = make([]string, len(s.Tokens))
	for _, tok := range s.Tokens {
		toks = append(toks, tok.Literal)
	}
	return toks
}

func NewStream() *Stream {
	return &Stream{
		Tokens: make([]*Token, 0),
		index:  0,
	}
}
