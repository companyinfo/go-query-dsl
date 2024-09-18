// Package ids provides structures and functions for building IDs queries.
// Returns documents based on their IDs.
// This query uses document IDs stored in the _id field.
package ids

// IDs represents an IDs query.
type IDs struct {
	IDs *Param `json:"ids"`
}

// Param represents the parameters for an ids query, including value and boost.
type Param struct {
	Values []int   `json:"values"`
	Boost  float64 `json:"boost,omitempty"`
}

// New creates a new IDs instance with the specified list of values.
func New(p *Param) *IDs {
	return &IDs{IDs: p}
}

// NewParam creates a new Param instance for the ids query.
func NewParam(value []int) *Param {
	return &Param{Values: value}
}

// SetBoost sets the boost for the ids query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}
