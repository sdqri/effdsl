package weightedavg_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2/aggregations"
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

func TestWeightedAvgAggregation_WithScriptValue(t *testing.T) {
	expectedBody := `{
        "weighted_avg": {
            "value": {
                "script": {
                    "source": "doc.grade.value + 1"
                }
            },
            "weight": {
                "field": "weight.combined"
            }
        }
    }`

	res := weightedavg.WeightedAvg(
		"weighted_grade",
		"",
		"weight.combined",
		weightedavg.WithValueScript(aggregations.Script{Source: "doc.grade.value + 1"}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestWeightedAvgAggregation_WithMissingValues(t *testing.T) {
	expectedBody := `{
        "weighted_avg": {
            "value": {
                "field": "grade",
                "missing": 2
            },
            "weight": {
                "field": "weight",
                "missing": 3
            }
        }
    }`

	res := weightedavg.WeightedAvg(
		"weighted_grade",
		"grade",
		"weight",
		weightedavg.WithValueMissing(2),
		weightedavg.WithWeightMissing(3),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
