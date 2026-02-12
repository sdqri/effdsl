package variablewidthhistogram_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	variablewidthhistogram "github.com/sdqri/effdsl/v2/aggregations/bucket/variablewidthhistogram"
)

func TestVariableWidthHistogramAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "variable_width_histogram": {
            "field": "price",
            "buckets": 5
        }
    }`

	res := variablewidthhistogram.VariableWidthHistogram("price_histogram", "price", 5)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
