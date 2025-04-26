package ast

import "github.com/neox5/texmax/tokenizer"

// Node represents a node in the LaTeX math abstract syntax tree.
type Node interface {
	// Pos returns the position of the first character of the node.
	Pos() tokenizer.Position
	// End returns the position of the character immediately after the node.
	End() tokenizer.Position

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
	Start    tokenizer.Position
	Elements []Node
}

func (n *ExpressionNode) Pos() tokenizer.Position { return n.Start }
func (n *ExpressionNode) End() tokenizer.Position {
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
	Start          tokenizer.Position
	LeftDelimiter  Node
	Content        Node
	RightDelimiter Node
}

func (n *DelimitedExpressionNode) Pos() tokenizer.Position { return n.Start }
func (n *DelimitedExpressionNode) End() tokenizer.Position {
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
	Start tokenizer.Position
	Value string
}

func (n *SymbolNode) Pos() tokenizer.Position { return n.Start }
func (n *SymbolNode) End() tokenizer.Position {
	end := n.Start
	end.Offset += len(n.Value)
	end.Column += len(n.Value)
	return end
}

func (n *SymbolNode) Accept(v Visitor) {
	v.VisitSymbolNode(n)
}

// NumberNode represents a numeric literal, e.g., "123".
type NumberNode struct {
	Start tokenizer.Position
	Value string
}

func (n *NumberNode) Pos() tokenizer.Position { return n.Start }
func (n *NumberNode) End() tokenizer.Position {
	end := n.Start
	end.Offset += len(n.Value)
	end.Column += len(n.Value)
	return end
}

func (n *NumberNode) Accept(v Visitor) {
	v.VisitNumberNode(n)
}

// OperatorNode represents an operator, e.g., "+", "-", "*".
type OperatorNode struct {
	Start tokenizer.Position
	Value string
}

func (n *OperatorNode) Pos() tokenizer.Position { return n.Start }
func (n *OperatorNode) End() tokenizer.Position {
	end := n.Start
	end.Offset += len(n.Value)
	end.Column += len(n.Value)
	return end
}

func (n *OperatorNode) Accept(v Visitor) {
	v.VisitOperatorNode(n)
}

// NonArgumentFunctionNode represents a mathematical function like \sin, \cos, \log, etc.
// These functions are rendered in upright Roman font with proper spacing, but don't
// take explicit arguments in LaTeX syntax (any following expression is implicitly an argument).
type NonArgumentFunctionNode struct {
	Start tokenizer.Position
	Name  string
}

func (n *NonArgumentFunctionNode) Pos() tokenizer.Position { return n.Start }
func (n *NonArgumentFunctionNode) End() tokenizer.Position {
	end := n.Start
	// End position is the start position plus the length of the name and backslash
	end.Offset += len(n.Name) + 1 // +1 for the backslash
	end.Column += len(n.Name) + 1
	return end
}

func (n *NonArgumentFunctionNode) Accept(v Visitor) {
	v.VisitNonArgumentFunctionNode(n)
}

// SpaceNode represents a space.
type SpaceNode struct {
	Start tokenizer.Position
	Value string
}

func (n *SpaceNode) Pos() tokenizer.Position { return n.Start }
func (n *SpaceNode) End() tokenizer.Position {
	end := n.Start
	end.Offset += len(n.Value)
	end.Column += len(n.Value)
	return end
}

func (n *SpaceNode) Accept(v Visitor) {
	v.VisitSpaceNode(n)
}

// DelimiterNode represents a visual math delimiter, such as "(" or "]".
type DelimiterNode struct {
	Start tokenizer.Position
	Value string
}

func (n *DelimiterNode) Pos() tokenizer.Position { return n.Start }
func (n *DelimiterNode) End() tokenizer.Position {
	end := n.Start
	end.Offset += len(n.Value)
	end.Column += len(n.Value)
	return end
}

func (n *DelimiterNode) Accept(v Visitor) {
	v.VisitDelimiterNode(n)
}

// --------------------
// Composite Nodes
// --------------------

// SuperscriptNode represents `base^exponent`.
type SuperscriptNode struct {
	Start    tokenizer.Position
	Base     Node
	Exponent Node
}

