// Package matchphraseprefix provides structures and functions for building match phrase prefix queries.
// Returns documents that contain the words of a provided text, in the same order as provided.
// The last term of the provided text is treated as a prefix, matching any words that begin with that term.
package matchphraseprefix

import "go.companyinfo.dev/go-query-dsl/query/full_text/match"

// MatchPhrasePrefix represents a match phrase prefix query.
type MatchPhrasePrefix struct {
	MatchPhrasePrefix map[string]*Param `json:"match_phrase_prefix"`
}

// Param represents the parameters for a match phrase prefix query, including query, analyzer, max expansions, and slop.
type Param struct {
	Query          string             `json:"query"`
	Analyzer       string             `json:"analyzer,omitempty"`
	MaxExpansions  int                `json:"max_expansions,omitempty"`
	Slop           int                `json:"slop,omitempty"`
	ZeroTermsQuery match.ZeroTermType `json:"zero_terms_query,omitempty"`
}

// New creates a new MatchPhrasePrefix instance with the specified field and match phrase prefix parameters.
func New(f string, mp *Param) *MatchPhrasePrefix {
	return &MatchPhrasePrefix{MatchPhrasePrefix: map[string]*Param{f: mp}}
}

// NewParam creates a new Param instance for match phrase prefix query parameters with the specified query.
func NewParam(q string) *Param {
	return &Param{Query: q}
}

// SetAnalyzer sets the analyzer for the match phrase prefix query parameter.
// The analyzer used to tokenize the query string text.
// Default is the index-time analyzer specified for the default_field.
// If no analyzer is specified for the default_field, the analyzer is the default analyzer for the index.
func (p *Param) SetAnalyzer(value string) *Param {
	p.Analyzer = value
	return p
}

// SetMaxExpansions sets the max expansions for the match phrase prefix query parameter.
// The maximum number of terms to which the query can expand.
// Fuzzy queries “expand to” a number of matching terms that are within the distance specified in fuzziness.
// Then OpenSearch tries to match those terms.
// Default is 50.
func (p *Param) SetMaxExpansions(value int) *Param {
	p.MaxExpansions = value
	return p
}

// SetSlop sets the slop for the match phrase prefix query parameter.
// Controls the degree to which words in a query can be misordered and still be considered a match.
// From the Lucene documentation: “The number of other words permitted between words in query phrase.
// For example, to switch the order of two words requires two moves (the first move places the words atop one another),
// so to permit reorderings of phrases, the slop must be at least two. A value of zero requires an exact match.”
func (p *Param) SetSlop(value int) *Param {
	p.Slop = value
	return p
}

// SetZeroTermsQuery sets the zero terms query value for the match phrase prefix query parameter.
// In some cases, the analyzer removes all terms from a query string.
// For example, the stop analyzer removes all terms from the string.
// In those cases, zero_terms_query specifies whether to match no documents (none) or all documents (all).
// Valid values are none and all.
// Default is none.
func (p *Param) SetZeroTermsQuery(value match.ZeroTermType) *Param {
	p.ZeroTermsQuery = value
	return p
}
