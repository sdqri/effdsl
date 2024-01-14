package objects

import (
	"encoding/json"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html
type WildcardQueryS struct {
	Field   string  `json:"-"`                 // (Required, object) Field you wish to search.
	Value   string  `json:"value"`             // (Required, string) Wildcard pattern for terms you wish to find in the provided <field>.
	Boost   float32 `json:"boost,omitempty"`   // (Optional, float) Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.
	Rewrite string  `json:"rewrite,omitempty"` // (Optional, string) Method used to rewrite the query. For valid values and more information, see the rewrite parameter.
}

func (wq WildcardQueryS) QueryInfo() string {
	return "Wildcard query"
}

func (wq WildcardQueryS) MarshalJSON() ([]byte, error) {
	type WildcardQuerySBase WildcardQueryS
	return json.Marshal(
		M{
			"wildcard": M{
				wq.Field: (WildcardQuerySBase)(wq),
			},
		},
	)
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-term-rewrite.html
type RewriteParameter string

const (
	RewriteParameterConstantScoreBlended  RewriteParameter = "constant_score_blended"
	RewriteParameterConstantScore         RewriteParameter = "constant_score"
	RewriteParameterConstantScoreBoolean  RewriteParameter = "constant_score_boolean"
	RewriteParameterScoringBoolean        RewriteParameter = "scoring_boolean"
	RewriteParameterTopTermsBlendedFreqsN RewriteParameter = "top_terms_blended_freqs_N"
	RewriteParameterTopTermsBoostN        RewriteParameter = "top_terms_boost_N"
	RewriteParameterTopTermsN             RewriteParameter = "top_terms_N"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html#wildcard-query-field-params
type wildcardQueryFieldParameters struct {
	/*
		(Optional, float) Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.

		You can use the boost parameter to adjust relevance scores for searches containing two or more queries.

		Boost values are relative to the default value of 1.0. A boost value between 0 and 1.0 decreases the relevance score. A value greater than 1.0 increases the relevance score.
	*/
	Boost float32

	/*
		(Optional, string) Method used to rewrite the query. For valid values and more information, see the rewrite parameter.
	*/
	RewriteParameter RewriteParameter
}

type WildcardQueryFieldParameter func(prms *wildcardQueryFieldParameters)

// WithBoost ...
func WithBoost(boost float32) WildcardQueryFieldParameter {
	return func(prms *wildcardQueryFieldParameters) {
		prms.Boost = boost
	}
}

// WithRewriteParameter ...
func WithRewriteParameter(rwp RewriteParameter) WildcardQueryFieldParameter {
	return func(prms *wildcardQueryFieldParameters) {
		prms.RewriteParameter = rwp
	}
}

// Returns documents that contain terms matching a wildcard pattern.
// [Wildcard query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html
func WildcardQuery(field string, value string, prms ...WildcardQueryFieldParameter) QueryResult {
	wildcardQuery := WildcardQueryS{
		Field: field,
		Value: value,
	}

	var parameters wildcardQueryFieldParameters
	for _, prm := range prms {
		prm(&parameters)
	}

	wildcardQuery.Boost = parameters.Boost
	wildcardQuery.Rewrite = string(parameters.RewriteParameter)

	return QueryResult{
		Ok:  wildcardQuery,
		Err: nil,
	}
}
