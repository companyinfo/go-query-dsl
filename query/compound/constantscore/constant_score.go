// Package constantscore wraps a filter query and returns every matching document with a relevance
// score equal to the boost parameter value.
package constantscore

// ConstantScore represents a constant score query.
type ConstantScore struct {
	ConstantScore Param `json:"constant_score"`
}

// Param represents the parameters for a constant score query.
type Param struct {
	Filter any     `json:"filter"`
	Boost  float64 `json:"boost,omitempty"`
}

// New creates a new Bool instance with initialized query parameters.
func New() *ConstantScore {
	return &ConstantScore{}
}

// SetFilter the filter query that a document must match to be returned in the results.
func (c *ConstantScore) SetFilter(value any) *ConstantScore {
	c.ConstantScore.Filter = value
	return c
}

// SetBoost sets the boost for the constant score query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (c *ConstantScore) SetBoost(value float64) *ConstantScore {
	c.ConstantScore.Boost = value
	return c
}
