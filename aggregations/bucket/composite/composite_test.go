package composite_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	composite "github.com/sdqri/effdsl/v2/aggregations/bucket/composite"
)

func TestCompositeAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "composite": {
            "sources": [
                {"date": {"date_histogram": {"field": "date", "calendar_interval": "month"}}},
                {"type": {"terms": {"field": "type"}}}
            ]
        }
    }`

	sources := []any{
		map[string]any{"date": map[string]any{"date_histogram": map[string]any{"field": "date", "calendar_interval": "month"}}},
		map[string]any{"type": map[string]any{"terms": map[string]any{"field": "type"}}},
	}

	res := composite.Composite("buckets", sources)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
