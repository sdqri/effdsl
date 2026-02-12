package missing_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	missing "github.com/sdqri/effdsl/v2/aggregations/bucket/missing"
)

func TestMissingAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "missing": {
            "field": "price"
        }
    }`

	res := missing.Missing("missing_price", "price")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
