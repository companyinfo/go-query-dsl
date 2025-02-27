package intervals

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IntervalsTestSuite struct {
	suite.Suite
}

func (m *IntervalsTestSuite) TestNewIntervals_Success() {
	require := m.Require()
	intervals := &Intervals{Intervals: map[string]*Param{"field": {}}}

	require.Equal(intervals, New("field", &Param{}))
}

func (m *IntervalsTestSuite) TestNewParam_Success() {
	require := m.Require()
	param := &Param{}

	require.Equal(param, NewParam())
}

func (m *IntervalsTestSuite) TestSetMatch_Success() {
	require := m.Require()
	param := &Param{Match: &Match{Query: "value"}}

	require.Equal(param, NewParam().SetMatch(NewParamMatch("value")))
}

func (m *IntervalsTestSuite) TestSetPrefix_Success() {
	require := m.Require()
	param := &Param{Prefix: &Prefix{Prefix: "pre"}}

	require.Equal(param, NewParam().SetPrefix(NewParamPrefix("pre")))
}

func (m *IntervalsTestSuite) TestSetWildcard_Success() {
	require := m.Require()
	param := &Param{Wildcard: &Wildcard{Pattern: "pre*"}}

	require.Equal(param, NewParam().SetWildcard(NewParamWildcard("pre*")))
}

func (m *IntervalsTestSuite) TestSetFuzzy_Success() {
	require := m.Require()
	param := &Param{Fuzzy: &Fuzzy{Term: "value"}}

	require.Equal(param, NewParam().SetFuzzy(NewParamFuzzy("value")))
}

func (m *IntervalsTestSuite) TestSetAllOf_Success() {
	require := m.Require()
	param := &Param{AllOf: &AllOf{Intervals: []*Param{}}}

	require.Equal(param, NewParam().SetAllOf(NewParamAllOf([]*Param{})))
}

func (m *IntervalsTestSuite) TestSetAnyOf_Success() {
	require := m.Require()
	param := &Param{AnyOf: &AnyOf{Intervals: []*Param{}}}

	require.Equal(param, NewParam().SetAnyOf(NewParamAnyOf([]*Param{})))
}

func (m *IntervalsTestSuite) TestMatchSetAnalyzer_Success() {
	require := m.Require()
	param := &Param{
		Match: &Match{
			Query:    "value",
			Analyzer: "standard",
		},
	}

	require.Equal(param, NewParam().SetMatch(NewParamMatch("value").SetAnalyzer("standard")))
}

func (m *IntervalsTestSuite) TestMatchSetFilter_Success() {
	require := m.Require()
	param := &Param{
		Match: &Match{
			Query:  "value",
			Filter: &Filter{},
		},
	}

	require.Equal(param, NewParam().SetMatch(NewParamMatch("value").SetFilter(NewFilter())))
}

func (m *IntervalsTestSuite) TestMatchSetMaxGaps_Success() {
	require := m.Require()
	param := &Param{
		Match: &Match{
			Query:   "value",
			MaxGaps: -1,
		},
	}

	require.Equal(param, NewParam().SetMatch(NewParamMatch("value").SetMaxGaps(-1)))
}

func (m *IntervalsTestSuite) TestMatchSetOrdered_Success() {
	require := m.Require()
	param := &Param{
		Match: &Match{
			Query:   "value",
			Ordered: false,
		},
	}

	require.Equal(param, NewParam().SetMatch(NewParamMatch("value").SetOrdered(false)))
}

func (m *IntervalsTestSuite) TestMatchSetUseField_Success() {
	require := m.Require()
	param := &Param{
		Match: &Match{
			Query:    "value",
			UseField: "field",
		},
	}

	require.Equal(param, NewParam().SetMatch(NewParamMatch("value").SetUseField("field")))
}

func (m *IntervalsTestSuite) TestFilterSetAfter_Success() {
	require := m.Require()
	filter := &Filter{
		After: &Param{Match: &Match{Query: "value"}},
	}

	require.Equal(filter, NewFilter().SetAfter(NewParam().SetMatch(NewParamMatch("value"))))
}

