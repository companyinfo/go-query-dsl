// Package boosting returns documents matching a positive query while reducing the relevance score of documents
// that also match a negative query.
package boosting

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/full_text/match"
)

type BoostingTestSuite struct {
	suite.Suite
}

func (b *BoostingTestSuite) TestNewBoosting_Success() {
	require := b.Require()
	boosting := &Boosting{}

	require.Equal(boosting, New())
}

func (b *BoostingTestSuite) TestSetPositive_Success() {
	require := b.Require()
	must := match.New("field_1", match.NewParam("q"))
	boosting := &Boosting{Boosting: Params{
		Positive: must,
	}}

	require.Equal(boosting, New().SetPositive(must))
}

func (b *BoostingTestSuite) TestSetNegative_Success() {
	require := b.Require()
	must := match.New("field_1", match.NewParam("q"))
	boosting := &Boosting{Boosting: Params{
		Negative: must,
	}}

	require.Equal(boosting, New().SetNegative(must))
}

func (b *BoostingTestSuite) TestSetNegativeBoost_Success() {
	require := b.Require()
	boosting := &Boosting{Boosting: Params{
		NegativeBoost: 1.0,
	}}

	require.Equal(boosting, New().SetNegativeBoost(1.0))
}

func TestBoosting(t *testing.T) {
	suite.Run(t, new(BoostingTestSuite))
}
