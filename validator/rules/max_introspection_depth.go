package rules

import (
	"github.com/vektah/gqlparser/v2/ast"
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

const maxListsDepth = 3

var MaxIntrospectionDepth = Rule{
	Name: "MaxIntrospectionDepth",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		// Counts the depth of list fields in "__Type" recursively and
		// returns `true` if the limit has been reached.
		observers.OnField(func(walker *Walker, field *ast.Field) {
			if field.Name == "__schema" || field.Name == "__type" {
				visitedFragments := make(map[string]bool)
				if checkDepthField(field, visitedFragments, 0) {
					addError(
						Message(`Maximum introspection depth exceeded`),
						At(field.Position),
					)
				}
				return
			}
		})
	},
}

func checkDepthSelectionSet(
	selectionSet ast.SelectionSet,
	visitedFragments map[string]bool,
	depth int,
) bool {
	_ = "STUB: not implemented"
	return false
}

func checkDepthField(field *ast.Field, visitedFragments map[string]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func checkDepthFragmentSpread(
	fragmentSpread *ast.FragmentSpread,
	visitedFragments map[string]bool,
	depth int,
) bool {
	_ = "STUB: not implemented"
	return false
}

// Fragment cycles are handled by `NoFragmentCyclesRule`.

// Missing fragments checks are handled by `KnownFragmentNamesRule`.

// Rather than following an immutable programming pattern which has
// significant memory and garbage collection overhead, we've opted to
// take a mutable approach for efficiency's sake. Importantly visiting a
// fragment twice is fine, so long as you don't do one visit inside the
// other.
