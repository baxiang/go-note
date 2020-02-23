package main

import (
	"errors"
	"fmt"
)

func main() {
	source := `
int s = 123
age >= 45
intA = 67
1+2*3/6
`
	lex := NewLexer([]rune(source))
	toks, err := lex.lex()
	if err != nil {
		panic(err)
	}
	for _, tok := range toks {
		fmt.Printf("%v\n", tok)
	}
	/*
		output:

		~~~~~~~~~~~~~
		token: {INT: 'int'}
		token: {Identifier: 's'}
		token: {EQ: '='}
		token: {IntLiteral: '123'}
		~~~~~~~~~~~~~
		token: {Identifier: 'age'}
		token: {GE: '>='}
		token: {IntLiteral: '45'}
		~~~~~~~~~~~~~
		token: {Identifier: 'intA'}
		token: {EQ: '='}
		token: {IntLiteral: '67'}
		~~~~~~~~~~~~~
		token: {IntLiteral: '1'}
		token: {Plus: '+'}
		token: {IntLiteral: '2'}
		token: {Star: '*'}
		token: {IntLiteral: '3'}
		token: {Slash: '/'}
		token: {IntLiteral: '6'}
		~~~~~~~~~~~~~
		token: {EOF: ''}
	*/
}

type TokenType int

const (
	EOF TokenType = iota
	INT
	DOUBLE
	STRING
	BOOL
	NIL
	NewLine

	Identifier

	GT
	GE
	EQ
	Assign
	Plus
	PlusPlus
	PlusEQ
	Minus
	MinusMinus
	MinusEQ
	Star
	StarEQ
	Slash
	SlashEQ

	IntLiteral
)

var tokenTypeStringMap = map[TokenType]string{
	EOF:     "EOF",
	INT:     "INT",
	DOUBLE:  "DOUBLE",
	STRING:  "STRING",
	BOOL:    "BOOL",
	NIL:     "NIL",
	NewLine: "~~~~~~~~~~~~~",

	Identifier: "Identifier",

	GT:         "GT",
	GE:         "GE",
	EQ:         "EQ",
	Assign:     "Assign",
	Plus:       "Plus",
	PlusPlus:   "PlusPlus",
	PlusEQ:     "PlusEQ",
	Minus:      "Minus",
	MinusMinus: "MinusMinus",
	MinusEQ:    "MinusEQ",
	Star:       "Star",
	StarEQ:     "StarEQ",
	Slash:      "Slash",
	SlashEQ:    "SlashEQ",

	IntLiteral: "IntLiteral",
}

var keywords = map[string]TokenType{
	"int":    INT,
	"double": DOUBLE,
	"string": STRING,
	"bool":   BOOL,
	"nil":    NIL,
}

func (t TokenType) String() string {
	s, ok := tokenTypeStringMap[t]
	if ok {
		return s
	}
	return "Unknown TokenType"
}

var (
	ErrorEOF = errors.New("reached EOF")
)

type Token struct {
	typ TokenType
	buf []rune
}

func (t *Token) push(r rune) {
	t.buf = append(t.buf, r)
}

func (t *Token) isEOF() bool {
	return t.typ == EOF
}

func (t *Token) String() string {
	if t.typ == NewLine {
		return fmt.Sprintf("%v", t.typ)
	}
	return fmt.Sprintf("token: {%v: '%s'}", t.typ, string(t.buf))
}

type Lexer struct {
	buf []rune
	idx int // always point to the next rune
}

func NewLexer(buf []rune) *Lexer {
	return &Lexer{
		buf: buf,
		idx: 0,
	}
}

func (l *Lexer) lex() ([]*Token, error) {
	toks := make([]*Token, 0)
	for {
		tok, err := l.lexToken()
		if err != nil {
			return nil, err
		}
		toks = append(toks, tok)
		if tok.isEOF() {
			break
		}
	}
	return toks, nil
}

func (l *Lexer) lexToken() (*Token, error) {
	l.skipWhitespace()
	ch, err := l.next()
	if err != nil {
		return &Token{
			typ: EOF,
		}, nil
	}
	switch {
	case isAlpha(ch):
		return l.lexIdentifier(ch), nil
	case ch == '>':
		return l.lexGT(ch), nil
	case ch == '=':
		return l.lexEQ(ch), nil
	case ch == '+':
		return l.lexPlus(ch), nil
	case ch == '-':
		return l.lexMinus(ch), nil
	case ch == '*':
		return l.lexStar(ch), nil
	case ch == '/':
		return l.lexSlash(ch), nil
	case ch == '\n':
		return &Token{
			typ: NewLine,
		}, nil
	case isDigit(ch):
		return l.lexDigital(ch), nil
	default:
		return nil, errors.New("not supported rune: " + string(ch))
	}
}

