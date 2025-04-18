package ast

// Visitor defines the interface for visiting AST nodes
type Visitor interface {
	// Visit methods for container nodes
	VisitExpressionNode(node *ExpressionNode)

	// Visit methods for leaf nodes
	VisitSymbolNode(node *SymbolNode)
	VisitNumberNode(node *NumberNode)
	VisitOperatorNode(node *OperatorNode)
	VisitNonArgumentFunctionNode(node *NonArgumentFunctionNode)
	VisitSpaceNode(node *SpaceNode)
	VisitDelimiterNode(node *DelimiterNode)

	// Visit methods for composite nodes
	VisitSuperscriptNode(node *SuperscriptNode)
	VisitSubscriptNode(node *SubscriptNode)
	VisitFractionNode(node *FractionNode)
	VisitIntegralNode(node *IntegralNode)
	VisitSqrtNode(node *SqrtNode)
}

// BaseVisitor provides default implementations for all Visitor methods.
// Embed this in your own visitors to avoid having to implement all methods.
type BaseVisitor struct{}

func (v *BaseVisitor) VisitExpressionNode(node *ExpressionNode) {
	for _, child := range node.Elements {
		child.Accept(v)
	}
}

func (v *BaseVisitor) VisitSymbolNode(node *SymbolNode)                           {}
func (v *BaseVisitor) VisitNumberNode(node *NumberNode)                           {}
func (v *BaseVisitor) VisitOperatorNode(node *OperatorNode)                       {}
func (v *BaseVisitor) VisitNonArgumentFunctionNode(node *NonArgumentFunctionNode) {}
func (v *BaseVisitor) VisitSpaceNode(node *SpaceNode)                             {}
func (v *BaseVisitor) VisitDelimiterNode(node *DelimiterNode)                     {}

func (v *BaseVisitor) VisitSuperscriptNode(node *SuperscriptNode) {
	node.Base.Accept(v)
	node.Exponent.Accept(v)
}

func (v *BaseVisitor) VisitSubscriptNode(node *SubscriptNode) {
	node.Base.Accept(v)
	node.Subscript.Accept(v)
}

func (v *BaseVisitor) VisitFractionNode(node *FractionNode) {
	node.Numerator.Accept(v)
	node.Denominator.Accept(v)
}

func (v *BaseVisitor) VisitIntegralNode(node *IntegralNode) {
	if node.LowerLimit != nil {
		node.LowerLimit.Accept(v)
	}
	if node.UpperLimit != nil {
		node.UpperLimit.Accept(v)
	}
}

func (v *BaseVisitor) VisitSqrtNode(node *SqrtNode) {
	if node.Index != nil {
		node.Index.Accept(v)
	}
	if node.Radicand != nil {
		node.Radicand.Accept(v)
	}
}
