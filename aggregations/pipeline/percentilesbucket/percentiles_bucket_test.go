package percentilesbucket_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	percentilesbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/percentilesbucket"
)

func TestPercentilesBucketAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "percentiles_bucket": {
            "buckets_path": "sales_per_month>sales",
            "percents": [
                25,
                50,
                75
            ]
        }
    }`

	res := percentilesbucket.PercentilesBucket(
		"percentiles_monthly_sales",
		"sales_per_month>sales",
		percentilesbucket.WithPercents([]float64{25, 50, 75}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
