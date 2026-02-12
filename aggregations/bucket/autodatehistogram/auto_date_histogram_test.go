package autodatehistogram_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	autodatehistogram "github.com/sdqri/effdsl/v2/aggregations/bucket/autodatehistogram"
)

func TestAutoDateHistogramAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "auto_date_histogram": {
            "field": "timestamp",
            "buckets": 10
        }
    }`

	res := autodatehistogram.AutoDateHistogram("sales_over_time", "timestamp", 10)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
