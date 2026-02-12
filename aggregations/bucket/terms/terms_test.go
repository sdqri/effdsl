package terms_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2/aggregations"
	terms "github.com/sdqri/effdsl/v2/aggregations/bucket/terms"
)

func TestTermsAggregation_NoOptions(t *testing.T) {
	expectedBody := `{
        "terms": {
            "field": "category"
        }
    }`

	res := terms.Terms("categories", "category")

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestTermsAggregation_WithOptions(t *testing.T) {
	expectedBody := `{
        "terms": {
            "field": "category",
            "size": 5,
            "shard_size": 100,
            "order": {"_key": "asc"},
            "min_doc_count": 1,
            "shard_min_doc_count": 0,
            "include": ".*sport.*",
            "exclude": ".*kids.*",
            "missing": "missing-category",
            "execution_hint": "map",
            "collect_mode": "breadth_first",
            "show_term_doc_count_error": true,
            "value_type": "string"
        }
    }`

	res := terms.Terms(
		"categories",
		"category",
		terms.WithSize(5),
		terms.WithShardSize(100),
		terms.WithOrder(map[string]any{"_key": "asc"}),
		terms.WithMinDocCount(1),
		terms.WithShardMinDocCount(0),
		terms.WithInclude(".*sport.*"),
		terms.WithExclude(".*kids.*"),
		terms.WithMissing("missing-category"),
		terms.WithExecutionHint("map"),
		terms.WithCollectMode("breadth_first"),
		terms.WithShowTermDocCountError(true),
		terms.WithValueType("string"),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestTermsAggregation_WithScript(t *testing.T) {
	expectedBody := `{
        "terms": {
            "field": "category",
            "script": {
                "source": "doc['category'].value"
            }
        }
    }`

	res := terms.Terms(
		"categories",
		"category",
		terms.WithScript(aggregations.Script{Source: "doc['category'].value"}),
	)

	assert.Nil(t, res.Err)
	assert.NotNil(t, res.Ok)

	jsonBody, err := json.Marshal(res.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
