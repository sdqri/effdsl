package geocentroid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	geocentroid "github.com/sdqri/effdsl/v2/aggregations/metrics/geocentroid"
)

func TestGeoCentroidAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "geo_centroid": {
            "field": "location"
        }
    }`

	res := geocentroid.GeoCentroid("centroid", "location")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
