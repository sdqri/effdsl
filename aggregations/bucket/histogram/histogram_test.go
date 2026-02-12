package histogram_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	histogram "github.com/sdqri/effdsl/v2/aggregations/bucket/histogram"
)

func TestHistogramAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "histogram": {
            "field": "price",
            "interval": 50
        }
    }`

	res := histogram.Histogram("prices", "price", 50)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
