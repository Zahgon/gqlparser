package rules

import (
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

func ruleFuncKnownArgumentNames(observers *Events, addError AddErrFunc, disableSuggestion bool) {
	_ = "STUB: not implemented"
	// A GraphQL field is only valid if all supplied arguments are defined by that field.
	return
}

var KnownArgumentNamesRule = Rule{
	Name: "KnownArgumentNames",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncKnownArgumentNames(observers, addError, false)
	},
}

var KnownArgumentNamesRuleWithoutSuggestions = Rule{
	Name: "KnownArgumentNamesWithoutSuggestions",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncKnownArgumentNames(observers, addError, true)
	},
}
