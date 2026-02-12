package randomsampler_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	randomsampler "github.com/sdqri/effdsl/v2/aggregations/bucket/randomsampler"
)

func TestRandomSamplerAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "random_sampler": {
            "probability": 0.1
        }
    }`

	res := randomsampler.RandomSampler("sample", 0.1)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
