package matchphrasequery

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

type MatchPhraseQueryS struct {
	Field          string    `json:"-"`                          // (Required, string) The field to search.
	Query          string    `json:"query"`                      // (Required, string) The query text to search for.
	Analyzer       string    `json:"analyzer,omitempty"`         // (Optional, string) Analyzer used to convert text in the query string into tokens.
	Slop           int       `json:"slop,omitempty"`             // (Optional, integer) Maximum number of positions allowed between matching tokens for phrases.
	ZeroTermsQuery ZeroTerms `json:"zero_terms_query,omitempty"` // (Optional, string) What to do when the analyzed text contains no terms. Valid values are "none" (default) or "all".
}

func (mq MatchPhraseQueryS) QueryInfo() string {
	return "Match phrase query"
}

func (mq MatchPhraseQueryS) MarshalJSON() ([]byte, error) {
	type MatchPhraseQueryBase MatchPhraseQueryS
	return json.Marshal(
		effdsl.M{
			"match_phrase": effdsl.M{
				mq.Field: (MatchPhraseQueryBase)(mq),
			},
		},
	)
}

type MatchPhraseQueryOption func(*MatchPhraseQueryS)

func WithAnalyzer(analyzer string) MatchPhraseQueryOption {
	return func(query *MatchPhraseQueryS) {
		query.Analyzer = analyzer
	}
}

func WithSlop(slop int) MatchPhraseQueryOption {
	return func(query *MatchPhraseQueryS) {
		query.Slop = slop
	}
}

type ZeroTerms string

const (
	None ZeroTerms = "none"
	All  ZeroTerms = "all"
)

func WithZeroTermsquery(zt ZeroTerms) MatchPhraseQueryOption {
	return func(params *MatchPhraseQueryS) {
		params.ZeroTermsQuery = zt
	}
}

// MatchPhraseQuery creates a match_phrase query for a given field and query text.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase.html
func MatchPhraseQuery(field, query string, opts ...MatchPhraseQueryOption) effdsl.QueryResult {
	matchQuery := MatchPhraseQueryS{
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
