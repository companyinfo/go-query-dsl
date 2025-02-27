// Package intervals provides structures and functions for building intervals queries.
// Returns documents based on the order and proximity of matching terms.
package intervals

// Match Matches analyzed text.
type Match struct {
	Query    string  `json:"query"`
	Analyzer string  `json:"analyzer,omitempty"`
	Filter   *Filter `json:"filter,omitempty"`
	MaxGaps  int     `json:"max_gaps,omitempty"`
	Ordered  bool    `json:"ordered,omitempty"`
	UseField string  `json:"use_field,omitempty"`
}

// NewParamMatch creates a new Match instance for the intervals query parameters.
func NewParamMatch(q string) *Match {
	return &Match{Query: q}
}

// SetAnalyzer sets the Analyzer for the match.
// The analyzer used to analyze the query text. Default is the analyzer specified for the <field>.
func (m *Match) SetAnalyzer(value string) *Match {
	m.Analyzer = value
	return m
}

// SetFilter sets the Filter for the match.
// A rule used to filter returned intervals.
func (m *Match) SetFilter(f *Filter) *Match {
	m.Filter = f
	return m
}

// SetMaxGaps sets the MaxGaps for the match.
// The maximum allowed number of positions between the matching terms.
// Terms further apart than max_gaps are not considered matches.
// If max_gaps is not specified or is set to -1, terms are considered matches regardless of their position.
// If max_gaps is set to 0, matching terms must appear next to each other.
// Default is -1.
func (m *Match) SetMaxGaps(value int) *Match {
	m.MaxGaps = value
	return m
}

// SetOrdered sets the Ordered for the match.
// Specifies whether matching terms must appear in their specified order.
// Default is false.
func (m *Match) SetOrdered(value bool) *Match {
	m.Ordered = value
	return m
}

// SetUseField sets the UseField for the match.
// Specifies to search this field instead of the top-level.
// Terms are analyzed using the search analyzer specified for this field.
// By specifying `use_field`, you can search across multiple fields as if they were all the same field.
// For example, if you index the same text into stemmed and unstemmed fields, you can search for stemmed tokens
// that are near unstemmed ones.
func (m *Match) SetUseField(value string) *Match {
	m.UseField = value
	return m
}
