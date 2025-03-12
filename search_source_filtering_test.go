package effdsl_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2"
)

func TestNewSourceFilter(t *testing.T) {
	expectedBody := `{"includes":["field1","field2"],"excludes":["field3","field4"]}`
	sourceFilter := effdsl.SourceFilter(
		effdsl.WithIncludes("field1", "field2"),
		effdsl.WithExcludes("field3", "field4"),
	)
	jsonBody, err := json.Marshal(sourceFilter)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
