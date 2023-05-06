package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
)

func TestNewRangeQueryWithNoOptions(t *testing.T) {
	rangeQueryResult := effdsl.RangeQuery("fake_field")
	err := rangeQueryResult.Err
	assert.NotNil(t, err)
}

func TestNewRangeQ1(t *testing.T) {
	expectedBody := `{"range":{"fake_field":{"gt":0,"lt":10}}}`
	rangeQueryResult := effdsl.RangeQuery("fake_field", effdsl.WithGT(0), effdsl.WithLT(10))
	err := rangeQueryResult.Err
	rangeQuery := rangeQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(rangeQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
