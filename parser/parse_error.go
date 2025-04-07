package parser

import "fmt"

type ParseError struct {
	Message string
	Pos     int
}

func (e ParseError) String() string {
	return fmt.Sprintf("%s at position %d", e.Message, e.Pos)
}

func (p *Parser) addError(msg string, pos int) {
	p.errors = append(p.errors, ParseError{msg, pos})
}
