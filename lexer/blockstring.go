package lexer

// blockStringValue produces the value of a block string from its parsed raw value, similar to
// Coffeescript's block string, Python's docstring trim or Ruby's strip_heredoc.
//
// This implements the GraphQL spec's BlockStringValue() static algorithm.
func blockStringValue(raw string) string { _ = "STUB: not implemented"; return "" }

func leadingWhitespace(str string) int { _ = "STUB: not implemented"; return 0 }

// this line is made up entirely of whitespace, its leading whitespace doesnt count.
