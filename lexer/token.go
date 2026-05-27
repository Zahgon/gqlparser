package lexer

import (
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	Invalid Type = iota
	EOF
	Bang
	Dollar
	Amp
	ParenL
	ParenR
	Spread
	Colon
	Equals
	At
	BracketL
	BracketR
	BraceL
	BraceR
	Pipe
	Name
	Int
	Float
	String
	BlockString
	Comment
)

func (t Type) Name() string { _ = "STUB: not implemented"; return "" }

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

// Kind represents a type of token. The types are predefined as constants.
type Type int

type Token struct {
	Kind  Type         // The token type.
	Value string       // The literal value consumed.
	Pos   ast.Position // The file and line this token was read from
}

func (t Token) String() string { _ = "STUB: not implemented"; return "" }
