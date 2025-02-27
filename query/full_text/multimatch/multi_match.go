// Package multimatch provides structures and functions for building multi-match queries.
// The multi_match query builds on the match query to allow multi-field queries
package multimatch

import "github.com/companyinfo/go-query-dsl/query/full_text/match"

// QueryType represents the type of multi-match query.
type QueryType string

// Possible values for QueryType, indicating different strategies for multi-match queries.
const (
	BestFields   QueryType = "best_fields"   // Use the best-performing field for each document.
	MostFields   QueryType = "most_fields"   // Sum the scores from all matching fields.
	CrossFields  QueryType = "cross_fields"  // Treat fields independently and combine their scores.
	Phrase       QueryType = "phrase"        // Match the full phrase as a single unit.
	PhrasePrefix QueryType = "phrase_prefix" // Match the prefix of a phrase.
	BoolPrefix   QueryType = "bool_prefix"   // Treat the query string as a boolean prefix query.
)

// MultiMatch represents a multi-match query.
type MultiMatch struct {
	MultiMatch *Param `json:"multi_match"`
}

// Param represents the parameters for a multi-match query, including query, analyzer, fuzziness, and more.
type Param struct {
	Query                           string             `json:"query"`
	Analyzer                        string             `json:"analyzer,omitempty"`
	Fuzziness                       string             `json:"fuzziness,omitempty"`
	FuzzyRewrite                    string             `json:"fuzzy_rewrite,omitempty"`
	ZeroTermsQuery                  string             `json:"zero_terms_query,omitempty"`
	Type                            QueryType          `json:"type,omitempty"`
	Operator                        match.OperatorType `json:"operator,omitempty"`
	Fields                          []string           `json:"fields,omitempty"`
	TieBreaker                      float64            `json:"tie_breaker,omitempty"`
	Slop                            int                `json:"slop,omitempty"`
	MinimumShouldMatch              int                `json:"minimum_should_match,omitempty"`
	MaxExpansions                   int                `json:"max_expansions,omitempty"`
	PrefixLength                    int                `json:"prefix_length,omitempty"`
	Boost                           float64            `json:"boost,omitempty"`
	Lenient                         bool               `json:"lenient,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool               `json:"auto_generate_synonyms_phrase_query,omitempty"`
	FuzzyTranspositions             bool               `json:"fuzzy_transpositions,omitempty"`
}

// New creates a new MultiMatch instance with the specified parameters.
func New(mmp *Param) *MultiMatch {
	return &MultiMatch{MultiMatch: mmp}
}

// NewParam creates a new Param instance for multi-match query parameters with the specified query.
func NewParam(q string) *Param {
	return &Param{Query: q}
}

// SetFields sets the fields for the multi-match query parameter.
// The list of fields in which to search. If you don’t provide the fields parameter,
// multi_match query searches the fields specified in the index.query.
// Default_field setting, which defaults to *.
func (p *Param) SetFields(fields []string) *Param {
	p.Fields = fields
	return p
}

// SetType sets the type of the multi-match query parameter.
// The multi-match query type. Valid values are best_fields, most_fields, cross_fields, phrase, phrase_prefix,
// bool_prefix.
// Default is best_fields.
func (p *Param) SetType(value QueryType) *Param {
	p.Type = value
	return p
}

// SetTieBreaker sets the tiebreaker value for the multi-match query parameter.
// A factor between 0 and 1.0 that is used to give more weight to documents that match multiple query clauses.
func (p *Param) SetTieBreaker(value float64) *Param {
	p.TieBreaker = value
	return p
}

// SetOperator sets the operator value for the multi-match query parameter.
// If the query string contains multiple search terms, whether all terms need to match (AND) or
// only one term needs to match (OR) for a document to be considered a match.
func (p *Param) SetOperator(value match.OperatorType) *Param {
	p.Operator = value
	return p
}

// SetAnalyzer sets the analyzer for the multi-match query parameter.
// The analyzer used to tokenize the query string text.
// Default is the index-time analyzer specified for the default_field.
// If no analyzer is specified for the default_field, the analyzer is the default analyzer for the index.
func (p *Param) SetAnalyzer(value string) *Param {
	p.Analyzer = value
	return p
}

// SetSlop sets the slop value for the multi-match query parameter.
// Controls the degree to which words in a query can be misordered and still be considered a match.
// From the Lucene documentation: “The number of other words permitted between words in query phrase.
// For example, to switch the order of two words requires two moves (the first move places the words atop one another),
// so to permit reorderings of phrases, the slop must be at least two. A value of zero requires an exact match.”
func (p *Param) SetSlop(value int) *Param {
	p.Slop = value
	return p
}

// SetFuzziness sets the fuzziness for the multi-match query parameter.
// The number of character edits (insert, delete, substitute) that it takes to change one word to another when
// determining whether a term matched a value.
// For example, the distance between wined and wind is 1.
// Valid values are non-negative integers or AUTO.
// The default, AUTO, chooses a value based on the length of each term and is a good choice for most use cases.
func (p *Param) SetFuzziness(value string) *Param {
	p.Fuzziness = value
	return p
}

// SetPrefixLength sets the prefix length for the multi-match query parameter.
// The number of leading characters that are not considered in fuzziness.
// Default is 0.
func (p *Param) SetPrefixLength(count int) *Param {
	p.PrefixLength = count
	return p
}

// SetBoost sets the boost value for the multi-match query parameter.
// Boosts the clause by the given multiplier. Useful for weighing clauses in compound queries.
// Values in the [0, 1] range decrease relevance, and values greater than 1 increase relevance.
// Default is 1.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetMinimumShouldMatch sets the minimum should match value for the multi-match query parameter.
// The maximum number of terms to which the query can expand.
// Fuzzy queries “expand to” a number of matching terms that are within the distance specified in fuzziness.
// Then OpenSearch tries to match those terms.
// Default is 50.
func (p *Param) SetMinimumShouldMatch(value int) *Param {
	p.MinimumShouldMatch = value
	return p
}

// SetMaxExpansions sets the max expansions value for the multi-match query parameter.
// The maximum number of terms to which the query can expand.
// Fuzzy queries “expand to” a number of matching terms that are within the distance specified in fuzziness.
// Then OpenSearch tries to match those terms.
// Default is 50.
func (p *Param) SetMaxExpansions(value int) *Param {
	p.MaxExpansions = value
	return p
}

// SetFuzzYRewrite sets the fuzzy rewrite value for the multi-match query parameter.
// Determines how OpenSearch rewrites the query.
// Valid values are constant_score, scoring_boolean, constant_score_boolean, top_terms_N, top_terms_boost_N,
// and top_terms_blended_freqs_N.
// If the fuzziness parameter is not 0, the query uses a fuzzy_rewrite method of
// top_terms_blended_freqs_${max_expansions} by default.
// Default is constant_score.
func (p *Param) SetFuzzYRewrite(value string) *Param {
	p.FuzzyRewrite = value
	return p
}

// SetZeroTermsQuery sets the zero terms query value for the multi-match query parameter.
// In some cases, the analyzer removes all terms from a query string.
// For example, the stop analyzer removes all terms from the string.
// In those cases, zero_terms_query specifies whether to match no documents (none) or all documents (all).
// Valid values are none and all.
// Default is none.
func (p *Param) SetZeroTermsQuery(value string) *Param {
	p.ZeroTermsQuery = value
	return p
}

// SetAutoGenerateSynonymsPhraseQuery sets the auto generate synonyms
// phrase query value for the multi-match query parameter.
// Specifies whether to create a match phrase query automatically for multi-term synonyms.
// Default is true.
func (p *Param) SetAutoGenerateSynonymsPhraseQuery(value bool) *Param {
	p.AutoGenerateSynonymsPhraseQuery = value
	return p
}

// SetFuzzyTranspositions sets the fuzzy transpositions value for the multi-match query parameter.
// Setting fuzzy_transpositions to true (default) adds swaps of adjacent characters to the insert,
// delete, and substitute operations of the fuzziness option.
// For example, the distance between wind and wnid is 1 if fuzzy_transpositions is true (swap “n” and “i”) and 2 if
// it is false (delete “n”, insert “n”).
// If fuzzy_transpositions is false, rewind and wnid have the same distance (2) from wind, despite the more
// human-centric opinion that wnid is an obvious typo.
// The default is a good choice for most use cases.
func (p *Param) SetFuzzyTranspositions(value bool) *Param {
	p.FuzzyTranspositions = value
	return p
}

// SetLenient sets the lenient value for the multi-match query parameter.
// Setting lenient to true ignores data type mismatches between the query and the document field.
// For example, a query string of "8.2" could match a field of type float.
// Default is false.
func (p *Param) SetLenient(value bool) *Param {
	p.Lenient = value
	return p
}
