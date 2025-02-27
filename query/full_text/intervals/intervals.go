// Package intervals provides structures and functions for building intervals queries.
// Returns documents based on the order and proximity of matching terms.
package intervals

// Intervals represents a intervals query.
type Intervals struct {
	Intervals map[string]*Param `json:"intervals"`
}

// Param represents the parameters for an intervals query.
type Param struct {
	Match    *Match    `json:"match,omitempty"`
	Prefix   *Prefix   `json:"prefix,omitempty"`
	Wildcard *Wildcard `json:"wildcard,omitempty"`
	Fuzzy    *Fuzzy    `json:"fuzzy,omitempty"`
	AllOf    *AllOf    `json:"all_of,omitempty"`
	AnyOf    *AnyOf    `json:"any_of,omitempty"`
}

// New creates a new intervals instance with the specified field and parameters.
func New(f string, p *Param) *Intervals {
	return &Intervals{Intervals: map[string]*Param{f: p}}
}

// NewParam creates a new Param instance for intervals query.
func NewParam() *Param {
	return &Param{}
}

// SetMatch sets the Match for the intervals query parameter.
// Matches analyzed text.
func (p *Param) SetMatch(m *Match) *Param {
	p.Match = m
	return p
}

// SetPrefix sets the Prefix for the intervals query parameter.
// Matches terms that start with a specified set of characters.
func (p *Param) SetPrefix(pre *Prefix) *Param {
	p.Prefix = pre
	return p
}

// SetWildcard sets the Wildcard for the intervals query parameter.
// Matches terms using a wildcard pattern.
func (p *Param) SetWildcard(w *Wildcard) *Param {
	p.Wildcard = w
	return p
}

// SetFuzzy sets the Fuzzy for the intervals query parameter.
// Matches terms that are similar to the provided term within a specified edit distance.
func (p *Param) SetFuzzy(f *Fuzzy) *Param {
	p.Fuzzy = f
	return p
}

// SetAllOf sets the AllOf for the intervals query parameter.
// Combines multiple rules using a conjunction (AND).
func (p *Param) SetAllOf(a *AllOf) *Param {
	p.AllOf = a
	return p
}

// SetAnyOf sets the AnyOf for the intervals query parameter.
// Combines multiple rules using a disjunction (OR).
func (p *Param) SetAnyOf(a *AnyOf) *Param {
	p.AnyOf = a
	return p
}
