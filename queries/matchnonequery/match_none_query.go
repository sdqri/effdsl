package matchnonequery

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

// MatchNoneQueryS represents the structure of the Match None query.
type MatchNoneQueryS struct{}

func (mq MatchNoneQueryS) QueryInfo() string {
	return "Match None query"
}

func (mq MatchNoneQueryS) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		effdsl.M{
			"match_none": struct{}{},
		},
	)
}

// MatchNoneQuery is the inverse of the match_all query, which matches no documents.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html
func MatchNoneQuery() effdsl.QueryResult {
	return effdsl.QueryResult{
		Ok:  MatchNoneQueryS{},
		Err: nil,
	}
}
