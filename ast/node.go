package ast

import "fmt"

// Node represents a node in the LaTeX math abstract syntax tree.
type Node interface {
	// Pos returns the position of the first character of the node.
	Pos() int
	// End returns the position of the character immediately after the node.
	End() int
	
	// VisitChildren visits all child nodes with the given visitor
	VisitChildren(v Visitor)
	// implementation of Stringer interface
	String() string
}

// --------------------
// Container Node
// --------------------

// ExpressionNode represents a sequence of LaTeX math expressions that form a coherent unit.
// This serves as a general-purpose container for multiple nodes and can be used
// at any level of the AST, including as the root.
type ExpressionNode struct {
	Start    int
	Elements []Node
}

func (n *ExpressionNode) Pos() int { return n.Start }
func (n *ExpressionNode) End() int {
	if len(n.Elements) == 0 {
		return n.Start
	}
	return n.Elements[len(n.Elements)-1].End()
}

func (n *ExpressionNode) VisitChildren(v Visitor) {
	v.EnterNode(n)
	for _, el := range n.Elements {
		Walk(v, el)
	}
	v.ExitNode(n)
}

func (n *ExpressionNode) String() string {
	return fmt.Sprintf("ExpressionNode (%d nodes)",len(n.Elements))
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

func (n *SymbolNode) VisitChildren(v Visitor) {
	// Leaf node, no children to visit
}

func (n *SymbolNode) String() string {
	return fmt.Sprintf("SymbolNode (%s)", n.Value)
}

// NumberNode represents a numeric literal, e.g., "123".
type NumberNode struct {
	Start int
	Value string
}

func (n *NumberNode) Pos() int { return n.Start }
func (n *NumberNode) End() int { return n.Start + len(n.Value) }

func (n *NumberNode) VisitChildren(v Visitor) {
	// Leaf node, no children to visit
}

func (n *NumberNode) String() string {
	return fmt.Sprintf("NumberNode (%s)", n.Value)
}

// OperatorNode represents an operator, e.g., "+", "-", "*".
type OperatorNode struct {
	Start int
	Value string
}

func (n *OperatorNode) Pos() int { return n.Start }
func (n *OperatorNode) End() int { return n.Start + len(n.Value) }

func (n *OperatorNode) VisitChildren(v Visitor) {
	// Leaf node, no children to visit
}

func (n *OperatorNode) String() string {
	return fmt.Sprintf("OperatorNode (%s)", n.Value)
}

// SpaceNode represents a space.
type SpaceNode struct {
	Start int
	Value string
}

func (n *SpaceNode) Pos() int { return n.Start }
func (n *SpaceNode) End() int { return n.Start + len(n.Value) }

func (n *SpaceNode) VisitChildren(v Visitor) {
	// Leaf node, no children to visit
}

func (n *SpaceNode) String() string {
	return "SpaceNode"
}

// DelimiterNode represents a visual math delimiter, such as "(" or "]".
type DelimiterNode struct {
	Start int
	Value string
}

func (n *DelimiterNode) Pos() int { return n.Start }
func (n *DelimiterNode) End() int { return n.Start + len(n.Value) }

func (n *DelimiterNode) VisitChildren(v Visitor) {
	// Leaf node, no children to visit
}

func (n *DelimiterNode) String() string {
	return fmt.Sprintf("DelimiterNode (%s)", n.Value)
}

// --------------------
// Composite Nodes
// --------------------

// SuperscriptNode represents `base^exponent`.
type SuperscriptNode struct {
	Start    int
	Base     Node
	Exponent Node
}

func (n *SuperscriptNode) Pos() int { return n.Start }
func (n *SuperscriptNode) End() int { return n.Exponent.End() }

func (n *SuperscriptNode) VisitChildren(v Visitor) {
	v.EnterNode(n)
	
	// Visit base
	Walk(v, n.Base)
	
	// Visit exponent
	Walk(v, n.Exponent)
	
	v.ExitNode(n)
}

func (n *SuperscriptNode) String() string {
	return "SuperscriptNode"
}

// SubscriptNode represents `base_subscript`.
type SubscriptNode struct {
	Start     int
	Base      Node
	Subscript Node
}

func (n *SubscriptNode) Pos() int { return n.Start }
func (n *SubscriptNode) End() int { return n.Subscript.End() }

func (n *SubscriptNode) VisitChildren(v Visitor) {
	v.EnterNode(n)
	
	// Visit base
	Walk(v, n.Base)
	
	// Visit subscript
	Walk(v, n.Subscript)
	
	v.ExitNode(n)
}

func (n *SubscriptNode) String() string {
	return "SubscriptNode"
}

// FractionNode represents a LaTeX `\frac{a}{b}`.
type FractionNode struct {
	Start       int
	Numerator   Node
	Denominator Node
}

func (n *FractionNode) Pos() int { return n.Start }
func (n *FractionNode) End() int { return n.Denominator.End() }

func (n *FractionNode) VisitChildren(v Visitor) {
	v.EnterNode(n)
	
	// Visit numerator
	Walk(v, n.Numerator)
	
	// Visit denominator
	Walk(v, n.Denominator)
	
	v.ExitNode(n)
}

func (n *FractionNode) String() string {
	return "FractionNode"
}

// IntegralNode represents a LaTeX \int command with optional limits.
type IntegralNode struct {
	Start      int
	LowerLimit Node
	UpperLimit Node
}

func (n *IntegralNode) Pos() int { return n.Start }
func (n *IntegralNode) End() int {
	// If there are limits, the end is the end of the last limit
	if n.UpperLimit != nil {
		return n.UpperLimit.End()
	}
	if n.LowerLimit != nil {
		return n.LowerLimit.End()
	}
	return n.Start + 4 // Length of "\int"
}

func (n *IntegralNode) VisitChildren(v Visitor) {
	v.EnterNode(n)
	
	// Visit lower limit if it exists
	if n.LowerLimit != nil {
		Walk(v, n.LowerLimit)
	}
	
	// Visit upper limit if it exists
	if n.UpperLimit != nil {
		Walk(v, n.UpperLimit)
	}
	
	v.ExitNode(n)
}

func (n *IntegralNode) String() string {
	return "IntegralNode"
}
