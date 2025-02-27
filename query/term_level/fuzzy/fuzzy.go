// Package fuzzy provides structures and functions for building fuzzy queries.
// Returns documents that contain terms similar to the search term, as measured by a Levenshtein edit distance.
package fuzzy

// Fuzzy searches for documents containing terms that are similar to the search term within the
// maximum allowed Levenshtein distance.
// The Levenshtein distance measures the number of one-character changes needed to change one term to another term.
type Fuzzy struct {
	Fuzzy map[string]*Param `json:"fuzzy"`
}

// Param represents the parameters for a fuzzy query, including value, boost, rewrite, fuzziness, max expansions,
// prefix length, and transpositions.
type Param struct {
	Value          string  `json:"value"`
	Boost          float64 `json:"boost,omitempty"`
	Fuzziness      string  `json:"fuzziness,omitempty"`
	MaxExpansions  int     `json:"max_expansions,omitempty"`
	PrefixLength   int     `json:"prefix_length,omitempty"`
	Rewrite        string  `json:"rewrite,omitempty"`
	Transpositions bool    `json:"transpositions,omitempty"`
}

// New creates a new Fuzzy instance.
func New(f string, fp *Param) *Fuzzy {
	return &Fuzzy{Fuzzy: map[string]*Param{f: fp}}
}

// NewParam creates a new Param instance for the fuzzy query.
// value is the term to search for in the field specified in field.
func NewParam(value string) *Param {
	return &Param{Value: value}
}

// SetBoost sets the boost for the fuzzy query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetFuzziness sets the fuzziness for the fuzzy query parameter.
// The number of character edits (insert, delete, substitute) needed to change one word to another when
// determining whether a term matched a value.
// For example, the distance between wined and wind is 1.
// The default, AUTO, chooses a value based on the length of each term and is a good choice for most use cases.
func (p *Param) SetFuzziness(value string) *Param {
	p.Fuzziness = value
	return p
}

// SetPrefixLength sets the prefix length for the fuzzy query parameter.
// The number of leading characters that are not considered in fuzziness.
// Default is 0.
func (p *Param) SetPrefixLength(value int) *Param {
	p.PrefixLength = value
	return p
}

// SetMaxExpansions sets the max expansions for the fuzzy query parameter.
// The maximum number of terms to which the query can expand.
// Fuzzy queries “expand to” a number of matching terms that are within the distance specified in fuzziness.
// Then OpenSearch tries to match those terms.
// Default is 50.
func (p *Param) SetMaxExpansions(value int) *Param {
	p.MaxExpansions = value
	return p
}

// SetRewrite sets the rewrite for the fuzzy query parameter.
// Determines how OpenSearch rewrites and scores multi-term queries.
// Valid values are constant_score, scoring_boolean, constant_score_boolean, top_terms_N, top_terms_boost_N,
// and top_terms_blended_freqs_N.
// Default is constant_score.
func (p *Param) SetRewrite(value string) *Param {
	p.Rewrite = value
	return p
}

// SetTranspositions sets the transpositions for the fuzzy query parameter.
// Specifies whether to allow transpositions of two adjacent characters (ab to ba) as edits.
// Default is true.
func (p *Param) SetTranspositions(value bool) *Param {
	p.Transpositions = value
	return p
}
