package iprange_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	iprange "github.com/sdqri/effdsl/v2/aggregations/bucket/iprange"
)

func TestIPRangeAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "ip_range": {
            "field": "ip",
            "ranges": [
                {"mask": "10.0.0.0/8"},
                {"from": "10.0.0.0", "to": "10.0.0.255"}
            ]
        }
    }`

	ranges := []iprange.IPRangeItem{
		{Mask: "10.0.0.0/8"},
		{From: "10.0.0.0", To: "10.0.0.255"},
	}

	res := iprange.IPRange("ip_ranges", "ip", ranges)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
