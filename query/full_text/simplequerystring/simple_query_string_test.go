package simplequerystring

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/companyinfo/go-query-dsl/query/full_text/match"
)

type SimpleQueryStringTestSuite struct {
	suite.Suite
	simpleQueryString *SimpleQueryString
	param             *Param
}

func (m *SimpleQueryStringTestSuite) SetupTest() {
	m.simpleQueryString = &SimpleQueryString{SimpleQueryString: &Param{}}
	m.param = &Param{Query: "query"}
}

func (m *SimpleQueryStringTestSuite) TestNewSimpleQueryString_Success() {
	require := m.Require()

	require.Equal(m.simpleQueryString, New(&Param{}))
}

func (m *SimpleQueryStringTestSuite) TestNewParam_Success() {
	require := m.Require()

	require.Equal(m.param, NewParam("query"))
}

func (m *SimpleQueryStringTestSuite) TestSetAnalyzeWildcard_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", AnalyzeWildcard: true}

	require.Equal(m.param, NewParam("query").SetAnalyzeWildcard(true))
}

func (m *SimpleQueryStringTestSuite) TestSetAnalyzer_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Analyzer: "default_field"}

	require.Equal(m.param, NewParam("query").SetAnalyzer("default_field"))
}

func (m *SimpleQueryStringTestSuite) TestSetAutoGenerateSynonymsPhraseQuery_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", AutoGenerateSynonymsPhraseQuery: true}

	require.Equal(m.param, NewParam("query").SetAutoGenerateSynonymsPhraseQuery(true))
}

func (m *SimpleQueryStringTestSuite) TestSetDefaultOperator_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", DefaultOperator: match.OrOperator}

	require.Equal(m.param, NewParam("query").SetDefaultOperator(match.OrOperator))
}

func (m *SimpleQueryStringTestSuite) TestSetFields_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Fields: []string{"field_1", "field_2"}}

	require.Equal(m.param, NewParam("query").SetFields([]string{"field_1", "field_2"}))
}

func (m *SimpleQueryStringTestSuite) TestSetFlags_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Flags: "ALL"}

	require.Equal(m.param, NewParam("query").SetFlags("ALL"))
}

func (m *SimpleQueryStringTestSuite) TestSetFuzzyMaxExpansions_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", FuzzyMaxExpansions: 50}

	require.Equal(m.param, NewParam("query").SetFuzzyMaxExpansions(50))
}

func (m *SimpleQueryStringTestSuite) TestSetFuzzyTranspositions_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", FuzzyTranspositions: true}

	require.Equal(m.param, NewParam("query").SetFuzzyTranspositions(true))
}

func (m *SimpleQueryStringTestSuite) TestSetFuzzyPrefixLength_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", FuzzyPrefixLength: 0}

	require.Equal(m.param, NewParam("query").SetFuzzyPrefixLength(0))
}

func (m *SimpleQueryStringTestSuite) TestSetLenient_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Lenient: true}

	require.Equal(m.param, NewParam("query").SetLenient(true))
}

func (m *SimpleQueryStringTestSuite) TestSetMinimumShouldMatch_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", MinimumShouldMatch: 2}

	require.Equal(m.param, NewParam("query").SetMinimumShouldMatch(2))
}

func (m *SimpleQueryStringTestSuite) TestSetQuoteFieldSuffix_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", QuoteFieldSuffix: ".exact"}

	require.Equal(m.param, NewParam("query").SetQuoteFieldSuffix(".exact"))
}

func TestSimpleQueryString(t *testing.T) {
	suite.Run(t, new(SimpleQueryStringTestSuite))
}
