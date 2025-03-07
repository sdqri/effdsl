package matchallquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
	"github.com/sdqri/effdsl/v2/utils"
)

type MatchAllQueryS struct {
	Boost *float64 `json:"boost,omitempty"` // (Optional, float) Floating point number used to decrease or increase the relevance scores of the query.
}

type matchAllQueryJsonS struct {
	MatchAllQuery matchAllQueryFieldParameters `json:"match_all"`
}

func (mq MatchAllQueryS) QueryInfo() string {
	return "Match all query"
}

func (mq MatchAllQueryS) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		matchAllQueryJsonS{
			MatchAllQuery: matchAllQueryFieldParameters{Boost: mq.Boost},
		},
	)
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html
type matchAllQueryFieldParameters struct {
	/*
		(Optional, float) Floating point number used to decrease or increase the relevance scores of the query. Defaults to 1.0.
		Boost values are relative to the default value of 1.0. A boost value between 0 and 1.0 decreases the relevance score. A value greater than 1.0 increases the relevance score.
	*/
	Boost *float64 `json:"boost,omitempty"`
}

type MatchAllQueryFieldParameter func(params *matchAllQueryFieldParameters)

func WithBoost(boost float64) MatchAllQueryFieldParameter {
	return func(params *matchAllQueryFieldParameters) {
		params.Boost = &boost
	}
}

// The most simple query, which matches all documents, giving them all a _score of 1.0.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html
func MatchAllQuery(params ...MatchAllQueryFieldParameter) effdsl.QueryResult {
	matchAllQuery := MatchAllQueryS{}

	var parameters matchAllQueryFieldParameters
	for _, prm := range params {
		prm(&parameters)
	}

	matchAllQuery, err := utils.CastStruct[matchAllQueryFieldParameters, MatchAllQueryS](parameters)

	return effdsl.QueryResult{
		Ok:  matchAllQuery,
		Err: err,
	}
}
