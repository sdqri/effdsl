package movingfn_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	movingfn "github.com/sdqri/effdsl/v2/aggregations/pipeline/movingfn"
)

func TestMovingFnAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "moving_fn": {
            "buckets_path": "the_sum",
            "window": 10,
            "script": "MovingFunctions.unweightedAvg(values)"
        }
    }`

	res := movingfn.MovingFn(
		"the_movfn",
		"the_sum",
		10,
		"MovingFunctions.unweightedAvg(values)",
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
