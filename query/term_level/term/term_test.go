package term

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TermTestSuite struct {
	suite.Suite
	term  *Term
	param *Param
}

func (t *TermTestSuite) SetupTest() {
	t.param = &Param{Value: "kimchy"}
	t.term = &Term{Term: map[string]*Param{"user.id": t.param}}
}

func (t *TermTestSuite) TestNewTerm_Success() {
	require := t.Require()

	require.Equal(t.term, New("user.id", &Param{Value: "kimchy"}))
}

func (t *TermTestSuite) TestNewParam_Success() {
	require := t.Require()

	require.Equal(t.param, NewParam("kimchy"))
}

func (t *TermTestSuite) TestSetBoost_Success() {
	require := t.Require()
	t.param = &Param{Value: "kimchy", Boost: 1.0}

	require.Equal(t.param, NewParam("kimchy").SetBoost(1.0))
}

func (t *TermTestSuite) TestSetCaseInsensitive_Success() {
	require := t.Require()
	t.param = &Param{Value: "kimchy", CaseInsensitive: true}

	require.Equal(t.param, NewParam("kimchy").SetCaseInsensitive(true))
}

func TestTerm(t *testing.T) {
	suite.Run(t, new(TermTestSuite))
}
