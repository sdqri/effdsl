package wildcardquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html
type WildcardQueryS struct {
	Field           string  `json:"-"`                          // (Required, object) Field you wish to search.
	Boost           float64 `json:"boost,omitempty"`            // (Optional, float) Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.
	CaseInsensitive bool    `json:"case_insensitive,omitempty"` // (Optional, bool) Whether the regular expression is case-insensitive.
	Rewrite         Rewrite `json:"rewrite,omitempty"`          // (Optional, string) Method used to rewrite the query. For valid values and more information, see the rewrite parameter.
	Value           string  `json:"value"`                      // (Required, string) Wildcard pattern for terms you wish to find in the provided <field>.
	Wildcard        string  `json:"wildcard,omitempty"`         // (Required, string) An alias for the value parameter. If you specify both value and wildcard, the query uses the last one in the request body.
}

func (wq WildcardQueryS) QueryInfo() string {
	return "Wildcard query"
}

func (wq WildcardQueryS) MarshalJSON() ([]byte, error) {
	type WildcardQuerySBase WildcardQueryS
	return json.Marshal(
		effdsl.M{
			"wildcard": effdsl.M{
				wq.Field: (WildcardQuerySBase)(wq),
			},
		},
	)
}

type WildcardQueryOption func(prms *WildcardQueryS)

func WithBoost(boost float64) WildcardQueryOption {
	return func(params *WildcardQueryS) {
		params.Boost = boost
	}
}

func WithCaseInsensitive() WildcardQueryOption {
	return func(params *WildcardQueryS) {
		params.CaseInsensitive = true
	}
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-term-rewrite.html
type Rewrite string

const (
	RewriteParameterConstantScoreBlended  Rewrite = "constant_score_blended"
	RewriteParameterConstantScore         Rewrite = "constant_score"
	RewriteParameterConstantScoreBoolean  Rewrite = "constant_score_boolean"
	RewriteParameterScoringBoolean        Rewrite = "scoring_boolean"
	RewriteParameterTopTermsBlendedFreqsN Rewrite = "top_terms_blended_freqs_N"
	RewriteParameterTopTermsBoostN        Rewrite = "top_terms_boost_N"
	RewriteParameterTopTermsN             Rewrite = "top_terms_N"
)

func WithRewrite(rewrite Rewrite) WildcardQueryOption {
	return func(params *WildcardQueryS) {
		params.Rewrite = rewrite
	}
}

// Returns documents that contain terms matching a wildcard pattern.
// [Wildcard query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html
func WildcardQuery(field string, value string, opts ...WildcardQueryOption) effdsl.QueryResult {
	wildcardQuery := WildcardQueryS{
		Field: field,
		Value: value,
	}

	for _, opt := range opts {
		opt(&wildcardQuery)
	}

	return effdsl.QueryResult{
		Ok:  wildcardQuery,
		Err: nil,
	}
}
