package effdsl_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2"
)

func TestNewSearchCollapse(t *testing.T) {
	expectedBody := `{"field":"fake_field"}`
	searchCollapse := effdsl.Collapse("fake_field")
	jsonBody, err := json.Marshal(searchCollapse)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
