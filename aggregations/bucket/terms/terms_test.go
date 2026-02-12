package terms_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	terms "github.com/sdqri/effdsl/v2/aggregations/bucket/terms"
)

func TestTermsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "terms": {
            "field": "category"
        }
    }`

	res := terms.Terms("categories", "category")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
