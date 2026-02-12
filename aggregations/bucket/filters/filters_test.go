package filters_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	filters "github.com/sdqri/effdsl/v2/aggregations/bucket/filters"
)

func TestFiltersAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "filters": {
            "filters": {
                "errors": {"term": {"status": "error"}},
                "warnings": {"term": {"status": "warning"}}
            }
        }
    }`

	filtersMap := map[string]any{
		"errors":   map[string]any{"term": map[string]any{"status": "error"}},
		"warnings": map[string]any{"term": map[string]any{"status": "warning"}},
	}

	res := filters.Filters("statuses", filtersMap)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
