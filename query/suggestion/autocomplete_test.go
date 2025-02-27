// Package suggestion accepts a list of suggestions and builds them into a finite-state transducer (FST),
// an optimized data structure that is essentially a graph.
// This data structure lives in memory and is optimized for fast prefix lookups.
package suggestion

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AutocompleteTestSuite struct {
	suite.Suite
}

func (s *AutocompleteTestSuite) TestNewAutocomplete_Success() {
	require := s.Require()
	autocomplete := &Autocomplete{Autocomplete: &Param{}}

	require.Equal(autocomplete, NewAutocomplete(&Param{}))
}

func (s *AutocompleteTestSuite) TestNewParam_Success() {
	require := s.Require()
	param := &Param{Prefix: "prefix", Completion: &Completion{Field: "field"}}

	require.Equal(param, NewAutocompleteParam("prefix", "field"))
}

func (s *AutocompleteTestSuite) TestSetSize_Success() {
	require := s.Require()
	param := &Param{Prefix: "prefix", Completion: &Completion{Field: "field", Size: 5}}

	require.Equal(param, NewAutocompleteParam("prefix", "field").SetSize(5))
}

func (s *AutocompleteTestSuite) TestSetFuzziness_Success() {
	require := s.Require()
	param := &Param{
		Prefix: "prefix",
		Completion: &Completion{
			Field: "field",
			Fuzzy: &Fuzzy{
				Fuzziness: "fuzzy",
			},
		},
	}

	require.Equal(param, NewAutocompleteParam("prefix", "field").SetFuzziness("fuzzy"))
}

func TestAutocomplete(t *testing.T) {
	suite.Run(t, new(AutocompleteTestSuite))
}
