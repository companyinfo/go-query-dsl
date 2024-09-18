package match

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MatchTestSuite struct {
	suite.Suite
	match *Match
	param *Param
}

func (m *MatchTestSuite) SetupTest() {
	m.match = &Match{Match: map[string]*Param{"field": {}}}
	m.param = &Param{Query: "value"}
}

func (m *MatchTestSuite) TestNewMatch_Success() {
	require := m.Require()

	require.Equal(m.match, New("field", &Param{}))
}

func (m *MatchTestSuite) TestNewParam_Success() {
	require := m.Require()

	require.Equal(m.param, NewParam("value"))
}

func (m *MatchTestSuite) TestSetAnalyzer_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", Analyzer: "standard"}

	require.Equal(m.param, NewParam("value").SetAnalyzer("standard"))
}

func (m *MatchTestSuite) TestSetFuzziness_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", Fuzziness: "AUTO"}

	require.Equal(m.param, NewParam("value").SetFuzziness("AUTO"))
}

func (m *MatchTestSuite) TestSetPrefixLength_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", PrefixLength: 0}

	require.Equal(m.param, NewParam("value").SetPrefixLength(0))
}

func (m *MatchTestSuite) TestSetMaxExpansions_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", MaxExpansions: 50}

	require.Equal(m.param, NewParam("value").SetMaxExpansions(50))
}

func (m *MatchTestSuite) TestSetOperator_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", Operator: OrOperator}

	require.Equal(m.param, NewParam("value").SetOperator(OrOperator))
}

func (m *MatchTestSuite) TestSetMinimumShouldMatch_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", MinimumShouldMatch: 1}

	require.Equal(m.param, NewParam("value").SetMinimumShouldMatch(1))
}

func (m *MatchTestSuite) TestSetFuzzYRewrite_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", FuzzyRewrite: "constant_score"}

	require.Equal(m.param, NewParam("value").SetFuzzyRewrite("constant_score"))
}

func (m *MatchTestSuite) TestSetZeroTermsQuery_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", ZeroTermsQuery: NoneZeroTermType}

	require.Equal(m.param, NewParam("value").SetZeroTermsQuery(NoneZeroTermType))
}

func (m *MatchTestSuite) TestSetBoost_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", Boost: 1}

	require.Equal(m.param, NewParam("value").SetBoost(1))
}

func (m *MatchTestSuite) TestSetAutoGenerateSynonymsPhraseQuery_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", AutoGenerateSynonymsPhraseQuery: true}

	require.Equal(m.param, NewParam("value").SetAutoGenerateSynonymsPhraseQuery(true))
}

func (m *MatchTestSuite) TestSetEnablePositionIncrements_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", EnablePositionIncrements: true}

	require.Equal(m.param, NewParam("value").SetEnablePositionIncrements(true))
}

func (m *MatchTestSuite) TestSetFuzzyTranspositions_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", FuzzyTranspositions: true}

	require.Equal(m.param, NewParam("value").SetFuzzyTranspositions(true))
}

func (m *MatchTestSuite) TestSetLenient_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", Lenient: true}

	require.Equal(m.param, NewParam("value").SetLenient(true))
}

func TestMatch(t *testing.T) {
	suite.Run(t, new(MatchTestSuite))
}
