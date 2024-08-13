package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
)

func TestQueryString(t *testing.T) {
	expectedBody := `{"query_string":{"query":"fake_query","fields":["field1","field2"]}}`
	queryStringResult := effdsl.QueryString("fake_query", effdsl.WithFields("field1", "field2"))
	err := queryStringResult.Err
	queryString := queryStringResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(queryString)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestQueryStringWithAnalyzeWildcard(t *testing.T) {
	expectedBody := `{"query_string":{"query":"fake_query","fields":["field1","field2"],"analyze_wildcard":true}}`
	queryStringResult := effdsl.QueryString("fake_query", effdsl.WithFields("field1", "field2"), effdsl.WithAnalyzeWildcard())
	err := queryStringResult.Err
	queryString := queryStringResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(queryString)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
