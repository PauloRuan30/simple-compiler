package parser

import (
	"fmt"
	"strconv"

	"github.com/yourusername/simple-compiler/pkg/lexer"
)

type Parser struct {
	l *lexer.Lexer
	
	currentToken lexer.Token
	peekToken    lexer.Token
	
	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	
	// Ler dois tokens para configurar currentToken e peekToken
	p.nextToken()
	p.nextToken()
	
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}
	
	for p.currentToken.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	
	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case lexer.LET:
		return p.parseLetStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *LetStatement {
	stmt := &LetStatement{Token: p.currentToken}
	
	// Espera um identificador após LET
	if !p.expectPeek(lexer.IDENTIFIER) {
		return nil
	}
	
	stmt.Name = &Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
	
	// Espera uma atribuição