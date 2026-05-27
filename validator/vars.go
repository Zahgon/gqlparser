package validator

import (
	"errors"
	"reflect"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

//nolint:staticcheck // We do not care about capitalized error strings
var ErrUnexpectedType = errors.New("Unexpected Type")

// VariableValues coerces and validates variable values.
func VariableValues(
	schema *ast.Schema,
	op *ast.OperationDefinition,
	variables map[string]any,
) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type varValidator struct {
	path   ast.Path
	schema *ast.Schema
}

func (v *varValidator) validateVarType(
	typ *ast.Type,
	val reflect.Value,
) (reflect.Value, *gqlerror.Error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// GraphQL spec says that non-null values should be coerced to an array when possible.
// Hence if the value is not a slice, we create a slice and add val to it.

// If the type is not null and we got a invalid value namely null/nil, then it's valid

// assume custom scalars are ok

// check for unknown fields

// allow null object field and skip it

func IsValidIntString(val reflect.Value, kind reflect.Kind) bool {
	_ = "STUB: not implemented"
	return false
}

func IsValidFloatString(val reflect.Value, kind reflect.Kind) bool {
	_ = "STUB: not implemented"
	return false
}
