package diversifiedsampler_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	diversifiedsampler "github.com/sdqri/effdsl/v2/aggregations/bucket/diversifiedsampler"
)

func TestDiversifiedSamplerAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "diversified_sampler": {
            "field": "user_id",
            "shard_size": 200
        }
    }`

	res := diversifiedsampler.DiversifiedSampler(
		"sample",
		diversifiedsampler.WithField("user_id"),
		diversifiedsampler.WithShardSize(200),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
