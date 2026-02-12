package valuecount_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2/aggregations"
	valuecount "github.com/sdqri/effdsl/v2/aggregations/metrics/valuecount"
)

func TestValueCountAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "value_count": {
            "field": "type"
        }
    }`

	res := valuecount.ValueCount("types_count", "type")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestValueCountAggregation_WithScript(t *testing.T) {
	expectedBody := `{
        "value_count": {
            "script": {
                "source": "doc['grade'].value"
            }
        }
    }`

	res := valuecount.ValueCount(
		"types_count",
		"",
		valuecount.WithScript(aggregations.Script{Source: "doc['grade'].value"}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
