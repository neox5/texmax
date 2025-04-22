package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

// parseDelimitedExpression parses expressions of the form \left( ... \right)
func (p *Parser) parseDelimitedExpression(startPos int) ast.Node {
	// Parse the left delimiter directly
	leftDelimiter := p.parseDelimiter()
	if leftDelimiter == nil {
		p.addError("expected delimiter after \\left", p.peek().Pos)
		return nil
	}

	// Parse the content between delimiters, stopping at \right
	content := p.parseExpression(func(t tokenizer.Token) bool {
		return t.Type == tokenizer.COMMAND && t.Value == "right"
	})

	// Ensure we have a \right command
	if p.peek().Type != tokenizer.COMMAND || p.peek().Value != "right" {
		p.addError("expected \\right to close \\left", p.peek().Pos)
		// Return just the content as we couldn't complete the delimited expression
		return content
	}
	p.next() // consume \right

	// Parse the right delimiter directly
	rightDelimiter := p.parseDelimiter()
	if rightDelimiter == nil {
		p.addError("expected delimiter after \\right", p.peek().Pos)
		// Return just the content as we couldn't complete the delimited expression
		return content
	}

	// Only create a DelimitedExpressionNode if all parts were successfully parsed
	return &ast.DelimitedExpressionNode{
		Start:          startPos,
		LeftDelimiter:  leftDelimiter,
		Content:        content,
		RightDelimiter: rightDelimiter,
	}
}
