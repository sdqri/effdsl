package rangequery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	rq "github.com/sdqri/effdsl/queries/rangequery"
)

func TestNewRangeQueryWithNoOptions(t *testing.T) {
	rangeQueryResult := rq.RangeQuery("fake_field")
	err := rangeQueryResult.Err
	assert.NotNil(t, err)
}

func TestNewRangeQ1(t *testing.T) {
	expectedBody := `{"range":{"fake_field":{"gt":0,"lt":10}}}`
	rangeQueryResult := rq.RangeQuery(
		"fake_field",
		rq.WithGT(0),
		rq.WithLT(10),
	)
	err := rangeQueryResult.Err
	rangeQuery := rangeQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(rangeQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
