package validator

import (
	. "github.com/vektah/gqlparser/v2/ast" //nolint:staticcheck // bad, yeah
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func LoadSchema(inputs ...*Source) (*Schema, error) { _ = "STUB: not implemented"; return nil, nil }

func ValidateSchemaDocument(sd *SchemaDocument) (*Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// While the spec says SDL must not (§3.5) explicitly define builtin
// scalars, it may (§3.13) define builtin directives. Here we check for
// that, and reject doubly-defined directives otherwise.

// the builtins
// In principle here we might want to validate that the
// directives are the same. But they might not be, if the
// server has an older spec than we do. (Plus, validating this
// is a lot of work.) So we just keep the first one we saw.
// That's an arbitrary choice, but in theory the only way it
// fails is if the server is using features newer than this
// version of gqlparser, in which case they're in trouble
// anyway.

// Inferred root operation type names should be performed only when a `schema` directive is
// **not** provided, when it is, `Mutation` and `Subscription` becomes valid types and are not
// assigned as a root operation on the schema.

func validateTypeDefinitions(schema *Schema) *gqlerror.Error { _ = "STUB: not implemented"; return nil }

func validateDirectiveDefinitions(schema *Schema) *gqlerror.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateDirective(schema *Schema, def *DirectiveDefinition) *gqlerror.Error {
	_ = "STUB: not implemented"
	return nil
}

// now, GraphQL spec doesn't have reserved directive name

func validateDefinition(schema *Schema, def *Definition) *gqlerror.Error {
	_ = "STUB: not implemented"
	return nil
}

// now, GraphQL spec doesn't have reserved field name

// GraphQL spec has reserved type names a lot!

func validateTypeRef(schema *Schema, typ *Type) *gqlerror.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateArgs(
	schema *Schema,
	args ArgumentDefinitionList,
	currentDirective *DirectiveDefinition,
) *gqlerror.Error {
	_ = "STUB: not implemented"
	return nil
}

// now, GraphQL spec doesn't have reserved argument name

func validateDirectives(
	schema *Schema,
	dirs DirectiveList,
	location DirectiveLocation,
	currentDirective *DirectiveDefinition,
) *gqlerror.Error {
	_ = "STUB: not implemented"
	return nil
}

// now, GraphQL spec doesn't have reserved directive name

func validateImplements(schema *Schema, def *Definition, intfName string) *gqlerror.Error {
	_ = "STUB: not implemented"
	// see validation rules at the bottom of
	// https://spec.graphql.org/October2021/#sec-Objects
	return nil
}

// validateTypeImplementsAncestors
// https://github.com/graphql/graphql-js/blob/47bd8c8897c72d3efc17ecb1599a95cee6bac5e8/src/type/validate.ts#L428
func validateTypeImplementsAncestors(
	schema *Schema,
	def *Definition,
	intfName string,
) *gqlerror.Error {
	_ = "STUB: not implemented"
	return nil
}

func containsString(slice []string, want string) bool { _ = "STUB: not implemented"; return false }

func isCovariant(schema *Schema, required, actual *Type) bool {
	_ = "STUB: not implemented"
	return false
}

func validateName(pos *Position, name string) *gqlerror.Error {
	_ = "STUB: not implemented"
	return nil
}

func isValidKind(kind DefinitionKind, valid ...DefinitionKind) bool {
	_ = "STUB: not implemented"
	return false
}

func kindList(kinds ...DefinitionKind) string { _ = "STUB: not implemented"; return "" }
