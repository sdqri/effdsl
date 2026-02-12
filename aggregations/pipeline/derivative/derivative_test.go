package derivative_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	derivative "github.com/sdqri/effdsl/v2/aggregations/pipeline/derivative"
)

func TestDerivativeAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "derivative": {
            "buckets_path": "sales"
        }
    }`

	res := derivative.Derivative("sales_deriv", "sales")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
