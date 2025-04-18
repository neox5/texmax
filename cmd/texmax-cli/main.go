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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] 'latex_expression'\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s '\\frac{a^2}{b}'\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	tokensOnly := flag.Bool("tokens", false, "Only show tokenization results")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Get input from command line arguments
	input := strings.Join(flag.Args(), " ")
	fmt.Printf("Input: %s\n\n", input)

	// Tokenize
	tokens := tokenizer.Tokenize(input)

	fmt.Println("Tokens:")
	for i, tok := range tokens {
		if tok.Type == tokenizer.EOF {
			fmt.Printf("%d: %s at position %d\n", i, tok.Type, tok.Pos)
		} else {
			fmt.Printf("%d: %s %q at position %d\n", i, tok.Type, tok.Value, tok.Pos)
		}
	}

	// If only tokens are requested, exit here
	if *tokensOnly {
		return
	}

	// Parse
	p := parser.New(tokens)
	root, errors := p.Parse()

	// Print errors if any
	if len(errors) > 0 {
		fmt.Println("\nParser errors:")
		for i, err := range errors {
			fmt.Printf("%d: %s\n", i, err)
		}
	}

	// Print AST using the GoLikePrinter
	fmt.Println("\nAST Structure:")
	visitor := ast.NewGoLikePrinter(os.Stdout)
	ast.Walk(visitor, root)
}
