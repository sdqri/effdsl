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
