package tokenizer

import "fmt"

// TokenType defines the type of a token from LaTeX input.
type TokenType int

const (
	// Special
	ILLEGAL TokenType = iota
	EOF
	SPACE

	// Core token types
	COMMAND     // e.g., \frac, \alpha
	SYMBOL      // e.g., x, y, z
	NUMBER      // e.g., 123
	OPERATOR    // e.g., +, -, =, *, /
	SUPERSCRIPT // ^
	SUBSCRIPT   // _
	LBRACE      // {
	RBRACE      // }
	DELIMITER   // ( ), [ ], | etc. (visual math delimiters)
	PERIOD      // .
)

// Token represents a single lexical token.
type Token struct {
	Type  TokenType // Type of the token
	Value string    // Literal value from input
	Pos   Position  // Byte offset in the input
}

// String returns a readable representation of the token.
func (t Token) String() string {
	return fmt.Sprintf("%s('%s' at line %d, col %d)",
		t.Type.String(),
		t.Value,
		t.Pos.Line,
		t.Pos.Column)
}

// String implements the fmt.Stringer interface for TokenType.
func (tt TokenType) String() string {
	switch tt {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case SPACE:
		return "SPACE"
	case COMMAND:
		return "COMMAND"
	case SYMBOL:
		return "SYMBOL"
	case NUMBER:
		return "NUMBER"
	case OPERATOR:
		return "OPERATOR"
	case SUPERSCRIPT:
		return "SUPERSCRIPT"
	case SUBSCRIPT:
		return "SUBSCRIPT"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case DELIMITER:
		return "DELIMITER"
	case PERIOD:
		return "PERIOD"
	default:
		return "UNKNOWN"
	}
}
