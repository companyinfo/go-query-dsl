// Package aggregation provides structures and functions for building aggregation.
// An aggregation summarizes your data as metrics, statistics, or other analytics.
package aggregation

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AggregationTestSuite struct {
	suite.Suite
}

func (a *AggregationTestSuite) TestNewAggregation_Success() {
	require := a.Require()
	aggs := &Aggregation{}

	require.Equal(aggs, New())
}

func (a *AggregationTestSuite) TestAddTerms_Success() {
	require := a.Require()
	aggs := &Aggregation{
		Terms: &TermsParam{
			Field: "field1",
			Size:  10,
		},
	}

	require.Equal(aggs, New().AddTerms("field1", 10))
}

func (a *AggregationTestSuite) TestAddAvg_Success() {
	require := a.Require()
	aggs := &Aggregation{
		Avg: &FieldParam{
			Field: "field1",
		},
	}

	require.Equal(aggs, New().AddAvg("field1"))
}

func (a *AggregationTestSuite) TestAddSum_Success() {
	require := a.Require()
	aggs := &Aggregation{
		Sum: &FieldParam{
			Field: "field1",
		},
	}

	require.Equal(aggs, New().AddSum("field1"))
}

func (a *AggregationTestSuite) TestAddMin_Success() {
	require := a.Require()
	aggs := &Aggregation{
		Min: &FieldParam{
			Field: "field1",
		},
	}

	require.Equal(aggs, New().AddMin("field1"))
}

func (a *AggregationTestSuite) TestAddMax_Success() {
	require := a.Require()
	aggs := &Aggregation{
		Max: &FieldParam{
			Field: "field1",
		},
	}

	require.Equal(aggs, New().AddMax("field1"))
}

func (a *AggregationTestSuite) TestAddHistogram_Success() {
	require := a.Require()
	aggs := &Aggregation{
		Histogram: &histogramParam{
			Field:    "field1",
			Interval: 10,
		},
	}

	require.Equal(aggs, New().AddHistogram("field1", 10))
}

func (a *AggregationTestSuite) TestAddDateHistogram_Success() {
	require := a.Require()
	aggs := &Aggregation{
		DateHistogram: &dateHistogramParam{
			Field:            "field1",
			CalendarInterval: "month",
		},
	}

	require.Equal(aggs, New().AddDateHistogram("field1", "month"))
}

func (a *AggregationTestSuite) TestAddRange_Success() {
	require := a.Require()
	aggs := &Aggregation{
		Range: &rangeParam{
			Field: "field1",
			Ranges: []map[string]any{
				{"from": 0, "to": 10},
				{"from": 10, "to": 20},
			},
		},
	}

	require.Equal(aggs, New().AddRange("field1", []map[string]any{
		{"from": 0, "to": 10},
		{"from": 10, "to": 20},
	}))
}

func (a *AggregationTestSuite) TestAddFilters_Success() {
	require := a.Require()
	aggs := &Aggregation{
		Filters: &filtersParam{
			Filters: map[string]map[string]any{
				"filter1": {
					"term": map[string]any{"field1": "value"},
				},
			},
		},
	}

	require.Equal(aggs, New().AddFilters(map[string]map[string]any{
		"filter1": {
			"term": map[string]any{"field1": "value"},
		},
	}))
}

func (a *AggregationTestSuite) TestAddNested_Success() {
	require := a.Require()
	aggs := &Aggregation{
		NestedAggregation: &nestedParam{
			Nested: &Aggregation{
				Terms: &TermsParam{
					Field: "field1",
					Size:  10,
				},
			},
		},
	}

	require.Equal(aggs, New().AddNested(New().AddTerms("field1", 10)))
}

func TestAggregation(t *testing.T) {
	suite.Run(t, new(AggregationTestSuite))
}
