package gqlerror

import (
	"github.com/vektah/gqlparser/v2/ast"
)

// Error is the standard graphql error type described in https://spec.graphql.org/draft/#sec-Errors
type Error struct {
	Err        error          `json:"-"`
	Message    string         `json:"message"`
	Path       ast.Path       `json:"path,omitempty"`
	Locations  []Location     `json:"locations,omitempty"`
	Extensions map[string]any `json:"extensions,omitempty"`
	Rule       string         `json:"-"`
}

func (err *Error) SetFile(file string) { _ = "STUB: not implemented"; return }

type Location struct {
	Line   int `json:"line,omitempty"`
	Column int `json:"column,omitempty"`
}

type List []*Error

func (err *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (err *Error) pathString() string { _ = "STUB: not implemented"; return "" }

func (err *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (err *Error) AsError() error { _ = "STUB: not implemented"; return nil }

func (errs List) Error() string { _ = "STUB: not implemented"; return "" }

func (errs List) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (errs List) As(target any) bool { _ = "STUB: not implemented"; return false }

func (errs List) Unwrap() []error { _ = "STUB: not implemented"; return nil }

func WrapPath(path ast.Path, err error) *Error { _ = "STUB: not implemented"; return nil }

func Wrap(err error) *Error { _ = "STUB: not implemented"; return nil }

func WrapIfUnwrapped(err error) *Error { _ = "STUB: not implemented"; return nil }

func Errorf(message string, args ...any) *Error { _ = "STUB: not implemented"; return nil }

func ErrorPathf(path ast.Path, message string, args ...any) *Error {
	_ = "STUB: not implemented"
	return nil
}

func ErrorPosf(pos *ast.Position, message string, args ...any) *Error {
	_ = "STUB: not implemented"
	return nil
}

func ErrorLocf(file string, line, col int, message string, args ...any) *Error {
	_ = "STUB: not implemented"
	return nil
}
