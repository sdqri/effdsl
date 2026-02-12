package scriptedmetric_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2/aggregations"
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

func TestScriptedMetricAggregation_WithScripts(t *testing.T) {
	expectedBody := `{
        "scripted_metric": {
            "init_script": {
                "source": "state.transactions = []"
            },
            "map_script": {
                "source": "state.transactions.add(doc.type.value == 'sale' ? doc.amount.value : -1 * doc.amount.value)"
            },
            "combine_script": {
                "source": "double profit = 0; for (t in state.transactions) { profit += t } return profit"
            },
            "reduce_script": {
                "source": "double profit = 0; for (a in states) { profit += a } return profit"
            }
        }
    }`

	res := scriptedmetric.ScriptedMetric(
		"profit",
		scriptedmetric.WithInitScript(aggregations.Script{Source: "state.transactions = []"}),
		scriptedmetric.WithMapScript(aggregations.Script{Source: "state.transactions.add(doc.type.value == 'sale' ? doc.amount.value : -1 * doc.amount.value)"}),
		scriptedmetric.WithCombineScript(aggregations.Script{Source: "double profit = 0; for (t in state.transactions) { profit += t } return profit"}),
		scriptedmetric.WithReduceScript(aggregations.Script{Source: "double profit = 0; for (a in states) { profit += a } return profit"}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestScriptedMetricAggregation_WithStoredScripts(t *testing.T) {
	res := scriptedmetric.ScriptedMetric(
		"profit",
		scriptedmetric.WithParams(map[string]any{"field": "amount"}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.Contains(t, string(jsonBody), "\"params\":{\"field\":\"amount\"}")
}
