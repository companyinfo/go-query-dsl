package exists

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExistsTestSuite struct {
	suite.Suite
	exists *Exists
	param  *Param
}

func (e *ExistsTestSuite) SetupSuite() {
	e.exists = &Exists{Exists: &Param{Field: "user.id"}}
	e.param = &Param{Field: "user.id"}
}

func (e *ExistsTestSuite) TestNewExists_Success() {
	require := e.Require()

	require.Equal(e.exists, New(e.param))
}

func (e *ExistsTestSuite) TestNewParam_Success() {
	require := e.Require()

	require.Equal(e.param, NewParam("user.id"))
}

func (e *ExistsTestSuite) TestSetBoost_Success() {
	require := e.Require()
	e.param = &Param{Field: "user.id", Boost: 1.0}

	require.Equal(e.param, NewParam("user.id").SetBoost(1.0))
}

func TestExists(t *testing.T) {
	suite.Run(t, new(ExistsTestSuite))
}
