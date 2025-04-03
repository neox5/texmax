package ast

type FractionNode struct {
	Start       int
	Numerator   Node
	Denominator Node
}

func (n *FractionNode) Pos() int {
	return n.Start
}

func (n *FractionNode) End() int {
	return n.Denominator.End()
}
