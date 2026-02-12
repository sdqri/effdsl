package cumulativecardinality_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	cumulativecardinality "github.com/sdqri/effdsl/v2/aggregations/pipeline/cumulativecardinality"
)

func TestCumulativeCardinalityAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "cumulative_cardinality": {
            "buckets_path": "distinct_users"
        }
    }`

	res := cumulativecardinality.CumulativeCardinality("total_new_users", "distinct_users")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
