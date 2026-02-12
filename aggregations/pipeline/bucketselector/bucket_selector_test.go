package bucketselector_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2/aggregations"
	bucketselector "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketselector"
)

func TestBucketSelectorAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "bucket_selector": {
            "buckets_path": {
                "total_cost": "sales_per_month>cost",
                "total_sales": "sales_per_month>sales"
            },
            "script": {
                "source": "params.total_sales > params.total_cost"
            }
        }
    }`

	res := bucketselector.BucketSelector(
		"only_profitable",
		map[string]string{
			"total_sales": "sales_per_month>sales",
			"total_cost":  "sales_per_month>cost",
		},
		aggregations.Script{Source: "params.total_sales > params.total_cost"},
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
