package maxbucket_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	maxbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/maxbucket"
)

func TestMaxBucketAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "max_bucket": {
            "buckets_path": "sales_per_month>sales"
        }
    }`

	res := maxbucket.MaxBucket("max_monthly_sales", "sales_per_month>sales")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
