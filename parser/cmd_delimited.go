package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

// parseDelimitedExpression parses expressions of the form \left( ... \right)
func (p *Parser) parseDelimitedExpression(pos tokenizer.Position) ast.Node {
	left := p.parseDelimiter()
	if left == nil {
		p.addError("expected delimiter after \\left", p.peek().Pos)
		return nil
	}

	// Parse the content between delimiters, stopping at \right
	expr := p.parseExpression(func(t tokenizer.Token) bool {
		return t.Type == tokenizer.COMMAND && t.Value == "right"
	})

	// Ensure we have a \right command
	if p.peek().Type != tokenizer.COMMAND || p.peek().Value != "right" {
		p.addError("expected \\right to close \\left", p.peek().Pos)

		return expr
	}
	p.next() // consume \right

	right := p.parseDelimiter()
	if right == nil {
		p.addError("expected delimiter after \\right", p.peek().Pos)

		return expr
	}

	return &ast.DelimitedExpressionNode{
		Start:          pos,
		LeftDelimiter:  left,
		Content:        expr,
		RightDelimiter: right,
	}
}
