package bucketcountkstest_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	bucketcountkstest "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketcountkstest"
)

func TestBucketCountKSTestAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "bucket_count_ks_test": {
            "buckets_path": "latency_ranges>_count",
            "alternative": [
                "less",
                "greater",
                "two_sided"
            ],
            "sampling_method": "upper_tail"
        }
    }`

	res := bucketcountkstest.BucketCountKSTest(
		"ks_test",
		"latency_ranges>_count",
		bucketcountkstest.WithAlternatives([]string{"less", "greater", "two_sided"}),
		bucketcountkstest.WithSamplingMethod("upper_tail"),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
