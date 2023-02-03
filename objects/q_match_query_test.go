package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestNewMatchQuery(t *testing.T) {
	expectedBody := `{"match":{"fake_field":{"query":"fake_query"}}}`
	matchQueryResult := objs.D.MatchQuery("fake_field", "fake_query")
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
