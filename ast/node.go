package ast

// Node represents a node in the LaTeX math abstract syntax tree.
type Node interface {
	// Pos returns the position of the first character of the node.
	Pos() int
	// End returns the position of the character immediately after the node.
	End() int

	// Accept accepts a visitor to this node.
	Accept(v Visitor)
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

func (n *ExpressionNode) Accept(v Visitor) {
	v.VisitExpressionNode(n)
}

// DelimitedExpressionNode represents a LaTeX expression enclosed in delimiters,
// typically created with \left and \right commands. These delimiters automatically
// adjust their size based on the height of the enclosed content.
// Examples: \left( ... \right), \left[ ... \right], \left\{ ... \right\}
type DelimitedExpressionNode struct {
	Start          int
	LeftDelimiter  Node
	Content        Node
	RightDelimiter Node
}

func (n *DelimitedExpressionNode) Pos() int { return n.Start }
func (n *DelimitedExpressionNode) End() int {
	// the end is after the right Delimiter
	return n.RightDelimiter.End()
}

func (n *DelimitedExpressionNode) Accept(v Visitor) {
	v.VisitDelimitedExpressionNode(n)
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

func (n *SymbolNode) Accept(v Visitor) {
	v.VisitSymbolNode(n)
}

// NumberNode represents a numeric literal, e.g., "123".
type NumberNode struct {
	Start int
	Value string
}

func (n *NumberNode) Pos() int { return n.Start }
func (n *NumberNode) End() int { return n.Start + len(n.Value) }

func (n *NumberNode) Accept(v Visitor) {
	v.VisitNumberNode(n)
}

// OperatorNode represents an operator, e.g., "+", "-", "*".
type OperatorNode struct {
	Start int
	Value string
}

func (n *OperatorNode) Pos() int { return n.Start }
func (n *OperatorNode) End() int { return n.Start + len(n.Value) }

func (n *OperatorNode) Accept(v Visitor) {
	v.VisitOperatorNode(n)
}

// NonArgumentFunctionNode represents a mathematical function like \sin, \cos, \log, etc.
// These functions are rendered in upright Roman font with proper spacing, but don't
// take explicit arguments in LaTeX syntax (any following expression is implicitly an argument).
type NonArgumentFunctionNode struct {
	Start int
	Name  string
}

func (n *NonArgumentFunctionNode) Pos() int { return n.Start }
func (n *NonArgumentFunctionNode) End() int {
	// End position is the start position plus the length of the name and backslash
	return n.Start + len(n.Name) + 1 // +1 for the backslash
}

func (n *NonArgumentFunctionNode) Accept(v Visitor) {
	v.VisitNonArgumentFunctionNode(n)
}

// SpaceNode represents a space.
type SpaceNode struct {
	Start int
	Value string
}

func (n *SpaceNode) Pos() int { return n.Start }
func (n *SpaceNode) End() int { return n.Start + len(n.Value) }

func (n *SpaceNode) Accept(v Visitor) {
	v.VisitSpaceNode(n)
}

// DelimiterNode represents a visual math delimiter, such as "(" or "]".
type DelimiterNode struct {
	Start int
	Value string
}

func (n *DelimiterNode) Pos() int { return n.Start }
func (n *DelimiterNode) End() int { return n.Start + len(n.Value) }

func (n *DelimiterNode) Accept(v Visitor) {
	v.VisitDelimiterNode(n)
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

func (n *SuperscriptNode) Accept(v Visitor) {
	v.VisitSuperscriptNode(n)
}

// SubscriptNode represents `base_subscript`.
type SubscriptNode struct {
	Start     int
	Base      Node
	Subscript Node
}

func (n *SubscriptNode) Pos() int { return n.Start }
func (n *SubscriptNode) End() int { return n.Subscript.End() }

func (n *SubscriptNode) Accept(v Visitor) {
	v.VisitSubscriptNode(n)
}

// FractionNode represents a LaTeX `\frac{a}{b}`.
type FractionNode struct {
	Start       int
	Numerator   Node
	Denominator Node
}

func (n *FractionNode) Pos() int { return n.Start }
func (n *FractionNode) End() int { return n.Denominator.End() }

func (n *FractionNode) Accept(v Visitor) {
	v.VisitFractionNode(n)
}

// LimitedOperatorNode represents operators like \int, \sum, \prod, \lim that can have
// limits (subscripts and/or superscripts).
type LimitedOperatorNode struct {
	Start      int
	Operator   string // "int", "sum", "prod", "lim", etc.
	LowerLimit Node
	UpperLimit Node
}

func (n *LimitedOperatorNode) Pos() int { return n.Start }
func (n *LimitedOperatorNode) End() int {
	// If there are limits, the end is the end of the last limit
	if n.UpperLimit != nil {
		return n.UpperLimit.End()
	}
	if n.LowerLimit != nil {
		return n.LowerLimit.End()
	}
	// Length of the operator backslash + name
	return n.Start + len(n.Operator) + 1
}

func (n *LimitedOperatorNode) Accept(v Visitor) {
	v.VisitLimitedOperatorNode(n)
}

// SqrtNode represents a LaTeX `\sqrt` command, optionally with an index.
// For example: \sqrt{x} for a square root, or \sqrt[n]{x} for an nth root.
type SqrtNode struct {
	Start    int
	Radicand Node // The expression under the radical
	Index    Node // Optional: The index for nth roots (as in \sqrt[n]{x})
}

func (n *SqrtNode) Pos() int { return n.Start }
func (n *SqrtNode) End() int {
	// If there's a radicand, the end is the end of the radicand
	if n.Radicand != nil {
		return n.Radicand.End()
	}
	// If there's no radicand (which shouldn't happen in valid LaTeX),
	// the end is the start plus the length of \sqrt
	return n.Start + 5 // Length of "\sqrt"
}

func (n *SqrtNode) Accept(v Visitor) {
	v.VisitSqrtNode(n)
}

// BinomNode represents a LaTeX `\binom{a}{b}` command for binomial coefficients.
type BinomNode struct {
	Start int
	Upper Node
	Lower Node
}

func (n *BinomNode) Pos() int { return n.Start }
func (n *BinomNode) End() int { return n.Lower.End() }

func (n *BinomNode) Accept(v Visitor) {
	v.VisitBinomNode(n)
}
