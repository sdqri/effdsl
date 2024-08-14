package querystring_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	qs "github.com/sdqri/effdsl/queries/querystring"
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
