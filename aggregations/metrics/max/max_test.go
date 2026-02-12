package max_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	max "github.com/sdqri/effdsl/v2/aggregations/metrics/max"
)

func TestMaxAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "max": {
            "field": "price"
        }
    }`

	res := max.Max("max_price", "price")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestMaxAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
        "max": {
            "field": "grade",
            "missing": 10
        }
    }`

	res := max.Max("grade_max", "grade", max.WithMissing(10))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
