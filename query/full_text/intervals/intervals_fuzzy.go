// Package intervals provides structures and functions for building intervals queries.
// Returns documents based on the order and proximity of matching terms.
package intervals

// Fuzzy represents the Fuzzy interval query
type Fuzzy struct {
	Term           string `json:"term"`
	Analyzer       string `json:"analyzer,omitempty"`
	Fuzziness      string `json:"fuzziness,omitempty"`
	PrefixLength   int    `json:"prefix_length,omitempty"`
	Transpositions bool   `json:"transpositions,omitempty"`
	UseField       string `json:"use_field,omitempty"`
}

// NewParamFuzzy creates a new Fuzzy instance for the intervals query parameters.
func NewParamFuzzy(t string) *Fuzzy {
	return &Fuzzy{Term: t}
}

// SetAnalyzer sets the Analyzer for the Fuzzy.
// The analyzer used to analyze the query text. Default is the analyzer specified for the <field>.
func (f *Fuzzy) SetAnalyzer(value string) *Fuzzy {
	f.Analyzer = value
	return f
}

// SetFuzziness sets the Fuzziness for the Fuzzy.
// The number of character edits (insert, delete, substitute) that it takes to change one word to another
// when determining whether a term matched a value.
// For example, the distance between wined and wind is 1.
// Valid values are non-negative integers or AUTO.
// The default, AUTO, chooses a value based on the length of each term and is a good choice for most use cases.
func (f *Fuzzy) SetFuzziness(value string) *Fuzzy {
	f.Fuzziness = value
	return f
}

// SetPrefixLength sets the PrefixLength for the Fuzzy.
// The number of beginning characters left unchanged for fuzzy matching.
// Default is 0.
func (f *Fuzzy) SetPrefixLength(value int) *Fuzzy {
	f.PrefixLength = value
	return f
}

// SetTranspositions sets the Transpositions for the Fuzzy.
// Setting transpositions to true (default) adds swaps of adjacent characters to the insert,
// delete, and substitute operations of the fuzziness option.
// For example, the distance between wind and wnid is 1 if transpositions is true (swap “n” and “i”) and 2 if
// it is false (delete “n”, insert “n”).
// If transpositions is false, rewind and wnid have the same distance (2) from wind, despite the
// more human-centric opinion that wnid is an obvious typo.
// The default is a good choice for most use cases.
func (f *Fuzzy) SetTranspositions(value bool) *Fuzzy {
	f.Transpositions = value
	return f
}

// SetUseField sets the UseField for the Fuzzy.
// Specifies to search this field instead of the top-level .
// The `prefix` is normalized using the search analyzer specified for this field, unless you specify an `analyzer`.
func (f *Fuzzy) SetUseField(value string) *Fuzzy {
	f.UseField = value
	return f
}
