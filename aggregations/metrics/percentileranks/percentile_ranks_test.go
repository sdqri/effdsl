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
