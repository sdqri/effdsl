package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
)

func TestNewExistsQuery(t *testing.T) {
	expectedBody := `{"exists":{"field":"fake_field"}}`
	existsQueryResult := effdsl.ExistsQuery("fake_field")
	err := existsQueryResult.Err
	existsQuery := existsQueryResult.Ok
	jsonBody, err := json.Marshal(existsQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