func (n *SuperscriptNode) Pos() tokenizer.Position { return n.Start }
func (n *SuperscriptNode) End() tokenizer.Position { return n.Exponent.End() }

func (n *SuperscriptNode) Accept(v Visitor) {
	v.VisitSuperscriptNode(n)
}

// SubscriptNode represents `base_subscript`.
type SubscriptNode struct {
	Start     tokenizer.Position
	Base      Node
	Subscript Node
}

func (n *SubscriptNode) Pos() tokenizer.Position { return n.Start }
func (n *SubscriptNode) End() tokenizer.Position { return n.Subscript.End() }

func (n *SubscriptNode) Accept(v Visitor) {
	v.VisitSubscriptNode(n)
}

// FractionNode represents a LaTeX `\frac{a}{b}`.
type FractionNode struct {
	Start       tokenizer.Position
	Numerator   Node
	Denominator Node
}

func (n *FractionNode) Pos() tokenizer.Position { return n.Start }
func (n *FractionNode) End() tokenizer.Position { return n.Denominator.End() }

func (n *FractionNode) Accept(v Visitor) {
	v.VisitFractionNode(n)
}

// LimitedOperatorNode represents operators like \int, \sum, \prod, \lim that can have
// limits (subscripts and/or superscripts).
type LimitedOperatorNode struct {
	Start      tokenizer.Position
	Operator   string // "int", "sum", "prod", "lim", etc.
	LowerLimit Node
	UpperLimit Node
}

func (n *LimitedOperatorNode) Pos() tokenizer.Position { return n.Start }
func (n *LimitedOperatorNode) End() tokenizer.Position {
	// If there are limits, the end is the end of the last limit
	if n.UpperLimit != nil {
		return n.UpperLimit.End()
	}
	if n.LowerLimit != nil {
		return n.LowerLimit.End()
	}
	// Length of the operator backslash + name
	end := n.Start
	end.Offset += len(n.Operator) + 1
	end.Column += len(n.Operator) + 1
	return end
}

func (n *LimitedOperatorNode) Accept(v Visitor) {
	v.VisitLimitedOperatorNode(n)
}

// SqrtNode represents a LaTeX `\sqrt` command, optionally with an index.
// For example: \sqrt{x} for a square root, or \sqrt[n]{x} for an nth root.
type SqrtNode struct {
	Start    tokenizer.Position
	Radicand Node // The expression under the radical
	Index    Node // Optional: The index for nth roots (as in \sqrt[n]{x})
}

func (n *SqrtNode) Pos() tokenizer.Position { return n.Start }
func (n *SqrtNode) End() tokenizer.Position {
	// If there's a radicand, the end is the end of the radicand
	if n.Radicand != nil {
		return n.Radicand.End()
	}
	// If there's no radicand (which shouldn't happen in valid LaTeX),
	// the end is the start plus the length of \sqrt
	end := n.Start
	end.Offset += 5 // Length of "\sqrt"
	end.Column += 5
	return end
}

func (n *SqrtNode) Accept(v Visitor) {
	v.VisitSqrtNode(n)
}

// BinomNode represents a LaTeX `\binom{a}{b}` command for binomial coefficients.
type BinomNode struct {
	Start tokenizer.Position
	Upper Node
	Lower Node
}

func (n *BinomNode) Pos() tokenizer.Position { return n.Start }
func (n *BinomNode) End() tokenizer.Position { return n.Lower.End() }

func (n *BinomNode) Accept(v Visitor) {
	v.VisitBinomNode(n)
}

// MatrixNode represents a LaTeX matrix environment containing a grid of elements
// arranged in rows and columns, such as in \begin{matrix}...\end{matrix}
type MatrixNode struct {
	Start tokenizer.Position
	Rows  [][]Node // A 2D array of nodes representing the cells
}

func (n *MatrixNode) Pos() tokenizer.Position { return n.Start }
func (n *MatrixNode) End() tokenizer.Position {
	if len(n.Rows) == 0 {
		return n.Start
	}
	lastRow := n.Rows[len(n.Rows)-1]
	if len(lastRow) == 0 {
		return n.Start
	}
	return lastRow[len(lastRow)-1].End()
}

func (n *MatrixNode) Accept(v Visitor) {
	v.VisitMatrixNode(n)
}
