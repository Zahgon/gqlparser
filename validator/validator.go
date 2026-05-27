package validator

import (
	"sort"

	//nolint:staticcheck // bad, yeah
	. "github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vektah/gqlparser/v2/validator/core"
	validatorrules "github.com/vektah/gqlparser/v2/validator/rules"
)

type (
	AddErrFunc  = core.AddErrFunc
	RuleFunc    = core.RuleFunc
	Rule        = core.Rule
	Events      = core.Events
	ErrorOption = core.ErrorOption
	Walker      = core.Walker
)

var (
	Message      = core.Message
	QuotedOrList = core.QuotedOrList
	OrList       = core.OrList
)

// Walk is an alias for core.Walk.
func Walk(schema *Schema, document *QueryDocument, observers *Events) {
	_ = "STUB: not implemented"
	return
}

var specifiedRules []Rule

func init() {
	// Initialize specifiedRules with default rules
	defaultRules := validatorrules.NewDefaultRules()
	for name, ruleFunc := range defaultRules.GetInner() {
		specifiedRules = append(specifiedRules, Rule{Name: name, RuleFunc: ruleFunc})
		// ensure initial default is in deterministic order
		sort.Sort(core.NameSorter(specifiedRules))
	}
}

// AddRule adds a rule to the rule set.
// ruleFunc is called once each time `Validate` is executed.
func AddRule(name string, ruleFunc RuleFunc) { _ = "STUB: not implemented"; return }

// RemoveRule removes an existing rule from the rule set
// if one of the same name exists.
// The rule set is global, so it is not safe for concurrent changes.
func RemoveRule(name string) {
	_ = "STUB: not implemented"
	//nolint:prealloc // using initialized with len(rules) produces a race condition
	return
}

// ReplaceRule replaces an existing rule from the rule set
// if one of the same name exists.
// If no match is found, it will add a new rule to the rule set.
// The rule set is global, so it is not safe for concurrent changes.
func ReplaceRule(name string, ruleFunc RuleFunc) { _ = "STUB: not implemented"; return }

//nolint:prealloc // using initialized with len(rules) produces a race condition

// Deprecated: use ValidateWithRules instead.
func Validate(schema *Schema, doc *QueryDocument, rules ...Rule) gqlerror.List {
	_ = "STUB: not implemented"
	return *new(gqlerror.List)
}

func ValidateWithRules(
	schema *Schema,
	doc *QueryDocument,
	rules *validatorrules.Rules,
) gqlerror.List {
	_ = "STUB: not implemented"
	return *new(gqlerror.List)
}

//nolint:prealloc // would require extra local refs for len

// ensure deterministic order evaluation
