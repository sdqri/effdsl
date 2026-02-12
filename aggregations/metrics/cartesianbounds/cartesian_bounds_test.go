package cartesianbounds_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	cartesianbounds "github.com/sdqri/effdsl/v2/aggregations/metrics/cartesianbounds"
)

func TestCartesianBoundsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "cartesian_bounds": {
            "field": "location"
        }
    }`

	res := cartesianbounds.CartesianBounds("viewport", "location")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
