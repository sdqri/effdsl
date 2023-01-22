package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestNewRangeQueryWithNoOptions(t *testing.T) {
	rangeQueryResult := objs.D.RangeQuery("fake_field")
	err := rangeQueryResult.Err
	assert.NotNil(t, err)
}

func TestNewRangeQ1(t *testing.T) {
	expectedBody := `{"range":{"fake_field":{"gt":0,"lt":10}}}`
	rangeQueryResult := objs.D.RangeQuery("fake_field", objs.D.WithGT(0), objs.D.WithLT(10))
	err := rangeQueryResult.Err
	rangeQuery := rangeQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(rangeQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
