package significantterms_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	significantterms "github.com/sdqri/effdsl/v2/aggregations/bucket/significantterms"
)

func TestSignificantTermsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "significant_terms": {
            "field": "category"
        }
    }`

	res := significantterms.SignificantTerms("significant_categories", "category")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
