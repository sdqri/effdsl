package normalize_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	normalize "github.com/sdqri/effdsl/v2/aggregations/pipeline/normalize"
)

func TestNormalizeAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "normalize": {
            "buckets_path": "sales",
            "method": "percent_of_sum",
            "format": "00.00%"
        }
    }`

	res := normalize.Normalize(
		"percent_of_total_sales",
		"sales",
		"percent_of_sum",
		normalize.WithFormat("00.00%"),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
