package matchphraseprefix_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	mppq "github.com/sdqri/effdsl/queries/matchphraseprefix"
)

func TestNewMatchPhrasePrefixQueryWithNoOptions(t *testing.T) {
	matchQueryResult := mppq.MatchPhrasePrefixQuery("message", "quick brown f")
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	expectedBody := `{"match_phrase_prefix":{"message":{"query":"quick brown f"}}}`
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhrasePrefixQueryWithAnalyzer(t *testing.T) {
	expectedBody := `{"match_phrase_prefix":{"message":{"query":"quick brown f","analyzer":"my_analyzer"}}}`
	matchQueryResult := mppq.MatchPhrasePrefixQuery(
		"message",
		"quick brown f",
		mppq.WithAnalyzer("my_analyzer"),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhrasePrefixQueryWithSlop(t *testing.T) {
	expectedBody := `{"match_phrase_prefix":{"message":{"query":"quick brown f","slop":2}}}`
	matchQueryResult := mppq.MatchPhrasePrefixQuery(
		"message",
		"quick brown f",
		mppq.WithSlop(2),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhrasePrefixQueryWithMaxExpansions(t *testing.T) {
	expectedBody := `{"match_phrase_prefix":{"message":{"query":"quick brown f","max_expansions":50}}}`
	matchQueryResult := mppq.MatchPhrasePrefixQuery(
		"message",
		"quick brown f",
		mppq.WithMaxExpansions(50),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhrasePrefixQueryWithZeroTermsQuery(t *testing.T) {
	expectedBody := `{"match_phrase_prefix":{"message":{"query":"quick brown f","zero_terms_query":"all"}}}`
	matchQueryResult := mppq.MatchPhrasePrefixQuery(
		"message",
		"quick brown f",
		mppq.WithZeroTermsQuery(mppq.All),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewMatchPhrasePrefixQueryWithAllOptions(t *testing.T) {
	expectedBody := `{"match_phrase_prefix":{"message":{"query":"quick brown f","analyzer":"my_analyzer","slop":2,"max_expansions":50,"zero_terms_query":"all"}}}`
	matchQueryResult := mppq.MatchPhrasePrefixQuery(
		"message",
		"quick brown f",
		mppq.WithAnalyzer("my_analyzer"),
		mppq.WithSlop(2),
		mppq.WithMaxExpansions(50),
		mppq.WithZeroTermsQuery(mppq.All),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
