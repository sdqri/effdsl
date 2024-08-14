package termsquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	tsq "github.com/sdqri/effdsl/queries/termsquery"
)

func TestNewTermsQuery(t *testing.T) {
	expectedBody := `{"terms":{"boost":10,"fake_term":["fake_value1","fake_value2"]}}`
	termsQueryResult := tsq.TermsQuery(
		"fake_term", []string{"fake_value1", "fake_value2"},
		tsq.WithTSQBoost(10),
	)
	err := termsQueryResult.Err
	termsQuery := termsQueryResult.Ok
	jsonBody, err := json.Marshal(termsQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
