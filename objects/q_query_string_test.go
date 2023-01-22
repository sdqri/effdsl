package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestQueryString(t *testing.T) {
	expectedBody := `{"query_string":{"query":"fake_query","fields":["field1","field2"]}}`
	queryStringResult := objs.D.QueryString("fake_query", objs.D.WithFields("field1", "field2"))
	err := queryStringResult.Err
	queryString := queryStringResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(queryString)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
