package gqlparser

import (
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vektah/gqlparser/v2/validator/rules"
)

func LoadSchema(str ...*ast.Source) (*ast.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MustLoadSchema(str ...*ast.Source) *ast.Schema { _ = "STUB: not implemented"; return nil }

// Deprecated: use LoadQueryWithRules instead.
func LoadQuery(schema *ast.Schema, str string) (*ast.QueryDocument, gqlerror.List) {
	_ = "STUB: not implemented"
	return nil, *new(gqlerror.List)
}

func LoadQueryWithRules(
	schema *ast.Schema,
	str string,
	rules *rules.Rules,
) (*ast.QueryDocument, gqlerror.List) {
	_ = "STUB: not implemented"
	return nil, *new(gqlerror.List)
}

// Deprecated: use MustLoadQueryWithRules instead.
func MustLoadQuery(schema *ast.Schema, str string) *ast.QueryDocument {
	_ = "STUB: not implemented"
	return nil
}

func MustLoadQueryWithRules(schema *ast.Schema, str string, rules *rules.Rules) *ast.QueryDocument {
	_ = "STUB: not implemented"
	return nil
}
