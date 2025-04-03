package ast

type CommandNode struct {
	Start     int
	Name      string
	Arguments []Node
}

func (n *CommandNode) Pos() int {
	return n.Start
}

func (n *CommandNode) End() int {
	end := n.Start + len(n.Name)
	for _, arg := range n.Arguments {
		if arg.End() > end {
			end = arg.End()
		}
	}
	return end
}
