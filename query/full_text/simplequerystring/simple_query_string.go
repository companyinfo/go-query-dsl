// Package simplequerystring provides structures and functions for building simple query string queries.
// Returns documents based on a provided query string, using a parser with a limited but fault-tolerant syntax.
package simplequerystring

import "go.companyinfo.dev/go-query-dsl/query/full_text/match"

// SimpleQueryString represents a simple query string query.
type SimpleQueryString struct {
	SimpleQueryString *Param `json:"simple_query_string"`
}

// Param represents the parameters for a simple query string query.
type Param struct {
	Query                           string             `json:"query"`
	AnalyzeWildcard                 bool               `json:"analyze_wildcard,omitempty"`
	Analyzer                        string             `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool               `json:"auto_generate_synonyms_phrase_query,omitempty"`
	DefaultOperator                 match.OperatorType `json:"default_operator,omitempty"`
	Fields                          []string           `json:"fields,omitempty"`
	Flags                           string             `json:"flags,omitempty"`
	FuzzyMaxExpansions              uint               `json:"fuzzy_max_expansions,omitempty"`
	FuzzyTranspositions             bool               `json:"fuzzy_transpositions,omitempty"`
	FuzzyPrefixLength               int                `json:"fuzzy_prefix_length,omitempty"`
	Lenient                         bool               `json:"lenient,omitempty"`
	MinimumShouldMatch              int                `json:"minimum_should_match,omitempty"`
	QuoteFieldSuffix                string             `json:"quote_field_suffix,omitempty"`
}

// New creates a new SimpleQueryString instance with the specified parameters.
func New(p *Param) *SimpleQueryString {
	return &SimpleQueryString{SimpleQueryString: p}
}

// NewParam creates a new Param instance for simple query string query parameter.
func NewParam(q string) *Param {
	return &Param{Query: q}
}

// SetAnalyzeWildcard sets the analysis wildcard value for the simple query string query parameter.
// Specifies whether OpenSearch should attempt to analyze wildcard terms.
// Default is false.
func (p *Param) SetAnalyzeWildcard(value bool) *Param {
	p.AnalyzeWildcard = value
	return p
}

// SetAnalyzer sets the analyzer for the simple query string query parameter.
// The analyzer used to tokenize the query string text.
// Default is the index-time analyzer specified for the default_field.
// If no analyzer is specified for the default_field, the analyzer is the default analyzer for the index.
func (p *Param) SetAnalyzer(value string) *Param {
	p.Analyzer = value
	return p
}

// SetAutoGenerateSynonymsPhraseQuery sets the auto generate synonyms
// phrase query value for the simple query string query parameter.
// Specifies whether to create a match phrase query automatically for multi-term synonyms.
// Default is true.
func (p *Param) SetAutoGenerateSynonymsPhraseQuery(value bool) *Param {
	p.AutoGenerateSynonymsPhraseQuery = value
	return p
}

// SetDefaultOperator sets the default operator value for the simple query string query parameter.
// If the query string contains multiple search terms, whether all terms need to match (AND) or only one
// term needs to match (OR) for a document to be considered a match.
func (p *Param) SetDefaultOperator(value match.OperatorType) *Param {
	p.DefaultOperator = value
	return p
}

// SetFields sets the fields for the simple query string query parameter.
// The list of fields to search (for example, "fields": ["title^4", "description"]).
// Supports wildcards. If unspecified, defaults to the index.query.
// Default_field setting, which defaults to ["*"].
func (p *Param) SetFields(fields []string) *Param {
	p.Fields = fields
	return p
}

// SetFlags sets the flags for the simple query string query parameter.
// A |-delimited string of flags to enable (for example, AND|OR|NOT).
// Default is ALL.
// You can explicitly set the value for default_field.
// For example, to return all titles, set it to "default_field": "title".
func (p *Param) SetFlags(value string) *Param {
	p.Flags = value
	return p
}

// SetFuzzyMaxExpansions sets the fuzzy max expansions value for the simple query string query parameter.
// The maximum number of terms to which the query can expand. Fuzzy queries “expand to” a number of matching terms
// that are within the distance specified in fuzziness.
// Then OpenSearch tries to match those terms.
// Default is 50.
func (p *Param) SetFuzzyMaxExpansions(value uint) *Param {
	p.FuzzyMaxExpansions = value
	return p
}

// SetFuzzyTranspositions sets the fuzzy transpositions value for the simple query string query parameter.
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

// SetFuzzyPrefixLength sets the fuzzy prefix length value for the simple query string query parameter.
// The number of beginning characters left unchanged for fuzzy matching. Default is 0.
func (p *Param) SetFuzzyPrefixLength(value int) *Param {
	p.FuzzyPrefixLength = value
	return p
}

// SetLenient sets the lenient value for the simple query string query parameter.
// Setting lenient to true ignores data type mismatches between the query and the document field.
// For example, a query string of "8.2" could match a field of type float.
// Default is false.
func (p *Param) SetLenient(value bool) *Param {
	p.Lenient = value
	return p
}

// SetMinimumShouldMatch sets the minimum should match value for the simple query string query parameter.
// If the query string contains multiple search terms, and you use the or operator, the number of terms that need
// to match for the document to be considered a match.
// For example, if minimum_should_match is 2, wind often rising does not match The Wind Rises.
// If minimum_should_match is 1, it matches.
func (p *Param) SetMinimumShouldMatch(value int) *Param {
	p.MinimumShouldMatch = value
	return p
}

// SetQuoteFieldSuffix sets the quote field suffix of the simple query string query parameter.
// This option supports searching for exact matches (surrounded with quotation marks) using a different
// analysis method than non-exact matches use.
func (p *Param) SetQuoteFieldSuffix(value string) *Param {
	p.QuoteFieldSuffix = value
	return p
}
