package stringstats_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	stringstats "github.com/sdqri/effdsl/v2/aggregations/metrics/stringstats"
)

func TestStringStatsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "string_stats": {
            "field": "message.keyword"
        }
    }`

	res := stringstats.StringStats("message_stats", "message.keyword")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestStringStatsAggregation_WithDistribution(t *testing.T) {
	expectedBody := `{
        "string_stats": {
            "field": "message.keyword",
            "show_distribution": true
        }
    }`

	res := stringstats.StringStats(
		"message_stats",
		"message.keyword",
		stringstats.WithShowDistribution(true),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestStringStatsAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
        "string_stats": {
            "field": "message.keyword",
            "missing": "[empty message]"
        }
    }`

	res := stringstats.StringStats(
		"message_stats",
		"message.keyword",
		stringstats.WithMissing("[empty message]"),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
