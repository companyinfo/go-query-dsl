package fuzzy

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FuzzyTestSuite struct {
	suite.Suite
	fuzzy *Fuzzy
	param *Param
}

func (f *FuzzyTestSuite) SetupTest() {
	f.fuzzy = &Fuzzy{Fuzzy: map[string]*Param{"field": {}}}
	f.param = &Param{Value: "value"}
}

func (f *FuzzyTestSuite) TestNewFuzzy_Success() {
	require := f.Require()

	require.Equal(f.fuzzy, New("field", &Param{}))
}

func (f *FuzzyTestSuite) TestNewParam_Success() {
	require := f.Require()

	require.Equal(f.param, NewParam("value"))
}

func (f *FuzzyTestSuite) TestSetBoost_Success() {
	require := f.Require()
	f.param = &Param{Value: "value", Boost: 1.0}

	require.Equal(f.param, NewParam("value").SetBoost(1.0))
}

func (f *FuzzyTestSuite) TestSetFuzziness_Success() {
	require := f.Require()
	f.param = &Param{Value: "value", Fuzziness: "AUTO"}

	require.Equal(f.param, NewParam("value").SetFuzziness("AUTO"))
}

func (f *FuzzyTestSuite) TestSetPrefixLength_Success() {
	require := f.Require()
	f.param = &Param{Value: "value", PrefixLength: 0}

	require.Equal(f.param, NewParam("value").SetPrefixLength(0))
}

func (f *FuzzyTestSuite) TestSetMaxExpansions_Success() {
	require := f.Require()
	f.param = &Param{Value: "value", MaxExpansions: 50}

	require.Equal(f.param, NewParam("value").SetMaxExpansions(50))
}

func (f *FuzzyTestSuite) TestSetTranspositions_Success() {
	require := f.Require()
	f.param = &Param{Value: "value", Transpositions: true}

	require.Equal(f.param, NewParam("value").SetTranspositions(true))
}

func (f *FuzzyTestSuite) TestSetRewrite_Success() {
	require := f.Require()
	f.param = &Param{Value: "value", Rewrite: "constant_score"}

	require.Equal(f.param, NewParam("value").SetRewrite("constant_score"))
}

func TestFuzzy(t *testing.T) {
	suite.Run(t, new(FuzzyTestSuite))
}
