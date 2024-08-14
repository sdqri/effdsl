package constantscore_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	cs "github.com/sdqri/effdsl/queries/constantscore"
	tq "github.com/sdqri/effdsl/queries/termquery"
)

func TestConstantScoreQuery_WithAllOptions(t *testing.T) {
	expectedBody := `{"constant_score":{"filter":{"term":{"user.id":{"value":"kimchy"}}},"boost":1.2}}`

	// Create the constant score query
	constantScoreQueryResult := cs.ConstantScoreQuery(
		tq.TermQuery("user.id", "kimchy"),
		1.2, // boost
	)

	// Check for errors
	err := constantScoreQueryResult.Err
	constantScoreQuery := constantScoreQueryResult.Ok

	// Marshal the query to JSON
	jsonBody, err := json.Marshal(constantScoreQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
