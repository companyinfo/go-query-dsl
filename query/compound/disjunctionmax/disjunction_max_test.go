// Package disjunctionmax A disjunction max (dis_max) query returns any document that matches one or more query clauses.
// For documents that match multiple query clauses, the relevance score is set to the highest relevance score from all
// matching query clauses.
package disjunctionmax

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/companyinfo/go-query-dsl/query/term_level/terms"
)

type DisjunctionMaxTestSuite struct {
	suite.Suite
}

func (d *DisjunctionMaxTestSuite) TestNewDisjunctionMax_Success() {
	require := d.Require()
	dm := &DisjunctionMax{
		DisjunctionMax: Param{
			Queries: terms.New("title", "kimchy"),
		},
	}

	require.Equal(dm, New(terms.New("title", "kimchy")))
}

func (d *DisjunctionMaxTestSuite) TestSetTieBreaker_Success() {
	require := d.Require()
	dm := &DisjunctionMax{DisjunctionMax: Param{
		Queries:    terms.New("title", "kimchy"),
		TieBreaker: 0.7,
	}}

	require.Equal(dm, New(terms.New("title", "kimchy")).SetTieBreaker(0.7))
}

func TestDisjunctionMax(t *testing.T) {
	suite.Run(t, new(DisjunctionMaxTestSuite))
}
