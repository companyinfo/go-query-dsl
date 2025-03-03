package multimatch

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/full_text/match"
)

type MultiMatchTestSuite struct {
	suite.Suite
	multiMatch *MultiMatch
	param      *Param
}

func (m *MultiMatchTestSuite) SetupTest() {
	m.multiMatch = &MultiMatch{MultiMatch: &Param{}}
	m.param = &Param{Query: "query"}
}

func (m *MultiMatchTestSuite) TestNewMultiMatch_Success() {
	require := m.Require()

	require.Equal(m.multiMatch, New(&Param{}))
}

func (m *MultiMatchTestSuite) TestNewParam_Success() {
	require := m.Require()

	require.Equal(m.param, NewParam("query"))
}

func (m *MultiMatchTestSuite) TestSetFields_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Fields: []string{"field_1", "field_2"}}

	require.Equal(m.param, NewParam("query").SetFields([]string{"field_1", "field_2"}))
}

func (m *MultiMatchTestSuite) TestSetType_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Type: BestFields}

	require.Equal(m.param, NewParam("query").SetType(BestFields))
}

func (m *MultiMatchTestSuite) TestSetTieBreaker_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", TieBreaker: 0.3}

	require.Equal(m.param, NewParam("query").SetTieBreaker(0.3))
}

func (m *MultiMatchTestSuite) TestSetOperator_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Operator: match.OrOperator}

	require.Equal(m.param, NewParam("query").SetOperator(match.OrOperator))
}

func (m *MultiMatchTestSuite) TestSetAnalyzer_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Analyzer: "standard"}

	require.Equal(m.param, NewParam("query").SetAnalyzer("standard"))
}

func (m *MultiMatchTestSuite) TestSetSlop_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Slop: 2}

	require.Equal(m.param, NewParam("query").SetSlop(2))
}

func (m *MultiMatchTestSuite) TestSetFuzziness_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Fuzziness: "AUTO"}

	require.Equal(m.param, NewParam("query").SetFuzziness("AUTO"))
}

func (m *MultiMatchTestSuite) TestSetPrefixLength_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", PrefixLength: 0}

	require.Equal(m.param, NewParam("query").SetPrefixLength(0))
}

func (m *MultiMatchTestSuite) TestSetBoost_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Boost: 1}

	require.Equal(m.param, NewParam("query").SetBoost(1))
}

func (m *MultiMatchTestSuite) TestSetMinimumShouldMatch_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", MinimumShouldMatch: 50}

	require.Equal(m.param, NewParam("query").SetMinimumShouldMatch(50))
}

func (m *MultiMatchTestSuite) TestSetMaxExpansions_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", MaxExpansions: 50}

	require.Equal(m.param, NewParam("query").SetMaxExpansions(50))
}

func (m *MultiMatchTestSuite) TestSetFuzzYRewrite_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", FuzzyRewrite: "constant_score"}

	require.Equal(m.param, NewParam("query").SetFuzzYRewrite("constant_score"))
}

func (m *MultiMatchTestSuite) TestSetZeroTermsQuery_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", ZeroTermsQuery: "none"}

	require.Equal(m.param, NewParam("query").SetZeroTermsQuery("none"))
}

func (m *MultiMatchTestSuite) TestSetAutoGenerateSynonymsPhraseQuery_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", AutoGenerateSynonymsPhraseQuery: true}

	require.Equal(m.param, NewParam("query").SetAutoGenerateSynonymsPhraseQuery(true))
}

func (m *MultiMatchTestSuite) TestSetFuzzyTranspositions_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", FuzzyTranspositions: true}

	require.Equal(m.param, NewParam("query").SetFuzzyTranspositions(true))
}

func (m *MultiMatchTestSuite) TestSetLenient_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Lenient: true}

	require.Equal(m.param, NewParam("query").SetLenient(true))
}

func TestMultiMatch(t *testing.T) {
	suite.Run(t, new(MultiMatchTestSuite))
}
