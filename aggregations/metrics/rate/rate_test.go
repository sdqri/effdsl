package rate_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	rate "github.com/sdqri/effdsl/v2/aggregations/metrics/rate"
)

func TestRateAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "rate": {
            "unit": "year"
        }
    }`

	res := rate.Rate("my_rate", rate.WithUnit("year"))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestRateAggregation_WithFieldAndUnit(t *testing.T) {
	expectedBody := `{
        "rate": {
            "field": "price",
            "unit": "day"
        }
    }`

	res := rate.Rate("avg_price", rate.WithField("price"), rate.WithUnit("day"))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestRateAggregation_WithValueCountMode(t *testing.T) {
	expectedBody := `{
        "rate": {
            "field": "price",
            "unit": "year",
            "mode": "value_count"
        }
    }`

	res := rate.Rate(
		"avg_number_of_sales_per_year",
		rate.WithField("price"),
		rate.WithUnit("year"),
		rate.WithMode("value_count"),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
