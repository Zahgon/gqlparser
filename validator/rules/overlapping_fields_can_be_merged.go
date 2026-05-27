package rules

import (
	"bytes"

	"github.com/vektah/gqlparser/v2/ast"
	//nolint:staticcheck // Validator rules each use dot imports for convenience.
	. "github.com/vektah/gqlparser/v2/validator/core"
)

var OverlappingFieldsCanBeMergedRule = Rule{
	Name: "OverlappingFieldsCanBeMerged",
	RuleFunc: func(observers *Events, addError AddErrFunc) {
		/**
		 * Algorithm:
		 *
		 * Conflicts occur when two fields exist in a query which will produce the same
		 * response name, but represent differing values, thus creating a conflict.
		 * The algorithm below finds all conflicts via making a series of comparisons
		 * between fields. In order to compare as few fields as possible, this makes
		 * a series of comparisons "within" sets of fields and "between" sets of fields.
		 *
		 * Given any selection set, a collection produces both a set of fields by
		 * also including all inline fragments, as well as a list of fragments
		 * referenced by fragment spreads.
		 *
		 * A) Each selection set represented in the document first compares "within" its
		 * collected set of fields, finding any conflicts between every pair of
		 * overlapping fields.
		 * Note: This is the *only time* that a the fields "within" a set are compared
		 * to each other. After this only fields "between" sets are compared.
		 *
		 * B) Also, if any fragment is referenced in a selection set, then a
		 * comparison is made "between" the original set of fields and the
		 * referenced fragment.
		 *
		 * C) Also, if multiple fragments are referenced, then comparisons
		 * are made "between" each referenced fragment.
		 *
		 * D) When comparing "between" a set of fields and a referenced fragment, first
		 * a comparison is made between each field in the original set of fields and
		 * each field in the referenced set of fields.
		 *
		 * E) Also, if any fragment is referenced in the referenced selection set,
		 * then a comparison is made "between" the original set of fields and the
		 * referenced fragment (recursively referring to step D).
		 *
		 * F) When comparing "between" two fragments, first a comparison is made between
		 * each field in the first referenced set of fields and each field in the the
		 * second referenced set of fields.
		 *
		 * G) Also, any fragments referenced by the first must be compared to the
		 * second, and any fragments referenced by the second must be compared to the
		 * first (recursively referring to step F).
		 *
		 * H) When comparing two fields, if both have selection sets, then a comparison
		 * is made "between" both selection sets, first comparing the set of fields in
		 * the first selection set with the set of fields in the second.
		 *
		 * I) Also, if any fragment is referenced in either selection set, then a
		 * comparison is made "between" the other set of fields and the
		 * referenced fragment.
		 *
		 * J) Also, if two fragments are referenced in both selection sets, then a
		 * comparison is made "between" the two fragments.
		 *
		 */

		m := &overlappingFieldsCanBeMergedManager{
			comparedFragmentPairs: pairSet{data: make(map[string]map[string]bool)},
		}

		observers.OnOperation(func(walker *Walker, operation *ast.OperationDefinition) {
			m.walker = walker
			conflicts := m.findConflictsWithinSelectionSet(operation.SelectionSet)
			for _, conflict := range conflicts {
				conflict.addFieldsConflictMessage(addError)
			}
		})
		observers.OnField(func(walker *Walker, field *ast.Field) {
			if walker.CurrentOperation == nil {
				// When checking both Operation and Fragment, errors are duplicated when processing
				// FragmentDefinition referenced from Operation
				return
			}
			m.walker = walker
			conflicts := m.findConflictsWithinSelectionSet(field.SelectionSet)
			for _, conflict := range conflicts {
				conflict.addFieldsConflictMessage(addError)
			}
		})
		observers.OnInlineFragment(func(walker *Walker, inlineFragment *ast.InlineFragment) {
			m.walker = walker
			conflicts := m.findConflictsWithinSelectionSet(inlineFragment.SelectionSet)
			for _, conflict := range conflicts {
				conflict.addFieldsConflictMessage(addError)
			}
		})
		observers.OnFragment(func(walker *Walker, fragment *ast.FragmentDefinition) {
			m.walker = walker
			conflicts := m.findConflictsWithinSelectionSet(fragment.SelectionSet)
			for _, conflict := range conflicts {
				conflict.addFieldsConflictMessage(addError)
			}
		})
	},
}

type pairSet struct {
	data map[string]map[string]bool
}

func (pairSet *pairSet) Add(
	a *ast.FragmentSpread,
	b *ast.FragmentSpread,
	areMutuallyExclusive bool,
) {
	_ = "STUB: not implemented"
	return
}

