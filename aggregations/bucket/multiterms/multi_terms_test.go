package multiterms_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	multiterms "github.com/sdqri/effdsl/v2/aggregations/bucket/multiterms"
)

func TestMultiTermsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "multi_terms": {
            "terms": [
                {"field": "category"},
                {"field": "brand"}
            ]
        }
    }`

	terms := []any{
		map[string]any{"field": "category"},
		map[string]any{"field": "brand"},
	}

	res := multiterms.MultiTerms("category_brand", terms)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
