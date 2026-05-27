package core

import (
	"context"

	"github.com/vektah/gqlparser/v2/ast"
)

type Events struct {
	operationVisitor []func(walker *Walker, operation *ast.OperationDefinition)
	field            []func(walker *Walker, field *ast.Field)
	fragment         []func(walker *Walker, fragment *ast.FragmentDefinition)
	inlineFragment   []func(walker *Walker, inlineFragment *ast.InlineFragment)
	fragmentSpread   []func(walker *Walker, fragmentSpread *ast.FragmentSpread)
	directive        []func(walker *Walker, directive *ast.Directive)
	directiveList    []func(walker *Walker, directives []*ast.Directive)
	value            []func(walker *Walker, value *ast.Value)
	variable         []func(walker *Walker, variable *ast.VariableDefinition)
}

func (o *Events) OnOperation(f func(walker *Walker, operation *ast.OperationDefinition)) {
	_ = "STUB: not implemented"
	return
}

func (o *Events) OnField(f func(walker *Walker, field *ast.Field)) {
	_ = "STUB: not implemented"
	return
}

func (o *Events) OnFragment(f func(walker *Walker, fragment *ast.FragmentDefinition)) {
	_ = "STUB: not implemented"
	return
}

func (o *Events) OnInlineFragment(f func(walker *Walker, inlineFragment *ast.InlineFragment)) {
	_ = "STUB: not implemented"
	return
}

func (o *Events) OnFragmentSpread(f func(walker *Walker, fragmentSpread *ast.FragmentSpread)) {
	_ = "STUB: not implemented"
	return
}

func (o *Events) OnDirective(f func(walker *Walker, directive *ast.Directive)) {
	_ = "STUB: not implemented"
	return
}

func (o *Events) OnDirectiveList(f func(walker *Walker, directives []*ast.Directive)) {
	_ = "STUB: not implemented"
	return
}

func (o *Events) OnValue(f func(walker *Walker, value *ast.Value)) {
	_ = "STUB: not implemented"
	return
}

func (o *Events) OnVariable(f func(walker *Walker, variable *ast.VariableDefinition)) {
	_ = "STUB: not implemented"
	return
}

func Walk(schema *ast.Schema, document *ast.QueryDocument, observers *Events) {
	_ = "STUB: not implemented"
	return
}

type Walker struct {
	Context   context.Context
	Observers *Events
	Schema    *ast.Schema
	Document  *ast.QueryDocument

	validatedFragmentSpreads map[string]bool
	CurrentOperation         *ast.OperationDefinition
}

func (w *Walker) walk() { _ = "STUB: not implemented"; return }

func (w *Walker) walkOperation(operation *ast.OperationDefinition) {
	_ = "STUB: not implemented"
	return
}

func (w *Walker) walkFragment(it *ast.FragmentDefinition) { _ = "STUB: not implemented"; return }

func (w *Walker) walkDirectives(
	parentDef *ast.Definition,
	directives []*ast.Directive,
	location ast.DirectiveLocation,
) {
	_ = "STUB: not implemented"
	return
}

func (w *Walker) walkValue(value *ast.Value) { _ = "STUB: not implemented"; return }

func (w *Walker) walkArgument(argDef *ast.ArgumentDefinition, arg *ast.Argument) {
	_ = "STUB: not implemented"
	return
}

func (w *Walker) walkSelectionSet(parentDef *ast.Definition, it ast.SelectionSet) {
	_ = "STUB: not implemented"
	return
}

func (w *Walker) walkSelection(parentDef *ast.Definition, it ast.Selection) {
	_ = "STUB: not implemented"
	return
}

// prevent infinite recursion
