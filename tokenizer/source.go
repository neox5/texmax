package tokenizer

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Position represents a position in the source code
type Position struct {
	Offset int // Byte offset in the input
	Line   int // Line number (1-indexed)
	Column int // Column number (1-indexed)
}

// Source provides input characters with position tracking
type Source struct {
	reader *bufio.Reader
	file   string   // Optional filename (moved from Position)
	pos    Position // Current position in the source
	ch     rune     // Current rune
	err    error    // Last error encountered

	// Peek state
	nextCh  rune     // Next character (if peeked)
	nextPos Position // Position of next character
	nextErr error    // Error from peek operation
	peeked  bool     // Whether we have peeked
}

// NewStringSource creates a new source from a string
func NewStringSource(content string) *Source {
	return &Source{
		reader: bufio.NewReader(strings.NewReader(content)),
		pos:    Position{Offset: 0, Line: 1, Column: 1},
	}
}

// NewFileSource creates a new source from a file
func NewFileSource(filename string) (*Source, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &Source{
		reader: bufio.NewReader(file),
		file:   filename,
		pos:    Position{Offset: 0, Line: 1, Column: 1},
	}, nil
}

// readRune reads the next rune and updates the state
func (s *Source) readRune() {
	if s.peeked {
		s.ch = s.nextCh
		s.err = s.nextErr
		s.pos = s.nextPos
		s.peeked = false
		return
	}

	ch, size, err := s.reader.ReadRune()
	s.ch = ch
	s.err = err

	if err == nil {
		// Update position
		if ch == '\n' {
			s.pos.Line++
			s.pos.Column = 1
		} else {
			s.pos.Column++
		}
		s.pos.Offset += size
	}
}

// Scan advances to the next rune and returns it
// Returns EOF (-1) at the end of the input
func (s *Source) Scan() rune {
	s.readRune()

	if s.err != nil {
		if s.err == io.EOF {
			return -1 // EOF marker
		}
		return 0 // Null character for other errors
	}

	return s.ch
}

// Peek returns the next rune without advancing
func (s *Source) Peek() rune {
	if s.peeked {
		return s.nextCh
	}

	// Save current state
	savedCh := s.ch
	savedPos := s.pos
	savedErr := s.err

	// Perform a scan
	s.readRune()

	// Store peeked state
	s.nextCh = s.ch
	s.nextPos = s.pos
	s.nextErr = s.err
	s.peeked = true

	// Restore original state
	s.ch = savedCh
	s.pos = savedPos
	s.err = savedErr

	if s.nextErr != nil && s.nextErr == io.EOF {
		return -1 // EOF marker
	}

	return s.nextCh
}

// File returns the filename associated with this source
func (s *Source) File() string {
	return s.file
}

// Position returns the position of the current rune
func (s *Source) Position() Position {
	return s.pos
}

// Char returns the current rune
// Note: Before the first Scan(), this will return zero value (0)
func (s *Source) Char() rune {
	return s.ch
}

// Err returns the last error that occurred
func (s *Source) Err() error {
	return s.err
}

// IsEOF returns true if the scanner has reached the end of input
func (s *Source) IsEOF() bool {
	return s.err == io.EOF
}
