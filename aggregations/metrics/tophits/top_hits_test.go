package tophits_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	tophits "github.com/sdqri/effdsl/v2/aggregations/metrics/tophits"
)

func TestTopHitsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "top_hits": {
            "size": 1
        }
    }`

	res := tophits.TopHits("top_sales_hits", tophits.WithSize(1))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
