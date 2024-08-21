package termquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	tq "github.com/sdqri/effdsl/queries/termquery"
)

func TestNewTermQuery(t *testing.T) {
	expectedBody := `{"term":{"fake_term":{"value":"fake_value","boost":2}}}`
	termQueryResult := tq.TermQuery("fake_term", "fake_value", tq.WithBoost(2))
	err := termQueryResult.Err
	termQuery := termQueryResult.Ok
	jsonBody, err := json.Marshal(termQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
