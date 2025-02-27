// Package matchphrase provides structures and functions for building match phrase queries.
// The match_phrase query analyzes the text and creates a phrase query out of the analyzed text.
package matchphrase

// MatchPhrase represents a match phrase query.
type MatchPhrase struct {
	MatchPhrase map[string]*Param `json:"match_phrase"`
}

// Param represents the parameters for a match phrase query, including query, analyzer, zero terms query, and slop.
type Param struct {
	Query          string `json:"query"`
	Analyzer       string `json:"analyzer,omitempty"`
	ZeroTermsQuery string `json:"zero_terms_query,omitempty"`
	Slop           int    `json:"slop,omitempty"`
}

// New creates a new MatchPhrase instance with the specified field and match phrase parameters.
func New(f string, mp *Param) *MatchPhrase {
	return &MatchPhrase{MatchPhrase: map[string]*Param{f: mp}}
}

// NewParam creates a new Param instance for match phrase query parameters with the specified query.
func NewParam(q string) *Param {
	return &Param{Query: q}
}

// SetAnalyzer sets the analyzer for the match phrase query parameter.
// The analyzer used to tokenize the query string text.
// Default is the index-time analyzer specified for the default_field.
// If no analyzer is specified for the default_field, the analyzer is the default analyzer for the index.
func (p *Param) SetAnalyzer(value string) *Param {
	p.Analyzer = value
	return p
}

// SetZeroTermsQuery sets the zero terms query for the match phrase query parameter.
// In some cases, the analyzer removes all terms from a query string.
// For example, the stop analyzer removes all terms from the string.
// In those cases, zero_terms_query specifies whether to match no documents (none) or all documents (all).
// Valid values are none and all.
// Default is none.
func (p *Param) SetZeroTermsQuery(value string) *Param {
	p.ZeroTermsQuery = value
	return p
}

// SetSlop sets the slop for the match phrase query parameter.
// Controls the degree to which words in a query can be misordered and still be considered a match.
// From the Lucene documentation: “The number of other words permitted between words in query phrase.
// For example, to switch the order of two words requires two moves (the first move places the words atop one another),
// so to permit reorderings of phrases, the slop must be at least two. A value of zero requires an exact match.”
func (p *Param) SetSlop(value int) *Param {
	p.Slop = value
	return p
}
