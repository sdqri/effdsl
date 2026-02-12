package extendedstatsbucket_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	extendedstatsbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/extendedstatsbucket"
)

func TestExtendedStatsBucketAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "extended_stats_bucket": {
            "buckets_path": "sales_per_month>sales"
        }
    }`

	res := extendedstatsbucket.ExtendedStatsBucket("stats_monthly_sales", "sales_per_month>sales")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
