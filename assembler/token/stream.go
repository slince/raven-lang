package token

import (
	"fmt"
	"strconv"
	"strings"
)

type Stream struct {
	index  int
	Tokens []*Token
}

func (s *Stream) Add(token ...*Token) {
	s.Tokens = append(s.Tokens, token...)
}

func (s *Stream) Eof() bool {
	return s.Current().Test(EOF)
}

func (s *Stream) Next() *Token {
	token := s.Tokens[s.index]
	s.index++
	return token
}

func (s *Stream) Look(number int) *Token {
	return s.Tokens[s.index+number]
}

func (s *Stream) Peek() *Token {
	return s.Look(1)
}

func (s *Stream) Current() *Token {
	return s.Tokens[s.index]
}

func (s *Stream) Count() int {
	return len(s.Tokens)
}

func (s *Stream) ExpectDirective() *Token {
	return s.Expect(directiveTypes...)
}

func (s *Stream) ExpectInstruction() *Token {
	return s.Expect(instructionTypes...)
}

func (s *Stream) Expect(kind ...Type) *Token {
	tok, err := s.expect(kind...)
	if err != nil {
		panic(err)
	}
	return tok
}

func (s *Stream) SkipIfTest(kind ...Type) *Token {
	tok := s.Current()
	if tok.Test(kind...) {
		s.Next()
		return tok
	}
	return nil
}

func (s *Stream) expect(kind ...Type) (tok *Token, err error) {
	tok = s.Current()
	if !tok.Test(kind...) {
		var values = make([]string, len(kind), cap(kind))
		var types = make([]string, len(kind), cap(kind))
		for _, item := range kind {
			types = append(types, strconv.Itoa(int(item)))
			values = append(values, ValueOf(item))
		}
		var msg = fmt.Sprintf("Unexpected token \"%d\" of value \"%s\" (\"%s\" expected with value \"%s\")", tok.Type, tok.Value, strings.Join(types, ","), strings.Join(values, ","))
		err = NewSyntaxError(msg, tok.Position)
		return
	}
	s.Next()
	return
}

func NewStream() *Stream {
	return &Stream{
		Tokens: []*Token{},
		index:  0,
	}
}

func (s *Stream) Dump() []string {
	s.index = 0
	var toks = make([]string, len(s.Tokens))
	for _, tok := range s.Tokens {
		toks = append(toks, tok.Value)
	}
	return toks
}
