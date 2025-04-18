package ast

import (
	"fmt"
	"io"
	"strings"
)

// PrintVisitor prints the AST with proper indentation in tree format
type PrintVisitor struct {
	Writer io.Writer
	Depth  int
}

// NewPrintVisitor creates a new PrintVisitor
func NewPrintVisitor(w io.Writer) *PrintVisitor {
	return &PrintVisitor{
		Writer: w,
		Depth:  0,
	}
}

// indent returns the indentation string for the current depth
func (v *PrintVisitor) indent() string {
	if v.Depth == 0 {
		return ""
	}

	// Use a simple pipe-based indentation
	// For every level, add a pipe and some spaces
	return strings.Repeat("  ", v.Depth-1) + "|- "
}

// printNode helper function to print a node with proper indentation
func (v *PrintVisitor) printNode(description string) {
	if v.Depth == 0 {
		fmt.Fprintf(v.Writer, "%s\n", description)
	} else {
		fmt.Fprintf(v.Writer, "%s%s\n", v.indent(), description)
	}
}

// Visit methods for container nodes
func (v *PrintVisitor) VisitExpressionNode(node *ExpressionNode) {
	v.printNode(fmt.Sprintf("ExpressionNode (%d nodes)", len(node.Elements)))
	v.Depth++
	
	for _, element := range node.Elements {
		element.Accept(v)
	}
	
	v.Depth--
}

// Visit methods for leaf nodes
func (v *PrintVisitor) VisitSymbolNode(node *SymbolNode) {
	v.printNode(fmt.Sprintf("SymbolNode (%s)", node.Value))
}

func (v *PrintVisitor) VisitNumberNode(node *NumberNode) {
	v.printNode(fmt.Sprintf("NumberNode (%s)", node.Value))
}

func (v *PrintVisitor) VisitOperatorNode(node *OperatorNode) {
	v.printNode(fmt.Sprintf("OperatorNode (%s)", node.Value))
}

func (v *PrintVisitor) VisitNonArgumentFunctionNode(node *NonArgumentFunctionNode) {
	v.printNode(fmt.Sprintf("NonArgumentFunctionNode (%s)", node.Name))
}

func (v *PrintVisitor) VisitSpaceNode(node *SpaceNode) {
	v.printNode("SpaceNode")
}

func (v *PrintVisitor) VisitDelimiterNode(node *DelimiterNode) {
	v.printNode(fmt.Sprintf("DelimiterNode (%s)", node.Value))
}

// Visit methods for composite nodes
func (v *PrintVisitor) VisitSuperscriptNode(node *SuperscriptNode) {
	v.printNode("SuperscriptNode")
	v.Depth++
	
	node.Base.Accept(v)
	node.Exponent.Accept(v)
	
	v.Depth--
}

func (v *PrintVisitor) VisitSubscriptNode(node *SubscriptNode) {
	v.printNode("SubscriptNode")
	v.Depth++
	
	node.Base.Accept(v)
	node.Subscript.Accept(v)
	
	v.Depth--
}

func (v *PrintVisitor) VisitFractionNode(node *FractionNode) {
	v.printNode("FractionNode")
	v.Depth++
	
	node.Numerator.Accept(v)
	node.Denominator.Accept(v)
	
	v.Depth--
}

func (v *PrintVisitor) VisitIntegralNode(node *IntegralNode) {
	v.printNode("IntegralNode")
	v.Depth++
	
	if node.LowerLimit != nil {
		v.printNode("LowerLimit:")
		v.Depth++
		node.LowerLimit.Accept(v)
		v.Depth--
	}
	
	if node.UpperLimit != nil {
		v.printNode("UpperLimit:")
		v.Depth++
		node.UpperLimit.Accept(v)
		v.Depth--
	}
	
	v.Depth--
}

func (v *PrintVisitor) VisitSqrtNode(node *SqrtNode) {
	if node.Index != nil {
		v.printNode("SqrtNode (with index)")
	} else {
		v.printNode("SqrtNode (square root)")
	}
	v.Depth++
	
	if node.Index != nil {
		v.printNode("Index:")
		v.Depth++
		node.Index.Accept(v)
		v.Depth--
	}
	
	if node.Radicand != nil {
		v.printNode("Radicand:")
		v.Depth++
		node.Radicand.Accept(v)
		v.Depth--
	}
	
	v.Depth--
}
