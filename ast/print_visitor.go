package ast

import (
	"fmt"
	"io"
	"strings"
)

// PrintVisitor prints the AST with proper indentation
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

// Visit processes a node
func (v *PrintVisitor) Visit(node Node) {
	// Get the indentation string
	indent := v.indent()
	
	// Print the node with its prefix
	if v.Depth == 0 {
		// Root node - no prefix
		fmt.Fprintf(v.Writer, "%s\n", node.String())
	} else {
		// Child node - with pipe prefix
		fmt.Fprintf(v.Writer, "%s%s\n", indent, node.String())
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

// EnterNode is called when entering a container node
func (v *PrintVisitor) EnterNode(node Node) {
	v.Depth++
}

// ExitNode is called when exiting a container node
func (v *PrintVisitor) ExitNode(node Node) {
	v.Depth--
}
