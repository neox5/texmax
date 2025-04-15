package parser

import (
	"fmt"

	"github.com/neox5/texmax/ast"
)

// List of recognized mathematical functions that don't take explicit arguments
var mathFunctions = map[string]bool{
	"sin":    true,
	"cos":    true,
	"tan":    true,
	"cot":    true,
	"sec":    true,
	"csc":    true,
	"log":    true,
	"ln":     true,
	"exp":    true,
	"arcsin": true,
	"arccos": true,
	"arctan": true,
	"sinh":   true,
	"cosh":   true,
	"tanh":   true,
	"lim":    true,
	"max":    true,
	"min":    true,
	"det":    true,
	"arg":    true,
	"mod":    true,
}

func (p *Parser) parseCommand() ast.Node {
	token := p.next() // consume the COMMAND token
	commandName := token.Value
	startPos := token.Pos

	// Check if this is a known math function
	if mathFunctions[commandName] {
		return &ast.NonArgumentFunctionNode{
			Start: startPos,
			Name:  commandName,
		}
	}

	switch commandName {
	case "frac":
		return p.parseFrac(startPos)
	case "int":
		return p.parseInt(startPos)
	case "sqrt":
		return p.parseSqrt(startPos)
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

func (p *Parser) parseInt(startPos int) ast.Node {
	// Parse optional limits (subscript and superscript)
	lowerLimit, upperLimit := p.parseLimits()

	// Create an IntegralNode with just the limits
	// The integrand will follow naturally in the parent expression
	return &ast.IntegralNode{
		Start:      startPos,
		LowerLimit: lowerLimit,
		UpperLimit: upperLimit,
	}
}

func (p *Parser) parseSqrt(startPos int) ast.Node {
    // Parse optional numeric index argument in square brackets
    index := p.parseOptionalArgument()
    
    // Parse required radicand in curly braces
    radicand := p.parseGroupedStrict()
    if radicand == nil {
        p.addError("expected radicand after \\sqrt", startPos)
        return nil
    }
    
    return &ast.SqrtNode{
        Start:    startPos,
        Radicand: radicand,
        Index:    index,
    }
}
