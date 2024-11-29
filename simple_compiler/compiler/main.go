package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"simple-compiler/pkg/lexer"
	"simple-compiler/pkg/parser"
)

func main() {
	// Lê o arquivo de exemplo
	sourceCode, err := ioutil.ReadFile("../../examples/sample_program.txt")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo: %v", err)
	}

	// Cria o lexer
	l := lexer.New(string(sourceCode))

	// Tokeniza o código fonte
	tokens, err := lexer.Tokenize(string(sourceCode))
	if err != nil {
		log.Fatalf("Erro na tokenização: %v", err)
	}

	// Imprime tokens
	fmt.Println("--- Tokens ---")
	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}

	// Cria o parser
	p := parser.New(l)

	// Parseia o programa
	program := p.ParseProgram()

	// Verifica erros no parsing
	if len(p.Errors()) > 0 {
		fmt.Println("Erros de parsing:")
		for _, msg := range p.Errors() {
			fmt.Println(msg)
		}
		return
	}

	// Imprime a AST
	fmt.Println("\n--- Árvore Sintática Abstrata (AST) ---")
	fmt.Println(program.String())
}
