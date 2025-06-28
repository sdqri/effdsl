package effdsl_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sdqri/effdsl/v2"

	mq "github.com/sdqri/effdsl/v2/queries/matchquery"
	cs "github.com/sdqri/effdsl/v2/suggesters/completionsuggester"
	ps "github.com/sdqri/effdsl/v2/suggesters/phrasesuggester"
	ts "github.com/sdqri/effdsl/v2/suggesters/termsuggester"
)

// Suggesters is deprecated and only supports term suggestion, Should be removed on v3
func TestDeprecatedSuggesters(t *testing.T) {
	query := effdsl.Suggesters("test",
		effdsl.TermSuggester(
			effdsl.Term("my-suggestion-1", "tring out Elasticsearch", "message"),
			effdsl.Term("my-suggestion-2", "tring out Elasticsearch", "message",
				effdsl.WithTermSuggestMode(effdsl.SuggestModeAlways),
				effdsl.WithTermSuggestAnalyzer("test"),
				effdsl.WithTermSuggestSize(1),
				effdsl.WithTermSuggestSort(effdsl.SortScore),
			),
		),
	)

	jsonBody, err := json.Marshal(query.Ok)
	require.NoError(t, err)

	const expected = `{
	  "suggest": {
		"my-suggestion-1": {
		  "text": "tring out Elasticsearch",
		  "term": {
			"field": "message"
		  }
		},
		"my-suggestion-2": {
		  "text": "tring out Elasticsearch",
		  "term": {
			"field": "message",
			"analyzer": "test",
			"size": 1,
			"sort": "score",
			"suggest_mode": "always"
		  }
		}
	  }
	}`

	assert.JSONEq(t, expected, string(jsonBody))
}

// Suggesters is deprecated and only supports term suggestion, Should be removed on v3
func TestSuggestersWithAllTypes(t *testing.T) {
	query, err := effdsl.Define(
		effdsl.WithSuggest(
			ts.TermSuggester(
				"term suggester",
				"tring out Elasticsearch",
				"message",
				ts.WithAnalyzer("test"),
				ts.WithSize(1),
				ts.WithSort(ts.ByScore),
				ts.WithMode(ts.Always),
			),
			ps.PhraseSuggester(
				"phrase suggester",
				"noble prize",
				"title.trigram",
				ps.WithGramSize(3),
				ps.WithRealWordErrorLikelihood(0.95),
				ps.WithConfidence(1.0),
				ps.WithMaxErrors(1),
				ps.WithSeparator(" "),
				ps.WithSize(1),
				ps.WithAnalyzer("trigram"),
				ps.WithShardSize(2),
				ps.WithHighlight("<em>", "</em>"),
				ps.WithCollate(mq.MatchQuery("title", ps.CURRENT_SUGGESTION), ps.WithPrune(), ps.WithParams(map[string]string{"field_name": "title"})),
				ps.WithLaplaceSmoothing(0.7),
				ps.WithDirectGenerator("title.trigram", ps.WithSuggestMode(ps.Always)),
			),
			cs.CompletionSuggester(
				"completion suggester",
				"nir",
				"suggest",
				cs.WithCompletionSuggesterFuzzy(
					cs.WithFuzzyFuzziness("AUTO"),
					cs.WithFuzzyTranspositions(true),
					cs.WithFuzzyMinLength(3),
					cs.WithFuzzyPrefixLength(2),
					cs.WithFuzzyUnicodeAware(true),
				),
			),
			cs.CompletionSuggester(
				"completion suggester with context",
				"restaurant",
				"location_suggest",
				cs.WithGeoContext(
					"office_nearby",
					40.7128, -74.0060,
					cs.WithGeoContextBoost(3.0),
					cs.WithGeoContextPrecision("5km"),
					cs.WithGeoContextNeighbours("brooklyn", "jersey_city"),
				),
			),
		),
	)

	require.NoError(t, err)

	jsonBody, err := json.Marshal(query)

	require.NoError(t, err)

	const expected = `{
	   "suggest":{
		  "completion suggester":{
			 "completion":{
				"field":"suggest",
				"fuzzy":{
				   "fuzziness":"AUTO",
				   "transpositions":true,
				   "min_length":3,
				   "prefix_length":2,
				   "unicode_aware":true
				}
			 },
			 "prefix":"nir"
		  },
		  "completion suggester with context":{
			 "completion":{
				"field":"location_suggest",
				"contexts":{
				   "office_nearby":{
					  "context":{
						 "lat":40.7128,
						 "lon":-74.006
					  },
					  "boost":3,
					  "precision":"5km",
					  "neighbours":[
						 "brooklyn",
						 "jersey_city"
					  ]
				   }
				}
			 },
			 "prefix":"restaurant"
		  },
		  "phrase suggester":{
			 "phrase":{
				"field":"title.trigram",
				"gram_size":3,
				"real_word_error_likelihood":0.95,
				"confidence":1,
				"max_errors":1,
				"separator":" ",
				"size":1,
				"analyzer":"trigram",
				"shard_size":2,
				"highlight":{
				   "pre_tag":"\u003cem\u003e",
				   "post_tag":"\u003c/em\u003e"
				},
				"collate":{
				   "query":{
					  "source":{
						 "match":{
							"title":{
							   "query":"{{suggestion}}"
							}
						 }
					  }
				   },
				   "params":{
					  "field_name":"title"
				   },
				   "prune":true
				},
				"smoothing":{
				   "laplace":{
					  "alpha":0.7
				   }
				},
				"direct_generator":[
				   {
					  "field":"title.trigram",
					  "suggest_mode":"always"
				   }
				]
			 },
			 "text":"noble prize"
		  },
		  "term suggester":{
			 "term":{
				"field":"message",
				"analyzer":"test",
				"size":1,
				"sort":"score",
				"suggest_mode":"always"
			 },
			 "text":"tring out Elasticsearch"
		  }
	   }
	}`

	assert.JSONEq(t, expected, string(jsonBody))
}
