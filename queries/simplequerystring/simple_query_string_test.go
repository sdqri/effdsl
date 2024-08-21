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
