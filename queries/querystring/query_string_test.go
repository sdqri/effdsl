package querystring_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	qs "github.com/sdqri/effdsl/v2/queries/querystring"
)

func TestQueryString(t *testing.T) {
	expectedBody := `{"query_string":{"query":"fake_query","fields":["field1","field2"]}}`
	queryStringResult := qs.QueryString("fake_query", qs.WithFields("field1", "field2"))
	err := queryStringResult.Err
	queryString := queryStringResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(queryString)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestQueryStringWithAnalyzeWildcard(t *testing.T) {
	expectedBody := `{"query_string":{"query":"fake_query","analyze_wildcard":true,"fields":["field1","field2"]}}`
	queryStringResult := qs.QueryString("fake_query", qs.WithFields("field1", "field2"), qs.WithAnalyzeWildcard())
	err := queryStringResult.Err
	queryString := queryStringResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(queryString)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestQueryStringAllOptions(t *testing.T) {
	expectedBody := `{
		"query_string": {
			"query": "fake_query",
			"default_field": "default_field",
			"allow_leading_wildcard": true,
			"analyze_wildcard": true,
			"analyzer": "standard",
			"auto_generate_synonyms_phrase_query": true,
			"boost": 1.5,
			"default_operator": "AND",
			"enable_position_increments": true,
			"fields": ["field1", "field2"],
			"fuzziness": "AUTO",
			"fuzzy_max_expansions": 50,
			"fuzzy_prefix_length": 2,
			"fuzzy_transpositions": true,
			"lenient": true,
			"max_determinized_states": 10000,
			"minimum_should_match": "2",
			"quote_analyzer": "quote_analyzer",
			"phrase_slop": 3,
			"quote_field_suffix": ".exact",
			"rewrite": "constant_score",
			"time_zone": "UTC"
		}
	}`
	queryStringResult := qs.QueryString(
		"fake_query",
		qs.WithDefaultField("default_field"),
		qs.WithAllowLeadingWildcard(),
		qs.WithAnalyzeWildcard(),
		qs.WithAnalyzer("standard"),
		qs.WithAutoGenerateSynonymsPhrase(true),
		qs.WithBoost(1.5),
		qs.WithDefaultOperator(qs.AND),
		qs.WithEnablePositionIncrements(true),
		qs.WithFields("field1", "field2"),
		qs.WithFuzziness("AUTO"),
		qs.WithFuzzyMaxExpansions(50),
		qs.WithFuzzyPrefixLength(2),
		qs.WithFuzzyTranspositions(true),
		qs.WithLenient(true),
		qs.WithMaxDeterminizedStates(10000),
		qs.WithMinimumShouldMatch("2"),
		qs.WithQuoteAnalyzer("quote_analyzer"),
		qs.WithPhraseSlop(3),
		qs.WithQuoteFieldSuffix(".exact"),
		qs.WithRewrite(qs.ConstantScore),
		qs.WithTimeZone("UTC"),
	)

	err := queryStringResult.Err
	queryString := queryStringResult.Ok
	assert.Nil(t, err)

	jsonBody, err := json.Marshal(queryString)
	assert.Nil(t, err)

	// Remove any unnecessary whitespace or newlines from expectedBody
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
