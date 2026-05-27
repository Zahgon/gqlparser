package parser

import (
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/lexer"
)

type parser struct {
	lexer lexer.Lexer
	err   error

	peeked    bool
	peekToken lexer.Token
	peekError error

	prev lexer.Token

	comment          *ast.CommentGroup
	commentConsuming bool

	tokenCount    int
	maxTokenLimit int
}

func (p *parser) SetMaxTokenLimit(maxToken int) { _ = "STUB: not implemented"; return }

func (p *parser) consumeComment() (*ast.Comment, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *parser) consumeCommentGroup() { _ = "STUB: not implemented"; return }

func (p *parser) peekPos() *ast.Position { _ = "STUB: not implemented"; return nil }

func (p *parser) peek() lexer.Token { _ = "STUB: not implemented"; return *new(lexer.Token) }

func (p *parser) error(tok lexer.Token, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) next() lexer.Token { _ = "STUB: not implemented"; return *new(lexer.Token) }

// Increment the token count before reading the next token

func (p *parser) expectKeyword(value string) (lexer.Token, *ast.CommentGroup) {
	_ = "STUB: not implemented"
	return *new(lexer.Token), nil
}

func (p *parser) expect(kind lexer.Type) (lexer.Token, *ast.CommentGroup) {
	_ = "STUB: not implemented"
	return *new(lexer.Token), nil
}

func (p *parser) skip(kind lexer.Type) bool { _ = "STUB: not implemented"; return false }

func (p *parser) unexpectedError() { _ = "STUB: not implemented"; return }

func (p *parser) unexpectedToken(tok lexer.Token) { _ = "STUB: not implemented"; return }

func (p *parser) many(start, end lexer.Type, cb func()) { _ = "STUB: not implemented"; return }

func (p *parser) some(start, end lexer.Type, cb func()) *ast.CommentGroup {
	_ = "STUB: not implemented"
	return nil
}
