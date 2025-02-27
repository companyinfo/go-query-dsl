package terms

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TermsTestSuite struct {
	suite.Suite
	termsAny    *Terms[any]
	termsLookup *Terms[*Lookup]
	lookup      *Lookup
}

func (t *TermsTestSuite) SetupTest() {
	t.termsAny = &Terms[any]{
		Terms: map[string]any{
			"user.id": []string{"12", "14"},
		},
	}
	t.lookup = &Lookup{
		Index: "classes",
		ID:    "102",
		Path:  "enrolled_students.id_list",
	}
	t.termsLookup = &Terms[*Lookup]{
		Terms: map[string]*Lookup{
			"student_id": t.lookup,
		},
	}
}

func (t *TermsTestSuite) TestNewTerms_Success() {
	require := t.Require()
	require.Equal(t.termsAny, New[any]("user.id", []string{"12", "14"}))
}

func (t *TermsTestSuite) TestNewLookup_Success() {
	require := t.Require()

	require.Equal(t.lookup, NewLookup("classes", "102", "enrolled_students.id_list"))
}

func (t *TermsTestSuite) TestSetBoost_Success() {
	require := t.Require()
	t.termsAny = &Terms[any]{
		Terms: map[string]any{
			"user.id": []string{"12", "14"},
			"boost":   1.0,
		},
	}

	require.Equal(t.termsAny, New[any]("user.id", []string{"12", "14"}).SetBoost(1.0))
}

func (t *TermsTestSuite) TestSetLookupBoost_Success() {
	require := t.Require()
	t.lookup.Boost = 1.0

	lookup := NewLookup("classes", "102", "enrolled_students.id_list").SetBoost(1.0)
	require.Equal(t.termsLookup, New[*Lookup]("student_id", lookup))
}

func TestTerms(t *testing.T) {
	suite.Run(t, new(TermsTestSuite))
}
