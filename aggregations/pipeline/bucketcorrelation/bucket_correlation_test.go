package bucketcorrelation_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	bucketcorrelation "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketcorrelation"
)

func TestBucketCorrelationAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "bucket_correlation": {
            "buckets_path": "latency_ranges>_count",
            "function": {
                "count_correlation": {
                    "indicator": {
                        "expectations": [
                            0,
                            52.5,
                            165,
                            335,
                            555
                        ],
                        "doc_count": 200
                    }
                }
            }
        }
    }`

	function := bucketcorrelation.CountCorrelation(
		[]float64{0, 52.5, 165, 335, 555},
		200,
		nil,
	)

	res := bucketcorrelation.BucketCorrelation(
		"bucket_correlation",
		"latency_ranges>_count",
		function,
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
