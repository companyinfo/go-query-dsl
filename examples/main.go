// Package main provides a simple example for using the query builder.
package main

import (
	"github.com/companyinfo/go-query-dsl/query"
	"github.com/companyinfo/go-query-dsl/query/aggregation"
	"github.com/companyinfo/go-query-dsl/query/compound/boolean"
	"github.com/companyinfo/go-query-dsl/query/full_text/match"
	"github.com/companyinfo/go-query-dsl/query/sort"
)

func main() {
	qb := query.New()

	qb.SetQuery(boolean.New().SetBoost(1.0).AddMust(match.New("first_name", match.NewParam("tom")))).
		SetFrom(1).
		SetSize(10).
		AddSort(sort.New("age", sort.NewParam().SetOrder(sort.AscOrder).SetMode(sort.MinMode))).
		AddAggregation("age_aggregation", aggregation.New().AddTerms("age", 20))

	if q, err := qb.Build(); err != nil {
		println("failed on building a query: ", err.Error())
	} else {
		println(string(q))
	}
}
