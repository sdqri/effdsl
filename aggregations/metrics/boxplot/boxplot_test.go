package boxplot_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	boxplot "github.com/sdqri/effdsl/v2/aggregations/metrics/boxplot"
)

func TestBoxplotAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "boxplot": {
            "field": "load_time"
        }
    }`

	res := boxplot.Boxplot("load_time_boxplot", "load_time")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestBoxplotAggregation_WithCompression(t *testing.T) {
	expectedBody := `{
        "boxplot": {
            "field": "load_time",
            "compression": 200
        }
    }`

	res := boxplot.Boxplot("load_time_boxplot", "load_time", boxplot.WithCompression(200))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestBoxplotAggregation_WithExecutionHint(t *testing.T) {
	expectedBody := `{
        "boxplot": {
            "field": "load_time",
            "execution_hint": "high_accuracy"
        }
    }`

	res := boxplot.Boxplot("load_time_boxplot", "load_time", boxplot.WithExecutionHint("high_accuracy"))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestBoxplotAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
        "boxplot": {
            "field": "grade",
            "missing": 10
        }
    }`

	res := boxplot.Boxplot("grade_boxplot", "grade", boxplot.WithMissing(10))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
