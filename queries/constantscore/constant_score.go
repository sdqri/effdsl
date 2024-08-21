package constantscore

import (
	"encoding/json"
	"fmt"

	"github.com/sdqri/effdsl/v2"
)

type ConstantScoreQueryS struct {
	Filter effdsl.Query `json:"filter"`          // (Required, query object) Filter query you wish to run. Any returned documents must match this query.
	Boost  float64      `json:"boost,omitempty"` // (Optional, float) Floating point number used as the constant relevance score for every document matching the filter query. Defaults to 1.0.
}

func (csq ConstantScoreQueryS) QueryInfo() string {
	return "Constant score query"
}

func (csq ConstantScoreQueryS) MarshalJSON() ([]byte, error) {
	type ConstantScoreQueryBase ConstantScoreQueryS
	return json.Marshal(
		map[string]any{
			"constant_score": (ConstantScoreQueryBase)(csq),
		},
	)
}

// ConstantScoreQuery wraps a filter query and returns every matching document with a relevance score equal to the boost parameter value.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-constant-score-query.html
func ConstantScoreQuery(filter effdsl.QueryResult, boost float64) effdsl.QueryResult {
	if filter.Err != nil {
		return effdsl.QueryResult{
			Ok:  nil,
			Err: fmt.Errorf("error in filter query: %w", filter.Err),
		}
	}

	constantScoreQuery := ConstantScoreQueryS{
		Filter: filter.Ok,
		Boost:  boost,
	}

	return effdsl.QueryResult{
		Ok:  constantScoreQuery,
		Err: nil,
	}
}
