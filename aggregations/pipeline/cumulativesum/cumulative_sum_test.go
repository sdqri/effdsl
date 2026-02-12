package cumulativesum_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	cumulativesum "github.com/sdqri/effdsl/v2/aggregations/pipeline/cumulativesum"
)

func TestCumulativeSumAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "cumulative_sum": {
            "buckets_path": "sales"
        }
    }`

	res := cumulativesum.CumulativeSum("cumulative_sales", "sales")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
