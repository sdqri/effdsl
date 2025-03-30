package effdsl_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2"
)

func TestTermsAggregation(t *testing.T) {
	expectedBody := `{"terms":{"field":"fake_field"}}`
	termsAggregation := effdsl.TermAggregation("fake_field")
	jsonBody, err := json.Marshal(termsAggregation)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestTermsAggregationWithSize(t *testing.T) {
	expectedBody := `{"terms":{"field":"fake_field","size":10}}`
	termsAggregation := effdsl.TermAggregation("fake_field", 10)
	jsonBody, err := json.Marshal(termsAggregation)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestStatsAggregation(t *testing.T) {
	expectedBody := `{"stats":{"field":"fake_field"}}`
	termsAggregation := effdsl.StatsAggregation("fake_field")
	jsonBody, err := json.Marshal(termsAggregation)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
