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
	"max":    true,
	"min":    true,
	"det":    true,
	"arg":    true,
	"mod":    true,
}

// List of operators that support limits
var limitedOperators = map[string]bool{
	"int":  true,
	"sum":  true,
	"prod": true,
	"lim":  true,
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

	// Check if this is a limited operator
	if limitedOperators[commandName] {
		return p.parseLimitedOperator(commandName, startPos)
	}

	switch commandName {
	case "frac":
		return p.parseFrac(startPos)
	case "sqrt":
		return p.parseSqrt(startPos)
	case "binom":
		return p.parseBinom(startPos)
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

func (p *Parser) parseLimitedOperator(operator string, startPos int) ast.Node {
	// Parse optional limits (subscript and superscript)
	lowerLimit, upperLimit := p.parseLimits()

	// For \lim, we should validate that it only has a lower limit
	if operator == "lim" && upperLimit != nil {
		p.addError("\\lim can only have a lower limit", upperLimit.Pos())
		upperLimit = nil // remove upperLimit to avoid any issues with \lim
	}

	// Create a LimitedOperatorNode with the limits
	return &ast.LimitedOperatorNode{
		Start:      startPos,
		Operator:   operator,
		LowerLimit: lowerLimit,
		UpperLimit: upperLimit,
	}
}

func (p *Parser) parseSqrt(startPos int) ast.Node {
	// Parse optional numeric index argument in square brackets
	index := p.parseOptionalArgument()

	// Parse radicand in curly braces or single token
	radicand := p.parseGroupedOrSingle()
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

func (p *Parser) parseBinom(startPos int) ast.Node {
	// A binomial coefficient requires two arguments: upper and lower
	upper := p.parseGroupedStrict()
	if upper == nil {
		p.addError("expected upper value after \\binom", startPos)
		return nil
	}

	lower := p.parseGroupedStrict()
	if lower == nil {
		p.addError("expected lower value after \\binom{...}", startPos)
		return nil
	}

	return &ast.BinomNode{
		Start: startPos,
		Upper: upper,
		Lower: lower,
	}
}
