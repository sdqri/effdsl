package movingpercentiles_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	movingpercentiles "github.com/sdqri/effdsl/v2/aggregations/pipeline/movingpercentiles"
)

func TestMovingPercentilesAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "moving_percentiles": {
            "buckets_path": "the_percentile",
            "window": 10
        }
    }`

	res := movingpercentiles.MovingPercentiles(
		"the_movperc",
		"the_percentile",
		10,
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
