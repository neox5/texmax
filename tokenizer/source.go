package tokenizer

import (
	"bufio"
	"os"
	"strings"
)

// Position represents a position in the source code
type Position struct {
	Offset int    // Byte offset in the input
	Line   int    // Line number (1-indexed)
	Column int    // Column number (1-indexed)
	File   string // Optional filename
}

// Source provides input characters with position tracking
type Source struct {
	reader   *bufio.Reader
	position int
	line     int
	column   int
	filename string
}

// NewStringSource creates a new source from a string
func NewStringSource(content string) *Source {
	return &Source{
		reader:   bufio.NewReader(strings.NewReader(content)),
		position: 0,
		line:     1,
		column:   1,
	}
}

// NewFileSource creates a new source from a file
func NewFileSource(filename string) (*Source, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	
	return &Source{
		reader:   bufio.NewReader(file),
		position: 0,
		line:     1,
		column:   1,
		filename: filename,
	}, nil
}

// Next returns the next rune along with its position and any error
// io.EOF will be returned when the end of the source is reached
func (s *Source) Next() (rune, Position, error) {
	pos := Position{
		Offset: s.position,
		Line:   s.line,
		Column: s.column,
		File:   s.filename,
	}
	
	r, size, err := s.reader.ReadRune()
	if err != nil {
		return 0, pos, err // Return the error directly (including io.EOF)
	}
	
	// Update position
	s.position += size
	if r == '\n' {
		s.line++
		s.column = 1
	} else {
		s.column++
	}
	
	return r, pos, nil
}
