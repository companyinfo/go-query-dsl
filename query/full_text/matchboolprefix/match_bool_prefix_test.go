package matchboolprefix

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/full_text/match"
)

type MatchBoolPrefixTestSuite struct {
	suite.Suite
	matchBoolPrefix *MatchBoolPrefix
	param           *Param
}

func (m *MatchBoolPrefixTestSuite) SetupTest() {
	m.matchBoolPrefix = &MatchBoolPrefix{MatchBoolPrefix: map[string]*Param{"field": {}}}
	m.param = &Param{Query: "value"}
}

func (m *MatchBoolPrefixTestSuite) TestNewMatchBoolPrefix_Success() {
	require := m.Require()

	require.Equal(m.matchBoolPrefix, New("field", &Param{}))
}

func (m *MatchBoolPrefixTestSuite) TestNewParam_Success() {
	require := m.Require()

	require.Equal(m.param, NewParam("value"))
}

func (m *MatchBoolPrefixTestSuite) TestSetAnalyzer_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", Analyzer: "standard"}

	require.Equal(m.param, NewParam("value").SetAnalyzer("standard"))
}

func (m *MatchBoolPrefixTestSuite) TestSetFuzziness_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", Fuzziness: "AUTO"}

	require.Equal(m.param, NewParam("value").SetFuzziness("AUTO"))
}

func (m *MatchBoolPrefixTestSuite) TestSetPrefixLength_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", PrefixLength: 0}

	require.Equal(m.param, NewParam("value").SetPrefixLength(0))
}

func (m *MatchBoolPrefixTestSuite) TestSetMaxExpansions_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", MaxExpansions: 50}

	require.Equal(m.param, NewParam("value").SetMaxExpansions(50))
}

func (m *MatchBoolPrefixTestSuite) TestSetOperator_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", Operator: match.OrOperator}

	require.Equal(m.param, NewParam("value").SetOperator(match.OrOperator))
}

func (m *MatchBoolPrefixTestSuite) TestSetMinimumShouldMatch_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", MinimumShouldMatch: 1}

	require.Equal(m.param, NewParam("value").SetMinimumShouldMatch(1))
}

func (m *MatchBoolPrefixTestSuite) TestSetFuzzYRewrite_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", FuzzyRewrite: "constant_score"}

	require.Equal(m.param, NewParam("value").SetFuzzyRewrite("constant_score"))
}

func (m *MatchBoolPrefixTestSuite) TestSetFuzzyTranspositions_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", FuzzyTranspositions: true}

	require.Equal(m.param, NewParam("value").SetFuzzyTranspositions(true))
}

func TestMatchBoolPrefix(t *testing.T) {
	suite.Run(t, new(MatchBoolPrefixTestSuite))
}
