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
