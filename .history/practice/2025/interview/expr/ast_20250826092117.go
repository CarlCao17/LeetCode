package interview

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenPlus
	TokenMinus
	TokenMul
	TokenDiv
	TokenLeftParen
	TokenRightParen
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value int
	Text  string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func (l *Lexer) NextToken() Token {
	for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
		l.pos++
	}

	if l.pos >= len(l.input) {
		return Token{Type: TokenEOF}
	}

	ch := l.input[l.pos]

	if unicode.IsDigit(rune(ch)) {
		return l.readNumber()
	}

	l.pos++
	switch ch {
	case '+':
		return Token{Type: TokenPlus, Text: "+"}
	case '-':
		return Token{Type: TokenMinus, Text: "-"}
	case '*':
		return Token{Type: TokenMul, Text: "*"}
	case '/':
		return Token{Type: TokenDiv, Text: "/"}
	case '(':
		return Token{Type: TokenLeftParen, Text: "("}
	case ')':
		return Token{Type: TokenRightParen, Text: ")"}
	default:
		panic(fmt.Sprintf("Unexpected character: %c", ch))
	}
}

func (l *Lexer) readNumber() Token {
	start := l.pos
	for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
		l.pos++
	}

	numStr := l.inpu[start:l.pos]
	value, err := strconv.Atoi(numStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid number: %s", numStr))
	}
}
