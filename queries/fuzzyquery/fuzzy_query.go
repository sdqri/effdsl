package fuzzyquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

// FuzzyQueryS represents a query that matches documents containing terms
// similar to the specified term with a given degree of similarity.
type FuzzyQueryS struct {
	Field          string  `json:"-"`                        // (Required, string) The field to search.
	Value          string  `json:"value"`                    // (Required, string) The term to search for.
	Fuzziness      string  `json:"fuzziness,omitempty"`      // (Optional, string) The degree of fuzziness allowed for the search term (e.g. "AUTO", "1", "2", etc.).
	MaxExpansions  int     `json:"max_expansions,omitempty"` // (Optional, int) The maximum number of terms to match.
	PrefixLength   int     `json:"prefix_length,omitempty"`  // (Optional, int) The number of initial characters that must match exactly.
	Transpositions bool    `json:"transpositions,omitempty"` // (Optional, bool) Whether to allow transpositions of two adjacent characters.
	Rewrite        Rewrite `json:"rewrite,omitempty"`        // (Optional, string) The method used to rewrite the query (e.g. "constant_score_boolean").
}

func (fq FuzzyQueryS) QueryInfo() string {
	return "Fuzzy query"
}

func (fq FuzzyQueryS) MarshalJSON() ([]byte, error) {
	type FuzzyQueryBase FuzzyQueryS
	return json.Marshal(
		effdsl.M{
			"fuzzy": effdsl.M{
				fq.Field: (FuzzyQueryBase)(fq),
			},
		},
	)
}

type FuzzyQueryOption func(*FuzzyQueryS)

func WithFuzziness(fuzziness string) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.Fuzziness = fuzziness
	}
}

func WithMaxExpansions(maxExpansions int) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.MaxExpansions = maxExpansions
	}
}

func WithPrefixLength(prefixLength int) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.PrefixLength = prefixLength
	}
}

func WithTranspositions(transpositions bool) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.Transpositions = transpositions
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

func WithRewrite(rewrite Rewrite) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.Rewrite = rewrite
	}
}

// FuzzyQuery returns documents that contain terms similar to the specified term with a given degree of similarity.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
func FuzzyQuery(field string, value string, opts ...FuzzyQueryOption) effdsl.QueryResult {
	fuzzyQuery := FuzzyQueryS{
		Field: field,
		Value: value,
	}
	for _, opt := range opts {
		opt(&fuzzyQuery)
	}
	return effdsl.QueryResult{
		Ok:  fuzzyQuery,
		Err: nil,
	}
}
