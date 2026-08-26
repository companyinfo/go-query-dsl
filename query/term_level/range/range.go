// Package ranging provides structures and functions for building range queries.
// Returns documents that contain terms within a provided range.
package ranging

import "errors"

// Value is the set of Go types a range bound (GT/GTE/LT/LTE) may hold. It covers every
// Elasticsearch field family a range query can target: numeric (long, integer, float, double, ...),
// term-like strings (keyword, ip, version), and dates via time.Time.
type Value any

// Range represents a range query, specifying range conditions for a field.
type Range[T Value] struct {
	Range map[string]*Param[T] `json:"range"`
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
type Param[T Value] struct {
	GT       *T           `json:"gt,omitempty"`
	GTE      *T           `json:"gte,omitempty"`
	LT       *T           `json:"lt,omitempty"`
	LTE      *T           `json:"lte,omitempty"`
	Format   string       `json:"format,omitempty"`
	Relation relationType `json:"relation,omitempty"`
	Timezone string       `json:"timezone,omitempty"`
	Boost    *float64     `json:"boost,omitempty"`
}

// New creates a new Range instance with the specified field and range parameters.
func New[T Value](f string, rp *Param[T]) *Range[T] {
	return &Range[T]{Range: map[string]*Param[T]{f: rp}}
}

// NewParam creates a new Param instance for range query parameters.
func NewParam[T Value]() *Param[T] {
	return &Param[T]{}
}

// SetGT sets the greater than (GT) condition in the range query parameter.
func (p *Param[T]) SetGT(value T) *Param[T] {
	p.GT = &value
	return p
}

// SetGTE sets the greater than or equal to (GTE) condition in the range query parameter.
func (p *Param[T]) SetGTE(value T) *Param[T] {
	p.GTE = &value
	return p
}

// SetLT sets the less than (LT) condition in the range query parameter.
func (p *Param[T]) SetLT(value T) *Param[T] {
	p.LT = &value
	return p
}

// SetLTE sets the less than or equal to (LTE) condition in the range query parameter.
func (p *Param[T]) SetLTE(value T) *Param[T] {
	p.LTE = &value
	return p
}

// SetFormat sets the format condition in the range query parameter.
// Date format used to convert date values in the query.
// Default is the field’s mapped format.
func (p *Param[T]) SetFormat(value string) *Param[T] {
	p.Format = value
	return p
}

// SetRelation sets the relation condition in the range query parameter.
// Indicates how the range query matches values for range fields. Valid values are:
// - INTERSECTS (default): Matches documents whose range field value intersects the range provided in the query.
// - CONTAINS: Matches documents whose range field value contains the entire range provided in the query.
// - WITHIN: Matches documents whose range field value is entirely within the range provided in the query.
func (p *Param[T]) SetRelation(value relationType) *Param[T] {
	p.Relation = value
	return p
}

// SetBoost sets the boost value in the range query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (p *Param[T]) SetBoost(value float64) *Param[T] {
	p.Boost = &value
	return p
}

// SetTimezone sets the timezone condition in the range query parameter.
// The time zone used to convert date values to UTC in the query.
// Valid values are ISO 8601 UTC offsets and IANA time zone IDs.
func (p *Param[T]) SetTimezone(value string) *Param[T] {
	p.Timezone = value
	return p
}

// Validate reports configuration mistakes that Elasticsearch would otherwise reject at query
// time, so callers can catch them while still building the request. It does not (and cannot,
// without a mapping lookup) check that the field itself accepts type T.
func (p *Param[T]) Validate() error {
	if p.GT != nil && p.GTE != nil {
		return errors.New("ranging: GT and GTE are mutually exclusive")
	}

	if p.LT != nil && p.LTE != nil {
		return errors.New("ranging: LT and LTE are mutually exclusive")
	}

	if p.GT == nil && p.GTE == nil && p.LT == nil && p.LTE == nil {
		return errors.New("ranging: at least one of GT, GTE, LT, LTE must be set")
	}

	return nil
}
