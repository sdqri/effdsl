package rareterms_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	rareterms "github.com/sdqri/effdsl/v2/aggregations/bucket/rareterms"
)

func TestRareTermsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "rare_terms": {
            "field": "category"
        }
    }`

	res := rareterms.RareTerms("rare_categories", "category")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
