package rules

import (
	"github.com/vektah/gqlparser/v2/validator/core"
)

// Rules manages GraphQL validation rules.
type Rules struct {
	rules        map[string]core.RuleFunc
	ruleNameKeys []string // for deterministic order
}

// NewRules creates a Rules instance with the specified rules.
func NewRules(rs ...core.Rule) *Rules { _ = "STUB: not implemented"; return nil }

// NewDefaultRules creates a Rules instance containing the default GraphQL validation rule set.
func NewDefaultRules() *Rules { _ = "STUB: not implemented"; return nil }

// AddRule adds a rule with the specified name and rule function to the rule set.
// If a rule with the same name already exists, it will not be added.
func (r *Rules) AddRule(name string, ruleFunc core.RuleFunc) { _ = "STUB: not implemented"; return }

// GetInner returns the internal rule map.
// If the map is not initialized, it returns an empty map.
// This returns a copy of the rules map, not the original map.
func (r *Rules) GetInner() map[string]core.RuleFunc { _ = "STUB: not implemented"; return nil }

// impossible nonsense, hopefully

// RemoveRule removes a rule with the specified name from the rule set.
// If no rule with the specified name exists, it does nothing.
func (r *Rules) RemoveRule(name string) { _ = "STUB: not implemented"; return }

// impossible nonsense, hopefully

// delete the name rule key

// ReplaceRule replaces a rule with the specified name with a new rule function.
// If no rule with the specified name exists, it does nothing.
func (r *Rules) ReplaceRule(name string, ruleFunc core.RuleFunc) { _ = "STUB: not implemented"; return }

// impossible nonsense, hopefully
