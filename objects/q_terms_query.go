package objects

import (
	"encoding/json"

	"github.com/sdqri/effdsl/utils"
)

type TermsQueryS struct {
	Field  string   `json:"-"`               // (Required, string) Field you wish to search.
	Values []string `json:"-"`               // (Required, array) Array of terms you wish to find in the provided field.
	Boost  *float64 `json:"boost,omitempty"` // (Optional, float) Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.
}

func (tq TermsQueryS) QueryInfo() string {
	return "Terms query"
}

func (tq TermsQueryS) MarshalJSON() ([]byte, error) {
	type TermsQueryBase TermsQueryS
	m, err := utils.CastStruct[TermsQueryBase, M](TermsQueryBase(tq))
	if err != nil {
		return nil, err
	}
	m[tq.Field] = tq.Values

	return json.Marshal(
		M{
			"range": m,
		},
	)
}

type TermsQueryOption func(*TermsQueryS)

//TODO: Fix repetitive methods with DefineType problem !important
func WithTSQBoost(boost float64) TermsQueryOption {
	return func(termsQuery *TermsQueryS) {
		termsQuery.Boost = &boost
	}
}

// Returns documents that contain one or more exact terms in a provided field.
// [Terms query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-query.html
func TermsQuery(field string, values []string, opts ...TermsQueryOption) QueryResult {
	termsQuery := TermsQueryS{
		Field:  field,
		Values: values,
	}
	for _, opt := range opts {
		opt(&termsQuery)
	}
	return QueryResult{
		Ok:  termsQuery,
		Err: nil,
	}
}
