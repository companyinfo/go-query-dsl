// Package ranging provides structures and functions for building range queries.
// Returns documents that contain terms within a provided range.
package ranging

// Range represents a range query, specifying range conditions for a field.
type Range struct {
	Range map[string]*Param `json:"range"`
}

type relationType string

const (
	// IntersectsRelationType Matches documents whose range field value intersects the range provided in the query.
	IntersectsRelationType relationType = "INTERSECTS"
	// ContainsRelationType Matches documents whose range field value contains the entire range provided in the query.
	ContainsRelationType relationType = "CONTAINS"
	// WithinRelationType Matches documents whose range field value is entirely within the range provided in the query.
	WithinRelationType relationType = "WITHIN"
)

// Param represents the parameters for a range query, including greater than (GT), greater than or equal to (GTE),
// less than (LT), less than or equal to (LTE), format, relation, timezone, and boost.
type Param struct {
	GT       string       `json:"gt,omitempty"`
	GTE      string       `json:"gte,omitempty"`
	LT       string       `json:"lt,omitempty"`
	LTE      string       `json:"lte,omitempty"`
	Format   string       `json:"format,omitempty"`
	Relation relationType `json:"relation,omitempty"`
	Timezone string       `json:"timezone,omitempty"`
	Boost    float64      `json:"boost,omitempty"`
}

// New creates a new Range instance with the specified field and range parameters.
func New(f string, rp *Param) *Range {
	return &Range{Range: map[string]*Param{f: rp}}
}

// NewParam creates a new Param instance for range query parameters.
func NewParam() *Param {
	return &Param{}
}

// SetGT sets the greater than (GT) condition in the range query parameter.
func (p *Param) SetGT(value string) *Param {
	p.GT = value
	return p
}

// SetGTE sets the greater than or equal to (GTE) condition in the range query parameter.
func (p *Param) SetGTE(value string) *Param {
	p.GTE = value
	return p
}

// SetLT sets the less than (LT) condition in the range query parameter.
func (p *Param) SetLT(value string) *Param {
	p.LT = value
	return p
}

// SetLTE sets the less than or equal to (LTE) condition in the range query parameter.
func (p *Param) SetLTE(value string) *Param {
	p.LTE = value
	return p
}

// SetFormat sets the format condition in the range query parameter.
// Date format used to convert date values in the query.
// Default is the field’s mapped format.
func (p *Param) SetFormat(value string) *Param {
	p.Format = value
	return p
}

// SetRelation sets the relation condition in the range query parameter.
// Indicates how the range query matches values for range fields. Valid values are:
// - INTERSECTS (default): Matches documents whose range field value intersects the range provided in the query.
// - CONTAINS: Matches documents whose range field value contains the entire range provided in the query.
// - WITHIN: Matches documents whose range field value is entirely within the range provided in the query.
func (p *Param) SetRelation(value relationType) *Param {
	p.Relation = value
	return p
}

// SetBoost sets the boost value in the range query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param) SetBoost(value float64) *Param {
	p.Boost = value
	return p
}

// SetTimezone sets the timezone condition in the range query parameter.
// The time zone used to convert date values to UTC in the query.
// Valid values are ISO 8601 UTC offsets and IANA time zone IDs.
func (p *Param) SetTimezone(value string) *Param {
	p.Timezone = value
	return p
}
