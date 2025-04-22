package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

// parseEnvironment parses LaTeX environments like \begin{env}...\end{env}
// The \begin command token has already been consumed by parseCommand
func (p *Parser) parseEnvironment(pos int) ast.Node {

	// Internal helper function to extract environment name from ExpressionNode.
	// It concatenates all symbol tokens to form the complete environment name.
	extractEnv := func(node ast.Node) string {
		if node == nil {
			return ""
		}

		if expr, ok := node.(*ast.ExpressionNode); ok {
			var name string
			for _, element := range expr.Elements {
				if symbol, ok := element.(*ast.SymbolNode); ok {
					name += symbol.Value
				}
			}
			return name
		}
		return ""
	}

	// Parse environment name using existing parseGroupedStrict function
	envNode := p.parseGroupedStrict()

	// Extract environment name from the node
	env := extractEnv(envNode)
	if env == "" {
		p.addError("empty environment name", pos)
		return nil
	}

	// Dispatch based on environment type
	var node ast.Node
	switch env {
	case "matrix":
		// Parse matrix content in a separate file
		node = p.parseMatrix(pos)
	default:
		p.addError("unsupported environment: "+env, pos)
		return nil
	}

	// Check for \end command
	if p.peek().Type != tokenizer.COMMAND || p.peek().Value != "end" {
		p.addError("expected \\end command", p.peek().Pos)
		return node // Return what we've parsed so far
	}
	p.next() // consume \end

	// Check that environment name matches using parseGroupedStrict
	endEnvNode := p.parseGroupedStrict()

	// Extract end environment name
	endEnv := extractEnv(endEnvNode)
	if endEnv != env {
		p.addError("environment name mismatch: expected \\end{"+env+"}, got \\end{"+endEnv+"}", p.peek().Pos)
	}

	return node
}
