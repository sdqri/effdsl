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

func TestCardinalityAggregation_WithPrecisionThreshold(t *testing.T) {
	expectedBody := `{
        "cardinality": {
            "field": "type",
            "precision_threshold": 100
        }
    }`

	res := cardinality.Cardinality("type_count", "type", cardinality.WithPrecisionThreshold(100))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCardinalityAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
        "cardinality": {
            "field": "tag",
            "missing": "N/A"
        }
    }`

	res := cardinality.Cardinality("tag_cardinality", "tag", cardinality.WithMissing("N/A"))

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
