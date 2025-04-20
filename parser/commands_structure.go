package parser

import "github.com/neox5/texmax/ast"

// parseFractionCommand parses a LaTeX fraction command: \frac{numerator}{denominator}
func (p *Parser) parseFractionCommand(startPos int) ast.Node {
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

// parseSqrtCommand parses a LaTeX square root command: \sqrt[n]{x} or \sqrt{x}
func (p *Parser) parseSqrtCommand(startPos int) ast.Node {
	// Parse optional index if present (e.g., \sqrt[3]{x} for cube root)
	index := p.parseOptionalArgument()

	// Parse the radicand (expression under the root)
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

// parseBinomCommand parses a LaTeX binomial coefficient: \binom{n}{k}
func (p *Parser) parseBinomCommand(startPos int) ast.Node {
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
