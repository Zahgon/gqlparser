package ast

func NonNullNamedType(named string, pos *Position) *Type { _ = "STUB: not implemented"; return nil }

func NamedType(named string, pos *Position) *Type { _ = "STUB: not implemented"; return nil }

func NonNullListType(elem *Type, pos *Position) *Type { _ = "STUB: not implemented"; return nil }

func ListType(elem *Type, pos *Position) *Type { _ = "STUB: not implemented"; return nil }

type Type struct {
	NamedType string
	Elem      *Type
	NonNull   bool
	Position  *Position `dump:"-" json:"-"`
}

func (t *Type) Name() string { _ = "STUB: not implemented"; return "" }

func (t *Type) String() string { _ = "STUB: not implemented"; return "" }

func (t *Type) IsCompatible(other *Type) bool { _ = "STUB: not implemented"; return false }

func (t *Type) Dump() string { _ = "STUB: not implemented"; return "" }
