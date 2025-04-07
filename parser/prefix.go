package parser

import "github.com/neox5/texmax/ast"

func (p *Parser) parseSymbol() ast.Node {
	t := p.next()
	return &ast.SymbolNode{Start: t.Pos, Value: t.Value}
}

func (p *Parser) parseNumber() ast.Node {
	t := p.next()
	return &ast.NumberNode{Start: t.Pos, Value: t.Value}
}

func (p *Parser) parseSpace() ast.Node {
	t := p.next()
	return &ast.SpaceNode{Start: t.Pos, Value: t.Value}
}
