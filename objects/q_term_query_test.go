package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestNewTermQuery(t *testing.T) {
	expectedBody := `{"term":{"fake_term":{"value":"fake_value","boost":2}}}`
	termQueryResult := objs.D.TermQuery("fake_term", "fake_value", objs.D.WithBoost(2))
	err := termQueryResult.Err
	termQuery := termQueryResult.Ok
	jsonBody, err := json.Marshal(termQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
