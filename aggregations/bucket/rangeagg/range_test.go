package rangeagg_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	rangeagg "github.com/sdqri/effdsl/v2/aggregations/bucket/rangeagg"
)

func TestRangeAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "range": {
            "field": "price",
            "ranges": [
                {"to": 50},
                {"from": 50, "to": 100},
                {"from": 100}
            ]
        }
    }`

	ranges := []rangeagg.RangeItem{
		{To: 50},
		{From: 50, To: 100},
		{From: 100},
	}

	res := rangeagg.Range("price_ranges", "price", ranges)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
