package parser

import (
	"fmt"

	"github.com/neox5/texmax/ast"
)

func (p *Parser) parseCommand() ast.Node {
	token := p.next() // consume the COMMAND token
	commandName := token.Value
	startPos := token.Pos

	switch commandName {
	case "frac":
		return p.parseFrac(startPos)
	default:
		p.addError(fmt.Sprintf("unsupported command: \\%s", commandName), startPos)
		return nil
	}
}

func (p *Parser) parseFrac(startPos int) ast.Node {
	// A fraction requires two arguments: numerator and denominator
	numerator := p.parseGroupedStrict()
	if numerator == nil {
		p.addError("expected numerator after \\frac", startPos)
		return nil
	}

	denominator := p.parseGroupedStrict()
	if denominator == nil {
		p.addError("expected denominator after \\frac{...}", startPos)
		return nil
	}

	return &ast.FractionNode{
		Start:       startPos,
		Numerator:   numerator,
		Denominator: denominator,
	}
}
