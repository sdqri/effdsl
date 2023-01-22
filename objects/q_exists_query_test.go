package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestNewExistsQuery(t *testing.T) {
	expectedBody := `{"exists":{"field":"fake_field"}}`
	existsQueryResult := objs.D.ExistsQuery("fake_field")
	err := existsQueryResult.Err
	existsQuery := existsQueryResult.Ok
	jsonBody, err := json.Marshal(existsQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
