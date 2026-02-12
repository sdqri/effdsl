package significanttext_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	significanttext "github.com/sdqri/effdsl/v2/aggregations/bucket/significanttext"
)

func TestSignificantTextAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "significant_text": {
            "field": "description"
        }
    }`

	res := significanttext.SignificantText("significant_description", "description")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
