// Package disjunctionmax A disjunction max (dis_max) query returns any document that matches one or more query clauses.
// For documents that match multiple query clauses, the relevance score is set to the highest relevance score from all
// matching query clauses.
package disjunctionmax

// DisjunctionMax represents a disjunction max query.
type DisjunctionMax struct {
	DisjunctionMax Param `json:"dis_max"`
}

// Param represents the parameters for a disjunction max query.
type Param struct {
	Queries    any     `json:"queries"`
	TieBreaker float64 `json:"tie_breaker,omitempty"`
}

// New creates a new Bool instance with initialized query parameters.
// queries is an array of one or more query clauses that are used to match documents.
// A document must match at least one query clause to be returned in the results.
// If a document matches multiple query clauses, the relevance score is set to the highest relevance score from all
// matching query clauses.
func New(queries any) *DisjunctionMax {
	return &DisjunctionMax{DisjunctionMax: Param{
		Queries: queries,
	}}
}

// SetTieBreaker sets the tie-breaker for the disjunction max query parameter.
// A floating-point factor between 0 and 1.0 that is used to give more weight to documents that match multiple query
// clauses. In this case, the relevance score of a document is calculated using the following algorithm:
// Take the highest relevance score from all matching query clauses, multiply the scores from all other matching clauses
// by the tie_breaker value, and add the relevance scores together, normalizing them.
// Default is 0 (which means only the highest score counts).
func (d *DisjunctionMax) SetTieBreaker(value float64) *DisjunctionMax {
	d.DisjunctionMax.TieBreaker = value
	return d
}
