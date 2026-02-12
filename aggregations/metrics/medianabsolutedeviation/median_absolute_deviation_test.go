package medianabsolutedeviation_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	medianabsolutedeviation "github.com/sdqri/effdsl/v2/aggregations/metrics/medianabsolutedeviation"
)

func TestMedianAbsoluteDeviationAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "median_absolute_deviation": {
            "field": "rating"
        }
    }`

	res := medianabsolutedeviation.MedianAbsoluteDeviation("review_variability", "rating")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
