package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/parser"
	"github.com/neox5/texmax/tokenizer"
)

func main() {
	// Define short flags only
	var (
		fileInput  string
		tokensOnly bool
	)

	flag.StringVar(&fileInput, "f", "", "Input file containing LaTeX expressions")
	flag.BoolVar(&tokensOnly, "t", false, "Only show tokenization results")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] 'latex_expression'\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintf(os.Stderr, "  %s '\\frac{a^2}{b}'\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -f input.tex\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
	flag.Parse()

	// Setup input source
	var src *tokenizer.Source
	var err error

	if fileInput != "" {
		src, err = tokenizer.NewFileSource(fileInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file '%s': %v\n", fileInput, err)
			os.Exit(1)
		}
		fmt.Printf("Reading from file: %s\n\n", fileInput)
	} else {
		if flag.NArg() < 1 {
			flag.Usage()
			os.Exit(1)
		}
		input := strings.Join(flag.Args(), " ")
		fmt.Printf("Input: %s\n\n", input)
		src = tokenizer.NewStringSource(input)
	}

	// Tokenize input
	tokens := tokenizer.Tokenize(src)

	fmt.Println("Tokens:")
	for i, tok := range tokens {
		if tok.Type == tokenizer.EOF {
			fmt.Printf("%d: %s at position %d\n", i, tok.Type, tok.Pos.Offset)
		} else {
			fmt.Printf("%d: %s %q at position %d\n", i, tok.Type, tok.Value, tok.Pos.Offset)
		}
	}

	// Exit early if only tokenization was requested
	if tokensOnly {
		return
	}

	// Parse tokens
	p := parser.New(tokens)
	root, errors := p.Parse()

	// Handle parse errors
	if len(errors) > 0 {
		fmt.Println("\nParser errors:")
		for i, err := range errors {
			fmt.Printf("%d: %s\n", i, err)
		}
	}

	// Print the AST
	fmt.Println("\nAST Structure:")
	visitor := ast.NewPrintVisitor(os.Stdout)
	ast.Walk(visitor, root)
}
