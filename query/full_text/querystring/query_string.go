// Package querystring provides structures and functions for building query string queries.
// Returns documents based on a provided query string, using a parser with a strict syntax.
package querystring

import "go.companyinfo.dev/go-query-dsl/query/full_text/match"

// QueryString represents a query_string query.
type QueryString struct {
	QueryString *Param `json:"query_string"`
}

// Param represents the parameters for a query_string query.
type Param struct {
	Query                           string             `json:"query"`
	AllowLeadingWildcard            bool               `json:"allow_leading_wildcard,omitempty"`
	AnalyzeWildcard                 bool               `json:"analyze_wildcard,omitempty"`
	Analyzer                        string             `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool               `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           float64            `json:"boost,omitempty"`
	DefaultFields                   string             `json:"default_fields,omitempty"`
	DefaultOperator                 match.OperatorType `json:"default_operator,omitempty"`
	EnablePositionIncrements        bool               `json:"enable_position_increments,omitempty"`
	Fields                          []string           `json:"fields,omitempty"`
	Fuzziness                       string             `json:"fuzziness,omitempty"`
	FuzzyMaxExpansions              uint               `json:"fuzzy_max_expansions,omitempty"`
	FuzzyTranspositions             bool               `json:"fuzzy_transpositions,omitempty"`
	Lenient                         bool               `json:"lenient,omitempty"`
	MaxDeterminizedStates           uint               `json:"max_determinized_states,omitempty"`
	MinimumShouldMatch              int                `json:"minimum_should_match,omitempty"`
	PhraseSlop                      int                `json:"phrase_slop,omitempty"`
	QuoteAnalyzer                   string             `json:"quote_analyzer,omitempty"`
	QuoteFieldSuffix                string             `json:"quote_field_suffix,omitempty"`
	Rewrite                         string             `json:"rewrite,omitempty"`
	TimeZone                        string             `json:"time_zone,omitempty"`
}

// New creates a new QueryString instance with the specified parameters.
func New(p *Param) *QueryString {
	return &QueryString{QueryString: p}
}

// NewParam creates a new Param instance for query-string query parameter.
func NewParam(q string) *Param {
	return &Param{Query: q}
}

// SetAllowLeadingWildcard sets the allowed leading wildcard value for the query-string query parameter.
// Specifies whether * and ? are allowed as first characters of a search term.
// Default is true.
func (p *Param) SetAllowLeadingWildcard(value bool) *Param {
	p.AllowLeadingWildcard = value
	return p
}

// SetAnalyzeWildcard sets the analysis wildcard value for the query-string query parameter.
// Specifies whether OpenSearch should attempt to analyze wildcard terms.
// Default is false.
func (p *Param) SetAnalyzeWildcard(value bool) *Param {
	p.AnalyzeWildcard = value
	return p
}

// SetAnalyzer sets the analyzer for the query-string query parameter.
// The analyzer used to tokenize the query string text.
// Default is the index-time analyzer specified for the default_field.
// If no analyzer is specified for the default_field, the analyzer is the default analyzer for the index.
func (p *Param) SetAnalyzer(value string) *Param {
	p.Analyzer = value
	return p
}

// SetAutoGenerateSynonymsPhraseQuery sets the auto generate synonyms
// phrase query value for the query-string query parameter.
// Specifies whether to create a match phrase query automatically for multi-term synonyms.
// Default is true.
func (p *Param) SetAutoGenerateSynonymsPhraseQuery(value bool) *Param {
	p.AutoGenerateSynonymsPhraseQuery = value
	return p
}

// SetBoost sets the boost value for the query-string query parameter.
// Boosts the clause by the given multiplier. Useful for weighing clauses in compound queries.
// Values in the [0, 1] range decrease relevance, and values greater than 1 increase relevance.
// Default is 1.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetDefaultFields sets the default fields for the query-string query parameter.
// The field in which to search if the field is not specified in the query string.
// Supports wildcards.
// Defaults to the value specified in the index.query.
// Default_field index setting. By default, the index.query.
// Default_field is *, which means extract all fields eligible for term query and filter the metadata fields.
// The extracted fields are combined into a query if the prefix is not specified.
// Eligible fields do not include nested documents.
// Searching all eligible fields could be a resource-intensive operation.
// The indices.query.bool.max_clause_count search setting defines the maximum value for the product of the number
// of fields and the number of terms that can be queried at one time.
// The default value for indices.query.bool.max_clause_count is 1,024.
func (p *Param) SetDefaultFields(fields string) *Param {
	p.DefaultFields = fields
	return p
}

// SetDefaultOperator sets the default operator value for the query-string query parameter.
// If the query string contains multiple search terms, whether all terms need to match (AND) or only one
// term needs to match (OR) for a document to be considered a match.
func (p *Param) SetDefaultOperator(value match.OperatorType) *Param {
	p.DefaultOperator = value
	return p
}

