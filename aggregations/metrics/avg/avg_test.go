package avg_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	aa "github.com/sdqri/effdsl/v2/aggregations/metrics/avg"
)

func TestAvgAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
		"avg": {
			"field": "grade"
		}
	}`

	avgResult := aa.Avg("avg_grade", "grade")

	assert.Nil(t, avgResult.Err)
	assert.NotNil(t, avgResult.Ok)

	jsonBody, err := json.Marshal(avgResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestAvgAggregation_WithMissing(t *testing.T) {
	expectedBody := `{
		"avg": {
			"field": "grade",
			"missing": 10
		}
	}`

	avgResult := aa.Avg("grade_avg", "grade", aa.WithMissing(10))

	assert.Nil(t, avgResult.Err)
	assert.NotNil(t, avgResult.Ok)

	jsonBody, err := json.Marshal(avgResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

// func TestAvgAggregation_AllOptions(t *testing.T) {
// 	script := aggregations.Script{
// 		Lang:   "painless",
// 		Source: "doc['price'] * params.factor",
// 		Params: map[string]any{"factor": 2},
// 	}
//
// 	// use your real Sum aggregation
// 	subAgg := aa.Sum("sum_price", "price")
//
// 	avgResult := aa.Avg(
// 		"avg_price",
// 		"price",
// 		aa.WithScript(script),
// 		aa.WithMissing(0),
// 		aa.WithFormat("0.0"),
// 		aa.WithValueType("double"),
// 		aa.WithMetaMap(map[string]any{"meta": "value"}),
// 		aa.WithSubAggregation(subAgg),
// 	)
//
// 	assert.Nil(t, avgResult.Err)
// 	assert.NotNil(t, avgResult.Ok)
//
// 	jsonBody, err := json.Marshal(avgResult.Ok)
// 	assert.Nil(t, err)
//
// 	expectedBody := `{
// 	  "avg": {
// 	    "field": "price",
// 	    "missing": 0,
// 	    "script": {
// 	      "lang": "painless",
// 	      "source": "doc['price'] * params.factor",
// 	      "params": { "factor": 2 }
// 	    },
// 	    "format": "0.0",
// 	    "value_type": "double"
// 	  },
// 	  "aggs": {
// 	    "sum_price": {
// 	      "sum": { "field": "price" }
// 	    }
// 	  },
// 	  "meta": {
// 	    "meta": "value"
// 	  }
// 	}`
//
// 	assert.JSONEq(t, expectedBody, string(jsonBody))
// }
