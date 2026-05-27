package parser

import (
	. "github.com/vektah/gqlparser/v2/ast" //nolint:staticcheck // bad, yeah
)

func ParseSchemas(inputs ...*Source) (*SchemaDocument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseSchema(source *Source) (*SchemaDocument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default value is unlimited

func ParseSchemasWithLimit(maxTokenLimit int, inputs ...*Source) (*SchemaDocument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseSchemaWithLimit(source *Source, maxTokenLimit int) (*SchemaDocument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 0 is unlimited

func (p *parser) parseSchemaDocument() *SchemaDocument { _ = "STUB: not implemented"; return nil }

// treat end of file comments

func (p *parser) parseDescription() descriptionWithComment {
	_ = "STUB: not implemented"
	return *new(descriptionWithComment)
}

func (p *parser) parseTypeSystemDefinition(description descriptionWithComment) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseSchemaDefinition(description descriptionWithComment) *SchemaDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseOperationTypeDefinition() *OperationTypeDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseScalarTypeDefinition(description descriptionWithComment) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseObjectTypeDefinition(description descriptionWithComment) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseImplementsInterfaces() []string { _ = "STUB: not implemented"; return nil }

// optional leading ampersand

func (p *parser) parseFieldsDefinition() (FieldList, *CommentGroup) {
	_ = "STUB: not implemented"
	return *new(FieldList), nil
}

func (p *parser) parseFieldDefinition() *FieldDefinition { _ = "STUB: not implemented"; return nil }

// peek to set p.comment

func (p *parser) parseArgumentDefs() ArgumentDefinitionList {
	_ = "STUB: not implemented"
	return *new(ArgumentDefinitionList)
}

func (p *parser) parseArgumentDef() *ArgumentDefinition { _ = "STUB: not implemented"; return nil }

// peek to set p.comment

func (p *parser) parseInputValueDef() *FieldDefinition { _ = "STUB: not implemented"; return nil }

// peek to set p.comment

func (p *parser) parseInterfaceTypeDefinition(description descriptionWithComment) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseUnionTypeDefinition(description descriptionWithComment) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseUnionMemberTypes() []string { _ = "STUB: not implemented"; return nil }

// optional leading pipe

func (p *parser) parseEnumTypeDefinition(description descriptionWithComment) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseEnumValuesDefinition() (EnumValueList, *CommentGroup) {
	_ = "STUB: not implemented"
	return *new(EnumValueList), nil
}

func (p *parser) parseEnumValueDefinition() *EnumValueDefinition {
	_ = "STUB: not implemented"
	return nil
}

// peek to set p.comment

func (p *parser) parseInputObjectTypeDefinition(description descriptionWithComment) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseInputFieldsDefinition() (FieldList, *CommentGroup) {
	_ = "STUB: not implemented"
	return *new(FieldList), nil
}

func (p *parser) parseTypeSystemExtension(doc *SchemaDocument) { _ = "STUB: not implemented"; return }

func (p *parser) parseSchemaExtension(comment *CommentGroup) *SchemaDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseScalarTypeExtension(comment *CommentGroup) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseObjectTypeExtension(comment *CommentGroup) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseInterfaceTypeExtension(comment *CommentGroup) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseUnionTypeExtension(comment *CommentGroup) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseEnumTypeExtension(comment *CommentGroup) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseInputObjectTypeExtension(comment *CommentGroup) *Definition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseDirectiveDefinition(description descriptionWithComment) *DirectiveDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseDirectiveLocations() []DirectiveLocation {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseDirectiveLocation() DirectiveLocation {
	_ = "STUB: not implemented"
	return *new(DirectiveLocation)
}

type descriptionWithComment struct {
	text    string
	comment *CommentGroup
}
