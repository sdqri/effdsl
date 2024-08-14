package termssetquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	tsq "github.com/sdqri/effdsl/queries/termssetquery"
)

func TestNewTermSetQuery_WithTermsOnly(t *testing.T) {
	expectedBody := `{"terms_set":{"fake_term":{"terms":["fake_value1","fake_value2"]}}}`
	termSetQueryResult := tsq.TermsSetQuery(
		"fake_term",
		[]string{"fake_value1", "fake_value2"},
	)
	err := termSetQueryResult.Err
	termSetQuery := termSetQueryResult.Ok
	jsonBody, err := json.Marshal(termSetQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewTermSetQuery_WithAllOptions(t *testing.T) {
	expectedBody := `{"terms_set":{"fake_term":{"terms":["fake_value1","fake_value2"],"minimum_should_match_field":"fake_field","minimum_should_match_script":"doc['fake_field'].size() \u003e= 2"}}}`
	termSetQueryResult := tsq.TermsSetQuery(
		"fake_term",
		[]string{"fake_value1", "fake_value2"},
		tsq.WithMinimumShouldMatchField("fake_field"),
		tsq.WithMinimumShouldMatchScript("doc['fake_field'].size() >= 2"),
	)
	err := termSetQueryResult.Err
	termSetQuery := termSetQueryResult.Ok
	jsonBody, err := json.Marshal(termSetQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
