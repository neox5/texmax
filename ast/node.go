package ast

// Node is the base interface for all AST nodes
type Node interface {
	// Position returns the node's starting position in source
	Position() int

	// Children returns the node's children for traversal
	Children() []Node

	// ToLaTeX renders the node back to LaTeX format
	ToLaTeX() string
}

// BaseNode implements the Position method of the Node interface
type BaseNode struct {
	Pos int // Source position
}

func (n *BaseNode) Position() int {
	return n.Pos
}

// SpaceNode represents whitespace in the input
type SpaceNode struct {
	BaseNode
	Value string
}

func (n *SpaceNode) Children() []Node {
	return nil
}

func (n *SpaceNode) ToLaTeX() string {
	return n.Value
}

// CommandNode represents a LaTeX command with arguments
type CommandNode struct {
	BaseNode
	Name      string
	Arguments []Node
}

func (n *CommandNode) Children() []Node {
	return n.Arguments
}

func (n *CommandNode) ToLaTeX() string {
	result := "\\" + n.Name
	for _, arg := range n.Arguments {
		result += "{" + arg.ToLaTeX() + "}"
	}
	return result
}

// SymbolNode represents a variable or constant
type SymbolNode struct {
	BaseNode
	Value string
}

func (n *SymbolNode) Children() []Node {
	return nil
}

func (n *SymbolNode) ToLaTeX() string {
	return n.Value
}

// NumberNode represents a numeric literal
type NumberNode struct {
	BaseNode
	Value string
}

func (n *NumberNode) Children() []Node {
	return nil
}

func (n *NumberNode) ToLaTeX() string {
	return n.Value
}

// OperatorNode represents a mathematical operator
type OperatorNode struct {
	BaseNode
	Value string
}

func (n *OperatorNode) Children() []Node {
	return nil
}

func (n *OperatorNode) ToLaTeX() string {
	return n.Value
}

// SuperscriptNode represents a base with a superscript
type SuperscriptNode struct {
	BaseNode
	Base     Node
	Exponent Node
}

func (n *SuperscriptNode) Children() []Node {
	return []Node{n.Base, n.Exponent}
}

func (n *SuperscriptNode) ToLaTeX() string {
	return n.Base.ToLaTeX() + "^" + n.Exponent.ToLaTeX()
}

// SubscriptNode represents a base with a subscript
type SubscriptNode struct {
	BaseNode
	Base      Node
	Subscript Node
}

func (n *SubscriptNode) Children() []Node {
	return []Node{n.Base, n.Subscript}
}

func (n *SubscriptNode) ToLaTeX() string {
	return n.Base.ToLaTeX() + "_" + n.Subscript.ToLaTeX()
}

// GroupNode represents a group of elements in braces
type GroupNode struct {
	BaseNode
	Elements []Node
}

func (n *GroupNode) Children() []Node {
	return n.Elements
}

func (n *GroupNode) ToLaTeX() string {
	result := "{"
	for _, elem := range n.Elements {
		result += elem.ToLaTeX()
	}
	result += "}"
	return result
}

// DelimiterNode represents a delimiter like parentheses, brackets, etc.
type DelimiterNode struct {
	BaseNode
	Value string
}

func (n *DelimiterNode) Children() []Node {
	return nil
}

func (n *DelimiterNode) ToLaTeX() string {
	return n.Value
}

// FractionNode represents a LaTeX fraction
type FractionNode struct {
	BaseNode
	Numerator   Node
	Denominator Node
}

func (n *FractionNode) Children() []Node {
	return []Node{n.Numerator, n.Denominator}
}

func (n *FractionNode) ToLaTeX() string {
	return "\\frac{" + n.Numerator.ToLaTeX() + "}{" + n.Denominator.ToLaTeX() + "}"
}
