// Package exists provides structures and functions for building exists queries.
// Returns documents that contain an indexed value for a field.
package exists

// Exists searches for documents that contain a specific field.
type Exists struct {
	Exists *Param `json:"exists"`
}

// Param represents the parameters for an exists query, including filed and boost.
// The query accepts the name of the field (<field>) as a top-level parameter.
type Param struct {
	Field string  `json:"field"`
	Boost float64 `json:"boost,omitempty"`
}

// New creates a new Exists instance with the specified field name.
func New(p *Param) *Exists {
	return &Exists{Exists: p}
}

// NewParam creates a new Param instance for the exists query.
// value is the name of the field.
func NewParam(value string) *Param {
	return &Param{Field: value}
}

// SetBoost sets the boost for the exists query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}
