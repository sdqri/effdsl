package geoline_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	geoline "github.com/sdqri/effdsl/v2/aggregations/metrics/geoline"
)

func TestGeoLineAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "geo_line": {
            "point": {
                "field": "my_location"
            }
        }
    }`

	res := geoline.GeoLine("line", "my_location")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
