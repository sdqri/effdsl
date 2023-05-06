package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
)

func TestNewTermSetQuery_WithTermsOnly(t *testing.T) {
	expectedBody := `{"terms_set":{"fake_term":{"terms":["fake_value1","fake_value2"]}}}`
	termSetQueryResult := effdsl.TermsSetQuery("fake_term", []string{"fake_value1", "fake_value2"})
	err := termSetQueryResult.Err
	termSetQuery := termSetQueryResult.Ok
	jsonBody, err := json.Marshal(termSetQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewTermSetQuery_WithAllOptions(t *testing.T) {
	expectedBody := `{"terms_set":{"fake_term":{"terms":["fake_value1","fake_value2"],"minimum_should_match_field":"fake_field","minimum_should_match_script":"doc['fake_field'].size() \u003e= 2"}}}`
	termSetQueryResult := effdsl.TermsSetQuery("fake_term", []string{"fake_value1", "fake_value2"}, effdsl.WithMinimumShouldMatchField("fake_field"), effdsl.WithMinimumShouldMatchScript("doc['fake_field'].size() >= 2"))
	err := termSetQueryResult.Err
	termSetQuery := termSetQueryResult.Ok
	jsonBody, err := json.Marshal(termSetQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
