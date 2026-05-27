package parser

import (
	. "github.com/vektah/gqlparser/v2/ast" //nolint:staticcheck // bad, yeah
)

func ParseQuery(source *Source) (*QueryDocument, error) { _ = "STUB: not implemented"; return nil, nil }

// 0 means unlimited

func ParseQueryWithTokenLimit(source *Source, maxTokenLimit int) (*QueryDocument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseQueryDocument() *QueryDocument { _ = "STUB: not implemented"; return nil }

func (p *parser) parseOperationDefinition() *OperationDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseOperationType() Operation { _ = "STUB: not implemented"; return *new(Operation) }

func (p *parser) parseVariableDefinitions() VariableDefinitionList {
	_ = "STUB: not implemented"
	return *new(VariableDefinitionList)
}

func (p *parser) parseVariableDefinition() *VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseVariable() string { _ = "STUB: not implemented"; return "" }

func (p *parser) parseOptionalSelectionSet() SelectionSet {
	_ = "STUB: not implemented"
	return *new(SelectionSet)
}

func (p *parser) parseRequiredSelectionSet() SelectionSet {
	_ = "STUB: not implemented"
	return *new(SelectionSet)
}

func (p *parser) parseSelection() Selection { _ = "STUB: not implemented"; return *new(Selection) }

func (p *parser) parseField() *Field { _ = "STUB: not implemented"; return nil }

func (p *parser) parseArguments(isConst bool) ArgumentList {
	_ = "STUB: not implemented"
	return *new(ArgumentList)
}

func (p *parser) parseArgument(isConst bool) *Argument { _ = "STUB: not implemented"; return nil }

func (p *parser) parseFragment() Selection { _ = "STUB: not implemented"; return *new(Selection) }

// "on"

func (p *parser) parseFragmentDefinition() *FragmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseFragmentName() string { _ = "STUB: not implemented"; return "" }

func (p *parser) parseValueLiteral(isConst bool) *Value { _ = "STUB: not implemented"; return nil }

func (p *parser) parseList(isConst bool) *Value { _ = "STUB: not implemented"; return nil }

func (p *parser) parseObject(isConst bool) *Value { _ = "STUB: not implemented"; return nil }

func (p *parser) parseObjectField(isConst bool) *ChildValue { _ = "STUB: not implemented"; return nil }

func (p *parser) parseDirectives(isConst bool) []*Directive { _ = "STUB: not implemented"; return nil }

func (p *parser) parseDirective(isConst bool) *Directive { _ = "STUB: not implemented"; return nil }

func (p *parser) parseTypeReference() *Type { _ = "STUB: not implemented"; return nil }

func (p *parser) parseName() string { _ = "STUB: not implemented"; return "" }
