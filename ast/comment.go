package ast

type Comment struct {
	Value    string
	Position *Position
}

func (c *Comment) Text() string { _ = "STUB: not implemented"; return "" }

type CommentGroup struct {
	List []*Comment
}

func (c *CommentGroup) Dump() string { _ = "STUB: not implemented"; return "" }
