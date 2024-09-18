// Package wildcard provides structures and functions for building wildcard queries.
// Returns documents that contain terms matching a wildcard pattern.
package wildcard

// Wildcard represents a wildcard query.
type Wildcard struct {
	Wildcard map[string]*Param `json:"wildcard"`
}

// Param represents the parameters for a wildcard query, including value, boost,
// rewrite, and case sensitivity.
type Param struct {
	Value           string  `json:"value"`
	Boost           float64 `json:"boost,omitempty"`
	CaseInsensitive bool    `json:"case_insensitive,omitempty"`
	Rewrite         string  `json:"rewrite,omitempty"`
}

// New creates a new wildcard instance.
func New(f string, p *Param) *Wildcard {
	return &Wildcard{Wildcard: map[string]*Param{f: p}}
}

// NewParam creates a new Param instance for wildcard query parameters with the specified value.
func NewParam(value string) *Param {
	return &Param{Value: value}
}

// SetBoost sets the boost for the wildcard query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetCaseInsensitive sets the case sensitivity for the wildcard query parameter.
// If true, allows case-insensitive matching of the value with the indexed field values.
// Default is false (case sensitivity is determined by the field’s mapping).
func (p *Param) SetCaseInsensitive(value bool) *Param {
	p.CaseInsensitive = value
	return p
}

// SetRewrite sets the rewrite condition for the wildcard query parameter.
// Determines how OpenSearch rewrites and scores multi-term queries.
// Valid values are constant_score, scoring_boolean, constant_score_boolean, top_terms_N,
// top_terms_boost_N, and top_terms_blended_freqs_N.
// Default is constant_score.
func (p *Param) SetRewrite(value string) *Param {
	p.Rewrite = value
	return p
}
