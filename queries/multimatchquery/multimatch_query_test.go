package multimatchquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	mmq "github.com/sdqri/effdsl/v2/queries/multimatchquery"
)

func TestMultiMatchQueryWithNoOptions(t *testing.T) {
	queryResult := mmq.MultiMatchQuery("quick brown fox")
	err := queryResult.Err
	query := queryResult.Ok
	assert.Nil(t, err)
	expectedBody := `{"multi_match":{"query":"quick brown fox"}}`
	jsonBody, err := json.Marshal(query)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestMultiMatchQueryWithFieldsAndType(t *testing.T) {
	queryResult := mmq.MultiMatchQuery(
		"quick brown fox",
		mmq.WithFields("title", "message"),
		mmq.WithType(mmq.BestFields),
	)
	err := queryResult.Err
	query := queryResult.Ok
	assert.Nil(t, err)
	expectedBody := `{"multi_match":{"query":"quick brown fox","fields":["title","message"],"type":"best_fields"}}`
	jsonBody, err := json.Marshal(query)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestMultiMatchQueryWithAllOptions(t *testing.T) {
	queryResult := mmq.MultiMatchQuery(
		"quick brown fox",
		mmq.WithFields("title", "message"),
		mmq.WithType(mmq.MostFields),
		mmq.WithOperator(mmq.OperatorAnd),
		mmq.WithAnalyzer("standard"),
		mmq.WithSlop(2),
		mmq.WithFuzziness("AUTO"),
		mmq.WithPrefixLength(1),
		mmq.WithMaxExpansions(10),
		mmq.WithMinimumShouldMatch("2"),
		mmq.WithTieBreaker(0.3),
		mmq.WithLenient(true),
		mmq.WithZeroTermsQuery(mmq.All),
		mmq.WithAutoGenerateSynonymsPhraseQuery(true),
	)
	err := queryResult.Err
	query := queryResult.Ok
	assert.Nil(t, err)
	expectedBody := `{"multi_match":{"query":"quick brown fox","fields":["title","message"],"type":"most_fields","operator":"and","analyzer":"standard","slop":2,"fuzziness":"AUTO","prefix_length":1,"max_expansions":10,"minimum_should_match":"2","tie_breaker":0.3,"lenient":true,"zero_terms_query":"all","auto_generate_synonyms_phrase_query":true}}`
	jsonBody, err := json.Marshal(query)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
