package ast

import (
	"bytes"
	"reflect"
)

// Dump turns ast into a stable string format for assertions in tests.
func Dump(i any) string { _ = "STUB: not implemented"; return "" }

type dumper struct {
	*bytes.Buffer
	indent int
}

type Dumpable interface {
	Dump() string
}

func (d *dumper) dump(v reflect.Value) { _ = "STUB: not implemented"; return }

func (d *dumper) writeIndent() { _ = "STUB: not implemented"; return }

func (d *dumper) nl() { _ = "STUB: not implemented"; return }

func typeName(t reflect.Type) string { _ = "STUB: not implemented"; return "" }

func (d *dumper) dumpArray(v reflect.Value) { _ = "STUB: not implemented"; return }

func (d *dumper) dumpStruct(v reflect.Value) { _ = "STUB: not implemented"; return }

func isZero(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// Never consider Bool field as zero value.
// Always include them in AST dump.

func (d *dumper) dumpPtr(v reflect.Value) { _ = "STUB: not implemented"; return }
