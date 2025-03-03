package combinedfields

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go.companyinfo.dev/go-query-dsl/query/full_text/match"
)

type CombinedFieldsTestSuite struct {
	suite.Suite
	combinedFields *CombinedFields
	param          *Param
}

func (c *CombinedFieldsTestSuite) SetupTest() {
	c.combinedFields = &CombinedFields{CombinedFields: map[string]*Param{"field": {}}}
	c.param = &Param{Query: "query", Fields: make([]string, 0)}
}

func (c *CombinedFieldsTestSuite) TestNewCombinedFields_Success() {
	require := c.Require()

	require.Equal(c.combinedFields, New("field", &Param{}))
}

func (c *CombinedFieldsTestSuite) TestNewParam_Success() {
	require := c.Require()

	require.Equal(c.param, NewParam("query", []string{}))
}

func (c *CombinedFieldsTestSuite) TestSetAutoGenerateSynonymsPhraseQuery_Success() {
	require := c.Require()
	c.param = &Param{
		Query:                           "value",
		Fields:                          make([]string, 0),
		AutoGenerateSynonymsPhraseQuery: true,
	}

	require.Equal(c.param, NewParam("value", []string{}).SetAutoGenerateSynonymsPhraseQuery(true))
}

func (c *CombinedFieldsTestSuite) TestSetOperator_Success() {
	require := c.Require()
	c.param = &Param{
		Query:    "value",
		Fields:   make([]string, 0),
		Operator: match.OrOperator,
	}

	require.Equal(c.param, NewParam("value", []string{}).SetOperator(match.OrOperator))
}

func (c *CombinedFieldsTestSuite) TestSetMinimumShouldMatch_Success() {
	require := c.Require()
	c.param = &Param{
		Query:              "value",
		Fields:             make([]string, 0),
		MinimumShouldMatch: 10,
	}

	require.Equal(c.param, NewParam("value", []string{}).SetMinimumShouldMatch(10))
}

func (c *CombinedFieldsTestSuite) TestSetZeroTermsQuery_Success() {
	require := c.Require()
	c.param = &Param{
		Query:          "value",
		Fields:         make([]string, 0),
		ZeroTermsQuery: match.NoneZeroTermType,
	}

	require.Equal(c.param, NewParam("value", []string{}).SetZeroTermsQuery(match.NoneZeroTermType))
}

func TestCombinedFields(t *testing.T) {
	suite.Run(t, new(CombinedFieldsTestSuite))
}
