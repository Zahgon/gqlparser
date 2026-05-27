package rules

import (
	"github.com/vektah/gqlparser/v2/ast"
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

func ruleFuncFieldsOnCorrectType(observers *Events, addError AddErrFunc, disableSuggestion bool) {
	_ = "STUB: not implemented"
	return
}

var FieldsOnCorrectTypeRule = Rule{
	Name: "FieldsOnCorrectType",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncFieldsOnCorrectType(observers, addError, false)
	},
}

var FieldsOnCorrectTypeRuleWithoutSuggestions = Rule{
	Name: "FieldsOnCorrectTypeWithoutSuggestions",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncFieldsOnCorrectType(observers, addError, true)
	},
}

// Go through all the implementations of type, as well as the interfaces
// that they implement. If any of those types include the provided field,
// suggest them, sorted by how often the type is referenced,  starting
// with Interfaces.
func getSuggestedTypeNames(walker *Walker, parent *ast.Definition, name string) []string {
	_ = "STUB: not implemented"
	return nil
}

// By employing a full slice expression (slice[low:high:max]),
// where max is set to the slice’s length,
// we ensure that appending elements results
// in a slice backed by a distinct array.
// This method prevents the shared array issue.
func concatSlice(first, second []string) []string { _ = "STUB: not implemented"; return nil }

// For the field name provided, determine if there are any similar field names
// that may be the result of a typo.
func getSuggestedFieldNames(parent *ast.Definition, name string) []string {
	_ = "STUB: not implemented"
	return nil
}
