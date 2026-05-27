package rules

import (
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

var NoFragmentCyclesRule = Rule{
	Name: "NoFragmentCycles",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		visitedFrags := make(map[string]bool)

		observers.OnFragment(func(walker *Walker, fragment *ast.FragmentDefinition) {
			var spreadPath []*ast.FragmentSpread
			spreadPathIndexByName := make(map[string]int)

			var recursive func(fragment *ast.FragmentDefinition)
			recursive = func(fragment *ast.FragmentDefinition) {
				if visitedFrags[fragment.Name] {
					return
				}

				visitedFrags[fragment.Name] = true

				spreadNodes := getFragmentSpreads(fragment.SelectionSet)
				if len(spreadNodes) == 0 {
					return
				}
				spreadPathIndexByName[fragment.Name] = len(spreadPath)

				for _, spreadNode := range spreadNodes {
					spreadName := spreadNode.Name

					cycleIndex, ok := spreadPathIndexByName[spreadName]

					spreadPath = append(spreadPath, spreadNode)
					if !ok {
						spreadFragment := walker.Document.Fragments.ForName(spreadName)
						if spreadFragment != nil {
							recursive(spreadFragment)
						}
					} else {
						cyclePath := spreadPath[cycleIndex : len(spreadPath)-1]
						var fragmentNames []string
						for _, fs := range cyclePath {
							fragmentNames = append(fragmentNames, fmt.Sprintf(`"%s"`, fs.Name))
						}
						var via string
						if len(fragmentNames) != 0 {
							via = fmt.Sprintf(" via %s", strings.Join(fragmentNames, ", "))
						}
						addError(
							Message(
								`Cannot spread fragment "%s" within itself%s.`,
								spreadName,
								via,
							),
							At(spreadNode.Position),
						)
					}

					spreadPath = spreadPath[:len(spreadPath)-1]
				}

				delete(spreadPathIndexByName, fragment.Name)
			}

			recursive(fragment)
		})
	},
}

func getFragmentSpreads(node ast.SelectionSet) []*ast.FragmentSpread {
	_ = "STUB: not implemented"
	return nil
}
