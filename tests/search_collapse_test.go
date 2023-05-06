package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl/objects"
)

func TestNewSearchCollapse(t *testing.T) {
	expectedBody := `{"field":"fake_field"}`
	searchCollapse := effdsl.Collapse("fake_field")
	jsonBody, err := json.Marshal(searchCollapse)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
