package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl/objects"
)

func TestNewTermsQuery(t *testing.T) {
	expectedBody := `{"terms":{"boost":10,"fake_term":["fake_value1","fake_value2"]}}`
	termsQueryResult := effdsl.TermsQuery("fake_term", []string{"fake_value1", "fake_value2"}, effdsl.WithTSQBoost(10))
	err := termsQueryResult.Err
	termsQuery := termsQueryResult.Ok
	jsonBody, err := json.Marshal(termsQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
