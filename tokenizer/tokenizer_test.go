package tokenizer_test

import (
	"testing"

	"github.com/neox5/texmax/tokenizer"
)

func TestTokenize(t *testing.T) {
	input := `\frac{a^2 + b_1}{c}`

	expected := []tokenizer.Token{
		{Type: tokenizer.COMMAND, Value: "frac"},
		{Type: tokenizer.LBRACE, Value: "{"},
		{Type: tokenizer.SYMBOL, Value: "a"},
		{Type: tokenizer.SUPERSCRIPT, Value: "^"},
		{Type: tokenizer.NUMBER, Value: "2"},
		{Type: tokenizer.OPERATOR, Value: "+"},
		{Type: tokenizer.SYMBOL, Value: "b"},
		{Type: tokenizer.SUBSCRIPT, Value: "_"},
		{Type: tokenizer.NUMBER, Value: "1"},
		{Type: tokenizer.RBRACE, Value: "}"},
		{Type: tokenizer.LBRACE, Value: "{"},
		{Type: tokenizer.SYMBOL, Value: "c"},
		{Type: tokenizer.RBRACE, Value: "}"},
		{Type: tokenizer.EOF, Value: ""},
	}

	tokens := tokenizer.Tokenize(input)

	if len(tokens) != len(expected) {
		t.Fatalf("token count mismatch: got %d, want %d", len(tokens), len(expected))
	}

	for i, tok := range tokens {
		exp := expected[i]
		if tok.Type != exp.Type || tok.Value != exp.Value {
			t.Errorf("token %d mismatch: got %+v, want %+v", i, tok, exp)
		}
	}
}
