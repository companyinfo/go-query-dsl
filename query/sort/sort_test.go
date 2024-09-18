package sort

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SortTestSuite struct {
	suite.Suite
	sort  Sort
	param *Param
}

func (s *SortTestSuite) SetupTest() {
	s.sort = Sort{"field": &Param{}}
	s.param = &Param{}
}

func (s *SortTestSuite) TestNewSort_Success() {
	require := s.Require()

	require.Equal(s.sort, New("field", &Param{}))
}

func (s *SortTestSuite) TestNewParam_Success() {
	require := s.Require()

	require.Equal(s.param, NewParam())
}

func (s *SortTestSuite) TestSetOrder_Success() {
	require := s.Require()
	s.param = &Param{Order: DescOrder}

	require.Equal(s.param, NewParam().SetOrder(DescOrder))
}

func (s *SortTestSuite) TestSetMod_Success() {
	require := s.Require()
	s.param = &Param{Mode: AvgMode}

	require.Equal(s.param, NewParam().SetMode(AvgMode))
}

func (s *SortTestSuite) TestSetNumericType_Success() {
	require := s.Require()
	s.param = &Param{NumericType: DoubleNumericType}

	require.Equal(s.param, NewParam().SetNumericType(DoubleNumericType))
}

func (s *SortTestSuite) TestSetNested_Success() {
	require := s.Require()
	filter := map[string]map[string]string{"term": {"name": "tom"}}
	s.param = &Param{Nested: &nested{
		Path:   "parent",
		Filter: filter,
	}}

	require.Equal(s.param, NewParam().SetNested("parent", filter))
}

func TestSort(t *testing.T) {
	suite.Run(t, new(SortTestSuite))
}
