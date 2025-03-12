package prefixquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

// PrefixQueryS represents a query that matches documents containing terms
// that start with a specific prefix in a given field.
type PrefixQueryS struct {
	Field           string  `json:"-"`                          // (Required, string) The field to search.
	Value           string  `json:"value"`                      // (Required, string) The prefix to match terms against.
	Rewrite         Rewrite `json:"rewrite,omitempty"`          // (Optional, string) The method used to rewrite the query.
	CaseInsensitive bool    `json:"case_insensitive,omitempty"` // (Optional, bool) Whether the query is case insensitive.
}

func (pq PrefixQueryS) QueryInfo() string {
	return "Prefix query"
}

func (pq PrefixQueryS) MarshalJSON() ([]byte, error) {
	type PrefixQueryBase PrefixQueryS
	return json.Marshal(
		effdsl.M{
			"prefix": effdsl.M{
				pq.Field: (PrefixQueryBase)(pq),
			},
		},
	)
}

type PrefixQueryOption func(*PrefixQueryS)

// Rewrite represents the type of rewrite to use in a fuzzy query.
type Rewrite string

const (
	ConstantScoreBlended  Rewrite = "constant_score_blended"
	ConstantScore         Rewrite = "constant_score"
	ConstantScoreBoolean  Rewrite = "constant_score_boolean"
	ScoringBoolean        Rewrite = "scoring_boolean"
	TopTermsBlendedFreqsN Rewrite = "top_terms_blended_freqs_N"
	TopTermsBoostN        Rewrite = "top_terms_boost_N"
	TopTermsN             Rewrite = "top_terms_N"
)

func WithRewrite(rewrite Rewrite) PrefixQueryOption {
	return func(params *PrefixQueryS) {
		params.Rewrite = rewrite
	}
}

func WithCaseInsensitive(caseInsensitive bool) PrefixQueryOption {
	return func(params *PrefixQueryS) {
		params.CaseInsensitive = caseInsensitive
	}
}

// PrefixQuery returns documents that contain terms starting with the specified prefix.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
func PrefixQuery(field string, value string, opts ...PrefixQueryOption) effdsl.QueryResult {
	prefixQuery := PrefixQueryS{
		Field: field,
		Value: value,
	}
	for _, opt := range opts {
		opt(&prefixQuery)
	}
	return effdsl.QueryResult{
		Ok:  prefixQuery,
		Err: nil,
	}
}
