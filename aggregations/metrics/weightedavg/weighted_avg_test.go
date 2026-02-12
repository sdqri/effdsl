package weightedavg_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	weightedavg "github.com/sdqri/effdsl/v2/aggregations/metrics/weightedavg"
)

func TestWeightedAvgAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "weighted_avg": {
            "value": {"field": "grade"},
            "weight": {"field": "weight"}
        }
    }`

	res := weightedavg.WeightedAvg("weighted_grade", "grade", "weight")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
