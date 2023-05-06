package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
)

func TestNewTermQuery(t *testing.T) {
	expectedBody := `{"term":{"fake_term":{"value":"fake_value","boost":2}}}`
	termQueryResult := effdsl.TermQuery("fake_term", "fake_value", effdsl.WithTQBoost(2))
	err := termQueryResult.Err
	termQuery := termQueryResult.Ok
	jsonBody, err := json.Marshal(termQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
