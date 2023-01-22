package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestNewSourceFilter(t *testing.T) {
	expectedBody := `{"includes":["field1","field2"],"excludes":["field3","field4"]}`
	sourceFilter := objs.SourceFilter(
		objs.D.WithIncludes("field1", "field2"),
		objs.D.WithExcludes("field3", "field4"),
	)
	jsonBody, err := json.Marshal(sourceFilter)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
