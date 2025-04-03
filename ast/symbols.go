package ast

type SymbolNode struct {
	Start int
	Value string
}

func (n *SymbolNode) Pos() int {
	return n.Start
}

func (n *SymbolNode) End() int {
	return n.Start + len(n.Value)
}

type NumberNode struct {
	Start int
	Value string
}

func (n *NumberNode) Pos() int {
	return n.Start
}

func (n *NumberNode) End() int {
	return n.Start + len(n.Value)
}

type SpaceNode struct {
	Start int
	Value string
}

func (n *SpaceNode) Pos() int {
	return n.Start
}

func (n *SpaceNode) End() int {
	return n.Start + len(n.Value)
}
