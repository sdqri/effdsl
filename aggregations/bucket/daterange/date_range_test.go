package daterange_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	daterange "github.com/sdqri/effdsl/v2/aggregations/bucket/daterange"
)

func TestDateRangeAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "date_range": {
            "field": "date",
            "ranges": [
                {"to": "now-1M/M"},
                {"from": "now-1M/M"}
            ]
        }
    }`

	ranges := []daterange.DateRangeItem{
		{To: "now-1M/M"},
		{From: "now-1M/M"},
	}

	res := daterange.DateRange("recent", "date", ranges)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
