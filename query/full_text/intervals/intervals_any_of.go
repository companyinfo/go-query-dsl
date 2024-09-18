// Package intervals provides structures and functions for building intervals queries.
// Returns documents based on the order and proximity of matching terms.
package intervals

// AnyOf represents the any_of interval query
type AnyOf struct {
	Intervals []*Param `json:"intervals"`
	Filter    *Filter  `json:"filter,omitempty"`
}

// NewParamAnyOf creates a new AnyOf instance for the intervals query parameters.
func NewParamAnyOf(p []*Param) *AnyOf {
	return &AnyOf{Intervals: p}
}

// SetFilter sets the Filter for the AnyOf.
// A rule used to filter returned intervals.
func (a *AnyOf) SetFilter(f *Filter) *AnyOf {
	a.Filter = f
	return a
}
