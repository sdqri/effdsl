package topmetrics_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	topmetrics "github.com/sdqri/effdsl/v2/aggregations/metrics/topmetrics"
)

func TestTopMetricsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "top_metrics": {
            "metrics": {"field": "m"},
            "sort": {"s": "desc"}
        }
    }`

	res := topmetrics.TopMetrics("tm", map[string]any{"field": "m"}, map[string]any{"s": "desc"})

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestTopMetricsAggregation_WithSizeAndMultipleMetrics(t *testing.T) {
	expectedBody := `{
        "top_metrics": {
            "metrics": [
                {"field": "m1"},
                {"field": "m2"}
            ],
            "sort": {"s": "desc"},
            "size": 3
        }
    }`

	res := topmetrics.TopMetrics(
		"tm",
		[]any{map[string]any{"field": "m1"}, map[string]any{"field": "m2"}},
		map[string]any{"s": "desc"},
		topmetrics.WithSize(3),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
