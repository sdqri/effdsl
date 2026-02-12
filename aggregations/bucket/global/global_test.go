package global_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	global "github.com/sdqri/effdsl/v2/aggregations/bucket/global"
)

func TestGlobalAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "global": {}
    }`

	res := global.Global("all_docs")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
