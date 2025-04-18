package tokenizer

import "unicode"

func Tokenize(input string) []Token {
	var tokens []Token
	var pos int
	runes := []rune(input)

	for i := 0; i < len(runes); {
		r := runes[i]

		// Track current position (as bytes, not runes)
		start := pos
		charLen := len(string(r)) // byte length of the current rune

		switch {
		// SPACE
		case unicode.IsSpace(r):
			tokens = append(tokens, Token{Type: SPACE, Value: " ", Pos: start})
			i++
			pos += charLen

		// COMMAND
		case r == '\\':
			startIdx := i + 1
			endIdx := startIdx

			for endIdx < len(runes) && unicode.IsLetter(runes[endIdx]) {
				endIdx++
			}
			cmd := string(runes[startIdx:endIdx])
			tokens = append(tokens, Token{Type: COMMAND, Value: cmd, Pos: start})
			i = endIdx
			pos += len(`\` + cmd)

		// NUMBER
		case unicode.IsDigit(r):
			startIdx := i
			endIdx := i

			for endIdx < len(runes) && unicode.IsDigit(runes[endIdx]) {
				endIdx++
			}
			number := string(runes[startIdx:endIdx])
			tokens = append(tokens, Token{Type: NUMBER, Value: number, Pos: start})
			i = endIdx
			pos += len(number)

		// SYMBOL
		case unicode.IsLetter(r):
			tokens = append(tokens, Token{Type: SYMBOL, Value: string(r), Pos: start})
			i++
			pos += charLen

		// OPERATOR
		case r == '+' || r == '-' || r == '*' || r == '/' || r == '=':
			tokens = append(tokens, Token{Type: OPERATOR, Value: string(r), Pos: start})
			i++
			pos += charLen

		// SUPERSCRIPT
		case r == '^':
			tokens = append(tokens, Token{Type: SUPERSCRIPT, Value: "^", Pos: start})
			i++
			pos += charLen

		// SUBSCRIPT
		case r == '_':
			tokens = append(tokens, Token{Type: SUBSCRIPT, Value: "_", Pos: start})
			i++
			pos += charLen

		// LBRACE
		case r == '{':
			tokens = append(tokens, Token{Type: LBRACE, Value: "{", Pos: start})
			i++
			pos += charLen

		// RBRACE
		case r == '}':
			tokens = append(tokens, Token{Type: RBRACE, Value: "}", Pos: start})
			i++
			pos += charLen

		// DELIMITERS: ( ) [ ] |
		case r == '(' || r == ')' || r == '[' || r == ']' || r == '|':
			tokens = append(tokens, Token{Type: DELIMITER, Value: string(r), Pos: start})
			i++
			pos += charLen

		// ILLEGAL
		default:
			tokens = append(tokens, Token{Type: ILLEGAL, Value: string(r), Pos: start})
			i++
			pos += charLen
		}
	}

	// Add EOF token at end
	tokens = append(tokens, Token{Type: EOF, Value: "", Pos: pos})
	return tokens
}
