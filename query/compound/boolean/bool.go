// Package boolean provides structures and functions for building bool queries.
package boolean

import "go.companyinfo.dev/go-query-dsl/query/full_text/match"

// Bool represents a bool query.
type Bool struct {
	Bool QueryParams `json:"bool"`
}

// QueryParams represents the parameters for a bool query, including must,
// must_not, should, filter, and minimum should match.
type QueryParams struct {
	Must               []any   `json:"must,omitempty"`
	MustNot            []any   `json:"must_not,omitempty"`
	Should             []any   `json:"should,omitempty"`
	Filter             []any   `json:"filter,omitempty"`
	MinimumShouldMatch int     `json:"minimum_should_match,omitempty"`
	Boost              float64 `json:"boost,omitempty"`
}

// Must represent the must clause in the bool query.
type Must struct {
	Must []*match.Match `json:"must"`
}

// New creates a new Bool instance with initialized query parameters.
func New() *Bool {
	return &Bool{Bool: QueryParams{
		Must:    make([]any, 0),
		MustNot: make([]any, 0),
		Should:  make([]any, 0),
		Filter:  make([]any, 0),
	}}
}

// AddMust appends a must clause to the bool query.
func (b *Bool) AddMust(m any) *Bool {
	b.Bool.Must = append(b.Bool.Must, m)
	return b
}

// AddMustNot appends a must_not clause to the bool query.
func (b *Bool) AddMustNot(m any) *Bool {
	b.Bool.MustNot = append(b.Bool.MustNot, m)
	return b
}

// AddShould appends a should clause to the bool query.
func (b *Bool) AddShould(m any) *Bool {
	b.Bool.Should = append(b.Bool.Should, m)
	return b
}

// AddFilter appends a filter clause to the bool query.
func (b *Bool) AddFilter(value any) *Bool {
	b.Bool.Filter = append(b.Bool.Filter, value)
	return b
}

// SetMinimumShouldMatch sets the minimum should match value for the bool query.
// If the query string contains multiple search terms, and you use the or operator, the number of terms that need
// to match for the document to be considered a match.
// For example, if minimum_should_match is 2, wind often rising does not match The Wind Rises.
// If minimum_should_match is 1, it matches.
func (b *Bool) SetMinimumShouldMatch(value int) *Bool {
	b.Bool.MinimumShouldMatch = value
	return b
}

// SetBoost sets the boost for the fuzzy query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (b *Bool) SetBoost(value float64) *Bool {
	b.Bool.Boost = value
	return b
}
