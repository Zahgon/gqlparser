package ast

type QueryDocument struct {
	Operations OperationList
	Fragments  FragmentDefinitionList
	Position   *Position `dump:"-" json:"-"`
	Comment    *CommentGroup
}

type SchemaDocument struct {
	Schema          SchemaDefinitionList
	SchemaExtension SchemaDefinitionList
	Directives      DirectiveDefinitionList
	Definitions     DefinitionList
	Extensions      DefinitionList
	Position        *Position `dump:"-" json:"-"`
	Comment         *CommentGroup
}

func (d *SchemaDocument) Merge(other *SchemaDocument) { _ = "STUB: not implemented"; return }

type Schema struct {
	Query            *Definition
	Mutation         *Definition
	Subscription     *Definition
	SchemaDirectives DirectiveList

	Types      map[string]*Definition
	Directives map[string]*DirectiveDefinition

	PossibleTypes map[string][]*Definition
	Implements    map[string][]*Definition

	Description string

	Comment *CommentGroup
}

// AddTypes is the helper to add types definition to the schema.
func (s *Schema) AddTypes(defs ...*Definition) { _ = "STUB: not implemented"; return }

func (s *Schema) AddPossibleType(name string, def *Definition) { _ = "STUB: not implemented"; return }

// GetPossibleTypes will enumerate all the definitions for a given interface or union.
func (s *Schema) GetPossibleTypes(def *Definition) []*Definition {
	_ = "STUB: not implemented"
	return nil
}

func (s *Schema) AddImplements(name string, iface *Definition) { _ = "STUB: not implemented"; return }

// GetImplements returns all the interface and union definitions that the given definition
// satisfies.
func (s *Schema) GetImplements(def *Definition) []*Definition {
	_ = "STUB: not implemented"
	return nil
}

type SchemaDefinition struct {
	Description    string
	Directives     DirectiveList
	OperationTypes OperationTypeDefinitionList
	Position       *Position `dump:"-" json:"-"`

	BeforeDescriptionComment *CommentGroup
	AfterDescriptionComment  *CommentGroup
	EndOfDefinitionComment   *CommentGroup
}

type OperationTypeDefinition struct {
	Operation Operation
	Type      string
	Position  *Position `dump:"-" json:"-"`
	Comment   *CommentGroup
}
