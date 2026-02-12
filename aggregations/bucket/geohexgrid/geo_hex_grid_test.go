package geohexgrid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	geohexgrid "github.com/sdqri/effdsl/v2/aggregations/bucket/geohexgrid"
)

func TestGeoHexGridAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "geohex_grid": {
            "field": "location",
            "precision": 5
        }
    }`

	res := geohexgrid.GeoHexGrid("hexes", "location", 5)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
