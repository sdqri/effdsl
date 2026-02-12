package geohashgrid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	geohashgrid "github.com/sdqri/effdsl/v2/aggregations/bucket/geohashgrid"
)

func TestGeoHashGridAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "geohash_grid": {
            "field": "location",
            "precision": 5
        }
    }`

	res := geohashgrid.GeoHashGrid("grid", "location", 5)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
