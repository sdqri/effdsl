package nested_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	nested "github.com/sdqri/effdsl/v2/aggregations/bucket/nested"
)

func TestNestedAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "nested": {
            "path": "comments"
        }
    }`

	res := nested.Nested("by_comment", "comments")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
