package parser

import "github.com/neox5/texmax/ast"

// nonArgumentFunctions maps LaTeX math function names (like \sin, \cos)
// that don't take explicit arguments in LaTeX syntax.
var nonArgumentFunctions = map[string]bool{
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

// isNonArgumentFunction checks if a command is a math function without explicit arguments.
func isNonArgumentFunction(name string) bool {
	_, ok := nonArgumentFunctions[name]
	return ok
}

// parseNonArgumentFunction creates a NonArgumentFunctionNode for functions like \sin, \cos
func (p *Parser) parseNonArgumentFunction(name string, startPos int) ast.Node {
	return &ast.NonArgumentFunctionNode{
		Start: startPos,
		Name:  name,
	}
}

// operators maps LaTeX big operators that can have upper/lower limits.
var operators = map[string]bool{
	"int":  true, // integral
	"sum":  true, // summation
	"prod": true, // product
	"lim":  true, // limit
}

// isOperator checks if a command is a big operator that can have limits.
func isOperator(name string) bool {
	_, ok := operators[name]
	return ok
}

// parseOperator parses big operators like \int, \sum, \prod, \lim
// that can have limits (subscripts/superscripts).
func (p *Parser) parseOperator(operatorName string, startPos int) ast.Node {
	lowerLimit, upperLimit := p.parseLimits()

	// For \lim, we should validate that it only has a lower limit
	if operatorName == "lim" && upperLimit != nil {
		p.addError("\\lim can only have a lower limit", upperLimit.Pos())
		upperLimit = nil // Ignore upper limit for \lim
	}

	return &ast.LimitedOperatorNode{
		Start:      startPos,
		Operator:   operatorName,
		LowerLimit: lowerLimit,
		UpperLimit: upperLimit,
	}
}
