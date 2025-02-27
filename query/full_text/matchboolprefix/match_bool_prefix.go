// Package matchboolprefix provides structures and functions for building match boolean prefix queries.
// A match_bool_prefix query analyzes its input and constructs a bool query from the terms.
package matchboolprefix

import "github.com/companyinfo/go-query-dsl/query/full_text/match"

// MatchBoolPrefix represents a match boolean prefix query.
type MatchBoolPrefix struct {
	MatchBoolPrefix map[string]*Param `json:"match_bool_prefix"`
}

// Param represents the parameters for a match boolean prefix query,
// including query, analyzer, fuzziness, and more.
type Param struct {
	Query               string             `json:"query"`
	Analyzer            string             `json:"analyzer,omitempty"`
	Fuzziness           string             `json:"fuzziness,omitempty"`
	FuzzyRewrite        string             `json:"fuzzy_rewrite,omitempty"`
	FuzzyTranspositions bool               `json:"fuzzy_transpositions,omitempty"`
	MaxExpansions       int                `json:"max_expansions,omitempty"`
	MinimumShouldMatch  int                `json:"minimum_should_match,omitempty"`
	Operator            match.OperatorType `json:"operator,omitempty"`
	PrefixLength        int                `json:"prefix_length,omitempty"`
}

// New creates a new MatchBoolPrefix instance with the specified field and matchBoolPrefix boolean prefix parameters.
func New(f string, p *Param) *MatchBoolPrefix {
	return &MatchBoolPrefix{MatchBoolPrefix: map[string]*Param{f: p}}
}

// NewParam creates a new Param instance for match boolean prefix query parameters with the specified query.
func NewParam(q string) *Param {
	return &Param{Query: q}
}

// SetAnalyzer sets the analyzer for the match boolean prefix query parameter.
// The analyzer used to tokenize the query string text. Default is the index-time analyzer
// specified for the default_field.
// If no analyzer is specified for the default_field, the analyzer is the default analyzer for the index.
func (p *Param) SetAnalyzer(value string) *Param {
	p.Analyzer = value
	return p
}

// SetOperator sets the operator for the match boolean prefix query parameter.
// If the query string contains multiple search terms, whether all terms need to match (and) or only one term needs
// to match (or) for a document to be considered a match.
// Valid values are or & and.
// Default is or.
func (p *Param) SetOperator(value match.OperatorType) *Param {
	p.Operator = value
	return p
}

// SetMinimumShouldMatch sets the minimum should matchBoolPrefix value
// for the match boolean prefix query parameter.
// If the query string contains multiple search terms, and you use the or operator, the number of terms that need
// to match for the document to be considered a match.
// For example, if minimum_should_match is 2, wind often rising does not match The Wind Rises.
// If minimum_should_match is 1, it matches.
func (p *Param) SetMinimumShouldMatch(value int) *Param {
	p.MinimumShouldMatch = value
	return p
}

// SetMaxExpansions sets the max expansions value for the match boolean prefix query parameter.
// The maximum number of terms to which the query can expand.
// Fuzzy queries “expand to” a number of matching terms that are within the distance specified in fuzziness.
// Then OpenSearch tries to match those terms.
// Default is 50.
func (p *Param) SetMaxExpansions(value int) *Param {
	p.MaxExpansions = value
	return p
}

// SetFuzziness sets the fuzziness for the match boolean prefix query parameter.
// The number of character edits (insert, delete, substitute) that it takes to change one word
// to another when determining whether a term matched a value.
// For example, the distance between wined and wind is 1.
// The default, AUTO, chooses a value based on the length of each term and is a good choice for most use cases.
func (p *Param) SetFuzziness(value string) *Param {
	p.Fuzziness = value
	return p
}

// SetFuzzyRewrite sets the fuzzy rewrite value for the match boolean prefix query parameter.
// Determines how OpenSearch rewrites the query. Valid values are constant_score, scoring_boolean,
// constant_score_boolean, top_terms_N, top_terms_boost_N, and top_terms_blended_freqs_N.
// If the fuzziness parameter is not 0, the query uses a fuzzy_rewrite method of
// top_terms_blended_freqs_${max_expansions} by default.
// Default is constant_score.
func (p *Param) SetFuzzyRewrite(value string) *Param {
	p.FuzzyRewrite = value
	return p
}

// SetPrefixLength sets the prefix length value for the match boolean prefix query parameter.
// The number of leading characters that are not considered in fuzziness.
// Default is 0.
func (p *Param) SetPrefixLength(value int) *Param {
	p.PrefixLength = value
	return p
}

// SetFuzzyTranspositions sets the fuzzy transpositions value for the match boolean prefix query parameter.
// Setting fuzzy_transpositions to true (default) adds swaps of adjacent characters to the insert, delete,
// and substitute operations of the fuzziness option.
// For example, the distance between wind and wnid is 1 if fuzzy_transpositions is true (swap “n” and “i”)
// and 2 if it is false (delete “n”, insert “n”).
// If fuzzy_transpositions is false, rewind and wnid have the same distance (2) from wind, despite the more
// human-centric opinion that wnid is an obvious typo. The default is a good choice for most use cases.
func (p *Param) SetFuzzyTranspositions(value bool) *Param {
	p.FuzzyTranspositions = value
	return p
}
