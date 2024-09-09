package rangequery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	rq "github.com/sdqri/effdsl/v2/queries/rangequery"
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

func TestRangeQueryAllOptions(t *testing.T) {
	expectedBody := `{
		"range": {
			"fake_field": {
				"gt": 10,
				"gte": 20,
				"lt": 30,
				"lte": 40,
				"format": "strict_date_optional_time",
				"relation": "INTERSECTS",
				"time_zone": "UTC",
				"boost": 1.5
			}
		}
	}`
	rangeQueryResult := rq.RangeQuery(
		"fake_field",
		rq.WithGT(10),
		rq.WithGTE(20),
		rq.WithLT(30),
		rq.WithLTE(40),
		rq.WithFormat("strict_date_optional_time"),
		rq.WithRelation(rq.INTERSECTS),
		rq.WithTimeZone("UTC"),
		rq.WithBoost(1.5),
	)
	err := rangeQueryResult.Err
	rangeQuery := rangeQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(rangeQuery)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
