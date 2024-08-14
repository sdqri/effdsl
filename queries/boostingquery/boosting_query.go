package boostingquery

import (
	"encoding/json"
	"fmt"

	"github.com/sdqri/effdsl"
)

type BoostingQueryS struct {
	Positive      effdsl.Query `json:"positive"`       // (Required, query object) Query you wish to run. Any returned documents must match this query.
	Negative      effdsl.Query `json:"negative"`       // (Required, query object) Query used to decrease the relevance score of matching documents.
	NegativeBoost float64      `json:"negative_boost"` // (Required, float) Floating point number between 0 and 1.0 used to decrease the relevance scores of documents matching the negative query.
}

func (bq BoostingQueryS) QueryInfo() string {
	return "Boosting query"
}

func (bq BoostingQueryS) MarshalJSON() ([]byte, error) {
	type BoostingQueryBase BoostingQueryS
	return json.Marshal(
		map[string]any{
			"boosting": (BoostingQueryBase)(bq),
		},
	)
}

// A query that returns documents matching a positive query while reducing the relevance score of documents that also match a negative query.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
func BoostingQuery(positive effdsl.QueryResult, negative effdsl.QueryResult, negativeBoost float64) effdsl.QueryResult {
	// Check for errors in the positive query
	if positive.Err != nil {
		return effdsl.QueryResult{
			Ok:  nil,
			Err: fmt.Errorf("error in positive query: %w", positive.Err),
		}
	}

	// Check for errors in the negative query
	if negative.Err != nil {
		return effdsl.QueryResult{
			Ok:  nil,
			Err: fmt.Errorf("error in negative query: %w", negative.Err),
		}
	}

	// Construct the BoostingQueryS struct
	boostingQuery := BoostingQueryS{
		Positive:      positive.Ok,
		Negative:      negative.Ok,
		NegativeBoost: negativeBoost,
	}

	// Return the constructed query with no error
	return effdsl.QueryResult{
		Ok:  boostingQuery,
		Err: nil,
	}
}
