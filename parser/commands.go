package parser

import (
	"fmt"

	"github.com/neox5/texmax/ast"
)

// parseCommand handles LaTeX commands (tokens that start with \)
func (p *Parser) parseCommand() ast.Node {
	token := p.next() // consume the COMMAND token
	cmd := token.Value
	pos := token.Pos

	// Check command type in order of likelihood/specificity
	switch {
	case isNonArgumentFunction(cmd):
		return p.parseNonArgumentFunction(cmd, pos)
	case isOperator(cmd):
		return p.parseOperator(cmd, pos)
	case isGreekLetter(cmd):
		return p.parseGreekLetter(cmd, pos)
	}

	// Handle specific command types with arguments
	switch cmd {
	case "frac":
		return p.parseFractionCommand(pos)
	case "sqrt":
		return p.parseSqrtCommand(pos)
	case "binom":
		return p.parseBinomCommand(pos)
	case "left":
		return p.parseDelimitedExpression(pos)
	case "right":
		// Handle \right outside of a \left...\right context
		p.addError("unexpected \\right without matching \\left", pos)
		return nil
	default:
		p.addError(fmt.Sprintf("unsupported command: \\%s", cmd), pos)
		return nil
	}
}
