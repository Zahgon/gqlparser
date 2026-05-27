package rules

import (
	"github.com/vektah/gqlparser/v2/ast"
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

var UniqueArgumentNamesRule = Rule{
	Name: "UniqueArgumentNames",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		observers.OnField(func(walker *Walker, field *ast.Field) {
			checkUniqueArgs(field.Arguments, addError)
		})

		observers.OnDirective(func(walker *Walker, directive *ast.Directive) {
			checkUniqueArgs(directive.Arguments, addError)
		})
	},
}

func checkUniqueArgs(args ast.ArgumentList, addError AddErrFunc) { _ = "STUB: not implemented"; return }
