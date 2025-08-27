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

	numStr := l.input[start:l.pos]
	value, err := strconv.Atoi(numStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid number: %s", numStr))
	}
	return Token{Type: TokenNumber, Value: value, Text: numStr}
}

type ASTNodeType int

const (
	NodeNumber ASTNodeType = iota
	NodeBinaryOp
)

type ASTNode struct {
	Type     ASTNodeType
	Value    int
	Operator TokenType
	Left     *ASTNode
	Right    *ASTNode
}

func (node *ASTNode) Evaluate() int {
	switch node.Type {
	case NodeNumber:
		return node.Value
	case NodeBinaryOp:
		left := node.Left.Evaluate()
		right := node.Right.Evaluate()

		switch node.Operator {
		case TokenPlus:
			return left + right
		case TokenMinus:
			return left - right
		case TokenMul:
			return left * right
		case TokenDiv:
			if right == 0 {
				panic("Division by zero")
			}
			return left / right
		default:
			panic(fmt.Sprintf("Unknown operator: %v", node.Operator))
		}
	default:
		panic(fmt.Sprintf("Unknown node type: %v", node.Type))
	}
}

type Parser struct {
	lexer        *Lexer
	currentToken Token
}

func NewParser(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.currentToken = lexer.NextToken()
	return parser
}

func (p *Parser) Parse() *ASTNode {
	return p.parseExpression()
}

// 3 * (5 * 2 + (4 - 1 + func1()))

// Expression := Literal | BinaryExpression | ParentExpression
// Literal := "" | Number
// BinaryExpression := Expression Op Expression
// ParentExpression := "(" Expression ")"

// Expression := Term [ Plus/Minus Term]*
// Term := Factor [ Mul/Div Factor]*
// Factor := Number | LeftParen Expression RightParen | Minus Factor
func (p *Parser) parseExpression() *ASTNode {
	node := p.parseTerm()

	for p.currentToken.Type == TokenPlus || p.currentToken.Type == TokenMinus {
		op := p.currentToken.Type
		p.consume(op)

		right := p.parseTerm()
		node = &ASTNode{
			Type:     NodeBinaryOp,
			Operator: op,
			Left:     node,
			Right:    right,
		}
	}
	return node
}

func (p *Parser) parseTerm() *ASTNode {
	node := p.parseFactor()

	for p.currentToken.Type == TokenMul || p.currentToken.Type == TokenDiv {
		op := p.currentToken.Type
		p.consume(op)

		right := p.parseFactor()
		node = &ASTNode{
			Type:     NodeBinaryOp,
			Operator: op,
			Left:     node,
			Right:    right,
		}
	}
	return node
}

func (p *Parser) parseFactor() *ASTNode {
	token := p.currentToken

	if token.Type == TokenNumber {
		p.consume(TokenNumber)
		return &ASTNode{
			Type:  NodeNumber,
			Value: token.Value,
		}
	} else if token.Type == TokenLeftParen {
		p.consume(TokenLeftParen)
		node := p.parseExpression()
		p.consume(TokenRightParen)
		return node
	} else if token.Type == TokenMinus {
		p.consume(TokenMinus)
		node := p.parseFactor()
		return &ASTNode{
			Type:     NodeBinaryOp,
			Operator: TokenMinus,
			Left:     &ASTNode{Type: NodeNumber, Value: 0},
			Right:    node,
		}
	}
	panic(fmt.Sprintf("Unexpected token: %v", token))
}

func (p *Parser) consume(expectedType TokenType) {
	if p.currentToken.Type == expectedType {
		p.currentToken = p.lexer.NextToken()
	} else {
		panic(fmt.Sprintf("Expected %v, got %v", expectedType, p.currentToken.Type))
	}
}
