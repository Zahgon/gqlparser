package rules

import (
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

func ruleFuncKnownTypeNames(observers *Events, addError AddErrFunc, disableSuggestion bool) {
	_ = "STUB: not implemented"
	return
}

var KnownTypeNamesRule = Rule{
	Name: "KnownTypeNames",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncKnownTypeNames(observers, addError, false)
	},
}

var KnownTypeNamesRuleWithoutSuggestions = Rule{
	Name: "KnownTypeNamesWithoutSuggestions",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncKnownTypeNames(observers, addError, true)
	},
}
