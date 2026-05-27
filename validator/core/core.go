package core

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type AddErrFunc func(options ...ErrorOption)

type RuleFunc func(observers *Events, addError AddErrFunc)

type Rule struct {
	Name     string
	RuleFunc RuleFunc
}

// NameSorter sorts Rules by name.
// usage: sort.Sort(core.NameSorter(specifiedRules))
type NameSorter []Rule

func (a NameSorter) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a NameSorter) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a NameSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type ErrorOption func(err *gqlerror.Error)