func (pairSet *pairSet) Has(
	a *ast.FragmentSpread,
	b *ast.FragmentSpread,
	areMutuallyExclusive bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// areMutuallyExclusive being false is a superset of being true,
// hence if we want to know if this PairSet "has" these two with no
// exclusivity, we have to ensure it was added as such.

type sequentialFieldsMap struct {
	// We can't use map[string][]*ast.Field. because map is not stable...
	seq  []string
	data map[string][]*ast.Field
}

type fieldIterateEntry struct {
	ResponseName string
	Fields       []*ast.Field
}

func (m *sequentialFieldsMap) Push(responseName string, field *ast.Field) {
	_ = "STUB: not implemented"
	return
}

func (m *sequentialFieldsMap) Get(responseName string) ([]*ast.Field, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *sequentialFieldsMap) Iterator() [][]*ast.Field { _ = "STUB: not implemented"; return nil }

func (m *sequentialFieldsMap) KeyValueIterator() []*fieldIterateEntry {
	_ = "STUB: not implemented"
	return nil
}

type conflictMessageContainer struct {
	Conflicts []*ConflictMessage
}

type ConflictMessage struct {
	Message      string
	ResponseName string
	Names        []string
	SubMessage   []*ConflictMessage
	Position     *ast.Position
}

func (m *ConflictMessage) String(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (m *ConflictMessage) addFieldsConflictMessage(addError AddErrFunc) {
	_ = "STUB: not implemented"
	return
}

type overlappingFieldsCanBeMergedManager struct {
	walker *Walker

	// per walker
	comparedFragmentPairs pairSet
	// cachedFieldsAndFragmentNames interface{}

	// per selectionSet
	comparedFragments map[string]bool
}

func (m *overlappingFieldsCanBeMergedManager) findConflictsWithinSelectionSet(
	selectionSet ast.SelectionSet,
) []*ConflictMessage {
	_ = "STUB: not implemented"
	return nil
}

// (A) Find find all conflicts "within" the fieldMap of this selection set.
// Note: this is the *only place* `collectConflictsWithin` is called.

// (B) Then collect conflicts between these fieldMap and those represented by
// each spread fragment name found.

// (C) Then compare this fragment with all other fragments found in this
// selection set to collect conflicts between fragments spread together.
// This compares each item in the list of fragment names to every other
// item in that same list (except for itself).

func (m *overlappingFieldsCanBeMergedManager) collectConflictsBetweenFieldsAndFragment(
	conflicts *conflictMessageContainer,
	areMutuallyExclusive bool,
	fieldsMap *sequentialFieldsMap,
	fragmentSpread *ast.FragmentSpread,
) {
	_ = "STUB: not implemented"
	return
}

// Do not compare a fragment's fieldMap to itself.

// (D) First collect any conflicts between the provided collection of fields
// and the collection of fields represented by the given fragment.

// (E) Then collect any conflicts between the provided collection of fields
// and any fragment names found in the given fragment.

func (m *overlappingFieldsCanBeMergedManager) collectConflictsBetweenFragments(
	conflicts *conflictMessageContainer,
	areMutuallyExclusive bool,
	fragmentSpreadA *ast.FragmentSpread,
	fragmentSpreadB *ast.FragmentSpread,
) {
	_ = "STUB: not implemented"
	return
}

// (F) First, collect all conflicts between these two collections of fields
// (not including any nested fragments).

// (G) Then collect conflicts between the first fragment and any nested
// fragments spread in the second fragment.

// (G) Then collect conflicts between the second fragment and any nested
// fragments spread in the first fragment.

func (m *overlappingFieldsCanBeMergedManager) findConflictsBetweenSubSelectionSets(
	areMutuallyExclusive bool,
	selectionSetA ast.SelectionSet,
	selectionSetB ast.SelectionSet,
) *conflictMessageContainer {
	_ = "STUB: not implemented"
	return nil
}

// (H) First, collect all conflicts between these two collections of field.

// (I) Then collect conflicts between the first collection of fields and
// those referenced by each fragment name associated with the second.

// (I) Then collect conflicts between the second collection of fields and
// those referenced by each fragment name associated with the first.

// (J) Also collect conflicts between any fragment names by the first and
// fragment names by the second. This compares each item in the first set of
// names to each item in the second set of names.

func (m *overlappingFieldsCanBeMergedManager) collectConflictsWithin(
	conflicts *conflictMessageContainer,
	fieldsMap *sequentialFieldsMap,
) {
	_ = "STUB: not implemented"
	return
}

func (m *overlappingFieldsCanBeMergedManager) collectConflictsBetween(
	conflicts *conflictMessageContainer,
	parentFieldsAreMutuallyExclusive bool,
	fieldsMapA *sequentialFieldsMap,
	fieldsMapB *sequentialFieldsMap,
) {
	_ = "STUB: not implemented"
	return
}

func (m *overlappingFieldsCanBeMergedManager) findConflict(
	parentFieldsAreMutuallyExclusive bool,
	fieldA *ast.Field,
	fieldB *ast.Field,
) *ConflictMessage {
	_ = "STUB: not implemented"
	return nil
}

// Two aliases must refer to the same field.

// Two field calls must have the same arguments.

// Collect and compare sub-fields. Use the same "visited fragment names" list
// for both collections so fields in a fragment reference are never
// compared to themselves.

func sameArguments(args1, args2 []*ast.Argument) bool { _ = "STUB: not implemented"; return false }

func sameValue(value1, value2 *ast.Value) bool { _ = "STUB: not implemented"; return false }

func doTypesConflict(walker *Walker, type1, type2 *ast.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func getFieldsAndFragmentNames(
	selectionSet ast.SelectionSet,
) (*sequentialFieldsMap, []*ast.FragmentSpread) {
	_ = "STUB: not implemented"
	return nil, nil
}
