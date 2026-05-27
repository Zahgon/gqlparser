package rules

import (
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

func ruleFuncScalarLeafs(observers *Events, addError AddErrFunc, disableSuggestion bool) {
	_ = "STUB: not implemented"
	return
}

var ScalarLeafsRule = Rule{
	Name: "ScalarLeafs",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncScalarLeafs(observers, addError, false)
	},
}

var ScalarLeafsRuleWithoutSuggestions = Rule{
	Name: "ScalarLeafsWithoutSuggestions",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		ruleFuncScalarLeafs(observers, addError, true)
	},
}
