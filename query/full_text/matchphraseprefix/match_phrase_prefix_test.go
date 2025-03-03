package matchphraseprefix

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/full_text/match"
)

type MatchPhrasePrefixTestSuite struct {
	suite.Suite
	matchPhrasePrefix *MatchPhrasePrefix
	param             *Param
}

func (m *MatchPhrasePrefixTestSuite) SetupTest() {
	m.matchPhrasePrefix = &MatchPhrasePrefix{MatchPhrasePrefix: map[string]*Param{"field": {}}}
	m.param = &Param{Query: "query"}
}

func (m *MatchPhrasePrefixTestSuite) TestNewMatchPhrasePrefix_Success() {
	require := m.Require()

	require.Equal(m.matchPhrasePrefix, New("field", &Param{}))
}

func (m *MatchPhrasePrefixTestSuite) TestNewParam_Success() {
	require := m.Require()

	require.Equal(m.param, NewParam("query"))
}

func (m *MatchPhrasePrefixTestSuite) TestSetAnalyzer_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Analyzer: "standard"}

	require.Equal(m.param, NewParam("query").SetAnalyzer("standard"))
}

func (m *MatchPhrasePrefixTestSuite) TestSetMaxExpansions_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", MaxExpansions: 50}

	require.Equal(m.param, NewParam("query").SetMaxExpansions(50))
}

func (m *MatchPhrasePrefixTestSuite) TestSetSlop_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Slop: 0}

	require.Equal(m.param, NewParam("query").SetSlop(0))
}

func (m *MatchPhrasePrefixTestSuite) TestSetZeroTermsQuery_Success() {
	require := m.Require()
	m.param = &Param{Query: "value", ZeroTermsQuery: match.NoneZeroTermType}

	require.Equal(m.param, NewParam("value").SetZeroTermsQuery(match.NoneZeroTermType))
}

func TestMatchPhrasePrefix(t *testing.T) {
	suite.Run(t, new(MatchPhrasePrefixTestSuite))
}
