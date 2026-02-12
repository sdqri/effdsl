package tophits_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	tophits "github.com/sdqri/effdsl/v2/aggregations/metrics/tophits"
)

func TestTopHitsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "top_hits": {
            "size": 1
        }
    }`

	res := tophits.TopHits("top_sales_hits", tophits.WithSize(1))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestTopHitsAggregation_AllOptions(t *testing.T) {
	expectedBody := `{
        "top_hits": {
            "from": 5,
            "size": 3,
            "sort": [
                {"date": {"order": "desc"}}
            ],
            "_source": {
                "includes": ["date", "price"]
            },
            "stored_fields": ["title", "date"],
            "docvalue_fields": ["date"],
            "script_fields": {
                "discounted_price": {
                    "script": {
                        "lang": "painless",
                        "source": "doc['price'].value * 0.8"
                    }
                }
            },
            "highlight": {
                "fields": {
                    "body": {}
                }
            },
            "explain": true,
            "track_scores": true,
            "version": true,
            "seq_no_primary_term": true,
            "fields": ["title", "date"]
        }
    }`

	res := tophits.TopHits(
		"top_sales_hits",
		tophits.WithFrom(5),
		tophits.WithSize(3),
		tophits.WithSort([]any{map[string]any{"date": map[string]any{"order": "desc"}}}),
		tophits.WithSource(map[string]any{"includes": []string{"date", "price"}}),
		tophits.WithStoredFields([]string{"title", "date"}),
		tophits.WithDocvalueFields([]any{"date"}),
		tophits.WithScriptFields(map[string]any{
			"discounted_price": map[string]any{
				"script": map[string]any{
					"lang":   "painless",
					"source": "doc['price'].value * 0.8",
				},
			},
		}),
		tophits.WithHighlight(map[string]any{"fields": map[string]any{"body": map[string]any{}}}),
		tophits.WithExplain(true),
		tophits.WithTrackScores(true),
		tophits.WithVersion(true),
		tophits.WithSeqNoPrimaryTerm(true),
		tophits.WithFields([]any{"title", "date"}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
