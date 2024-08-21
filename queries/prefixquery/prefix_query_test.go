package prefixquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	pq "github.com/sdqri/effdsl/v2/queries/prefixquery"
)

func TestPrefixQuery_WithNoOptions(t *testing.T) {
	expectedBody := `{"prefix":{"user.id":{"value":"ki"}}}`
	prefixQueryResult := pq.PrefixQuery("user.id", "ki")
	err := prefixQueryResult.Err
	prefixQuery := prefixQueryResult.Ok
	jsonBody, err := json.Marshal(prefixQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
