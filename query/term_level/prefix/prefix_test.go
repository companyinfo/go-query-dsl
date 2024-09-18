package prefix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PrefixTestSuite struct {
	suite.Suite
	prefix *Prefix
	param  *Param
}

func (p *PrefixTestSuite) SetupTest() {
	p.prefix = &Prefix{Prefix: map[string]*Param{"field": {}}}
	p.param = &Param{Value: "value"}
}

func (p *PrefixTestSuite) TestNewPrefix_Success() {
	require := p.Require()

	require.Equal(p.prefix, New("field", &Param{}))
}

func (p *PrefixTestSuite) TestNewParam_Success() {
	require := p.Require()

	require.Equal(p.param, NewParam("value"))
}

func (p *PrefixTestSuite) TestSetBoost_Success() {
	require := p.Require()
	p.param = &Param{Value: "value", Boost: 1.0}

	require.Equal(p.param, NewParam("value").SetBoost(1.0))
}

func (p *PrefixTestSuite) TestSetCaseInsensitive_Success() {
	require := p.Require()
	p.param = &Param{Value: "value", CaseInsensitive: true}

	require.Equal(p.param, NewParam("value").SetCaseInsensitive(true))
}

func (p *PrefixTestSuite) TestSetRewrite_Success() {
	require := p.Require()
	p.param = &Param{Value: "value", Rewrite: "constant_score"}

	require.Equal(p.param, NewParam("value").SetRewrite("constant_score"))
}

func TestPrefix(t *testing.T) {
	suite.Run(t, new(PrefixTestSuite))
}
