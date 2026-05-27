package ast

func UnmarshalSelectionSet(b []byte) (SelectionSet, error) {
	_ = "STUB: not implemented"
	return *new(SelectionSet), nil
}

func (f *FragmentDefinition) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (f *InlineFragment) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (f *OperationDefinition) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (f *Field) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
