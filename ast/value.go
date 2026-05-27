package ast

type ValueKind int

const (
	Variable ValueKind = iota
	IntValue
	FloatValue
	StringValue
	BlockValue
	BooleanValue
	NullValue
	EnumValue
	ListValue
	ObjectValue
)

type Value struct {
	Raw      string
	Children ChildValueList
	Kind     ValueKind
	Position *Position `dump:"-" json:"-"`
	Comment  *CommentGroup

	// Require validation
	Definition             *Definition
	VariableDefinition     *VariableDefinition
	ExpectedType           *Type
	ExpectedTypeHasDefault bool
}

type ChildValue struct {
	Name     string
	Value    *Value
	Position *Position `dump:"-" json:"-"`
	Comment  *CommentGroup
}

func (v *Value) Value(vars map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (v *Value) String() string { _ = "STUB: not implemented"; return "" }

func (v *Value) Dump() string { _ = "STUB: not implemented"; return "" }
