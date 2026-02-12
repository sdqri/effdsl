package parent_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	parent "github.com/sdqri/effdsl/v2/aggregations/bucket/parent"
)

func TestParentAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "parent": {
            "type": "parent"
        }
    }`

	res := parent.Parent("to_parent", "parent")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
