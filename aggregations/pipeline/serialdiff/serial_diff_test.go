package serialdiff_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	serialdiff "github.com/sdqri/effdsl/v2/aggregations/pipeline/serialdiff"
)

func TestSerialDiffAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "serial_diff": {
            "buckets_path": "the_sum",
            "lag": 30
        }
    }`

	res := serialdiff.SerialDiff(
		"thirtieth_difference",
		"the_sum",
		serialdiff.WithLag(30),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
