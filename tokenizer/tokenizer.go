package tokenizer

import "unicode"

// Tokenize creates tokens from a source
func Tokenize(src *Source) []Token {
	var tokens []Token

	// Initial scan to get the first character
	src.Scan()

	for !src.IsEOF() {
		pos := src.Position()
		ch := src.Char()

		var tok Token

		switch {
		case unicode.IsSpace(ch):
			tok = Token{Type: SPACE, Value: " ", Pos: pos}
			src.Scan()

		case ch == '\\':
			// Handle commands
			src.Scan() // Skip backslash
			if src.IsEOF() {
				tok = Token{Type: ILLEGAL, Value: "\\", Pos: pos}
			} else if unicode.IsLetter(src.Char()) {
				// Letter command
				var cmd string
				for !src.IsEOF() && unicode.IsLetter(src.Char()) {
					cmd += string(src.Char())
					src.Scan()
				}
				tok = Token{Type: COMMAND, Value: cmd, Pos: pos}
			} else if !unicode.IsDigit(src.Char()) {
				// Symbol command
				tok = Token{Type: COMMAND, Value: string(src.Char()), Pos: pos}
				src.Scan()
			} else {
				// Invalid command
				tok = Token{Type: ILLEGAL, Value: "\\" + string(src.Char()), Pos: pos}
				src.Scan()
			}

		case unicode.IsDigit(ch):
			// Number
			var num string
			for !src.IsEOF() && unicode.IsDigit(src.Char()) {
				num += string(src.Char())
				src.Scan()
			}
			tok = Token{Type: NUMBER, Value: num, Pos: pos}

		case unicode.IsLetter(ch):
			tok = Token{Type: SYMBOL, Value: string(ch), Pos: pos}
			src.Scan()

		case ch == '^':
			tok = Token{Type: SUPERSCRIPT, Value: "^", Pos: pos}
			src.Scan()

		case ch == '_':
			tok = Token{Type: SUBSCRIPT, Value: "_", Pos: pos}
			src.Scan()

		case ch == '{':
			tok = Token{Type: LBRACE, Value: "{", Pos: pos}
			src.Scan()

		case ch == '}':
			tok = Token{Type: RBRACE, Value: "}", Pos: pos}
			src.Scan()

		case ch == '(' || ch == ')' || ch == '[' || ch == ']' || ch == '|':
			tok = Token{Type: DELIMITER, Value: string(ch), Pos: pos}
			src.Scan()

		case ch == '.':
			tok = Token{Type: PERIOD, Value: ".", Pos: pos}
			src.Scan()

		case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '=' || ch == '&':
			tok = Token{Type: OPERATOR, Value: string(ch), Pos: pos}
			src.Scan()

		default:
			tok = Token{Type: ILLEGAL, Value: string(ch), Pos: pos}
			src.Scan()
		}

		tokens = append(tokens, tok)
	}

	// Add EOF token
	tokens = append(tokens, Token{
		Type:  EOF,
		Value: "",
		Pos:   src.Position(),
	})

	return tokens
}
