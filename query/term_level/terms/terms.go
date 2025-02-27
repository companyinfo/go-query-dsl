// Package terms provides structures and functions for building terms queries.
// Returns documents that contain one or more exact terms in a provided field.
package terms

// Terms represents a terms query, allowing multiple values for a field.
type Terms[T any] struct {
	Terms map[string]T `json:"terms"`
}

// Lookup represents the parameters for a terms lookup query.
type Lookup struct {
	Index   string  `json:"index"`
	ID      string  `json:"id"`
	Path    string  `json:"path"`
	Routing string  `json:"routing,omitempty"`
	Boost   float64 `json:"boost,omitempty"`
}

// New creates a new Terms instance with the specified field and values.
func New[T any](f string, value T) *Terms[T] {
	return &Terms[T]{Terms: map[string]T{f: value}}
}

// NewLookup creates a new Lookup instance.
func NewLookup(index, id, path string) *Lookup {
	return &Lookup{
		Index: index,
		ID:    id,
		Path:  path,
	}
}

// SetBoost sets the boost for the terms query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (t *Terms[T]) SetBoost(value T) *Terms[T] {
	t.Terms["boost"] = value
	return t
}

// SetRouting sets the routing for the terms lookup query parameter.
// Custom routing value of the document from which to fetch field values.
func (l *Lookup) SetRouting(value string) *Lookup {
	l.Routing = value
	return l
}

// SetBoost sets the boost for the terms lookup query parameter.
// A floating-point value that specifies the weight of this field toward the relevance score.
// Values above 1.0 increase the field’s relevance.
// Values between 0.0 and 1.0 decrease the field’s relevance.
// Default is 1.0.
func (l *Lookup) SetBoost(value float64) *Lookup {
	l.Boost = value
	return l
}
