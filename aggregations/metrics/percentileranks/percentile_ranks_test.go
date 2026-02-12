package percentileranks_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	percentileranks "github.com/sdqri/effdsl/v2/aggregations/metrics/percentileranks"
)

func TestPercentileRanksAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "percentile_ranks": {
            "field": "load_time",
            "values": [500, 600]
        }
    }`

	res := percentileranks.PercentileRanks("load_time_ranks", "load_time", []float64{500, 600})

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPercentileRanksAggregation_WithKeyedFalse(t *testing.T) {
	expectedBody := `{
        "percentile_ranks": {
            "field": "load_time",
            "values": [500, 600],
            "keyed": false
        }
    }`

	res := percentileranks.PercentileRanks(
		"load_time_ranks",
		"load_time",
		[]float64{500, 600},
		percentileranks.WithKeyed(false),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPercentileRanksAggregation_WithHDR(t *testing.T) {
	expectedBody := `{
        "percentile_ranks": {
            "field": "load_time",
            "values": [500, 600],
            "hdr": {
                "number_of_significant_value_digits": 3
            }
        }
    }`

	res := percentileranks.PercentileRanks(
		"load_time_ranks",
		"load_time",
		[]float64{500, 600},
		percentileranks.WithHDR(3),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPercentileRanksAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
        "percentile_ranks": {
            "field": "load_time",
            "values": [500, 600],
            "missing": 10
        }
    }`

	res := percentileranks.PercentileRanks(
		"load_time_ranks",
		"load_time",
		[]float64{500, 600},
		percentileranks.WithMissing(10),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
