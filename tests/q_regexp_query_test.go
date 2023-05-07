package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
)

func TestRegexpQuery_WithNoOptions(t *testing.T) {
	expectedBody := `{"regexp":{"fake_field":{"value":"fake_value"}}}`
	regexpQueryResult := effdsl.RegexpQuery("fake_field", "fake_value")
	err := regexpQueryResult.Err
	regexpQuery := regexpQueryResult.Ok
	jsonBody, err := json.Marshal(regexpQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestRegexpQuery_WithAllOptions(t *testing.T) {
	expectedBody := `{"regexp":{"fake_field":{"value":"fake_value","flags":"fake_flags","case_insensitive":true,"max_determinized_states":100,"rewrite":"fake_rewrite"}}}`
	regexpQueryResult := effdsl.RegexpQuery(
		"fake_field",
		"fake_value",
		effdsl.WithFlags("fake_flags"),
		effdsl.WithCaseInsensitive(),
		effdsl.WithMaxDeterminizedStates(100),
		effdsl.WithRQRewrite("fake_rewrite"))
	err := regexpQueryResult.Err
	regexpQuery := regexpQueryResult.Ok
	jsonBody, err := json.Marshal(regexpQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
