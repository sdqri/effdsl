package min_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	min "github.com/sdqri/effdsl/v2/aggregations/metrics/min"
)

func TestMinAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "min": {
            "field": "price"
        }
    }`

	res := min.Min("min_price", "price")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestMinAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
        "min": {
            "field": "grade",
            "missing": 10
        }
    }`

	res := min.Min("grade_min", "grade", min.WithMissing(10))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
