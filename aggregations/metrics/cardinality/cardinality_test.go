package cardinality_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	cardinality "github.com/sdqri/effdsl/v2/aggregations/metrics/cardinality"
)

func TestCardinalityAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "cardinality": {
            "field": "type"
        }
    }`

	res := cardinality.Cardinality("type_count", "type")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
