package parser

import (
	"fmt"

	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

// parseMatrix parses a LaTeX matrix environment
func (p *Parser) parseMatrix(pos tokenizer.Position) *ast.MatrixNode {
	matrix := &ast.MatrixNode{
		Start: pos,
		Rows:  [][]ast.Node{},
	}

	// Initialize variables
	row := []ast.Node{}
	colCount := -1
	j := 0

	// Helper function to check column consistency
	checkRowLength := func(row []ast.Node, j int) {
		if colCount == -1 {
			colCount = len(row)
			return
		}

		if len(row) == colCount {
			return
		}

		if len(row) > colCount {
			p.addError(fmt.Sprintf("row %d has too many columns (%d instead of %d)", j+1, len(row), colCount), p.peek().Pos)
			return
		}

		p.addError(fmt.Sprintf("row %d has too few columns (%d instead of %d)", j+1, len(row), colCount), p.peek().Pos)
	}

	for {
		if p.peek().Type == tokenizer.COMMAND && p.peek().Value == "end" {
			break
		}

		// Parse cell content
		cell := p.parseExpression(func(t tokenizer.Token) bool {
			return t.Type == tokenizer.AMPERSAND || t.Type == tokenizer.BACKSLASH
		})

		row = append(row, cell)

		// Handle column separator
		if p.peek().Type == tokenizer.AMPERSAND {
			p.next() // consume &
			continue
		}

		// Handle row separator
		if p.peek().Type == tokenizer.BACKSLASH {
			p.next() // consume \\

			// Check if this is an invalid row separator before \end
			if p.peek().Type == tokenizer.COMMAND && p.peek().Value == "end" {
				p.addError("invalid row separator (\\) before end of matrix", p.peek().Pos)
			}

			checkRowLength(row, j)
			matrix.Rows = append(matrix.Rows, row)
			row = []ast.Node{} // Start new row
			j++
			continue
		}
	}

	// Add the last row if not empty
	if len(row) > 0 {
		checkRowLength(row, j)
		matrix.Rows = append(matrix.Rows, row)
	}

	return matrix
}
