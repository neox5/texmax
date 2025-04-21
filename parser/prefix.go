package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

func (p *Parser) parseSymbol() ast.Node {
	t := p.next()
	return &ast.SymbolNode{Start: t.Pos, Value: t.Value}
}

func (p *Parser) parseNumber() ast.Node {
	t := p.next()
	return &ast.NumberNode{Start: t.Pos, Value: t.Value}
}

func (p *Parser) parseSpace() ast.Node {
	t := p.next()
	return &ast.SpaceNode{Start: t.Pos, Value: t.Value}
}

func (p *Parser) parseOperatorSymbol() ast.Node {
	t := p.next()
	return &ast.OperatorNode{Start: t.Pos, Value: t.Value}
}

// parseDelimiter parses a delimiter token or delimiter command
func (p *Parser) parseDelimiter() ast.Node {
	token := p.peek()
	startPos := token.Pos

	switch token.Type {
	case tokenizer.DELIMITER:
		// Regular delimiter like (, ), [, ], |
		p.next() // consume the delimiter token
		return &ast.DelimiterNode{
			Start: startPos,
			Value: token.Value,
		}

	case tokenizer.COMMAND:
		// Command delimiter like \{, \}, \langle, \rangle, etc.
		p.next() // consume the command token
		value := token.Value

		// Store delimiter commands using standardized names
		switch value {
		// Empty delimiter
		case ".":
			return &ast.DelimiterNode{Start: startPos, Value: "."}

		// Braces
		case "{":
			return &ast.DelimiterNode{Start: startPos, Value: "{"}
		case "}":
			return &ast.DelimiterNode{Start: startPos, Value: "}"}

		// Vertical bar
		case "lvert", "rvert":
			return &ast.DelimiterNode{Start: startPos, Value: "|"}
		// Norm (double vertical bar)
		case "|", "lVert", "rVert":
			return &ast.DelimiterNode{Start: startPos, Value: "||"}

		// Angle brackets
		case "langle":
			return &ast.DelimiterNode{Start: startPos, Value: "langle"}
		case "rangle":
			return &ast.DelimiterNode{Start: startPos, Value: "rangle"}

		// Floor
		case "lfloor":
			return &ast.DelimiterNode{Start: startPos, Value: "lfloor"}
		case "rfloor":
			return &ast.DelimiterNode{Start: startPos, Value: "rfloor"}

		// Ceiling
		case "lceil":
			return &ast.DelimiterNode{Start: startPos, Value: "lceil"}
		case "rceil":
			return &ast.DelimiterNode{Start: startPos, Value: "rceil"}

		// Backslash
		case "backslash":
			return &ast.DelimiterNode{Start: startPos, Value: "backslash"}
		}
	}

	// If we reach here, it wasn't a valid delimiter
	return nil
}