func (m *IntervalsTestSuite) TestFilterSetBefore_Success() {
	require := m.Require()
	filter := &Filter{
		Before: &Param{Match: &Match{Query: "value"}},
	}

	require.Equal(filter, NewFilter().SetBefore(NewParam().SetMatch(NewParamMatch("value"))))
}

func (m *IntervalsTestSuite) TestFilterSetContainedBy_Success() {
	require := m.Require()
	filter := &Filter{
		ContainedBy: &Param{Match: &Match{Query: "value"}},
	}

	require.Equal(filter, NewFilter().SetContainedBy(NewParam().SetMatch(NewParamMatch("value"))))
}

func (m *IntervalsTestSuite) TestFilterSetContaining_Success() {
	require := m.Require()
	filter := &Filter{
		Containing: &Param{Match: &Match{Query: "value"}},
	}

	require.Equal(filter, NewFilter().SetContaining(NewParam().SetMatch(NewParamMatch("value"))))
}

func (m *IntervalsTestSuite) TestFilterSetNotContainedBy_Success() {
	require := m.Require()
	filter := &Filter{
		NotContainedBy: &Param{Match: &Match{Query: "value"}},
	}

	require.Equal(filter, NewFilter().SetNotContainedBy(NewParam().SetMatch(NewParamMatch("value"))))
}

func (m *IntervalsTestSuite) TestFilterSetNotContaining_Success() {
	require := m.Require()
	filter := &Filter{
		NotContaining: &Param{Match: &Match{Query: "value"}},
	}

	require.Equal(filter, NewFilter().SetNotContaining(NewParam().SetMatch(NewParamMatch("value"))))
}

func (m *IntervalsTestSuite) TestFilterSetNotOverlapping_Success() {
	require := m.Require()
	filter := &Filter{
		NotOverlapping: &Param{Match: &Match{Query: "value"}},
	}

	require.Equal(filter, NewFilter().SetNotOverlapping(NewParam().SetMatch(NewParamMatch("value"))))
}

func (m *IntervalsTestSuite) TestFilterSetOverlapping_Success() {
	require := m.Require()
	filter := &Filter{
		Overlapping: &Param{Match: &Match{Query: "value"}},
	}

	require.Equal(filter, NewFilter().SetOverlapping(NewParam().SetMatch(NewParamMatch("value"))))
}

func (m *IntervalsTestSuite) TestFilterSetScript_Success() {
	require := m.Require()
	filter := &Filter{
		Script: &Script{
			Source: "interval.start > 5 && interval.end < 8 && interval.gaps == 0",
		},
	}

	require.Equal(filter, NewFilter().
		SetScript(NewFilterScript("interval.start > 5 && interval.end < 8 && interval.gaps == 0")))
}

func (m *IntervalsTestSuite) TestScriptSetLang_Success() {
	require := m.Require()
	script := &Script{
		Source: "interval.start > 5 && interval.end < 8 && interval.gaps == 0",
		Lang:   "painless",
	}

	require.Equal(script, NewFilterScript("interval.start > 5 && interval.end < 8 && interval.gaps == 0").
		SetLang("painless"))
}

func (m *IntervalsTestSuite) TestScriptSetParams_Success() {
	require := m.Require()
	script := &Script{
		Source: "interval.start > 5 && interval.end < 8 && interval.gaps == 0",
		Params: map[string]any{"my_modifier": 2},
	}

	require.Equal(script, NewFilterScript("interval.start > 5 && interval.end < 8 && interval.gaps == 0").
		SetParams(map[string]any{"my_modifier": 2}))
}

func (m *IntervalsTestSuite) TestPrefixSetAnalyzer_Success() {
	require := m.Require()
	prefix := &Prefix{
		Prefix:   "pre",
		Analyzer: "standard",
	}

	require.Equal(prefix, NewParamPrefix("pre").
		SetAnalyzer("standard"))
}

