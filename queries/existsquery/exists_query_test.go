package existsquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	eq "github.com/sdqri/effdsl/v2/queries/existsquery"
)

func TestNewExistsQuery(t *testing.T) {
	expectedBody := `{"exists":{"field":"fake_field"}}`
	existsQueryResult := eq.ExistsQuery("fake_field")
	err := existsQueryResult.Err
	existsQuery := existsQueryResult.Ok
	jsonBody, err := json.Marshal(existsQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