func (l *Lexer) lexDigital(ch rune) *Token {
	tok := &Token{
		typ: IntLiteral,
		buf: []rune{ch},
	}
	for {
		ch, err := l.seek()
		if err != nil {
			return tok
		}
		if isDigit(ch) {
			l.mustNext()
			tok.push(ch)
			continue
		}
		return tok
	}
}

func (l *Lexer) lexIdentifier(ch rune) *Token {
	tok := &Token{
		typ: Identifier,
		buf: []rune{ch},
	}
	for {
		ch, err := l.seek()
		if err != nil {
			return tok
		}
		if isAlpha(ch) || isDigit(ch) {
			l.mustNext()
			tok.push(ch)
			continue
		}
		if typ, ok := keywords[string(tok.buf)]; ok {
			tok.typ = typ
		}
		return tok
	}
}

func (l *Lexer) lexEQ(ch rune) *Token {
	tok := &Token{
		typ: EQ,
		buf: []rune{ch},
	}
	ch, err := l.seek()
	if err != nil {
		return tok
	}
	if ch == '=' {
		tok.typ = Assign
		tok.buf = append(tok.buf, ch)
		l.mustNext()
	}
	return tok
}

func (l *Lexer) lexGT(ch rune) *Token {
	tok := &Token{
		typ: GT,
		buf: []rune{ch},
	}
	ch, err := l.seek()
	if err != nil {
		return tok
	}
	if ch == '=' {
		tok.typ = GE
		tok.buf = append(tok.buf, ch)
		l.mustNext()
	}
	return tok
}

func (l *Lexer) lexPlus(ch rune) *Token {
	tok := &Token{
		typ: Plus,
		buf: []rune{ch},
	}
	ch, err := l.seek()
	if err != nil {
		return tok
	}
	switch ch {
	case '+':
		tok.typ = PlusPlus
		tok.buf = append(tok.buf, ch)
		l.mustNext()
	case '=':
		tok.typ = PlusEQ
		tok.buf = append(tok.buf, ch)
		l.mustNext()
	}
	return tok
}

func (l *Lexer) lexMinus(ch rune) *Token {
	tok := &Token{
		typ: Minus,
		buf: []rune{ch},
	}
	ch, err := l.seek()
	if err != nil {
		return tok
	}
	switch ch {
	case '-':
		tok.typ = MinusMinus
		tok.buf = append(tok.buf, ch)
		l.mustNext()
	case '=':
		tok.typ = MinusEQ
		tok.buf = append(tok.buf, ch)
		l.mustNext()
	}
	return tok
}

func (l *Lexer) lexStar(ch rune) *Token {
	tok := &Token{
		typ: Star,
		buf: []rune{ch},
	}
	ch, err := l.seek()
	if err != nil {
		return tok
	}
	if ch == '=' {
		tok.typ = StarEQ
		tok.buf = append(tok.buf, ch)
		l.mustNext()
	}
	return tok
}

func (l *Lexer) lexSlash(ch rune) *Token {
	tok := &Token{
		typ: Slash,
		buf: []rune{ch},
	}
	ch, err := l.seek()
	if err != nil {
		return tok
	}
	if ch == '=' {
		tok.typ = SlashEQ
		tok.buf = append(tok.buf, ch)
		l.mustNext()
	}
	return tok
}
func (l *Lexer) skipWhitespace() {
	for {
		ch, err := l.seek()
		if err != nil {
			return
		}
		if isWhitespace(ch) {
			l.mustNext()
			continue
		}
		break
	}
}

func (l *Lexer) next() (rune, error) {
	ch, err := l.seek()
	if err != nil {
		return 0, err
	}
	l.idx++
	return ch, nil
}

func (l *Lexer) seek() (rune, error) {
	if l.idx > len(l.buf)-1 {
		return 0, ErrorEOF
	}
	return l.buf[l.idx], nil
}

func (l *Lexer) mustNext() rune {
	l.idx++
	return l.buf[l.idx-1]
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

