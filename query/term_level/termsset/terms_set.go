// Package termsset provides structures and functions for building terms set queries.
// Returns documents that contain a minimum number of exact terms in a provided field.
package termsset

// TermsSet represents a single terms set query.
type TermsSet struct {
	TermsSet map[string]*Param `json:"terms_set"`
}

// Param represents the parameters for a terms set query, including the terms, boost, minimum should match field,
// and minimum should match script.
type Param struct {
	Terms                    []string `json:"terms"`
	Boost                    float64  `json:"boost,omitempty"`
	MinimumShouldMatchField  string   `json:"minimum_should_match_field,omitempty"`
	MinimumShouldMatchScript *source  `json:"minimum_should_match_script,omitempty"`
}

// source specifies the minimum number of terms a document should match with a script.
type source struct {
	Source string `json:"source,omitempty"`
}

// New creates a new Terms set instance with the specified field and parameter.
func New(f string, tp *Param) *TermsSet {
	return &TermsSet{TermsSet: map[string]*Param{f: tp}}
}

// NewParam creates a new Param instance.
func NewParam(value []string) *Param {
	return &Param{Terms: value}
}

// SetBoost sets the boost for the terms set query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetMinimumShouldMatchField sets the minimum_should_match_field for the terms set query parameter.
// The name of the numeric field that specifies the number of matching terms required in order to return a
// document in the results.
func (p *Param) SetMinimumShouldMatchField(value string) *Param {
	p.MinimumShouldMatchField = value
	return p
}

// SetMinimumShouldMatchScript sets the minimum_should_match_script for the terms set query parameter.
// A script that returns the number of matching terms required in order to return a document in the results.
func (p *Param) SetMinimumShouldMatchScript(value string) *Param {
	p.MinimumShouldMatchScript = &source{
		Source: value,
	}
	return p
}
