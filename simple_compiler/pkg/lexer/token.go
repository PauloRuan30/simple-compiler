package lexer

import "fmt"

// TokenType representa o tipo de um token
type TokenType string

// Token representa um token léxico
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Implementa String para facilitar impressão
func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Literal: %s, Line: %d, Column: %d}", 
		t.Type, t.Literal, t.Line, t.Column)
}

// Definição de tipos de tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identificadores
	IDENTIFIER = "IDENTIFIER"
	
	// Literais
	INTEGER = "INTEGER"
	FLOAT   = "FLOAT"

	// Operadores
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	// Delimitadores
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Palavras-chave
	LET    = "LET"
	RETURN = "RETURN"
)

// Palavras reservadas
var Keywords = map[string]TokenType{
	"let":    LET,
	"return": RETURN,
}

// Verifica se é uma palavra-chave
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := Keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}