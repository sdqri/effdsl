package idsquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	iq "github.com/sdqri/effdsl/v2/queries/idsquery"
)

func TestIDsQuery(t *testing.T) {
	expectedBody := `{"ids":{"values":["1","4","100"]}}`

	// Create an instance of IDsQuery with the test values
	idsQueryResult := iq.IDsQuery("1", "4", "100")
	err := idsQueryResult.Err
	idsQuery := idsQueryResult.Ok

	// Convert the query to JSON
	jsonBody, err := json.Marshal(idsQuery)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