func (m *IntervalsTestSuite) TestPrefixSetUseField_Success() {
	require := m.Require()
	prefix := &Prefix{
		Prefix:   "pre",
		UseField: "field",
	}

	require.Equal(prefix, NewParamPrefix("pre").
		SetUseField("field"))
}

func (m *IntervalsTestSuite) TestWildcardSetAnalyzer_Success() {
	require := m.Require()
	wildcard := &Wildcard{
		Pattern:  "pre*",
		Analyzer: "standard",
	}

	require.Equal(wildcard, NewParamWildcard("pre*").
		SetAnalyzer("standard"))
}

func (m *IntervalsTestSuite) TestWildcardSetUseField_Success() {
	require := m.Require()
	wildcard := &Wildcard{
		Pattern:  "pre*",
		UseField: "field",
	}

	require.Equal(wildcard, NewParamWildcard("pre*").
		SetUseField("field"))
}

func (m *IntervalsTestSuite) TestFuzzySetAnalyzer_Success() {
	require := m.Require()
	fuzzy := &Fuzzy{
		Term:     "value",
		Analyzer: "standard",
	}

	require.Equal(fuzzy, NewParamFuzzy("value").
		SetAnalyzer("standard"))
}

func (m *IntervalsTestSuite) TestFuzzySetFuzziness_Success() {
	require := m.Require()
	fuzzy := &Fuzzy{
		Term:      "value",
		Fuzziness: "AUTO",
	}

	require.Equal(fuzzy, NewParamFuzzy("value").
		SetFuzziness("AUTO"))
}

func (m *IntervalsTestSuite) TestFuzzySetPrefixLength_Success() {
	require := m.Require()
	fuzzy := &Fuzzy{
		Term:         "value",
		PrefixLength: 0,
	}

	require.Equal(fuzzy, NewParamFuzzy("value").
		SetPrefixLength(0))
}

func (m *IntervalsTestSuite) TestFuzzySetTranspositions_Success() {
	require := m.Require()
	fuzzy := &Fuzzy{
		Term:           "value",
		Transpositions: true,
	}

	require.Equal(fuzzy, NewParamFuzzy("value").
		SetTranspositions(true))
}

func (m *IntervalsTestSuite) TestFuzzySetUseField_Success() {
	require := m.Require()
	fuzzy := &Fuzzy{
		Term:     "value",
		UseField: "field",
	}

	require.Equal(fuzzy, NewParamFuzzy("value").
		SetUseField("field"))
}

func (m *IntervalsTestSuite) TestAllOfSetFilter_Success() {
	require := m.Require()
	allOf := &AllOf{
		Intervals: []*Param{},
		Filter: &Filter{
			After: &Param{Match: &Match{Query: "value"}},
		},
	}

	require.Equal(allOf, NewParamAllOf([]*Param{}).
		SetFilter(NewFilter().SetAfter(NewParam().SetMatch(NewParamMatch("value")))))
}

func (m *IntervalsTestSuite) TestAllOfSetOrdered_Success() {
	require := m.Require()
	allOf := &AllOf{
		Intervals: []*Param{},
		Ordered:   false,
	}

	require.Equal(allOf, NewParamAllOf([]*Param{}).
		SetOrdered(false))
}

func (m *IntervalsTestSuite) TestAllOfSetMaxGaps_Success() {
	require := m.Require()
	allOf := &AllOf{
		Intervals: []*Param{},
		MaxGaps:   -1,
	}

	require.Equal(allOf, NewParamAllOf([]*Param{}).
		SetMaxGaps(-1))
}

func (m *IntervalsTestSuite) TestAnyOfSetFilter_Success() {
	require := m.Require()
	anyOf := &AnyOf{
		Intervals: []*Param{},
		Filter: &Filter{
			Before: &Param{Match: &Match{Query: "value"}},
		},
	}

	require.Equal(anyOf, NewParamAnyOf([]*Param{}).
		SetFilter(NewFilter().SetBefore(NewParam().SetMatch(NewParamMatch("value")))))
}

func TestIntervals(t *testing.T) {
	suite.Run(t, new(IntervalsTestSuite))
}
