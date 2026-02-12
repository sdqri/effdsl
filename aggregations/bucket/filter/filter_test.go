package filter_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	filter "github.com/sdqri/effdsl/v2/aggregations/bucket/filter"
)

func TestFilterAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "filter": {
            "filter": {"term": {"status": "error"}}
        }
    }`

	res := filter.Filter("errors", map[string]any{"term": map[string]any{"status": "error"}})

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
