package statsbucket_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	statsbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/statsbucket"
)

func TestStatsBucketAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "stats_bucket": {
            "buckets_path": "sales_per_month>sales"
        }
    }`

	res := statsbucket.StatsBucket("stats_monthly_sales", "sales_per_month>sales")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
