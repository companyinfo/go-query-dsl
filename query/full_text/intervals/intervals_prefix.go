// Package intervals provides structures and functions for building intervals queries.
// Returns documents based on the order and proximity of matching terms.
package intervals

// Prefix represents the Prefix interval query
type Prefix struct {
	Prefix   string `json:"prefix"`
	Analyzer string `json:"analyzer,omitempty"`
	UseField string `json:"use_field,omitempty"`
}

// NewParamPrefix creates a new Prefix instance for the intervals query parameters.
func NewParamPrefix(p string) *Prefix {
	return &Prefix{Prefix: p}
}

// SetAnalyzer sets the Analyzer for the Prefix.
// The analyzer used to analyze the query text. Default is the analyzer specified for the <field>.
func (p *Prefix) SetAnalyzer(value string) *Prefix {
	p.Analyzer = value
	return p
}

// SetUseField sets the UseField for the Prefix.
// Specifies to search this field instead of the top-level .
// The `prefix` is normalized using the search analyzer specified for this field, unless you specify an `analyzer`.
func (p *Prefix) SetUseField(value string) *Prefix {
	p.UseField = value
	return p
}
