package ast

type GroupNode struct {
	Start    int
	Elements []Node
}

func (n *GroupNode) Pos() int {
	return n.Start
}

func (n *GroupNode) End() int {
	if len(n.Elements) == 0 {
		return n.Start
	}
	return n.Elements[len(n.Elements)-1].End()
}

type DelimiterNode struct {
	Start int
	Value string
}

func (n *DelimiterNode) Pos() int {
	return n.Start
}

func (n *DelimiterNode) End() int {
	return n.Start + len(n.Value)
}
