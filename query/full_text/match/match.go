// Package match provides structures and functions for building match queries.
// The match query is the standard query for performing a full-text search, including options for fuzzy matching.
package match

// Match represents a match query.
type Match struct {
	Match map[string]*Param `json:"match"`
}

type (
	// OperatorType represents the operator type for match queries.
	OperatorType string
	// ZeroTermType indicates whether no documents are returned if the analyzer removes all tokens.
	ZeroTermType string
)

// Possible values for OperatorType, indicating different logical operators for match queries.
const (
	AndOperator      OperatorType = "and"  // Use the logical AND operator.
	OrOperator       OperatorType = "or"   // Use the logical OR operator.
	NoneZeroTermType ZeroTermType = "none" // No documents are returned if the analyzer removes all tokens.
	AllZeroTermType  ZeroTermType = "all"  // Returns all documents.
)

// Param represents the parameters for a match query, including query, analyzer, fuzziness, and more.
type Param struct {
	Query                           string       `json:"query"`
	Analyzer                        string       `json:"analyzer,omitempty"`
	Fuzziness                       string       `json:"fuzziness,omitempty"`
	FuzzyRewrite                    string       `json:"fuzzy_rewrite,omitempty"`
	ZeroTermsQuery                  ZeroTermType `json:"zero_terms_query,omitempty"`
	Operator                        OperatorType `json:"operator,omitempty"`
	MinimumShouldMatch              int          `json:"minimum_should_match,omitempty"`
	MaxExpansions                   int          `json:"max_expansions,omitempty"`
	PrefixLength                    int          `json:"prefix_length,omitempty"`
	Boost                           float64      `json:"boost,omitempty"`
	Lenient                         bool         `json:"lenient,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool         `json:"auto_generate_synonyms_phrase_query,omitempty"`
	EnablePositionIncrements        bool         `json:"enable_position_increments,omitempty"`
	FuzzyTranspositions             bool         `json:"fuzzy_transpositions,omitempty"`
}

// New creates a new Match instance with the specified field and match parameters.
func New(f string, p *Param) *Match {
	return &Match{Match: map[string]*Param{f: p}}
}

// NewParam creates a new Param instance for match query parameters with the specified query.
func NewParam(q string) *Param {
	return &Param{Query: q}
}

// SetAnalyzer sets the analyzer for the match query parameter.
// The analyzer used to tokenize the query string text.
// Default is the index-time analyzer specified for the default_field.
// If no analyzer is specified for the default_field, the analyzer is the default analyzer for the index.
func (p *Param) SetAnalyzer(value string) *Param {
	p.Analyzer = value
	return p
}

// SetOperator sets the operator for the match query parameter.
// If the query string contains multiple search terms, whether all terms need to match (AND) or
// only one term needs to match (OR) for a document to be considered a match.
func (p *Param) SetOperator(value OperatorType) *Param {
	p.Operator = value
	return p
}

// SetMinimumShouldMatch sets the minimum should match value for the match query parameter.
// If the query string contains multiple search terms, and you use the or operator, the number of terms that need
// to match for the document to be considered a match.
// For example, if minimum_should_match is 2, wind often rising does not match The Wind Rises.
// If minimum_should_match is 1, it matches.
func (p *Param) SetMinimumShouldMatch(value int) *Param {
	p.MinimumShouldMatch = value
	return p
}

// SetMaxExpansions sets the max expansions value for the match query parameter.
// The maximum number of terms to which the query can expand.
// Fuzzy queries “expand to” a number of matching terms that are within the distance specified in fuzziness.
// Then OpenSearch tries to match those terms.
// Default is 50.
func (p *Param) SetMaxExpansions(value int) *Param {
	p.MaxExpansions = value
	return p
}

// SetFuzziness sets the fuzziness for the match query parameter.
// The number of character edits (insert, delete, substitute) that it takes to change one word to another when
// determining whether a term matched a value.
// For example, the distance between wined and wind is 1.
// Valid values are non-negative integers or AUTO.
// The default, AUTO, chooses a value based on the length of each term and is a good choice for most use cases.
func (p *Param) SetFuzziness(value string) *Param {
	p.Fuzziness = value
	return p
}

// SetFuzzyRewrite sets the fuzzy rewrite value for the match query parameter.
// Determines how OpenSearch rewrites the query.
// Valid values are constant_score, scoring_boolean, constant_score_boolean, top_terms_N, top_terms_boost_N,
// and top_terms_blended_freqs_N.
// If the fuzziness parameter is not 0, the query uses a fuzzy_rewrite method of
// top_terms_blended_freqs_${max_expansions} by default.
// Default is constant_score.
func (p *Param) SetFuzzyRewrite(value string) *Param {
	p.FuzzyRewrite = value
	return p
}

// SetZeroTermsQuery sets the zero terms query value for the match query parameter.
// In some cases, the analyzer removes all terms from a query string.
// For example, the stop analyzer removes all terms from the string.
// In those cases, zero_terms_query specifies whether to match no documents (none) or all documents (all).
// Valid values are none and all.
// Default is none.
func (p *Param) SetZeroTermsQuery(value ZeroTermType) *Param {
	p.ZeroTermsQuery = value
	return p
}

// SetPrefixLength sets the prefix length value for the match query parameter.
// The number of leading characters that are not considered in fuzziness.
// Default is 0.
func (p *Param) SetPrefixLength(value int) *Param {
	p.PrefixLength = value
	return p
}

// SetBoost sets the boost value for the match query parameter.
// Boosts the clause by the given multiplier. Useful for weighing clauses in compound queries.
// Values in the [0, 1] range decrease relevance, and values greater than 1 increase relevance.
// Default is 1.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetAutoGenerateSynonymsPhraseQuery sets the auto generate synonyms phrase query value for the match query parameter.
// Specifies whether to create a match phrase query automatically for multi-term synonyms.
// Default is true.
func (p *Param) SetAutoGenerateSynonymsPhraseQuery(value bool) *Param {
	p.AutoGenerateSynonymsPhraseQuery = value
	return p
}

// SetEnablePositionIncrements sets the enable position increments value for the match query parameter.
// When true, resulting queries are aware of position increments.
// This setting is useful when the removal of stop words leaves an unwanted “gap” between terms.
// Default is true.
func (p *Param) SetEnablePositionIncrements(value bool) *Param {
	p.EnablePositionIncrements = value
	return p
}

// SetFuzzyTranspositions sets the fuzzy transpositions value for the match query parameter.
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

// SetLenient sets the lenient value for the match query parameter.
// Setting lenient to true ignores data type mismatches between the query and the document field.
// For example, a query string of "8.2" could match a field of type float.
// Default is false.
func (p *Param) SetLenient(value bool) *Param {
	p.Lenient = value
	return p
}
