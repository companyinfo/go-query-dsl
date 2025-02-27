package termsset

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TermsSetTestSuite struct {
	suite.Suite
	termsSet *TermsSet
	param    *Param
}

func (t *TermsSetTestSuite) SetupTest() {
	t.param = &Param{Terms: []string{"c++", "java"}}
	t.termsSet = &TermsSet{TermsSet: map[string]*Param{"programming_languages": t.param}}
}

func (t *TermsSetTestSuite) TestNewTermsSet_Success() {
	require := t.Require()

	require.Equal(t.termsSet, New("programming_languages", &Param{Terms: []string{"c++", "java"}}))
}

func (t *TermsSetTestSuite) TestNewParam_Success() {
	require := t.Require()

	require.Equal(t.param, NewParam([]string{"c++", "java"}))
}

func (t *TermsSetTestSuite) TestSetBoost_Success() {
	require := t.Require()
	t.param = &Param{Terms: []string{"c++", "java"}, Boost: 1.0}

	require.Equal(t.param, NewParam([]string{"c++", "java"}).SetBoost(1.0))
}

func (t *TermsSetTestSuite) TestSetMinimumShouldMatchField_Success() {
	require := t.Require()
	t.param = &Param{Terms: []string{"c++", "java"}, MinimumShouldMatchField: "required_matches"}

	require.Equal(t.param, NewParam([]string{"c++", "java"}).SetMinimumShouldMatchField("required_matches"))
}

func (t *TermsSetTestSuite) TestSetMinimumShouldMatchScript_Success() {
	require := t.Require()
	t.param = &Param{Terms: []string{"c++", "java"}, MinimumShouldMatchScript: &source{Source: "required_matches"}}

	require.Equal(t.param, NewParam([]string{"c++", "java"}).SetMinimumShouldMatchScript("required_matches"))
}

func TestTermsSet(t *testing.T) {
	suite.Run(t, new(TermsSetTestSuite))
}
