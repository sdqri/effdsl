package bucketscript_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2/aggregations"
	bucketscript "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketscript"
)

func TestBucketScriptAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "bucket_script": {
            "buckets_path": {
                "total_cost": "sales_per_month>cost",
                "total_sales": "sales_per_month>sales"
            },
            "script": {
                "source": "params.total_sales - params.total_cost"
            }
        }
    }`

	res := bucketscript.BucketScript(
		"sales_profit",
		map[string]string{
			"total_sales": "sales_per_month>sales",
			"total_cost":  "sales_per_month>cost",
		},
		aggregations.Script{Source: "params.total_sales - params.total_cost"},
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
