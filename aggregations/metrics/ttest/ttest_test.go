package ttest_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	ttest "github.com/sdqri/effdsl/v2/aggregations/metrics/ttest"
)

func TestTTestAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "t_test": {
            "a": {"field": "startup_time_before"},
            "b": {"field": "startup_time_after"}
        }
    }`

	res := ttest.TTest("startup_time_ttest", "startup_time_before", "startup_time_after")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
