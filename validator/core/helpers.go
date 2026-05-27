package core

import (
	"github.com/vektah/gqlparser/v2/ast"
)

func Message(msg string, args ...any) ErrorOption {
	_ = "STUB: not implemented"
	return *new(ErrorOption)
}

func At(position *ast.Position) ErrorOption { _ = "STUB: not implemented"; return *new(ErrorOption) }

func SuggestListQuoted(prefix, typed string, suggestions []string) ErrorOption {
	_ = "STUB: not implemented"
	return *new(ErrorOption)
}

func SuggestListUnquoted(prefix, typed string, suggestions []string) ErrorOption {
	_ = "STUB: not implemented"
	return *new(ErrorOption)
}

func Suggestf(suggestion string, args ...any) ErrorOption {
	_ = "STUB: not implemented"
	return *new(ErrorOption)
}

// Given [ A, B, C ] return '"A", "B", or "C"'.
func QuotedOrList(items ...string) string { _ = "STUB: not implemented"; return "" }

// Given [ A, B, C ] return 'A, B, or C'.
func OrList(items ...string) string { _ = "STUB: not implemented"; return "" }

// Given an invalid input string and a list of valid options, returns a filtered
// list of valid options sorted based on their similarity with the input.
func SuggestionList(input string, options []string) []string { _ = "STUB: not implemented"; return nil }

func calcThreshold(a string) (threshold int) {
	_ = "STUB: not implemented"
	// the logic is copied from here
	// https://github.com/graphql/graphql-js/blob/47bd8c8897c72d3efc17ecb1599a95cee6bac5e8/src/jsutils/suggestionList.ts#L14
	return 0
}

// Computes the lexical distance between strings A and B.
//
// The "distance" between two strings is given by counting the minimum number
// of edits needed to transform string A into string B. An edit can be an
// insertion, deletion, or substitution of a single character, or a swap of two
// adjacent characters.
//
// Includes a custom alteration from Damerau-Levenshtein to treat case changes
// as a single edit which helps identify mis-cased values with an edit distance
// of 1.
//
// This distance can be useful for detecting typos in input or sorting.
func lexicalDistance(a, b string) int { _ = "STUB: not implemented"; return 0 }

// Any case change counts as a single edit
