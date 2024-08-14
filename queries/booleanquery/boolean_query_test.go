package booleanquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
	bq "github.com/sdqri/effdsl/queries/booleanquery"
)

func TestBoolQuery(t *testing.T) {
	expectedBody := `{"bool":{"must":[{"fake_query1":"fake_value1"},{"fake_query2":"fake_value2"}],"filter":[{"fake_filter1":"fake_value1"},{"fake_filter2":"fake_value2"}],"should":[{"fake_query1":"fake_value1"},{"fake_query2":"fake_value2"}],"must_not":[{"fake_query1":"fake_value1"},{"fake_query2":"fake_value2"}]}}`
	boolQueryResult := bq.BoolQuery(
		bq.Must(effdsl.MockQuery(effdsl.M{"fake_query1": "fake_value1"})),
		bq.Must(effdsl.MockQuery(effdsl.M{"fake_query2": "fake_value2"})),
		bq.Filter(effdsl.MockQuery(effdsl.M{"fake_filter1": "fake_value1"})),
		bq.Filter(effdsl.MockQuery(effdsl.M{"fake_filter2": "fake_value2"})),
		bq.Should(effdsl.MockQuery(effdsl.M{"fake_query1": "fake_value1"})),
		bq.Should(effdsl.MockQuery(effdsl.M{"fake_query2": "fake_value2"})),
		bq.MustNot(effdsl.MockQuery(effdsl.M{"fake_query1": "fake_value1"})),
		bq.MustNot(effdsl.MockQuery(effdsl.M{"fake_query2": "fake_value2"})),
	)
	err := boolQueryResult.Err
	boolQuery := boolQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(boolQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
