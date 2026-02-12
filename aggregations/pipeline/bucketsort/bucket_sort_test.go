package bucketsort_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	bucketsort "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketsort"
)

func TestBucketSortAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "bucket_sort": {
            "sort": [
                {
                    "total_sales": {
                        "order": "desc"
                    }
                }
            ],
            "size": 3
        }
    }`

	res := bucketsort.BucketSort(
		"sales_bucket_sort",
		bucketsort.WithSort([]any{
			map[string]any{"total_sales": map[string]any{"order": "desc"}},
		}),
		bucketsort.WithSize(3),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
