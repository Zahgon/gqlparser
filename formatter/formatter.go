package formatter

import (
	"io"

	"github.com/vektah/gqlparser/v2/ast"
)

type Formatter interface {
	FormatSchema(schema *ast.Schema)
	FormatSchemaDocument(doc *ast.SchemaDocument)
	FormatQueryDocument(doc *ast.QueryDocument)
}

//nolint:revive // Ignore "stuttering" name formatter.FormatterOption
type FormatterOption func(*formatter)

// WithIndent uses the given string for indenting block bodies in the output,
// instead of the default, `"\t"`.
func WithIndent(indent string) FormatterOption {
	_ = "STUB: not implemented"
	return *new(FormatterOption)
}

// WithComments includes comments from the source/AST in the formatted output.
func WithComments() FormatterOption { _ = "STUB: not implemented"; return *new(FormatterOption) }

// WithBuiltin includes builtin fields/directives/etc from the source/AST in the formatted output.
func WithBuiltin() FormatterOption { _ = "STUB: not implemented"; return *new(FormatterOption) }

// WithNonIntrospectionBuiltin includes builtin fields/directives/etc from the
// source/AST in the formatted output, but excludes the introspection types and
// fields that starts with "__".
func WithNonIntrospectionBuiltin() FormatterOption {
	_ = "STUB: not implemented"
	return *new(FormatterOption)
}

// WithoutDescription excludes GQL description from the source/AST in the formatted output.
func WithoutDescription() FormatterOption { _ = "STUB: not implemented"; return *new(FormatterOption) }

// WithCompacted enables compacted output, which removes all unnecessary whitespace.
func WithCompacted() FormatterOption { _ = "STUB: not implemented"; return *new(FormatterOption) }

func NewFormatter(w io.Writer, options ...FormatterOption) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}

type formatter struct {
	writer io.Writer

	indent            string
	indentSize        int
	emitBuiltin       bool
	emitIntrospection bool
	emitComments      bool
	omitDescription   bool
	compacted         bool

	padNext  bool
	lineHead bool
}

func (f *formatter) writeString(s string) { _ = "STUB: not implemented"; return }

func (f *formatter) writeIndent() *formatter { _ = "STUB: not implemented"; return nil }

func (f *formatter) WriteNewline() *formatter { _ = "STUB: not implemented"; return nil }

func (f *formatter) WriteWord(word string) *formatter { _ = "STUB: not implemented"; return nil }

func (f *formatter) WriteString(s string) *formatter { _ = "STUB: not implemented"; return nil }

func (f *formatter) WriteDescription(s string) *formatter { _ = "STUB: not implemented"; return nil }

func (f *formatter) IncrementIndent() { _ = "STUB: not implemented"; return }

func (f *formatter) DecrementIndent() { _ = "STUB: not implemented"; return }

func (f *formatter) NoPadding() *formatter { _ = "STUB: not implemented"; return nil }

func (f *formatter) NeedPadding() *formatter { _ = "STUB: not implemented"; return nil }

func (f *formatter) FormatSchema(schema *ast.Schema) { _ = "STUB: not implemented"; return }

// Schema definition is omitted from output, but it has
// directives. Output them as the schema extension to not loose
// them

func (f *formatter) FormatSchemaDocument(doc *ast.SchemaDocument) {
	_ = "STUB: not implemented"
	// TODO emit by position based order
	return
}

// doc.Comment is end of file comment, so emit last

func (f *formatter) FormatQueryDocument(doc *ast.QueryDocument) {
	_ = "STUB: not implemented"
	// TODO emit by position based order
	return
}

func (f *formatter) FormatSchemaDefinitionList(lists ast.SchemaDefinitionList, extension bool) {
	_ = "STUB: not implemented"
	return
}

// Don't output empty schema definition block for extensions

// Return true if schema definitions is empty (besides directives), false otherwise
func (f *formatter) IsSchemaDefinitionsEmpty(lists ast.SchemaDefinitionList) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *formatter) FormatSchemaDefinition(def *ast.SchemaDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatOperationTypeDefinitionList(lists ast.OperationTypeDefinitionList) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatOperationTypeDefinition(def *ast.OperationTypeDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatFieldList(fieldList ast.FieldList, endOfDefComment *ast.CommentGroup) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatFieldDefinition(field *ast.FieldDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatArgumentDefinitionList(lists ast.ArgumentDefinitionList) {
	_ = "STUB: not implemented"
	return
}

// Skip emitting (insignificant) comma in case it is the
// last argument, or we printed a new line in its definition.

func (f *formatter) FormatArgumentDefinition(def *ast.ArgumentDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatDirectiveLocation(location ast.DirectiveLocation) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatDirectiveDefinitionList(lists ast.DirectiveDefinitionList) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatDirectiveDefinition(def *ast.DirectiveDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatDefinitionList(lists ast.DefinitionList, extend bool) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatDefinition(def *ast.Definition, extend bool) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatEnumValueList(
	lists ast.EnumValueList,
	endOfDefComment *ast.CommentGroup,
) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatEnumValueDefinition(def *ast.EnumValueDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatOperationList(lists ast.OperationList) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatOperationDefinition(def *ast.OperationDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatDirectiveList(lists ast.DirectiveList) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatDirective(dir *ast.Directive) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatArgumentList(lists ast.ArgumentList) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatArgument(arg *ast.Argument) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatFragmentDefinitionList(lists ast.FragmentDefinitionList) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatFragmentDefinition(def *ast.FragmentDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatVariableDefinitionList(lists ast.VariableDefinitionList) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatVariableDefinition(def *ast.VariableDefinition) {
	_ = "STUB: not implemented"
	return
}

// TODO https://github.com/vektah/gqlparser/v2/issues/102
//   VariableDefinition : Variable : Type DefaultValue? Directives[Const]?

func (f *formatter) FormatSelectionSet(sets ast.SelectionSet) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatSelection(selection ast.Selection) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatField(field *ast.Field) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatFragmentSpread(spread *ast.FragmentSpread) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatInlineFragment(inline *ast.InlineFragment) {
	_ = "STUB: not implemented"
	return
}

func (f *formatter) FormatType(t *ast.Type) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatValue(value *ast.Value) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatCommentGroup(group *ast.CommentGroup) { _ = "STUB: not implemented"; return }

func (f *formatter) FormatComment(comment *ast.Comment) { _ = "STUB: not implemented"; return }
