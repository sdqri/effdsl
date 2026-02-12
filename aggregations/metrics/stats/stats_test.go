package stats_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	stats "github.com/sdqri/effdsl/v2/aggregations/metrics/stats"
)

func TestStatsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "stats": {
            "field": "grade"
        }
    }`

	res := stats.Stats("grades_stats", "grade")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
