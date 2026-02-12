package geobounds_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	geobounds "github.com/sdqri/effdsl/v2/aggregations/metrics/geobounds"
)

func TestGeoBoundsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "geo_bounds": {
            "field": "location"
        }
    }`

	res := geobounds.GeoBounds("viewport", "location")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
