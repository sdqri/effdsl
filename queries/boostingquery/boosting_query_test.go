package boostingquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	bq "github.com/sdqri/effdsl/v2/queries/boostingquery"
	tq "github.com/sdqri/effdsl/v2/queries/termquery"
)

func TestBoostingQuery(t *testing.T) {
	expectedBody := `{"boosting":{"positive":{"term":{"text":{"value":"apple"}}},"negative":{"term":{"text":{"value":"pie tart fruit crumble tree"}}},"negative_boost":0.5}}`

	// Create the boosting query
	boostingQueryResult := bq.BoostingQuery(
		tq.TermQuery("text", "apple"),
		tq.TermQuery("text", "pie tart fruit crumble tree"),
		0.5, // negative_boost
	)

	// Check for errors
	err := boostingQueryResult.Err
	boostingQuery := boostingQueryResult.Ok

	// Marshal the query to JSON
	jsonBody, err := json.Marshal(boostingQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
