package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestNewTermsQuery(t *testing.T) {
	expectedBody := `{"terms":{"fake_term":["fake_value1","fake_value2"]}}`
	termsQueryResult := objs.D.TermsQuery("fake_term", []string{"fake_value1", "fake_value2"})
	err := termsQueryResult.Err
	termsQuery := termsQueryResult.Ok
	jsonBody, err := json.Marshal(termsQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

// TODO: Implement TestNewTermsQuery_WithAllOption
