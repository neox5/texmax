package ast

import (
	"fmt"
	"io"
	"strings"
)

// GoLikePrinter prints the AST in a manner similar to Go's AST format
type GoLikePrinter struct {
	Writer io.Writer
	Depth  int
}

// NewGoLikePrinter creates a new GoLikePrinter
func NewGoLikePrinter(w io.Writer) *GoLikePrinter {
	return &GoLikePrinter{
		Writer: w,
		Depth:  0,
	}
}

// indent returns the proper indentation string for the current depth
func (p *GoLikePrinter) indent() string {
	return strings.Repeat("  ", p.Depth)
}

// increaseDepth increases the indentation depth
func (p *GoLikePrinter) increaseDepth() {
	p.Depth++
}

// decreaseDepth decreases the indentation depth
func (p *GoLikePrinter) decreaseDepth() {
	if p.Depth > 0 {
		p.Depth--
	}
}

// printIndent prints the current indentation level
func (p *GoLikePrinter) printIndent() {
	fmt.Fprint(p.Writer, p.indent())
}

// Visit methods for container nodes
func (p *GoLikePrinter) VisitExpressionNode(node *ExpressionNode) {
	p.printIndent()
	fmt.Fprintf(p.Writer, "*ast.ExpressionNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Elements: []ast.Node (len = %d) {\n", len(node.Elements))
	p.increaseDepth()

	for i, element := range node.Elements {
		p.printIndent()
		fmt.Fprintf(p.Writer, "%d: ", i)
		element.Accept(p)
	}

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

// Visit methods for leaf nodes
func (p *GoLikePrinter) VisitSymbolNode(node *SymbolNode) {
	fmt.Fprintf(p.Writer, "*ast.SymbolNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Value: %q\n", node.Value)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitNumberNode(node *NumberNode) {
	fmt.Fprintf(p.Writer, "*ast.NumberNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Value: %q\n", node.Value)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitOperatorNode(node *OperatorNode) {
	fmt.Fprintf(p.Writer, "*ast.OperatorNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Value: %q\n", node.Value)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitNonArgumentFunctionNode(node *NonArgumentFunctionNode) {
	fmt.Fprintf(p.Writer, "*ast.NonArgumentFunctionNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Name: %q\n", node.Name)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitSpaceNode(node *SpaceNode) {
	fmt.Fprintf(p.Writer, "*ast.SpaceNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Value: %q\n", node.Value)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitDelimiterNode(node *DelimiterNode) {
	fmt.Fprintf(p.Writer, "*ast.DelimiterNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Value: %q\n", node.Value)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

// Visit methods for composite nodes
func (p *GoLikePrinter) VisitSuperscriptNode(node *SuperscriptNode) {
	fmt.Fprintf(p.Writer, "*ast.SuperscriptNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Base: ")
	node.Base.Accept(p)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Exponent: ")
	node.Exponent.Accept(p)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitSubscriptNode(node *SubscriptNode) {
	fmt.Fprintf(p.Writer, "*ast.SubscriptNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Base: ")
	node.Base.Accept(p)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Subscript: ")
	node.Subscript.Accept(p)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitFractionNode(node *FractionNode) {
	fmt.Fprintf(p.Writer, "*ast.FractionNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Numerator: ")
	node.Numerator.Accept(p)

	p.printIndent()
	fmt.Fprintf(p.Writer, "Denominator: ")
	node.Denominator.Accept(p)

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitIntegralNode(node *IntegralNode) {
	fmt.Fprintf(p.Writer, "*ast.IntegralNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	if node.LowerLimit != nil {
		fmt.Fprintf(p.Writer, "LowerLimit: ")
		node.LowerLimit.Accept(p)
	} else {
		fmt.Fprintf(p.Writer, "LowerLimit: nil\n")
	}

	p.printIndent()
	if node.UpperLimit != nil {
		fmt.Fprintf(p.Writer, "UpperLimit: ")
		node.UpperLimit.Accept(p)
	} else {
		fmt.Fprintf(p.Writer, "UpperLimit: nil\n")
	}

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}

func (p *GoLikePrinter) VisitSqrtNode(node *SqrtNode) {
	fmt.Fprintf(p.Writer, "*ast.SqrtNode {\n")
	p.increaseDepth()

	p.printIndent()
	fmt.Fprintf(p.Writer, "Start: %d\n", node.Start)

	p.printIndent()
	if node.Index != nil {
		fmt.Fprintf(p.Writer, "Index: ")
		node.Index.Accept(p)
	} else {
		fmt.Fprintf(p.Writer, "Index: nil\n")
	}

	p.printIndent()
	if node.Radicand != nil {
		fmt.Fprintf(p.Writer, "Radicand: ")
		node.Radicand.Accept(p)
	} else {
		fmt.Fprintf(p.Writer, "Radicand: nil\n")
	}

	p.decreaseDepth()
	p.printIndent()
	fmt.Fprintf(p.Writer, "}\n")
}
