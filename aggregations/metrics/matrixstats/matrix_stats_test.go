package matrixstats_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	matrixstats "github.com/sdqri/effdsl/v2/aggregations/metrics/matrixstats"
)

func TestMatrixStatsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "matrix_stats": {
            "fields": ["poverty", "income"]
        }
    }`

	res := matrixstats.MatrixStats("statistics", []string{"poverty", "income"})

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestMatrixStatsAggregation_WithMissingAndMode(t *testing.T) {
	expectedBody := `{
        "matrix_stats": {
            "fields": ["poverty", "income"],
            "missing": {"poverty": 0},
            "mode": "avg"
        }
    }`

	res := matrixstats.MatrixStats(
		"statistics",
		[]string{"poverty", "income"},
		matrixstats.WithMissing(map[string]any{"poverty": 0}),
		matrixstats.WithMode("avg"),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
