package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
)

func TestFuzzyQuery_WithNoOptions(t *testing.T) {
	expectedBody := `{"fuzzy":{"fake_field":{"value":"fake_value"}}}`
	fuzzyQueryResult := effdsl.FuzzyQuery("fake_field", "fake_value")
	err := fuzzyQueryResult.Err
	fuzzyQuery := fuzzyQueryResult.Ok
	jsonBody, err := json.Marshal(fuzzyQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestFuzzyQuery_WithAllOptions(t *testing.T) {
	expectedBody := `{"fuzzy":{"fake_field":{"value":"fake_value","fuzziness":"AUTO","prefix_length":1,"max_expansions":10,"rewrite":"fake_rewrite","transpositions":true}}}`
	fuzzyQueryResult := effdsl.FuzzyQuery(
		"fake_field",
		"fake_value",
		effdsl.WithFuzziness("AUTO"),
		effdsl.WithPrefixLength(1),
		effdsl.WithMaxExpansions(10),
		effdsl.WithTranspositions(true),
		effdsl.WithFQRewrite("fake_rewrite"),
	)
	err := fuzzyQueryResult.Err
	fuzzyQuery := fuzzyQueryResult.Ok
	jsonBody, err := json.Marshal(fuzzyQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
