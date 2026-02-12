package scriptedmetric_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	scriptedmetric "github.com/sdqri/effdsl/v2/aggregations/metrics/scriptedmetric"
)

func TestScriptedMetricAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "scripted_metric": {}
    }`

	res := scriptedmetric.ScriptedMetric("profit")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
