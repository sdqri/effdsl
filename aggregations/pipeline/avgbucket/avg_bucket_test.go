package avgbucket_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	avgbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/avgbucket"
)

func TestAvgBucketAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "avg_bucket": {
            "buckets_path": "sales_per_month>sales"
        }
    }`

	res := avgbucket.AvgBucket("avg_monthly_sales", "sales_per_month>sales")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
