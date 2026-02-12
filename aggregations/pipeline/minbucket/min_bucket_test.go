package minbucket_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	minbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/minbucket"
)

func TestMinBucketAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "min_bucket": {
            "buckets_path": "sales_per_month>sales"
        }
    }`

	res := minbucket.MinBucket("min_monthly_sales", "sales_per_month>sales")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
