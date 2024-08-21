package matchphraseprefix

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

type MatchPhrasePrefixQueryS struct {
	Field          string    `json:"-"`                          // (Required, string) The field to search.
	Query          string    `json:"query"`                      // (Required, string) The query text to search for.
	Analyzer       string    `json:"analyzer,omitempty"`         // (Optional, string) Analyzer used to convert text in the query string into tokens.
	Slop           int       `json:"slop,omitempty"`             // (Optional, integer) Maximum number of positions allowed between matching tokens for phrases.
	MaxExpansions  int       `json:"max_expansions,omitempty"`   // (Optional, integer) Maximum number of terms to which the last provided term will expand.
	ZeroTermsQuery ZeroTerms `json:"zero_terms_query,omitempty"` // (Optional, string) What to do when the analyzed text contains no terms. Valid values are "none" (default) or "all".
}

func (mq MatchPhrasePrefixQueryS) QueryInfo() string {
	return "Match phrase prefix query"
}

func (mq MatchPhrasePrefixQueryS) MarshalJSON() ([]byte, error) {
	type MatchPhrasePrefixQueryBase MatchPhrasePrefixQueryS
	return json.Marshal(
		effdsl.M{
			"match_phrase_prefix": effdsl.M{
				mq.Field: (MatchPhrasePrefixQueryBase)(mq),
			},
		},
	)
}

type MatchPhrasePrefixQueryOption func(*MatchPhrasePrefixQueryS)

func WithAnalyzer(analyzer string) MatchPhrasePrefixQueryOption {
	return func(query *MatchPhrasePrefixQueryS) {
		query.Analyzer = analyzer
	}
}

func WithSlop(slop int) MatchPhrasePrefixQueryOption {
	return func(query *MatchPhrasePrefixQueryS) {
		query.Slop = slop
	}
}

func WithMaxExpansions(maxExpansions int) MatchPhrasePrefixQueryOption {
	return func(query *MatchPhrasePrefixQueryS) {
		query.MaxExpansions = maxExpansions
	}
}

func WithZeroTermsQuery(zeroTermsQuery ZeroTerms) MatchPhrasePrefixQueryOption {
	return func(query *MatchPhrasePrefixQueryS) {
		query.ZeroTermsQuery = zeroTermsQuery
	}
}

type ZeroTerms string

const (
	None ZeroTerms = "none"
	All  ZeroTerms = "all"
)

// MatchPhrasePrefixQuery creates a match phrase prefix query for a given field and query text.
//
// For more details, see the official Elasticsearch documentation:
// https://elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase-prefix.html
func MatchPhrasePrefixQuery(field, query string, opts ...MatchPhrasePrefixQueryOption) effdsl.QueryResult {
	matchQuery := MatchPhrasePrefixQueryS{
		Field: field,
		Query: query,
	}
	for _, opt := range opts {
		opt(&matchQuery)
	}
	return effdsl.QueryResult{
		Ok:  matchQuery,
		Err: nil,
	}
}
