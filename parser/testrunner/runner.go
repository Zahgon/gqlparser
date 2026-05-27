package testrunner

import (
	"testing"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Features map[string][]Spec

type Spec struct {
	Name   string
	Input  string
	Error  *gqlerror.Error
	Tokens []Token
	AST    string
}

type Token struct {
	Kind   string
	Value  string
	Start  int
	End    int
	Line   int
	Column int
	Src    string
}

func (t Token) String() string { _ = "STUB: not implemented"; return "" }

func Test(t *testing.T, filename string, f func(t *testing.T, input string) Spec) {
	_ = "STUB: not implemented"
	return
}
