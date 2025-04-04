package ast

// Node represents a node in the LaTeX math abstract syntax tree.
type Node interface {
	// Pos returns the position of the first character of the node.
	Pos() int
	// End returns the position of the character immediately after the node.
	End() int
}

// --------------------
// Root Node
// --------------------

// MathExpressionNode represents a complete LaTeX math expression.
type MathExpressionNode struct {
	Start    int
	Elements []Node
}

func (n *MathExpressionNode) Pos() int { return n.Start }
func (n *MathExpressionNode) End() int {
	if len(n.Elements) == 0 {
		return n.Start
	}
	return n.Elements[len(n.Elements)-1].End()
}

// --------------------
// Leaf Nodes
// --------------------

// SymbolNode represents a single letter or variable, e.g., "x".
type SymbolNode struct {
	Start int
	Value string
}

func (n *SymbolNode) Pos() int { return n.Start }
func (n *SymbolNode) End() int { return n.Start + len(n.Value) }

// NumberNode represents a numeric literal, e.g., "123".
type NumberNode struct {
	Start int
	Value string
}

func (n *NumberNode) Pos() int { return n.Start }
func (n *NumberNode) End() int { return n.Start + len(n.Value) }

// OperatorNode represents an operator, e.g., "+", "-", "*".
type OperatorNode struct {
	Start int
	Value string
}

func (n *OperatorNode) Pos() int { return n.Start }
func (n *OperatorNode) End() int { return n.Start + len(n.Value) }

// SpaceNode represents a space or spacing command.
type SpaceNode struct {
	Start int
	Value string
}

func (n *SpaceNode) Pos() int { return n.Start }
func (n *SpaceNode) End() int { return n.Start + len(n.Value) }

// DelimiterNode represents a visual math delimiter, such as "(" or "]".
type DelimiterNode struct {
	Start int
	Value string
}

func (n *DelimiterNode) Pos() int { return n.Start }
func (n *DelimiterNode) End() int { return n.Start + len(n.Value) }

// --------------------
// Composite Nodes
// --------------------

// ArgumentGroupNode represents a braced `{...}` group.
type ArgumentGroupNode struct {
	Start    int
	Elements []Node
}

func (n *ArgumentGroupNode) Pos() int { return n.Start }
func (n *ArgumentGroupNode) End() int {
	if len(n.Elements) == 0 {
		return n.Start
	}
	return n.Elements[len(n.Elements)-1].End()
}

// CommandNode represents a LaTeX command with arguments, e.g., `\frac{a}{b}`.
type CommandNode struct {
	Start     int
	Name      string
	Arguments []Node
}

func (n *CommandNode) Pos() int { return n.Start }
func (n *CommandNode) End() int {
	end := n.Start + len(n.Name)
	for _, arg := range n.Arguments {
		if arg.End() > end {
			end = arg.End()
		}
	}
	return end
}

// SuperscriptNode represents `base^exponent`.
type SuperscriptNode struct {
	Start    int
	Base     Node
	Exponent Node
}

func (n *SuperscriptNode) Pos() int { return n.Start }
func (n *SuperscriptNode) End() int { return n.Exponent.End() }

// SubscriptNode represents `base_subscript`.
type SubscriptNode struct {
	Start     int
	Base      Node
	Subscript Node
}

func (n *SubscriptNode) Pos() int { return n.Start }
func (n *SubscriptNode) End() int { return n.Subscript.End() }

// FractionNode represents a LaTeX `\frac{a}{b}`.
type FractionNode struct {
	Start       int
	Numerator   Node
	Denominator Node
}

func (n *FractionNode) Pos() int { return n.Start }
func (n *FractionNode) End() int { return n.Denominator.End() }
