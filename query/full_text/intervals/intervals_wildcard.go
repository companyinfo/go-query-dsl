// Package intervals provides structures and functions for building intervals queries.
// Returns documents based on the order and proximity of matching terms.
package intervals

// Wildcard represents the Wildcard interval query
type Wildcard struct {
	Pattern  string `json:"pattern"`
	Analyzer string `json:"analyzer,omitempty"`
	UseField string `json:"use_field,omitempty"`
}

// NewParamWildcard creates a new Wildcard instance for the intervals query parameters.
func NewParamWildcard(p string) *Wildcard {
	return &Wildcard{Pattern: p}
}

// SetAnalyzer sets the Analyzer for the Wildcard.
// The analyzer used to analyze the query text. Default is the analyzer specified for the <field>.
func (w *Wildcard) SetAnalyzer(value string) *Wildcard {
	w.Analyzer = value
	return w
}

// SetUseField sets the UseField for the Wildcard.
// Specifies to search this field instead of the top-level .
// The `prefix` is normalized using the search analyzer specified for this field, unless you specify an `analyzer`.
func (w *Wildcard) SetUseField(value string) *Wildcard {
	w.UseField = value
	return w
}
