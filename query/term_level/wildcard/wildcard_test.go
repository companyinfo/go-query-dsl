package wildcard

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WildcardTestSuite struct {
	suite.Suite
	wildcard *Wildcard
	param    *Param
}

func (r *WildcardTestSuite) SetupTest() {
	r.param = &Param{Value: "ki*y"}
	r.wildcard = &Wildcard{Wildcard: map[string]*Param{"user.id": r.param}}
}

func (r *WildcardTestSuite) TestNewWildcard_Success() {
	require := r.Require()

	require.Equal(r.wildcard, New("user.id", r.param))
}

func (r *WildcardTestSuite) TestNewParam_Success() {
	require := r.Require()

	require.Equal(r.param, NewParam("ki*y"))
}

func (r *WildcardTestSuite) TestSetBoost_Success() {
	require := r.Require()
	r.param = &Param{Value: "ki*y", Boost: 1.0}

	require.Equal(r.param, NewParam("ki*y").SetBoost(1.0))
}

func (r *WildcardTestSuite) TestSetCaseInsensitive_Success() {
	require := r.Require()
	r.param = &Param{Value: "ki*y", CaseInsensitive: true}

	require.Equal(r.param, NewParam("ki*y").SetCaseInsensitive(true))
}

func (r *WildcardTestSuite) TestSetRewrite_Success() {
	require := r.Require()
	r.param = &Param{Value: "ki*y", Rewrite: "constant_score"}

	require.Equal(r.param, NewParam("ki*y").SetRewrite("constant_score"))
}

func TestWildcard(t *testing.T) {
	suite.Run(t, new(WildcardTestSuite))
}