// SetEnablePositionIncrements sets the enable position increments value for the query-string query parameter.
// When true, resulting queries are aware of position increments.
// This setting is useful when the removal of stop words leaves an unwanted “gap” between terms.
// Default is true.
func (p *Param) SetEnablePositionIncrements(value bool) *Param {
	p.EnablePositionIncrements = value
	return p
}

// SetFields sets the fields for the query-string query parameter.
// The list of fields to search (for example, "fields": ["title^4", "description"]).
// Supports wildcards. If unspecified, defaults to the index.query.
// Default_field setting, which defaults to ["*"].
func (p *Param) SetFields(fields []string) *Param {
	p.Fields = fields
	return p
}

// SetFuzziness sets the fuzziness for the query-string query parameter.
// The number of character edits (insert, delete, substitute) that it takes to change one word to another when
// determining whether a term matched a value.
// For example, the distance between wined and wind is 1.
// Valid values are non-negative integers or AUTO.
// The default, AUTO, chooses a value based on the length of each term and is a good choice for most use cases.
func (p *Param) SetFuzziness(value string) *Param {
	p.Fuzziness = value
	return p
}

// SetFuzzyMaxExpansions sets the fuzzy max expansions value for the query-string query parameter.
// The maximum number of terms to which the query can expand. Fuzzy queries “expand to” a number of matching terms
// that are within the distance specified in fuzziness.
// Then OpenSearch tries to match those terms.
// Default is 50.
func (p *Param) SetFuzzyMaxExpansions(value uint) *Param {
	p.FuzzyMaxExpansions = value
	return p
}

// SetFuzzyTranspositions sets the fuzzy transpositions value for the query-string query parameter.
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

// SetLenient sets the lenient value for the query-string query parameter.
// Setting lenient to true ignores data type mismatches between the query and the document field.
// For example, a query string of "8.2" could match a field of type float.
// Default is false.
func (p *Param) SetLenient(value bool) *Param {
	p.Lenient = value
	return p
}

// SetMaxDeterminizedStates sets the max determized states value for the query-string query parameter.
// Lucene converts a regular expression to an automaton with a number of determinized states.
// This parameter specifies the maximum number of automaton states the query requires.
// Use this parameter to prevent high resource consumption.
// To run complex regular expressions, you may need to increase the value of this parameter.
// Default is 10,000.
func (p *Param) SetMaxDeterminizedStates(value uint) *Param {
	p.MaxDeterminizedStates = value
	return p
}

// SetMinimumShouldMatch sets the minimum should match value for the query-string query parameter.
// If the query string contains multiple search terms, and you use the or operator, the number of terms that need
// to match for the document to be considered a match.
// For example, if minimum_should_match is 2, wind often rising does not match The Wind Rises.
// If minimum_should_match is 1, it matches.
func (p *Param) SetMinimumShouldMatch(value int) *Param {
	p.MinimumShouldMatch = value
	return p
}

// SetPhraseSlop sets the phrase slop value for the query-string query parameter.
// The maximum number of words that are allowed between the matched words.
// If phrase_slop is 2, a maximum of two words is allowed between matched words in a phrase.
// Transposed words have a slop of 2.
// Default is 0 (an exact phrase match where matched words must be next to each other).
func (p *Param) SetPhraseSlop(value int) *Param {
	p.PhraseSlop = value
	return p
}

// SetQuoteAnalyzer sets the quote analyzer of the query-string query parameter.
// The analyzer used to tokenize quoted text in the query string.
// Overrides the analyzer parameter for quoted text.
// Default is the search_quote_analyzer specified for the default_field.
func (p *Param) SetQuoteAnalyzer(value string) *Param {
	p.QuoteAnalyzer = value
	return p
}

// SetQuoteFieldSuffix sets the quote field suffix of the query-string query parameter.
// This option supports searching for exact matches (surrounded with quotation marks) using a different
// analysis method than non-exact matches use.
func (p *Param) SetQuoteFieldSuffix(value string) *Param {
	p.QuoteFieldSuffix = value
	return p
}

// SetRewrite sets the rewrite of the query-string query parameter.
// Determines how OpenSearch rewrites and scores multi-term queries.
// Valid values are constant_score, scoring_boolean, constant_score_boolean, top_terms_N, top_terms_boost_N,
// and top_terms_blended_freqs_N.
// Default is constant_score.
func (p *Param) SetRewrite(value string) *Param {
	p.Rewrite = value
	return p
}

// SetTimeZone sets the timezone of the query-string query parameter.
// Specifies the number of hours to offset the desired time zone from UTC.
// You need to indicate the time zone offset number if the query string contains a date range.
// For example, set time_zone": "-08:00" for a query with a date range such as
// "query": "wind rises release_date[2012-01-01 TO 2014-01-01]").
// The default time zone format used to specify number of offset hours is UTC.
func (p *Param) SetTimeZone(value string) *Param {
	p.TimeZone = value
	return p
}
