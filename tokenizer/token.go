package tokenizer

// TokenType defines the type of a token form LaTeX input.
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
)

type Token struct {
	Type  TokenType // Type of the Token
	Value string    // Literal value from input
	Pos   int    // Byte offset in the input
}

// String returns a readable representation of the token.
func (t Token) String() string {
	return tokenTypeToString(t.Type) + "('" + t.Value + "')"
}

// tokenTypeToString converts TokenType to string for debugging or testing
func tokenTypeToString(tt TokenType) string {
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
	default:
		return "UNKNOWN"
	}
}
