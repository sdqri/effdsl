package datehistogram_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	datehistogram "github.com/sdqri/effdsl/v2/aggregations/bucket/datehistogram"
)

func TestDateHistogramAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "date_histogram": {
            "field": "date",
            "calendar_interval": "month"
        }
    }`

	res := datehistogram.DateHistogram("sales_over_time", "date", datehistogram.WithCalendarInterval("month"))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
