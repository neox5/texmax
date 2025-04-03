package ast

type SuperscriptNode struct {
	Start    int
	Base     Node
	Exponent Node
}

func (n *SuperscriptNode) Pos() int {
	return n.Start
}

func (n *SuperscriptNode) End() int {
	return n.Exponent.End()
}

type SubscriptNode struct {
	Start     int
	Base      Node
	Subscript Node
}

func (n *SubscriptNode) Pos() int {
	return n.Start
}

func (n *SubscriptNode) End() int {
	return n.Subscript.End()
}
