package lexer

import (
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Lexer turns graphql request and schema strings into tokens.
type Lexer struct {
	*ast.Source
	// An offset into the string in bytes
	start int
	// An offset into the string in runes
	startRunes int
	// An offset into the string in bytes
	end int
	// An offset into the string in runes
	endRunes int
	// the current line number
	line int
	// An offset into the string in rune
	lineStartRunes int
}

func New(src *ast.Source) Lexer { _ = "STUB: not implemented"; return *new(Lexer) }

// take one rune from input and advance end.
func (s *Lexer) peek() (rune, int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *Lexer) makeToken(kind Type) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

func (s *Lexer) makeValueToken(kind Type, value string) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

func (s *Lexer) makeError(format string, args ...any) (Token, *gqlerror.Error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

// ReadToken gets the next token from the source starting at the given position.
//
// This skips over whitespace and comments until it finds the next lexable
// token, then lexes punctuators immediately or calls the appropriate helper
// function for more complicated tokens.
func (s *Lexer) ReadToken() (Token, error) { _ = "STUB: not implemented"; return *new(Token), nil }

// ws reads from body starting at startPosition until it finds a non-whitespace
// or commented character, and updates the token end to include all whitespace.
func (s *Lexer) ws() { _ = "STUB: not implemented"; return }

// skip the following newline if its there

// byte order mark, given ws is hot path we aren't relying on the unicode package here.

// readComment from the input
//
// #[\u0009\u0020-\uFFFF]*.
func (s *Lexer) readComment() (Token, error) { _ = "STUB: not implemented"; return *new(Token), nil }

// SourceCharacter but not LineTerminator

// readNumber from the input, either a float
// or an int depending on whether a decimal point appears.
//
// Int:   -?(0|[1-9][0-9]*)
// Float: -?(0|[1-9][0-9]*)(\.[0-9]+)?((E|e)(+|-)?[0-9]+)?
func (s *Lexer) readNumber() (Token, error) {
	_ = "STUB: not implemented"

	// backup to the first digit
	return *new(Token), nil
}

// acceptByte if it matches any of given bytes, returning true if it found anything.
func (s *Lexer) acceptByte(bytes ...uint8) bool { _ = "STUB: not implemented"; return false }

// acceptDigits from the input, returning the number of digits it found.
func (s *Lexer) acceptDigits() int { _ = "STUB: not implemented"; return 0 }

// describeNext peeks at the input and returns a human readable string. This should will alloc
// and should only be used in errors.
func (s *Lexer) describeNext() string { _ = "STUB: not implemented"; return "" }

// readString from the input
//
// "([^"\\\u000A\u000D]|(\\(u[0-9a-fA-F]{4}|["\\/bfnrt])))*".
func (s *Lexer) readString() (Token, error) {
	_ = "STUB: not implemented"
	return *

	// this buffer is lazily created only if there are escape characters.
	new(Token), nil
}

// skip the opening quote

// skip unicode overhead if we are in the ascii range

// the token should not include the quotes in its value, but should cover them in its
// position

// skip the close quote

// readBlockString from the input
//
// """("?"?(\\"""|\\(?!=""")|[^"\\]))*""".
func (s *Lexer) readBlockString() (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

// skip the opening quote

// Closing triple quote (""")

// Count consecutive quotes

// If we have at least 3 quotes, use the last 3 as the closing quote

// Add any extra quotes to the buffer (except the last 3)

// SourceCharacter

// skip unicode overhead if we are in the ascii range

func unhex(b string) (v rune, ok bool) { _ = "STUB: not implemented"; return 0, false }

// readName from the input
//
// [_A-Za-z][_0-9A-Za-z]*.
func (s *Lexer) readName() (Token, error) { _ = "STUB: not implemented"; return *new(Token), nil }
