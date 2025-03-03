// Package query provides a flexible and convenient way to build queries in Go.
package query

import (
	"encoding/json"
	"fmt"

	"go.companyinfo.dev/go-query-dsl/query/aggregation"
	"go.companyinfo.dev/go-query-dsl/query/sort"
	"go.companyinfo.dev/go-query-dsl/query/suggestion"
)

// Builder represents an query builder instance.
type Builder struct {

	// Sort specifies the sorting criteria for the query.
	Sort []sort.Sort `json:"sort,omitempty"`

	// Query represents the search query.
	Query any `json:"query,omitempty"`

	// Aggregations contains the aggregations for the query.
	Aggregations map[string]*aggregation.Aggregation `json:"aggs,omitempty"`

	// Suggestion represents query suggestions.
	Suggestion *suggestion.Autocomplete `json:"suggest,omitempty"`

	// From specifies the starting index for paginated results.
	From int `json:"from,omitempty"`

	// Size specifies the number of results to return.
	Size int `json:"size"`

	// Fields retrieves values for the specific fields.
	Fields []string `json:"fields,omitempty"`

	// Source accesses the original data that was passed at index time.
	Source any `json:"_source,omitempty"`
}

// New creates a new instance of Builder.
func New() *Builder {
	return &Builder{
		Sort:         make([]sort.Sort, 0),
		Aggregations: make(map[string]*aggregation.Aggregation),
		Fields:       make([]string, 0),
	}
}

// SetQuery sets the search query.
func (qb *Builder) SetQuery(q any) *Builder {
	qb.Query = q
	return qb
}

// SetFrom sets the starting index for paginated results.
func (qb *Builder) SetFrom(value int) *Builder {
	qb.From = value
	return qb
}

// SetSize sets the number of results to return for the query.
func (qb *Builder) SetSize(value int) *Builder {
	qb.Size = value
	return qb
}

// AddFields sets the specific fields on response.
func (qb *Builder) AddFields(values []string) *Builder {
	qb.Fields = append(qb.Fields, values...)
	return qb
}

// SetSource sets the access to the original data.
func (qb *Builder) SetSource(value bool) *Builder {
	qb.Source = value
	return qb
}

// AddSort adds a sorting criteria to the query.
func (qb *Builder) AddSort(value sort.Sort) *Builder {
	qb.Sort = append(qb.Sort, value)
	return qb
}

// AddAggregation adds an aggregation to the query.
func (qb *Builder) AddAggregation(name string, ag *aggregation.Aggregation) *Builder {
	if qb.Aggregations == nil {
		qb.Aggregations = make(map[string]*aggregation.Aggregation)
	}

	qb.Aggregations[name] = ag
	return qb
}

// AddSuggestion adds query suggestions to the query.
func (qb *Builder) AddSuggestion(sg *suggestion.Autocomplete) *Builder {
	qb.Suggestion = sg
	return qb
}

// Build converts the query to JSON.
func (qb *Builder) Build() ([]byte, error) {
	// Convert the query to JSON
	query, err := json.Marshal(qb)
	if err != nil {
		return nil, fmt.Errorf("error on converting the query to JSON: %w", err)
	}

	return query, nil
}

// String returns the JSON representation of the query.
func (qb *Builder) String() string {
	query, _ := qb.Build()
	return string(query)
}
