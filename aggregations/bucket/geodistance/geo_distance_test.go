package geodistance_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	geodistance "github.com/sdqri/effdsl/v2/aggregations/bucket/geodistance"
)

func TestGeoDistanceAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "geo_distance": {
            "field": "location",
            "origin": "52.3760, 4.894",
            "ranges": [
                {"to": 100},
                {"from": 100, "to": 300},
                {"from": 300}
            ]
        }
    }`

	to100 := 100.0
	from100 := 100.0
	to300 := 300.0
	from300 := 300.0

	ranges := []geodistance.GeoDistanceRange{
		{To: &to100},
		{From: &from100, To: &to300},
		{From: &from300},
	}

	res := geodistance.GeoDistance("rings", "location", "52.3760, 4.894", ranges)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
