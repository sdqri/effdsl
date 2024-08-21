package regexpquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl"
)

// RegexpQueryS represents a query that matches documents containing terms
// with a specified regular expression.
type RegexpQueryS struct {
	Field                 string  `json:"-"`                                 // (Required, string) The field to search.
	Value                 string  `json:"value"`                             // (Required, string) The regular expression pattern to match against the field.
	Flags                 string  `json:"flags,omitempty"`                   // (Optional, string) Additional matching options for the regular expression.
	CaseInsensitive       bool    `json:"case_insensitive,omitempty"`        // (Optional, bool) Whether the regular expression is case-insensitive.
	MaxDeterminizedStates int     `json:"max_determinized_states,omitempty"` // (Optional, int) The maximum number of automaton states required for the query. Lower values will reduce memory usage but increase query time.
	Rewrite               Rewrite `json:"rewrite,omitempty"`                 // (Optional, string) The method used to rewrite the query. Can be "constant_score_boolean", "constant_score_filter", "scoring_boolean", "top_terms_boost_N" (where N is the number of top terms), "top_terms_N" (where N is the number of top terms), "random_access_N" (where N is the maximum number of matching terms).
}

func (rq RegexpQueryS) QueryInfo() string {
	return "Regexp query"
}

func (rq RegexpQueryS) MarshalJSON() ([]byte, error) {
	type RegexpQueryBase RegexpQueryS
	return json.Marshal(
		effdsl.M{
			"regexp": effdsl.M{
				rq.Field: (RegexpQueryBase)(rq),
			},
		},
	)
}

type RegexpQueryOption func(*RegexpQueryS)

func WithFlags(flags string) RegexpQueryOption {
	return func(regexpQuery *RegexpQueryS) {
		regexpQuery.Flags = flags
	}
}

func WithCaseInsensitive() RegexpQueryOption {
	return func(regexpQuery *RegexpQueryS) {
		regexpQuery.CaseInsensitive = true
	}
}

func WithMaxDeterminizedStates(states int) RegexpQueryOption {
	return func(regexpQuery *RegexpQueryS) {
		regexpQuery.MaxDeterminizedStates = states
	}
}

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

func WithRewrite(rewrite Rewrite) RegexpQueryOption {
	return func(rq *RegexpQueryS) {
		rq.Rewrite = rewrite
	}
}

// Returns documents that contain terms matching a regular expression.
// [Regexp query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-regexp-query.html
func RegexpQuery(field string, value string, opts ...RegexpQueryOption) effdsl.QueryResult {
	regexpQuery := RegexpQueryS{
		Field: field,
		Value: value,
	}
	for _, opt := range opts {
		opt(&regexpQuery)
	}
	return effdsl.QueryResult{
		Ok:  regexpQuery,
		Err: nil,
	}
}
