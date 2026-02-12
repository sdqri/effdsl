package extendedstats_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	extendedstats "github.com/sdqri/effdsl/v2/aggregations/metrics/extendedstats"
)

func TestExtendedStatsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "extended_stats": {
            "field": "grade"
        }
    }`

	res := extendedstats.ExtendedStats("grades_stats", "grade")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
