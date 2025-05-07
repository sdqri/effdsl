package multimatchquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

type MultiMatchQueryS struct {
	Query                           string         `json:"query"`                                         // (Required, string) Query text to search for.
	Fields                          []string       `json:"fields,omitempty"`                              // (Optional, []string) Fields to search.
	Type                            MultiMatchType `json:"type,omitempty"`                                // (Optional, string) Type of multi match query.
	Operator                        Operator       `json:"operator,omitempty"`                            // (Optional, string) Operator used to combine terms.
	Analyzer                        string         `json:"analyzer,omitempty"`                            // (Optional, string) Analyzer to use for the query.
	Slop                            int            `json:"slop,omitempty"`                                // (Optional, integer) Maximum positions allowed between matching tokens.
	Fuzziness                       string         `json:"fuzziness,omitempty"`                           // (Optional, string) Fuzziness for fuzzy matching.
	PrefixLength                    int            `json:"prefix_length,omitempty"`                       // (Optional, integer) Prefix length for fuzzy matching.
	MaxExpansions                   int            `json:"max_expansions,omitempty"`                      // (Optional, integer) Max expansions for fuzzy queries.
	MinimumShouldMatch              string         `json:"minimum_should_match,omitempty"`                // (Optional, string) Minimum should match constraint.
	TieBreaker                      float64        `json:"tie_breaker,omitempty"`                         // (Optional, float) Tie breaker for scoring.
	Lenient                         bool           `json:"lenient,omitempty"`                             // (Optional, bool) Leniency for parsing query.
	ZeroTermsQuery                  ZeroTerms      `json:"zero_terms_query,omitempty"`                    // (Optional, string) Handling zero terms.
	AutoGenerateSynonymsPhraseQuery bool           `json:"auto_generate_synonyms_phrase_query,omitempty"` // (Optional, bool)
}

func (mq MultiMatchQueryS) QueryInfo() string {
	return "Multi match query"
}

func (mq MultiMatchQueryS) MarshalJSON() ([]byte, error) {
	type MultiMatchQueryBase MultiMatchQueryS
	return json.Marshal(
		effdsl.M{
			"multi_match": (MultiMatchQueryBase)(mq),
		},
	)
}

type MultiMatchQueryOption func(*MultiMatchQueryS)

func WithFields(fields ...string) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.Fields = fields
	}
}

func WithType(t MultiMatchType) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.Type = t
	}
}

func WithOperator(op Operator) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.Operator = op
	}
}

func WithAnalyzer(analyzer string) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.Analyzer = analyzer
	}
}

func WithSlop(slop int) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.Slop = slop
	}
}

func WithFuzziness(fuzziness string) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.Fuzziness = fuzziness
	}
}

func WithPrefixLength(prefixLength int) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.PrefixLength = prefixLength
	}
}

func WithMaxExpansions(maxExpansions int) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.MaxExpansions = maxExpansions
	}
}

func WithMinimumShouldMatch(msm string) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.MinimumShouldMatch = msm
	}
}

func WithTieBreaker(tieBreaker float64) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.TieBreaker = tieBreaker
	}
}

func WithLenient(lenient bool) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.Lenient = lenient
	}
}

func WithZeroTermsQuery(ztq ZeroTerms) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.ZeroTermsQuery = ztq
	}
}

func WithAutoGenerateSynonymsPhraseQuery(agspq bool) MultiMatchQueryOption {
	return func(query *MultiMatchQueryS) {
		query.AutoGenerateSynonymsPhraseQuery = agspq
	}
}

type Operator string

const (
	OperatorOr  Operator = "or"
	OperatorAnd Operator = "and"
)

type MultiMatchType string

const (
	BestFields   MultiMatchType = "best_fields"
	MostFields   MultiMatchType = "most_fields"
	CrossFields  MultiMatchType = "cross_fields"
	Phrase       MultiMatchType = "phrase"
	PhrasePrefix MultiMatchType = "phrase_prefix"
	BoolPrefix   MultiMatchType = "bool_prefix"
)

type ZeroTerms string

const (
	None ZeroTerms = "none"
	All  ZeroTerms = "all"
)

// MultiMatchQuery creates a multi_match query for provided query text and optional settings.
func MultiMatchQuery(query string, opts ...MultiMatchQueryOption) effdsl.QueryResult {
	multiMatchQuery := MultiMatchQueryS{
		Query: query,
	}
	for _, opt := range opts {
		opt(&multiMatchQuery)
	}
	return effdsl.QueryResult{
		Ok:  multiMatchQuery,
		Err: nil,
	}
}
