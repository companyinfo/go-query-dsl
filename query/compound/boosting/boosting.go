// Package boosting returns documents matching a positive query while reducing the relevance score of documents
// that also match a negative query.
package boosting

// Boosting represents a boosting query.
type Boosting struct {
	Boosting Params `json:"boosting"`
}

// Params represents the parameters for a boosting query.
type Params struct {
	Positive      any     `json:"positive"`
	Negative      any     `json:"negative"`
	NegativeBoost float64 `json:"negative_boost"`
}

// New creates a new Boosting instance with initialized query parameters.
func New() *Boosting {
	return &Boosting{}
}

// SetPositive sets the query you wish to run. Any returned documents must match this query.
func (b *Boosting) SetPositive(p any) *Boosting {
	b.Boosting.Positive = p
	return b
}

// SetNegative sets the query used to decrease the relevance score of matching documents.
func (b *Boosting) SetNegative(n any) *Boosting {
	b.Boosting.Negative = n
	return b
}

// SetNegativeBoost sets floating point number between 0 and 1.0 used to decrease the
// relevance scores of documents matching the negative query.
func (b *Boosting) SetNegativeBoost(boost float64) *Boosting {
	b.Boosting.NegativeBoost = boost
	return b
}
