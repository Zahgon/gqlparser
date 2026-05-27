package ast

type FieldList []*FieldDefinition

func (l FieldList) ForName(name string) *FieldDefinition { _ = "STUB: not implemented"; return nil }

type EnumValueList []*EnumValueDefinition

func (l EnumValueList) ForName(name string) *EnumValueDefinition {
	_ = "STUB: not implemented"
	return nil
}

type DirectiveList []*Directive

func (l DirectiveList) ForName(name string) *Directive { _ = "STUB: not implemented"; return nil }

func (l DirectiveList) ForNames(name string) []*Directive { _ = "STUB: not implemented"; return nil }

type OperationList []*OperationDefinition

func (l OperationList) ForName(name string) *OperationDefinition {
	_ = "STUB: not implemented"
	return nil
}

type FragmentDefinitionList []*FragmentDefinition

func (l FragmentDefinitionList) ForName(name string) *FragmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

type VariableDefinitionList []*VariableDefinition

func (l VariableDefinitionList) ForName(name string) *VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

type ArgumentList []*Argument

func (l ArgumentList) ForName(name string) *Argument { _ = "STUB: not implemented"; return nil }

type ArgumentDefinitionList []*ArgumentDefinition

func (l ArgumentDefinitionList) ForName(name string) *ArgumentDefinition {
	_ = "STUB: not implemented"
	return nil
}

type SchemaDefinitionList []*SchemaDefinition

type DirectiveDefinitionList []*DirectiveDefinition

func (l DirectiveDefinitionList) ForName(name string) *DirectiveDefinition {
	_ = "STUB: not implemented"
	return nil
}

type DefinitionList []*Definition

func (l DefinitionList) ForName(name string) *Definition { _ = "STUB: not implemented"; return nil }

type OperationTypeDefinitionList []*OperationTypeDefinition

func (l OperationTypeDefinitionList) ForType(name string) *OperationTypeDefinition {
	_ = "STUB: not implemented"
	return nil
}

type ChildValueList []*ChildValue

func (v ChildValueList) ForName(name string) *Value { _ = "STUB: not implemented"; return nil }
