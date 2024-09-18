// Package intervals provides structures and functions for building intervals queries.
// Returns documents based on the order and proximity of matching terms.
package intervals

// Filter rule is used to restrict the results.
type Filter struct {
	After          *Param  `json:"after,omitempty"`
	Before         *Param  `json:"before,omitempty"`
	ContainedBy    *Param  `json:"contained_by,omitempty"`
	Containing     *Param  `json:"containing,omitempty"`
	NotContainedBy *Param  `json:"not_contained_by,omitempty"`
	NotContaining  *Param  `json:"not_containing,omitempty"`
	NotOverlapping *Param  `json:"not_overlapping,omitempty"`
	Overlapping    *Param  `json:"overlapping,omitempty"`
	Script         *Script `json:"script,omitempty"`
}

// NewFilter creates a new Filter instance for the intervals query parameters.
func NewFilter() *Filter {
	return &Filter{}
}

// SetAfter sets the After for the Filter.
// A query used to return intervals that follow an interval specified in the filter rule.
func (f *Filter) SetAfter(p *Param) *Filter {
	f.After = p
	return f
}

// SetBefore sets the Before for the Filter.
// A query used to return intervals that are before an interval specified in the filter rule.
func (f *Filter) SetBefore(p *Param) *Filter {
	f.Before = p
	return f
}

// SetContainedBy sets the ContainedBy for the Filter.
// A query used to return intervals contained by an interval specified in the filter rule.
func (f *Filter) SetContainedBy(p *Param) *Filter {
	f.ContainedBy = p
	return f
}

// SetContaining sets the Containing for the Filter.
// A query used to return intervals that contain an interval specified in the filter rule.
func (f *Filter) SetContaining(p *Param) *Filter {
	f.Containing = p
	return f
}

// SetNotContainedBy sets the NotContainedBy for the Filter.
// A query used to return intervals that are not contained by an interval specified in the filter rule.
func (f *Filter) SetNotContainedBy(p *Param) *Filter {
	f.NotContainedBy = p
	return f
}

// SetNotContaining sets the NotContaining for the Filter.
// A query used to return intervals that do not contain an interval specified in the filter rule.
func (f *Filter) SetNotContaining(p *Param) *Filter {
	f.NotContaining = p
	return f
}

// SetNotOverlapping sets the NotOverlapping for the Filter.
// A query used to return intervals that do not overlap with an interval specified in the filter rule.
func (f *Filter) SetNotOverlapping(p *Param) *Filter {
	f.NotOverlapping = p
	return f
}

// SetOverlapping sets the Overlapping for the Filter.
// A query used to return intervals that overlap with an interval specified in the filter rule.
func (f *Filter) SetOverlapping(p *Param) *Filter {
	f.Overlapping = p
	return f
}

// SetScript sets the Script for the Filter.
// A script used to match documents. This script must return true or false.
func (f *Filter) SetScript(s *Script) *Filter {
	f.Script = s
	return f
}

// Script represents a Script Filter.
type Script struct {
	Source string         `json:"source"`
	Lang   string         `json:"lang,omitempty"`
	Params map[string]any `json:"params,omitempty"`
}

// NewFilterScript creates a new Script instance for the Filter.
func NewFilterScript(s string) *Script {
	return &Script{Source: s}
}

// SetLang sets the lang for the script.
// Specifies the language the script is written in. Defaults to painless.
func (s *Script) SetLang(value string) *Script {
	s.Lang = value
	return s
}

// SetParams sets the params for the script.
// Specifies any named parameters that are passed into the script as variables.
func (s *Script) SetParams(value map[string]any) *Script {
	s.Params = value
	return s
}
