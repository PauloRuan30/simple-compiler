package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

// Lexer representa o analisador léxico
type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
	line         int
	column       int
}

// New cria um novo lexer
func New(input string) *Lexer {
	l := &Lexer{
		input:    []rune(input),
		line:     1,
		column:   1,
	}
	l.readChar()
	return l
}

// Métodos internos do lexer...
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	
	// Acompanha linha e coluna
	if l.ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

// NextToken retorna o próximo token
func (l *Lexer) NextToken() Token {
	var tok Token
	
	l.skipWhitespace()
	
	tok.Line = l.line
	tok.Column = l.column
	
	switch l.ch {
	case '=':
		tok = Token{Type: ASSIGN, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '+':
		tok = Token{Type: PLUS, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '-':
		tok = Token{Type: MINUS, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '*':
		tok = Token{Type: MULTIPLY, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '/':
		tok = Token{Type: DIVIDE, Literal: string(l.ch), Line: l.line, Column: l.column}
	case ';':
		tok = Token{Type: SEMICOLON, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '(':
		tok = Token{Type: LPAREN, Literal: string(l.ch), Line: l.line, Column: l.column}
	case ')':
		tok = Token{Type: RPAREN, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '{':
		tok = Token{Type: LBRACE, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '}':
		tok = Token{Type: RBRACE, Literal: string(l.ch), Line: l.line, Column: l.column}
	case 0:
		tok.Type = EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdentifier(tok.Literal)
			tok.Line = l.line
			tok.Column = l.column
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = INTEGER
			tok.Line = l.line
			tok.Column = l.column
			return tok
		} else {
			tok = Token{
				Type:    ILLEGAL,
				Literal: string(l.ch),
				Line:    l.line,
				Column:  l.column,
			}
		}
	}
	
	l.readChar()
	return tok
}

// Funções auxiliares
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

// Tokenize converte todo input em uma fatia de tokens
func Tokenize(input string) ([]Token, error) {
	l := New(input)
	var tokens []Token
	
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		
		if tok.Type == EOF {
			break
		}
		
		if tok.Type == ILLEGAL {
			return nil, fmt.Errorf("token ilegal na linha %d, coluna %d: %s", 
				tok.Line, tok.Column, tok.Literal)
		}
	}
	
	return tokens, nil
}