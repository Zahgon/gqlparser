package ast

import (
	"encoding/json"
)

var _ json.Unmarshaler = (*Path)(nil)

type Path []PathElement

type PathElement interface {
	isPathElement()
}

var (
	_ PathElement = PathIndex(0)
	_ PathElement = PathName("")
)

func (path Path) String() string { _ = "STUB: not implemented"; return "" }

func (path *Path) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type PathIndex int

func (PathIndex) isPathElement() { _ = "STUB: not implemented"; return }

type PathName string

func (PathName) isPathElement() { _ = "STUB: not implemented"; return }
