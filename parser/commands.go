package parser

import (
	"fmt"

	"github.com/neox5/texmax/ast"
)

// parseCommand handles LaTeX commands (tokens that start with \)
func (p *Parser) parseCommand() ast.Node {
	token := p.next() // consume the COMMAND token
	commandName := token.Value
	startPos := token.Pos

	// Check command type in order of likelihood/specificity
	switch {
	case isNonArgumentFunction(commandName):
		return p.parseNonArgumentFunction(commandName, startPos)
	case isOperator(commandName):
		return p.parseOperator(commandName, startPos)
	case isGreekLetter(commandName):
		return p.parseGreekLetter(commandName, startPos)
	}

	// Handle specific command types with arguments
	switch commandName {
	case "frac":
		return p.parseFractionCommand(startPos)
	case "sqrt":
		return p.parseSqrtCommand(startPos)
	case "binom":
		return p.parseBinomCommand(startPos)
	default:
		p.addError(fmt.Sprintf("unsupported command: \\%s", commandName), startPos)
		return nil
	}
}
