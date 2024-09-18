// Package aggregation provides structures and functions for building aggregation.
// An aggregation summarizes your data as metrics, statistics, or other analytics.
package aggregation

// Aggregation represents an aggregation query.
type Aggregation struct {
	Terms             *TermsParam         `json:"terms,omitempty"`
	Avg               *FieldParam         `json:"avg,omitempty"`
	Sum               *FieldParam         `json:"sum,omitempty"`
	Min               *FieldParam         `json:"min,omitempty"`
	Max               *FieldParam         `json:"max,omitempty"`
	Histogram         *histogramParam     `json:"histogram,omitempty"`
	DateHistogram     *dateHistogramParam `json:"date_histogram,omitempty"`
	Range             *rangeParam         `json:"range,omitempty"`
	Filters           *filtersParam       `json:"filters,omitempty"`
	NestedAggregation *nestedParam        `json:"aggs,omitempty"`
}

// TermsParam represents the terms parameters for the aggregation.
type TermsParam struct {
	Field string `json:"field"`
	Size  uint   `json:"size,omitempty"`
}

// FieldParam represents the field and value for the aggregation.
type FieldParam struct {
	Field string `json:"field"`
}

// histogramParam represents the histogram parameters for the aggregation.
type histogramParam struct {
	Field    string `json:"field"`
	Interval int    `json:"interval"`
}

// dateHistogramParam represents the date histogram parameters for the aggregation.
type dateHistogramParam struct {
	Field            string `json:"field"`
	CalendarInterval string `json:"calendar_interval"`
}

// rangeParam represents the range parameters for the aggregation.
type rangeParam struct {
	Field  string `json:"field"`
	Ranges any    `json:"ranges"`
}

// filtersParam represents the filters parameters for the aggregation.
type filtersParam struct {
	Filters any `json:"filters"`
}

// nestedParam represents the nested parameters for the aggregation.
type nestedParam struct {
	Nested *Aggregation `json:"nested"`
}

// New creates a new Aggregation instance.
func New() *Aggregation {
	return &Aggregation{}
}

// AddTerms adds a Terms aggregation to the query.
func (a *Aggregation) AddTerms(field string, size uint) *Aggregation {
	a.Terms = &TermsParam{
		Field: field,
		Size:  size,
	}

	return a
}

// AddAvg adds an Average aggregation to the query.
func (a *Aggregation) AddAvg(field string) *Aggregation {
	a.Avg = &FieldParam{Field: field}
	return a
}

// AddSum adds a Sum aggregation to the query.
func (a *Aggregation) AddSum(field string) *Aggregation {
	a.Sum = &FieldParam{Field: field}
	return a
}

// AddMin adds a Minimum aggregation to the query.
func (a *Aggregation) AddMin(field string) *Aggregation {
	a.Min = &FieldParam{Field: field}
	return a
}

// AddMax adds a Maximum aggregation to the query.
func (a *Aggregation) AddMax(field string) *Aggregation {
	a.Max = &FieldParam{Field: field}
	return a
}

// AddHistogram adds a Histogram aggregation to the query.
func (a *Aggregation) AddHistogram(field string, interval int) *Aggregation {
	a.Histogram = &histogramParam{
		Field:    field,
		Interval: interval,
	}
	return a
}

// AddDateHistogram adds a Date Histogram aggregation to the query.
func (a *Aggregation) AddDateHistogram(field, interval string) *Aggregation {
	a.DateHistogram = &dateHistogramParam{
		Field:            field,
		CalendarInterval: interval,
	}
	return a
}

// AddRange adds a Range aggregation to the query.
func (a *Aggregation) AddRange(field string, ranges any) *Aggregation {
	a.Range = &rangeParam{
		Field:  field,
		Ranges: ranges,
	}
	return a
}

// AddFilters adds a Filters aggregation to the query.
func (a *Aggregation) AddFilters(filters any) *Aggregation {
	a.Filters = &filtersParam{
		Filters: filters,
	}
	return a
}

// AddNested adds a Nested aggregation to the query.
func (a *Aggregation) AddNested(subAggregation *Aggregation) *Aggregation {
	a.NestedAggregation = &nestedParam{Nested: subAggregation}
	return a
}
