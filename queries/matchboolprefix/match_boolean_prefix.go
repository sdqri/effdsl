package matchboolprefix

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

type MatchBoolPrefixQueryS struct {
	Field    string `json:"-"`                  // (Required, string) The field to search.
	Query    string `json:"query"`              // (Required, string) The query text to search for.
	Analyzer string `json:"analyzer,omitempty"` // (Optional, string) The analyzer to use for the query text.
}

func (mq MatchBoolPrefixQueryS) QueryInfo() string {
	return "Match bool prefix query"
}

func (mq MatchBoolPrefixQueryS) MarshalJSON() ([]byte, error) {
	type MatchBoolPrefixQueryBase MatchBoolPrefixQueryS
	return json.Marshal(
		effdsl.M{
			"match_bool_prefix": effdsl.M{
				mq.Field: (MatchBoolPrefixQueryBase)(mq),
			},
		},
	)
}

type MatchBoolPrefixQueryOption func(*MatchBoolPrefixQueryS)

func WithAnalyzer(analyzer string) MatchBoolPrefixQueryOption {
	return func(query *MatchBoolPrefixQueryS) {
		query.Analyzer = analyzer
	}
}

// MatchBoolPrefixQuery analyzes its input and constructs a bool query from the terms. Each term except the last is used in a term query. The last term is used in a prefix query.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-bool-prefix-query.html
func MatchBoolPrefixQuery(field, query string, opts ...MatchBoolPrefixQueryOption) effdsl.QueryResult {
	matchQuery := MatchBoolPrefixQueryS{
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
