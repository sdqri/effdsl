package matchquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mq "github.com/sdqri/effdsl/v2/queries/matchquery"
)

func Test_MatchQueryS_MarshalJSON(t *testing.T) {
	q := mq.MatchQuery("field_name", "some match query",
		mq.WithOperator(mq.AND),
		mq.WithFuzzinessParameter(mq.FuzzinessAUTO),
	)

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"match":{"field_name":{"query":"some match query","fuzziness":"AUTO","operator":"AND"}}}`
	require.Equal(t, expected, string(body))
}

func TestMatchQueryAllOptions(t *testing.T) {
	expectedBody := `{
		"match": {
			"field_name": {
				"query": "some match query",
				"analyzer": "standard",
				"auto_generate_synonyms_phrase_query": true,
				"boost": 2.5,
				"fuzziness": "AUTO",
				"max_expansions": 100,
				"prefix_length": 3,
				"fuzzy_transpositions": false,
				"fuzzy_rewrite": "top_terms_boost_N",
				"lenient": true,
				"operator": "AND",
				"minimum_should_match": "2",
				"zero_terms_query": "all"
			}
		}
	}`

	matchQueryResult := mq.MatchQuery(
		"field_name",
		"some match query",
		mq.WithAnalyzer("standard"),
		mq.WithAutoGenerateSynonymsPhrase(true),
		mq.WithBoost(2.5),
		mq.WithFuzzinessParameter(mq.FuzzinessAUTO),
		mq.WithMaxExpansions(100),
		mq.WithPrefixLength(3),
		mq.WithFuzzyTranspositions(false),
		mq.WithFuzzyRewrite(mq.TopTermsBoostN),
		mq.WithLenient(),
		mq.WithOperator(mq.AND),
		mq.MinimumShouldMatch("2"),
		mq.WithZeroTermsquery(mq.All),
	)

	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
