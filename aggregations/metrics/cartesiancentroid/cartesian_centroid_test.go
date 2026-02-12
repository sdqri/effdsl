package cartesiancentroid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	cartesiancentroid "github.com/sdqri/effdsl/v2/aggregations/metrics/cartesiancentroid"
)

func TestCartesianCentroidAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "cartesian_centroid": {
            "field": "location"
        }
    }`

	res := cartesiancentroid.CartesianCentroid("centroid", "location")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
