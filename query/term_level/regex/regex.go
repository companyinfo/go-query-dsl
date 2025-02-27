// Package regex provides structures and functions for building regex queries.
// Returns documents that contain terms matching a regular expression.
package regex

// Regex represents a regex query.
type Regex struct {
	Regex map[string]*Param `json:"regex"`
}

// Param represents the parameters for a regex query, including value, boost, flags,
// max determinized states, rewrite, and case sensitivity.
type Param struct {
	Value                 string  `json:"value"`
	Boost                 float64 `json:"boost,omitempty"`
	CaseInsensitive       bool    `json:"case_insensitive,omitempty"`
	Flags                 string  `json:"flags,omitempty"`
	MaxDeterminizedStates uint    `json:"max_determinized_states,omitempty"`
	Rewrite               string  `json:"rewrite,omitempty"`
}

// New creates a new regex instance.
func New(f string, pp *Param) *Regex {
	return &Regex{Regex: map[string]*Param{f: pp}}
}

// NewParam creates a new Param instance for regex query parameters with the specified value.
func NewParam(value string) *Param {
	return &Param{Value: value}
}

// SetBoost sets the boost for the regex query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetCaseInsensitive sets the case sensitivity for the regex query parameter.
// If true, allows case-insensitive matching of the value with the indexed field values.
// Default is false (case sensitivity is determined by the field’s mapping).
func (p *Param) SetCaseInsensitive(value bool) *Param {
	p.CaseInsensitive = value
	return p
}

// SetFlags sets the flags for the regex query parameter.
// Enables optional operators for Lucene’s regular expression engine.
func (p *Param) SetFlags(value string) *Param {
	p.Flags = value
	return p
}

// SetMaxDeterminizedStates sets the max determized states value for the regex query parameter.
// Lucene converts a regular expression to an automaton with a number of determinized states.
// This parameter specifies the maximum number of automaton states the query requires.
// Use this parameter to prevent high resource consumption.
// To run complex regular expressions, you may need to increase the value of this parameter.
// Default is 10,000.
func (p *Param) SetMaxDeterminizedStates(value uint) *Param {
	p.MaxDeterminizedStates = value
	return p
}

// SetRewrite sets the rewrite condition for the regex query parameter.
// Determines how OpenSearch rewrites and scores multi-term queries.
// Valid values are constant_score, scoring_boolean, constant_score_boolean, top_terms_N,
// top_terms_boost_N, and top_terms_blended_freqs_N.
// Default is constant_score.
func (p *Param) SetRewrite(value string) *Param {
	p.Rewrite = value
	return p
}
