package parser

import (
	"fmt"

	"github.com/neox5/texmax/tokenizer"
)

type ParseError struct {
	Message string
	Pos     int
	Line    int
	Column  int
}

func (e ParseError) String() string {
	return fmt.Sprintf("%s at line %d, column %d", e.Message, e.Line, e.Column)
}

func (p *Parser) addError(msg string, pos tokenizer.Position) {
	p.errors = append(p.errors, ParseError{
		Message: msg,
		Pos:     pos.Offset,
		Line:    pos.Line,
		Column:  pos.Column,
	})
}
