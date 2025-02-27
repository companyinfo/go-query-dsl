package regex

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RegexTestSuite struct {
	suite.Suite
	regex *Regex
	param *Param
}

func (r *RegexTestSuite) SetupTest() {
	r.regex = &Regex{Regex: map[string]*Param{"field": {}}}
	r.param = &Param{Value: "value"}
}

func (r *RegexTestSuite) TestNewRegex_Success() {
	require := r.Require()

	require.Equal(r.regex, New("field", &Param{}))
}

func (r *RegexTestSuite) TestNewParam_Success() {
	require := r.Require()

	require.Equal(r.param, NewParam("value"))
}

func (r *RegexTestSuite) TestSetBoost_Success() {
	require := r.Require()
	r.param = &Param{Value: "value", Boost: 1.0}

	require.Equal(r.param, NewParam("value").SetBoost(1.0))
}

func (r *RegexTestSuite) TestSetCaseInsensitive_Success() {
	require := r.Require()
	r.param = &Param{Value: "value", CaseInsensitive: true}

	require.Equal(r.param, NewParam("value").SetCaseInsensitive(true))
}

func (r *RegexTestSuite) TestSetFlags_Success() {
	require := r.Require()
	r.param = &Param{Value: "value", Flags: "ALL"}

	require.Equal(r.param, NewParam("value").SetFlags("ALL"))
}

func (r *RegexTestSuite) TestSetMaxDeterminizedStates_Success() {
	require := r.Require()
	r.param = &Param{Value: "value", MaxDeterminizedStates: 10000}

	require.Equal(r.param, NewParam("value").SetMaxDeterminizedStates(10000))
}

func (r *RegexTestSuite) TestSetRewrite_Success() {
	require := r.Require()
	r.param = &Param{Value: "value", Rewrite: "constant_score"}

	require.Equal(r.param, NewParam("value").SetRewrite("constant_score"))
}

func TestRegex(t *testing.T) {
	suite.Run(t, new(RegexTestSuite))
}
