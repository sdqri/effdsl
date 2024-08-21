package matchboolprefix_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	mbpq "github.com/sdqri/effdsl/v2/queries/matchboolprefix"
)

func TestNewMatchBoolPrefixQueryWithAnalyzer(t *testing.T) {
	expectedBody := `{"match_bool_prefix":{"message":{"query":"quick brown f","analyzer":"keyword"}}}`
	matchQueryResult := mbpq.MatchBoolPrefixQuery(
		"message",
		"quick brown f",
		mbpq.WithAnalyzer("keyword"),
	)
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
