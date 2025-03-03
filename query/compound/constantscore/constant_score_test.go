// Package constantscore wraps a filter query and returns every matching document with a relevance
// score equal to the boost parameter value.
package constantscore

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/term_level/term"
)

type ConstantScoreTestSuite struct {
	suite.Suite
}

func (c *ConstantScoreTestSuite) TestNewConstantScore_Success() {
	require := c.Require()
	cs := &ConstantScore{}

	require.Equal(cs, New())
}

func (c *ConstantScoreTestSuite) TestSetFilter_Success() {
	require := c.Require()
	cs := &ConstantScore{ConstantScore: Param{
		Filter: term.New("field_1", term.NewParam("kimchy")),
	}}

	require.Equal(cs, New().SetFilter(term.New("field_1", term.NewParam("kimchy"))))
}

func (c *ConstantScoreTestSuite) TestSetBoost_Success() {
	require := c.Require()
	cs := &ConstantScore{ConstantScore: Param{
		Boost: 1.0,
	}}

	require.Equal(cs, New().SetBoost(1.0))
}

func TestConstantScore(t *testing.T) {
	suite.Run(t, new(ConstantScoreTestSuite))
}
