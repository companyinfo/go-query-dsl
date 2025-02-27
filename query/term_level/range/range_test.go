package ranging

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RangeTestSuite struct {
	suite.Suite
	ranges *Range
	param  *Param
}

func (r *RangeTestSuite) SetupTest() {
	r.ranges = &Range{Range: map[string]*Param{"date": {}}}
	r.param = &Param{}
}

func (r *RangeTestSuite) TestNewRange_Success() {
	require := r.Require()

	require.Equal(r.ranges, New("date", &Param{}))
}

func (r *RangeTestSuite) TestNewParam_Success() {
	require := r.Require()

	require.Equal(r.param, NewParam())
}

func (r *RangeTestSuite) TestSetGT_Success() {
	require := r.Require()
	r.param = &Param{GT: "2022-04-17T06:00:00"}

	require.Equal(r.param, NewParam().SetGT("2022-04-17T06:00:00"))
}

func (r *RangeTestSuite) TestSetGTE_Success() {
	require := r.Require()
	r.param = &Param{GTE: "2022-04-17T06:00:00"}

	require.Equal(r.param, NewParam().SetGTE("2022-04-17T06:00:00"))
}

func (r *RangeTestSuite) TestSetLT_Success() {
	require := r.Require()
	r.param = &Param{LT: "2022-04-17T06:00:00"}

	require.Equal(r.param, NewParam().SetLT("2022-04-17T06:00:00"))
}

func (r *RangeTestSuite) TestSetLTE_Success() {
	require := r.Require()
	r.param = &Param{LTE: "2022-04-17T06:00:00"}

	require.Equal(r.param, NewParam().SetLTE("2022-04-17T06:00:00"))
}

func (r *RangeTestSuite) TestSetFormat_Success() {
	require := r.Require()
	r.param = &Param{Format: "epoch_millis"}

	require.Equal(r.param, NewParam().SetFormat("epoch_millis"))
}

func (r *RangeTestSuite) TestSetRelation_Success() {
	require := r.Require()
	r.param = &Param{Relation: IntersectsRelationType}

	require.Equal(r.param, NewParam().SetRelation(IntersectsRelationType))
}

func (r *RangeTestSuite) TestSetBoost_Success() {
	require := r.Require()
	r.param = &Param{Boost: 1.0}

	require.Equal(r.param, NewParam().SetBoost(1.0))
}

func (r *RangeTestSuite) TestSetTimezone_Success() {
	require := r.Require()
	r.param = &Param{Timezone: "-04:00"}

	require.Equal(r.param, NewParam().SetTimezone("-04:00"))
}

func TestRange(t *testing.T) {
	suite.Run(t, new(RangeTestSuite))
}
