package rules

import (
	"github.com/vektah/gqlparser/v2/ast"
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

func ruleFuncValuesOfCorrectType(observers *Events, addError AddErrFunc, disableSuggestion bool) {
	_ = "STUB: not implemented"
	return
}

// Skip custom validating scalars

var ValuesOfCorrectTypeRule = Rule{
	Name: "ValuesOfCorrectType",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncValuesOfCorrectType(observers, addError, false)
	},
}

var ValuesOfCorrectTypeRuleWithoutSuggestions = Rule{
	Name: "ValuesOfCorrectTypeWithoutSuggestions",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncValuesOfCorrectType(observers, addError, true)
	},
}

func unexpectedTypeMessage(addError AddErrFunc, v *ast.Value) { _ = "STUB: not implemented"; return }

func unexpectedTypeMessageOnly(v *ast.Value) ErrorOption {
	_ = "STUB: not implemented"
	return *new(ErrorOption)
}

// case "Enum":
// 		return Message(`Enum "%s" cannot represent non-enum value: %s`, v.ExpectedType.String(),
// v.String())
