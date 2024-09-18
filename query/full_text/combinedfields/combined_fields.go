// Package combinedfields provides structures and functions for building combined fields queries.
// The combined_fields query supports searching multiple text fields as if
// their contents had been indexed into one combined field.
package combinedfields

import "github.com/companyinfo/go-query-dsl/query/full_text/match"

// CombinedFields represents a combined fields query.
type CombinedFields struct {
	CombinedFields map[string]*Param `json:"combined_fields"`
}

// Param represents the parameters for a combined fields query.
type Param struct {
	Fields                          []string           `json:"fields"`
	Query                           string             `json:"query"`
	AutoGenerateSynonymsPhraseQuery bool               `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Operator                        match.OperatorType `json:"operator,omitempty"`
	MinimumShouldMatch              int                `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery                  match.ZeroTermType `json:"zero_terms_query,omitempty"`
}

// New creates a new CombinedFields instance with the specified field and combined fields parameters.
func New(f string, mp *Param) *CombinedFields {
	return &CombinedFields{CombinedFields: map[string]*Param{f: mp}}
}

// NewParam creates a new Param instance for combined fields query parameters with the specified query and fields.
func NewParam(q string, fields []string) *Param {
	return &Param{
		Query:  q,
		Fields: fields,
	}
}

// SetAutoGenerateSynonymsPhraseQuery sets the auto generate synonyms phrase query value
// for the combined fields query parameter.
// Specifies whether to create a match phrase query automatically for multi-term synonyms.
// Default is true.
func (p *Param) SetAutoGenerateSynonymsPhraseQuery(value bool) *Param {
	p.AutoGenerateSynonymsPhraseQuery = value
	return p
}

// SetOperator sets the operator for the combined fields query parameter.
// If the query string contains multiple search terms, whether all terms need to match (AND) or
// only one term needs to match (OR) for a document to be considered a match.
func (p *Param) SetOperator(value match.OperatorType) *Param {
	p.Operator = value
	return p
}

// SetMinimumShouldMatch sets the minimum should match value for the combined fields query parameter.
// If the query string contains multiple search terms, and you use the or operator, the number of terms that need
// to match for the document to be considered a match.
// For example, if minimum_should_match is 2, wind often rising does not match The Wind Rises.
// If minimum_should_match is 1, it matches.
func (p *Param) SetMinimumShouldMatch(value int) *Param {
	p.MinimumShouldMatch = value
	return p
}

// SetZeroTermsQuery sets the zero terms query value for the combined fields query parameter.
// In some cases, the analyzer removes all terms from a query string.
// For example, the stop analyzer removes all terms from the string.
// In those cases, zero_terms_query specifies whether to match no documents (none) or all documents (all).
// Valid values are none and all.
// Default is none.
func (p *Param) SetZeroTermsQuery(value match.ZeroTermType) *Param {
	p.ZeroTermsQuery = value
	return p
}
