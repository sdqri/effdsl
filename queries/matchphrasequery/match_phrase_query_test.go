package matchphrasequery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	mpq "github.com/sdqri/effdsl/queries/matchphrasequery"
)

func TestNewMatchPhraseQueryWithNoOptions(t *testing.T) {
	matchQueryResult := mpq.MatchPhraseQuery("message", "this is a test")
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	expectedBody := `{"match_phrase":{"message":{"query":"this is a test"}}}`
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhraseQueryWithAnalyzer(t *testing.T) {
	expectedBody := `{"match_phrase":{"message":{"query":"this is a test","analyzer":"my_analyzer"}}}`
	matchQueryResult := mpq.MatchPhraseQuery(
		"message",
		"this is a test",
		mpq.WithAnalyzer("my_analyzer"),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhraseQueryWithSlop(t *testing.T) {
	expectedBody := `{"match_phrase":{"message":{"query":"this is a test","slop":2}}}`
	matchQueryResult := mpq.MatchPhraseQuery(
		"message",
		"this is a test",
		mpq.WithSlop(2),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhraseQueryWithZeroTermsQuery(t *testing.T) {
	expectedBody := `{"match_phrase":{"message":{"query":"this is a test","zero_terms_query":"all"}}}`
	matchQueryResult := mpq.MatchPhraseQuery(
		"message",
		"this is a test",
		mpq.WithZeroTermsquery(mpq.All),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhraseQueryWithAllOptions(t *testing.T) {
	expectedBody := `{"match_phrase":{"message":{"query":"this is a test","analyzer":"my_analyzer","slop":2,"zero_terms_query":"all"}}}`
	matchQueryResult := mpq.MatchPhraseQuery(
		"message",
		"this is a test",
		mpq.WithAnalyzer("my_analyzer"),
		mpq.WithSlop(2),
		mpq.WithZeroTermsquery(mpq.All),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
