// Package boolean provides structures and functions for building bool queries.
package boolean

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/full_text/match"
	"go.companyinfo.dev/go-query-dsl/query/term_level/term"
)

type BoolTestSuite struct {
	suite.Suite
	bool *Bool
}

func (b *BoolTestSuite) SetupTest() {
	b.bool = &Bool{Bool: QueryParams{
		Must:    make([]any, 0),
		MustNot: make([]any, 0),
		Should:  make([]any, 0),
		Filter:  make([]any, 0),
	}}
}

func (b *BoolTestSuite) TestNewBool_Success() {
	require := b.Require()

	require.Equal(b.bool, New())
}

func (b *BoolTestSuite) TestAddMust_Success() {
	require := b.Require()
	must := match.New("field_1", match.NewParam("q"))
	b.bool.AddMust(must)

	require.Equal(must, b.bool.Bool.Must[0])
	require.Equal(1, len(b.bool.Bool.Must))
}

func (b *BoolTestSuite) TestAddMustNot_Success() {
	require := b.Require()
	mustNot := match.New("field_1", match.NewParam("q"))
	b.bool.AddMustNot(mustNot)

	require.Equal(mustNot, b.bool.Bool.MustNot[0])
	require.Equal(1, len(b.bool.Bool.MustNot))
}

func (b *BoolTestSuite) TestAddShould_Success() {
	require := b.Require()
	should := match.New("field_1", match.NewParam("q"))
	b.bool.AddShould(should)

	require.Equal(should, b.bool.Bool.Should[0])
	require.Equal(1, len(b.bool.Bool.Should))
}

func (b *BoolTestSuite) TestAddFilter_Success() {
	require := b.Require()
	filter := term.New("field_1", term.NewParam("kimchy"))
	b.bool.AddFilter(filter)

	require.Equal(filter, b.bool.Bool.Filter[0])
	require.Equal(1, len(b.bool.Bool.Filter))
}

func (b *BoolTestSuite) TestSetMinimumShouldMatch_Success() {
	require := b.Require()
	b.bool.SetMinimumShouldMatch(10)

	require.Equal(10, b.bool.Bool.MinimumShouldMatch)
}

func (b *BoolTestSuite) TestSetBoost_Success() {
	require := b.Require()
	b.bool.SetBoost(1.0)

	require.Equal(1.0, b.bool.Bool.Boost)
}

func TestBool(t *testing.T) {
	suite.Run(t, new(BoolTestSuite))
}
