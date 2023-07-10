package lexer

//The package should have a path relative to the module source name
import (
	"fmt"
	"lexer/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New returns a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
		fmt.Println(l.ch)
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	fmt.Println(l.ch)
	switch l.ch {
	case '=': // Assignment operator
		tok = newToken(token.ASSIGN, l.ch)
	case ';': // Semicolon
		tok = newToken(token.SEMICOLON, l.ch)
	case '(': // Left parenthesis
		tok = newToken(token.LPAREN, l.ch)
	case ')': // Right parenthesis
		tok = newToken(token.RPAREN, l.ch)
	case ',': // Comma
		tok = newToken(token.COMMA, l.ch)
	case '+': // Plus operator
		tok = newToken(token.PLUS, l.ch)
	case '{': // Left brace
		tok = newToken(token.LBRACE, l.ch)
	case '}': // Right brace
		tok = newToken(token.RBRACE, l.ch)
	case 0: // End of file
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
