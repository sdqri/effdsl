package percentiles_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	percentiles "github.com/sdqri/effdsl/v2/aggregations/metrics/percentiles"
)

func TestPercentilesAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "percentiles": {
            "field": "load_time",
            "percents": [95, 99, 99.9]
        }
    }`

	res := percentiles.Percentiles("load_time_outlier", "load_time", []float64{95, 99, 99.9})

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
