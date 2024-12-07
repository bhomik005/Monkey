package lexer

import "monkey/token"

/**
 * Potential improvements in lexer
 * It support the ASCII characters which are 8 bits wide
 * To support the Unicode and UTF-8 characters, we would need to change teh ch from byte (8 bits) to rune (32 bits wide)
 */
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination -
}

/**
 * creates and returns a pointer to the lexer object
 */
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/**
 * reads the next input character and increment the readPosition pointer
 */
func (l *Lexer) readChar() {
	// checks if the readPosition pointer has reached the end of file
	if l.readPosition >= len(l.input) {
		l.ch = 0
		// if it is not - l.ch stores the value of the next char
	} else {
		l.ch = l.input[l.readPosition]
	}
	// make l.position pointing to the read position
	l.position = l.readPosition
	// increments the l.readPosition (pointing to the next char after the current char)
	l.readPosition += 1
}

/**
 * it looks the current char and returns the appropriate token
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

/**
 * returns the token struct by passing the token type and the current char
 */
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
