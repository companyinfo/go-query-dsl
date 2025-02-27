package ids

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IDsTestSuite struct {
	suite.Suite
	ids   *IDs
	param *Param
}

func (i *IDsTestSuite) SetupSuite() {
	i.param = &Param{Values: []int{1}}
	i.ids = &IDs{IDs: i.param}
}

func (i *IDsTestSuite) TestNewIDs_Success() {
	require := i.Require()

	require.Equal(i.ids, New(i.param))
}

func (i *IDsTestSuite) TestNewParam_Success() {
	require := i.Require()

	require.Equal(i.param, NewParam([]int{1}))
}

func (i *IDsTestSuite) TestSetBoost_Success() {
	require := i.Require()
	i.param = &Param{Values: []int{1}, Boost: 1.0}

	require.Equal(i.param, NewParam([]int{1}).SetBoost(1.0))
}

func TestIDs(t *testing.T) {
	suite.Run(t, new(IDsTestSuite))
}
