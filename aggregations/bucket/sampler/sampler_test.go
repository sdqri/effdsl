package sampler_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	sampler "github.com/sdqri/effdsl/v2/aggregations/bucket/sampler"
)

func TestSamplerAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "sampler": {
            "shard_size": 200
        }
    }`

	res := sampler.Sampler("sample", sampler.WithShardSize(200))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
