package adjacencymatrix_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	adjacencymatrix "github.com/sdqri/effdsl/v2/aggregations/bucket/adjacencymatrix"
)

func TestAdjacencyMatrixAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "adjacency_matrix": {
            "filters": {
                "group_a": {"term": {"tags": "foo"}},
                "group_b": {"term": {"tags": "bar"}}
            }
        }
    }`

	filters := map[string]any{
		"group_a": map[string]any{"term": map[string]any{"tags": "foo"}},
		"group_b": map[string]any{"term": map[string]any{"tags": "bar"}},
	}

	res := adjacencymatrix.AdjacencyMatrix("interactions", filters)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
