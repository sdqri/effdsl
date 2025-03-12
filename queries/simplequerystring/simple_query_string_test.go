package simplequerystring_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	sqs "github.com/sdqri/effdsl/v2/queries/simplequerystring"
)

func TestQueryString(t *testing.T) {
	expectedBody := `{"simple_query_string":{"query":"\"fried eggs\" +(eggplant | potato) -frittata","fields":["title^5","body"],"default_operator":"AND"}}`
	queryStringResult := sqs.SimpleQueryString(
		`"fried eggs" +(eggplant | potato) -frittata`,
		sqs.WithFields("title^5", "body"),
		sqs.WithDefaultOperator(sqs.AND),
	)
	err := queryStringResult.Err
	queryString := queryStringResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(queryString)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestSimpleQueryStringAllOptions(t *testing.T) {
	expectedBody := `{
		"simple_query_string": {
			"query": "fake_query",
			"fields": ["field1", "field2"],
			"default_operator": "AND",
			"analyze_wildcard": true,
			"analyzer": "fake_analyzer",
			"auto_generate_synonyms_phrase_query": true,
			"flags": "ALL",
			"fuzzy_max_expansions": 50,
			"fuzzy_prefix_length": 2,
			"fuzzy_transpositions": true,
			"lenient": true,
			"minimum_should_match": "2",
			"quote_field_suffix": ".suffix"
		}
	}`
	queryStringResult := sqs.SimpleQueryString(
		"fake_query",
		sqs.WithFields("field1", "field2"),
		sqs.WithDefaultOperator(sqs.AND),
		sqs.WithAnalyzeWildcard(),
		sqs.WithAnalyzer("fake_analyzer"),
		sqs.WithAutoGenerateSynonymsPhrase(true),
		sqs.WithFlags("ALL"),
		sqs.WithFuzzyMaxExpansions(50),
		sqs.WithFuzzyPrefixLength(2),
		sqs.WithFuzzyTranspositions(true),
		sqs.WithLenient(true),
		sqs.WithMinimumShouldMatch("2"),
		sqs.WithQuoteFieldSuffix(".suffix"),
	)
	err := queryStringResult.Err
	queryString := queryStringResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(queryString)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
