// Package suggestion accepts a list of suggestions and builds them into a finite-state transducer (FST),
// an optimized data structure that is essentially a graph.
// This data structure lives in memory and is optimized for fast prefix lookups.
package suggestion

// Autocomplete represents the autocomplete query.
type Autocomplete struct {
	Autocomplete *Param `json:"autocomplete"`
}

// Param represents the parameters for an autocomplete query, including prefix, completion, and fuzzy options.
type Param struct {
	Prefix     string      `json:"prefix"`
	Completion *Completion `json:"completion"`
}

// Completion represents the completion details for an autocomplete query, including field, size, and fuzzy options.
type Completion struct {
	Field string `json:"field"`
	Size  int    `json:"size,omitempty"`
	Fuzzy *Fuzzy `json:"fuzzy,omitempty"`
}

// Fuzzy represents the fuzzy options for an autocomplete query.
type Fuzzy struct {
	Fuzziness string `json:"fuzziness,omitempty"`
}

// NewAutocomplete creates a new Autocomplete instance with the specified parameters.
func NewAutocomplete(p *Param) *Autocomplete {
	return &Autocomplete{Autocomplete: p}
}

// NewAutocompleteParam creates a new Param instance with the specified prefix and an empty completion.
func NewAutocompleteParam(prefix, field string) *Param {
	return &Param{Prefix: prefix, Completion: &Completion{
		Field: field,
	}}
}

// SetSize sets the size of the completion in an autocomplete query.
// An integer that specifies the maximum number of returned suggestions.
// Default is 5.
func (p *Param) SetSize(value int) *Param {
	p.Completion.Size = value
	return p
}

// SetFuzziness sets the fuzziness for the completion in the autocomplete query.
// The number of character edits (insert, delete, substitute) that it takes to change one word to another
// when determining whether a term matched a value.
// For example, the distance between wined and wind is 1.
// Valid values are non-negative integers or AUTO.
// The default, AUTO, chooses a value based on the length of each term and is a good choice for most use cases.
func (p *Param) SetFuzziness(value string) *Param {
	p.Completion.Fuzzy = &Fuzzy{Fuzziness: value}
	return p
}
