// Package sort provides structures and functions for building sort parameters in queries.
package sort

type (
	// Order defines the ordering type.
	Order string
	// Mode defines the mode type.
	Mode string
	// NumericType defines the numeric type.
	NumericType string
)

const (
	// DescOrder is descending order.
	DescOrder Order = "desc"
	// AscOrder is ascending order.
	AscOrder Order = "asc"

	// MinMode picks the lowest value.
	MinMode Mode = "min"
	// MaxMode picks the highest value.
	MaxMode Mode = "max"
	// SumMode uses the sum of all values as sort value.
	SumMode Mode = "sum"
	// AvgMode uses the average of all values as sort value.
	AvgMode Mode = "avg"
	// MedianMode uses the median of all values as sort value.
	MedianMode Mode = "median"

	// DoubleNumericType forces double type for all indices.
	DoubleNumericType NumericType = "double"
	// LongNumericType forces long type for all indices.
	LongNumericType NumericType = "long"
	// DateNumericType forces date type for all indices.
	DateNumericType NumericType = "date"
	// DateNanosNumericType forces date nanos type for all indices.
	DateNanosNumericType NumericType = "date_nanos"
)

// Sort represents the sorting criteria for a query.
type Sort map[string]*Param

// nested sorts by fields that are inside one or more nested objects.
type nested struct {
	Path   string `json:"path,omitempty"`
	Filter any    `json:"filter,omitempty"`
}

// Param represents the parameters for sorting, including order and mode.
type Param struct {
	Order       Order       `json:"order,omitempty"`
	Mode        Mode        `json:"mode,omitempty"`
	NumericType NumericType `json:"numeric_type,omitempty"`
	Nested      *nested     `json:"nested,omitempty"`
}

// New creates a new Sort instance with the specified field and sorting parameters.
func New(field string, param *Param) Sort {
	return map[string]*Param{field: param}
}

// NewParam creates a new Param instance.
func NewParam() *Param {
	return &Param{}
}

// SetOrder sets the order for sorting (asc or desc) in the sort parameter.
func (p *Param) SetOrder(o Order) *Param {
	p.Order = o
	return p
}

// SetMode sets the mode for sorting in the sort parameter.
func (p *Param) SetMode(m Mode) *Param {
	p.Mode = m
	return p
}

// SetNumericType casts the values from one type to another.
func (p *Param) SetNumericType(n NumericType) *Param {
	p.NumericType = n
	return p
}

// SetNested sets sorting by a nested field.
func (p *Param) SetNested(path string, filter any) *Param {
	p.Nested = &nested{
		Path:   path,
		Filter: filter,
	}
	return p
}
