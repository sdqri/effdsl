package fuzzyquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	fq "github.com/sdqri/effdsl/queries/fuzzyquery"
)

func TestFuzzyQuery_WithNoOptions(t *testing.T) {
	expectedBody := `{"fuzzy":{"fake_field":{"value":"fake_value"}}}`
	fuzzyQueryResult := fq.FuzzyQuery("fake_field", "fake_value")
	err := fuzzyQueryResult.Err
	fuzzyQuery := fuzzyQueryResult.Ok
	jsonBody, err := json.Marshal(fuzzyQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestFuzzyQuery_WithAllOptions(t *testing.T) {
	expectedBody := `{"fuzzy":{"fake_field":{"value":"fake_value","fuzziness":"AUTO","max_expansions":10,"prefix_length":1,"transpositions":true,"rewrite":"constant_score"}}}`
	fuzzyQueryResult := fq.FuzzyQuery(
		"fake_field",
		"fake_value",
		fq.WithFuzziness("AUTO"),
		fq.WithPrefixLength(1),
		fq.WithMaxExpansions(10),
		fq.WithTranspositions(true),
		fq.WithRewrite(fq.ConstantScore),
	)
	err := fuzzyQueryResult.Err
	fuzzyQuery := fuzzyQueryResult.Ok
	jsonBody, err := json.Marshal(fuzzyQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
