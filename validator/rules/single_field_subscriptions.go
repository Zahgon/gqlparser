package rules

import (
	"strconv"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

var SingleFieldSubscriptionsRule = Rule{
	Name: "SingleFieldSubscriptions",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		observers.OnOperation(func(walker *Walker, operation *ast.OperationDefinition) {
			if walker.Schema.Subscription == nil || operation.Operation != ast.Subscription {
				return
			}

			fields := retrieveTopFieldNames(operation.SelectionSet)

			name := "Anonymous Subscription"
			if operation.Name != "" {
				name = `Subscription ` + strconv.Quote(operation.Name)
			}

			if len(fields) > 1 {
				addError(
					Message(`%s must select only one top level field.`, name),
					At(fields[1].position),
				)
			}

			for _, field := range fields {
				if strings.HasPrefix(field.name, "__") {
					addError(
						Message(`%s must not select an introspection top level field.`, name),
						At(field.position),
					)
				}
			}
		})
	},
}

type topField struct {
	name     string
	position *ast.Position
}

func retrieveTopFieldNames(selectionSet ast.SelectionSet) []*topField {
	_ = "STUB: not implemented"
	return nil
}
