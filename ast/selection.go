package ast

type SelectionSet []Selection

type Selection interface {
	isSelection()
	GetPosition() *Position
}

func (*Field) isSelection()          { _ = "STUB: not implemented"; return }
func (*FragmentSpread) isSelection() { _ = "STUB: not implemented"; return }
func (*InlineFragment) isSelection() { _ = "STUB: not implemented"; return }

func (f *Field) GetPosition() *Position          { _ = "STUB: not implemented"; return nil }
func (s *FragmentSpread) GetPosition() *Position { _ = "STUB: not implemented"; return nil }
func (f *InlineFragment) GetPosition() *Position { _ = "STUB: not implemented"; return nil }

type Field struct {
	Alias        string
	Name         string
	Arguments    ArgumentList
	Directives   DirectiveList
	SelectionSet SelectionSet
	Position     *Position `dump:"-" json:"-"`
	Comment      *CommentGroup

	// Require validation
	Definition       *FieldDefinition
	ObjectDefinition *Definition
}

type Argument struct {
	Name     string
	Value    *Value
	Position *Position `dump:"-" json:"-"`
	Comment  *CommentGroup
}

func (f *Field) ArgumentMap(vars map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
