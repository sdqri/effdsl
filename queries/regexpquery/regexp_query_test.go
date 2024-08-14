package regexpquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	rq "github.com/sdqri/effdsl/queries/regexpquery"
)

func TestRegexpQuery_WithNoOptions(t *testing.T) {
	expectedBody := `{"regexp":{"fake_field":{"value":"fake_value"}}}`
	regexpQueryResult := rq.RegexpQuery("fake_field", "fake_value")
	err := regexpQueryResult.Err
	regexpQuery := regexpQueryResult.Ok
	jsonBody, err := json.Marshal(regexpQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestRegexpQuery_WithAllOptions(t *testing.T) {
	expectedBody := `{"regexp":{"fake_field":{"value":"fake_value","flags":"fake_flags","case_insensitive":true,"max_determinized_states":100,"rewrite":"fake_rewrite"}}}`
	regexpQueryResult := rq.RegexpQuery(
		"fake_field",
		"fake_value",
		rq.WithFlags("fake_flags"),
		rq.WithCaseInsensitive(),
		rq.WithMaxDeterminizedStates(100),
		rq.WithRQRewrite("fake_rewrite"))
	err := regexpQueryResult.Err
	regexpQuery := regexpQueryResult.Ok
	jsonBody, err := json.Marshal(regexpQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
