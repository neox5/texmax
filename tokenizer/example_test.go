package tokenizer_test

import (
	"fmt"

	"github.com/neox5/texmax/tokenizer"
)

func ExampleTokenize() {
	input := `\frac{a^2}{b}`

	tokens := tokenizer.Tokenize(input)
	for _, tok := range tokens {
		fmt.Printf("%s: %q\n", tok.Type, tok.Value)
	}

	// Output:
	// COMMAND: "frac"
	// LBRACE: "{"
	// SYMBOL: "a"
	// SUPERSCRIPT: "^"
	// NUMBER: "2"
	// RBRACE: "}"
	// LBRACE: "{"
	// SYMBOL: "b"
	// RBRACE: "}"
	// EOF: ""
}
