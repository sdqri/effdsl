package children_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	children "github.com/sdqri/effdsl/v2/aggregations/bucket/children"
)

func TestChildrenAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "children": {
            "type": "child"
        }
    }`

	res := children.Children("to_children", "child")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
