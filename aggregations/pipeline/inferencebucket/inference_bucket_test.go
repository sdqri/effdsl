package inferencebucket_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	inferencebucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/inferencebucket"
)

func TestInferenceBucketAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "inference": {
            "model_id": "malicious_clients_model",
            "buckets_path": {
                "bytes_sum": "bytes_sum",
                "response_count": "responses_total",
                "url_dc": "url_dc"
            }
        }
    }`

	res := inferencebucket.InferenceBucket(
		"malicious_client_ip",
		"malicious_clients_model",
		map[string]string{
			"response_count": "responses_total",
			"url_dc":         "url_dc",
			"bytes_sum":      "bytes_sum",
		},
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
