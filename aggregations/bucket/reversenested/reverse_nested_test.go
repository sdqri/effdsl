package reversenested_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	reversenested "github.com/sdqri/effdsl/v2/aggregations/bucket/reversenested"
)

func TestReverseNestedAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "reverse_nested": {}
    }`

	res := reversenested.ReverseNested("back_to_root")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
