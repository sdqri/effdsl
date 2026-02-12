package changepoint_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	changepoint "github.com/sdqri/effdsl/v2/aggregations/pipeline/changepoint"
)

func TestChangePointAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "change_point": {
            "buckets_path": "date>avg"
        }
    }`

	res := changepoint.ChangePoint("change_points_avg", "date>avg")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
