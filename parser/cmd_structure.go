package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

// parseFractionCommand parses a LaTeX fraction command: \frac{numerator}{denominator}
func (p *Parser) parseFractionCommand(pos tokenizer.Position) ast.Node {
	numerator := p.parseGroupedStrict()
	if numerator == nil {
		p.addError("expected numerator after \\frac", pos)
		return nil
	}

	denominator := p.parseGroupedStrict()
	if denominator == nil {
		p.addError("expected denominator after \\frac{...}", pos)
		return nil
	}

	return &ast.FractionNode{
		Start:       pos,
		Numerator:   numerator,
		Denominator: denominator,
	}
}

// parseSqrtCommand parses a LaTeX square root command: \sqrt[n]{x} or \sqrt{x}
func (p *Parser) parseSqrtCommand(pos tokenizer.Position) ast.Node {
	// Parse optional index if present (e.g., \sqrt[3]{x} for cube root)
	index := p.parseOptionalArgument()

	// Parse the radicand (expression under the root)
	radicand := p.parseGroupedOrSingle()
	if radicand == nil {
		p.addError("expected radicand after \\sqrt", pos)
		return nil
	}

	return &ast.SqrtNode{
		Start:    pos,
		Radicand: radicand,
		Index:    index,
	}
}

// parseBinomCommand parses a LaTeX binomial coefficient: \binom{n}{k}
func (p *Parser) parseBinomCommand(pos tokenizer.Position) ast.Node {
	upper := p.parseGroupedStrict()
	if upper == nil {
		p.addError("expected upper value after \\binom", pos)
		return nil
	}

	lower := p.parseGroupedStrict()
	if lower == nil {
		p.addError("expected lower value after \\binom{...}", pos)
		return nil
	}

	return &ast.BinomNode{
		Start: pos,
		Upper: upper,
		Lower: lower,
	}
}
