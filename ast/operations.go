package ast

type OperatorNode struct {
	Start int
	Value string
}

func (n *OperatorNode) Pos() int {
	return n.Start
}

func (n *OperatorNode) End() int {
	return n.Start + len(n.Value)
}
