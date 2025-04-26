package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

// List of Greek letter commands in LaTeX
var greekLetters = map[string]string{
	// Lowercase Greek letters
	"alpha":   "α",
	"beta":    "β",
	"gamma":   "γ",
	"delta":   "δ",
	"epsilon": "ε",
	"zeta":    "ζ",
	"eta":     "η",
	"theta":   "θ",
	"iota":    "ι",
	"kappa":   "κ",
	"lambda":  "λ",
	"mu":      "μ",
	"nu":      "ν",
	"xi":      "ξ",
	"omicron": "ο",
	"pi":      "π",
	"rho":     "ρ",
	"sigma":   "σ",
	"tau":     "τ",
	"upsilon": "υ",
	"phi":     "φ",
	"chi":     "χ",
	"psi":     "ψ",
	"omega":   "ω",

	// Uppercase Greek letters
	"Alpha":   "Α",
	"Beta":    "Β",
	"Gamma":   "Γ",
	"Delta":   "Δ",
	"Epsilon": "Ε",
	"Zeta":    "Ζ",
	"Eta":     "Η",
	"Theta":   "Θ",
	"Iota":    "Ι",
	"Kappa":   "Κ",
	"Lambda":  "Λ",
	"Mu":      "Μ",
	"Nu":      "Ν",
	"Xi":      "Ξ",
	"Omicron": "Ο",
	"Pi":      "Π",
	"Rho":     "Ρ",
	"Sigma":   "Σ",
	"Tau":     "Τ",
	"Upsilon": "Υ",
	"Phi":     "Φ",
	"Chi":     "Χ",
	"Psi":     "Ψ",
	"Omega":   "Ω",
}

// IsGreekLetter checks if the given command name represents a Greek letter
func isGreekLetter(name string) bool {
	_, ok := greekLetters[name]
	return ok
}

// parseGreek parses a Greek letter command and returns a SymbolNode
func (p *Parser) parseGreekLetter(commandName string, pos tokenizer.Position) ast.Node {
	// Look up the Unicode representation of the Greek letter
	if symbol, ok := greekLetters[commandName]; ok {
		return &ast.SymbolNode{
			Start: pos,
			Value: symbol,
		}
	}

	// This shouldn't happen if IsGreekLetter is checked first
	return nil
}
