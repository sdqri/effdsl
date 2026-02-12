package geotilegrid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	geotilegrid "github.com/sdqri/effdsl/v2/aggregations/bucket/geotilegrid"
)

func TestGeoTileGridAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "geotile_grid": {
            "field": "location",
            "precision": 7
        }
    }`

	res := geotilegrid.GeoTileGrid("tiles", "location", 7)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
