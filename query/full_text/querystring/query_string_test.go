package querystring

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/full_text/match"
)

type QueryStringTestSuite struct {
	suite.Suite
	queryString *QueryString
	param       *Param
}

func (m *QueryStringTestSuite) SetupTest() {
	m.queryString = &QueryString{QueryString: &Param{}}
	m.param = &Param{Query: "query"}
}

func (m *QueryStringTestSuite) TestNewQueryString_Success() {
	require := m.Require()

	require.Equal(m.queryString, New(&Param{}))
}

func (m *QueryStringTestSuite) TestNewParam_Success() {
	require := m.Require()

	require.Equal(m.param, NewParam("query"))
}

func (m *QueryStringTestSuite) TestSetAllowLeadingWildcard_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", AllowLeadingWildcard: true}

	require.Equal(m.param, NewParam("query").SetAllowLeadingWildcard(true))
}

func (m *QueryStringTestSuite) TestSetAnalyzeWildcard_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", AnalyzeWildcard: true}

	require.Equal(m.param, NewParam("query").SetAnalyzeWildcard(true))
}

func (m *QueryStringTestSuite) TestSetAnalyzer_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Analyzer: "default_field"}

	require.Equal(m.param, NewParam("query").SetAnalyzer("default_field"))
}

func (m *QueryStringTestSuite) TestSetAutoGenerateSynonymsPhraseQuery_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", AutoGenerateSynonymsPhraseQuery: true}

	require.Equal(m.param, NewParam("query").SetAutoGenerateSynonymsPhraseQuery(true))
}

func (m *QueryStringTestSuite) TestSetBoost_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Boost: 1}

	require.Equal(m.param, NewParam("query").SetBoost(1))
}

func (m *QueryStringTestSuite) TestSetDefaultFields_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", DefaultFields: "name"}

	require.Equal(m.param, NewParam("query").SetDefaultFields("name"))
}

func (m *QueryStringTestSuite) TestSetDefaultOperator_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", DefaultOperator: match.OrOperator}

	require.Equal(m.param, NewParam("query").SetDefaultOperator(match.OrOperator))
}

func (m *QueryStringTestSuite) TestSetEnablePositionIncrements_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", EnablePositionIncrements: true}

	require.Equal(m.param, NewParam("query").SetEnablePositionIncrements(true))
}

func (m *QueryStringTestSuite) TestSetFields_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Fields: []string{"field_1", "field_2"}}

	require.Equal(m.param, NewParam("query").SetFields([]string{"field_1", "field_2"}))
}

func (m *QueryStringTestSuite) TestSetFuzziness_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Fuzziness: "AUTO"}

	require.Equal(m.param, NewParam("query").SetFuzziness("AUTO"))
}

func (m *QueryStringTestSuite) TestSetFuzzyMaxExpansions_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", FuzzyMaxExpansions: 50}

	require.Equal(m.param, NewParam("query").SetFuzzyMaxExpansions(50))
}

func (m *QueryStringTestSuite) TestSetFuzzyTranspositions_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", FuzzyTranspositions: true}

	require.Equal(m.param, NewParam("query").SetFuzzyTranspositions(true))
}

func (m *QueryStringTestSuite) TestSetLenient_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Lenient: true}

	require.Equal(m.param, NewParam("query").SetLenient(true))
}

func (m *QueryStringTestSuite) TestSetMaxDeterminizedStates_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", MaxDeterminizedStates: 10000}

	require.Equal(m.param, NewParam("query").SetMaxDeterminizedStates(10000))
}

func (m *QueryStringTestSuite) TestSetMinimumShouldMatch_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", MinimumShouldMatch: 2}

	require.Equal(m.param, NewParam("query").SetMinimumShouldMatch(2))
}

func (m *QueryStringTestSuite) TestSetPhraseSlop_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", PhraseSlop: 2}

	require.Equal(m.param, NewParam("query").SetPhraseSlop(2))
}

func (m *QueryStringTestSuite) TestSetQuoteAnalyzer_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", QuoteAnalyzer: "search_quote_analyzer"}

	require.Equal(m.param, NewParam("query").SetQuoteAnalyzer("search_quote_analyzer"))
}

func (m *QueryStringTestSuite) TestSetQuoteFieldSuffix_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", QuoteFieldSuffix: ".exact"}

	require.Equal(m.param, NewParam("query").SetQuoteFieldSuffix(".exact"))
}

func (m *QueryStringTestSuite) TestSetRewrite_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Rewrite: "constant_score"}

	require.Equal(m.param, NewParam("query").SetRewrite("constant_score"))
}

func (m *QueryStringTestSuite) TestSetTimeZone_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", TimeZone: "-08:00"}

	require.Equal(m.param, NewParam("query").SetTimeZone("-08:00"))
}

func TestQueryString(t *testing.T) {
	suite.Run(t, new(QueryStringTestSuite))
}
