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

func TestPercentilesAggregation_WithKeyedFalse(t *testing.T) {
	expectedBody := `{
        "percentiles": {
            "field": "load_time",
            "keyed": false
        }
    }`

	res := percentiles.Percentiles(
		"load_time_outlier",
		"load_time",
		nil,
		percentiles.WithKeyed(false),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPercentilesAggregation_WithTDigestCompression(t *testing.T) {
	expectedBody := `{
        "percentiles": {
            "field": "load_time",
            "tdigest": {
                "compression": 200
            }
        }
    }`

	res := percentiles.Percentiles(
		"load_time_outlier",
		"load_time",
		nil,
		percentiles.WithTDigest(map[string]any{"compression": 200}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPercentilesAggregation_WithTDigestExecutionHint(t *testing.T) {
	expectedBody := `{
        "percentiles": {
            "field": "load_time",
            "tdigest": {
                "execution_hint": "high_accuracy"
            }
        }
    }`

	res := percentiles.Percentiles(
		"load_time_outlier",
		"load_time",
		nil,
		percentiles.WithTDigest(map[string]any{"execution_hint": "high_accuracy"}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPercentilesAggregation_WithHDR(t *testing.T) {
	expectedBody := `{
        "percentiles": {
            "field": "load_time",
            "percents": [95, 99, 99.9],
            "hdr": {
                "number_of_significant_value_digits": 3
            }
        }
    }`

	res := percentiles.Percentiles(
		"load_time_outlier",
		"load_time",
		[]float64{95, 99, 99.9},
		percentiles.WithHDR(3),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPercentilesAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
        "percentiles": {
            "field": "grade",
            "missing": 10
        }
    }`

	res := percentiles.Percentiles(
		"grade_percentiles",
		"grade",
		nil,
		percentiles.WithMissing(10),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
