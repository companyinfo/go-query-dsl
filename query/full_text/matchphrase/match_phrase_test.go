package matchphrase

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MatchPhraseTestSuite struct {
	suite.Suite
	matchPhrase *MatchPhrase
	param       *Param
}

func (m *MatchPhraseTestSuite) SetupTest() {
	m.matchPhrase = &MatchPhrase{MatchPhrase: map[string]*Param{"field": {}}}
	m.param = &Param{Query: "query"}
}

func (m *MatchPhraseTestSuite) TestNewMatchPhrase_Success() {
	require := m.Require()

	require.Equal(m.matchPhrase, New("field", &Param{}))
}

func (m *MatchPhraseTestSuite) TestNewParam_Success() {
	require := m.Require()

	require.Equal(m.param, NewParam("query"))
}

func (m *MatchPhraseTestSuite) TestSetAnalyzer_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Analyzer: "standard"}

	require.Equal(m.param, NewParam("query").SetAnalyzer("standard"))
}

func (m *MatchPhraseTestSuite) TestSetZeroTermsQuery_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", ZeroTermsQuery: "none"}

	require.Equal(m.param, NewParam("query").SetZeroTermsQuery("none"))
}

func (m *MatchPhraseTestSuite) TestSetSlop_Success() {
	require := m.Require()
	m.param = &Param{Query: "query", Slop: 0}

	require.Equal(m.param, NewParam("query").SetSlop(0))
}

func TestMatchPhrase(t *testing.T) {
	suite.Run(t, new(MatchPhraseTestSuite))
}
