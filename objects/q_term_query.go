package objects

import (
	"encoding/json"
)

type TermQueryS struct {
	Field string   `json:"-"`               //(Required, object) Field you wish to search.
	Value string   `json:"value"`           //(Required, string) Term you wish to find in the provided <field>. To return a document, the term must exactly match the field value, including whitespace and capitalization.
	Boost *float64 `json:"boost,omitempty"` //(Optional, float) Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.
}

func (tq TermQueryS) QueryInfo() string {
	return "Term query"
}

func (tq TermQueryS) MarshalJSON() ([]byte, error) {
	type TermQueryAlias TermQueryS
	return json.Marshal(
		M{
			"term": M{
				tq.Field: (TermQueryAlias)(tq),
			},
		},
	)
}

type TermQueryOption func(*TermQueryS)

func (f DefineType) WithBoost(boost float64) TermQueryOption {
	return func(termQuery *TermQueryS) {
		termQuery.Boost = &boost
	}
}

func (f DefineType) TermQuery(field string, value string, opts ...TermQueryOption) QueryResult {
	termQuery := TermQueryS{
		Field: field,
		Value: value,
	}
	for _, opt := range opts {
		opt(&termQuery)
	}
	return QueryResult{
		Ok:  termQuery,
		Err: nil,
	}
}
