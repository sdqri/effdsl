package sum_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	sum "github.com/sdqri/effdsl/v2/aggregations/metrics/sum"
)

func TestSumAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
		"sum": {
			"field": "price"
		}
	}`

	res := sum.Sum("sum_price", "price")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestSumAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
		"sum": {
			"field": "price",
			"missing": 100
		}
	}`

	res := sum.Sum("hat_prices", "price", sum.WithMissing(100))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
