// Package term provides structures and functions for building term queries.
// Returns documents that contain an exact term in a provided field.
package term

// Term represents a single term query.
type Term struct {
	Term map[string]*Param `json:"term"`
}

// Param represents the parameters for a term query, including the value and case sensitivity.
type Param struct {
	Value           any     `json:"value"`
	Boost           float64 `json:"boost,omitempty"`
	CaseInsensitive bool    `json:"case_insensitive,omitempty"`
}

// New creates a new Term instance with the specified field and parameter.
func New(f string, tp *Param) *Term {
	return &Term{Term: map[string]*Param{f: tp}}
}

// NewParam creates a new Param instance.
func NewParam(value any) *Param {
	return &Param{Value: value}
}

// SetBoost sets the boost for the term query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetCaseInsensitive sets the case sensitivity for the term query parameter.
// If true, allows case-insensitive matching of the value with the indexed field values.
// Default is false (case sensitivity is determined by the field’s mapping).
func (p *Param) SetCaseInsensitive(value bool) *Param {
	p.CaseInsensitive = value
	return p
}
