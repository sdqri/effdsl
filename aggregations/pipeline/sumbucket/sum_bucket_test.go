package sumbucket_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	sumbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/sumbucket"
)

func TestSumBucketAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "sum_bucket": {
            "buckets_path": "sales_per_month>sales"
        }
    }`

	res := sumbucket.SumBucket("sum_monthly_sales", "sales_per_month>sales")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
