// Package query provides a flexible and convenient way to build queries in Go.
package query

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/aggregation"
	"go.companyinfo.dev/go-query-dsl/query/sort"
	"go.companyinfo.dev/go-query-dsl/query/suggestion"
)

type QueryTestSuite struct {
	suite.Suite
	queryBuilder *Builder
}

func (q *QueryTestSuite) SetupTest() {
	q.queryBuilder = &Builder{}
}

func (q *QueryTestSuite) TestNewQueryBuilder_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Sort:         make([]sort.Sort, 0),
		Aggregations: map[string]*aggregation.Aggregation{},
		Fields:       make([]string, 0),
	}

	require.Equal(q.queryBuilder, New())
}

func (q *QueryTestSuite) TestSetQuery_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Sort:         make([]sort.Sort, 0),
		Aggregations: map[string]*aggregation.Aggregation{},
		Query:        "query",
		Fields:       make([]string, 0),
	}

	require.Equal(q.queryBuilder, New().SetQuery("query"))
}

func (q *QueryTestSuite) TestSetFrom_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Sort:         make([]sort.Sort, 0),
		Aggregations: map[string]*aggregation.Aggregation{},
		From:         1,
		Fields:       make([]string, 0),
	}

	require.Equal(q.queryBuilder, New().SetFrom(1))
}

func (q *QueryTestSuite) TestSetSize_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Sort:         make([]sort.Sort, 0),
		Aggregations: map[string]*aggregation.Aggregation{},
		Size:         10,
		Fields:       make([]string, 0),
	}

	require.Equal(q.queryBuilder, New().SetSize(10))
}

func (q *QueryTestSuite) TestAddFields_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Sort:         make([]sort.Sort, 0),
		Aggregations: map[string]*aggregation.Aggregation{},
		Fields:       []string{"name"},
	}

	require.Equal(q.queryBuilder, New().AddFields([]string{"name"}))
}

func (q *QueryTestSuite) TestSetSource_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Sort:         make([]sort.Sort, 0),
		Aggregations: map[string]*aggregation.Aggregation{},
		Source:       false,
		Fields:       make([]string, 0),
	}

	require.Equal(q.queryBuilder, New().SetSource(false))
}

func (q *QueryTestSuite) TestAddSort_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Sort: []sort.Sort{
			map[string]*sort.Param{
				"field": {
					Order: "desc",
				},
			},
		},
		Aggregations: map[string]*aggregation.Aggregation{},
		Fields:       make([]string, 0),
	}

	require.Equal(q.queryBuilder, New().AddSort(sort.New("field", sort.NewParam().SetOrder("desc"))))
}

func (q *QueryTestSuite) TestAddAggregation_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Aggregations: map[string]*aggregation.Aggregation{
			"name": {
				Terms: &aggregation.TermsParam{
					Field: "field",
					Size:  10,
				},
			},
		},
		Sort:   make([]sort.Sort, 0),
		Fields: make([]string, 0),
	}

	require.Equal(
		q.queryBuilder,
		New().AddAggregation("name", aggregation.New().AddTerms("field", 10)),
	)
}

func (q *QueryTestSuite) TestAddAutocomplete_Success() {
	require := q.Require()
	q.queryBuilder = &Builder{
		Suggestion: &suggestion.Autocomplete{
			Autocomplete: &suggestion.Param{
				Prefix:     "prefix",
				Completion: &suggestion.Completion{Field: "field"},
			},
		},
		Sort:         make([]sort.Sort, 0),
		Aggregations: map[string]*aggregation.Aggregation{},
		Fields:       make([]string, 0),
	}

	require.Equal(
		q.queryBuilder,
		New().AddSuggestion(suggestion.NewAutocomplete(suggestion.NewAutocompleteParam("prefix", "field"))),
	)
}

func (q *QueryTestSuite) TestBuild_Success() {
	require := q.Require()
	expectedJSONQuery := `{
					"sort":[
						{"field":{
							"order":"desc"
							}
						}],
					"query":"query",
					"aggs":{
						"name":{
							"terms":{
								"field":"field",
								"size": 10
							}
						}
					},
					"suggest":{
						"autocomplete":{
							"prefix":"prefix",
							"completion":{
								"field":"field"
							}
						}
					},
					"from":1,
					"size":10
					}`

	q.queryBuilder = New().
		SetQuery("query").
		SetFrom(1).
		SetSize(10).
		AddSort(sort.New("field", sort.NewParam().SetOrder("desc"))).
		AddAggregation("name", aggregation.New().AddTerms("field", 10)).
		AddSuggestion(suggestion.NewAutocomplete(suggestion.NewAutocompleteParam("prefix", "field")))
	query, err := q.queryBuilder.Build()

	require.NoError(err)
	require.JSONEq(expectedJSONQuery, string(query))
}

func TestQueryBuilder(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
