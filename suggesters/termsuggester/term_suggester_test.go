package termsuggester_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	ts "github.com/sdqri/effdsl/v2/suggesters/termsuggester"
)

func TestTermSuggester_WithNoOptions(t *testing.T) {
	expectedBody := `{
		"text": "tring out Elasticsearch",
		"term": {
			"field": "message"
		}
	}`

	termSuggesterResult := ts.TermSuggester(
		"my-suggestion-1",
		"tring out Elasticsearch",
		"message", // required base field
	)

	err := termSuggesterResult.Err
	suggestQuery := termSuggesterResult.Ok
	jsonBody, err := json.Marshal(suggestQuery)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestTermSuggesterAllOptions(t *testing.T) {
	expectedBody := `{
		"text": "tring out Elasticsearch",
		"term": {
			"field": "message",
			"analyzer": "test",
			"size": 1,
			"sort": "score",
			"suggest_mode": "always"
		}
	}`

	termSuggesterResult := ts.TermSuggester(
		"my-suggestion-2",
		"tring out Elasticsearch",
		"message",
		ts.WithAnalyzer("test"),
		ts.WithSize(1),
		ts.WithSort(ts.ByScore),
		ts.WithMode(ts.Always),
	)

	err := termSuggesterResult.Err
	matchQuery := termSuggesterResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
