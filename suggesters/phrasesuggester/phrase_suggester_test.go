package phrasesuggester_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	mq "github.com/sdqri/effdsl/v2/queries/matchquery"
	ps "github.com/sdqri/effdsl/v2/suggesters/phrasesuggester"
)

func TestPhraseSuggester_WithNoOptions(t *testing.T) {
	expectedBody := `{
		"text": "noble prize",
		"phrase": {
			"field": "title.trigram"
		}
	}`

	phraseSuggesterResult := ps.PhraseSuggester(
		"simple-phrase",
		"noble prize",
		"title.trigram", // required base field
	)

	err := phraseSuggesterResult.Err
	suggestQuery := phraseSuggesterResult.Ok
	jsonBody, err := json.Marshal(suggestQuery)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPhraseSuggester_WithLaplaceSmoothing(t *testing.T) {
	expectedBody := `{
		"text": "obel prize",
		"phrase": {
			"field": "title.trigram",
			"size": 1,
			"smoothing": {
				"laplace": {
					"alpha": 0.7
				}
			}
		}
	}`

	phraseSuggesterResult := ps.PhraseSuggester(
		"simple_phrase",
		"obel prize",
		"title.trigram",
		ps.WithSize(1),
		ps.WithLaplaceSmoothing(0.7),
	)

	err := phraseSuggesterResult.Err
	suggestQuery := phraseSuggesterResult.Ok
	jsonBody, err := json.Marshal(suggestQuery)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestPhraseSuggester_AllOptions(t *testing.T) {
	expectedBody := `{
		"text": "noble prize",
		"phrase": {
			"field": "title.trigram",
			"gram_size": 3,
			"real_word_error_likelihood": 0.95,
			"confidence": 1.0,
			"max_errors": 1,
			"separator": " ",
			"size": 1,
			"analyzer": "trigram",
			"shard_size": 2,
			"highlight": {
				"pre_tag": "<em>",
				"post_tag": "</em>"
			},
			"collate": {
				"query": {
					"source": {
						"match": {
							"title": {
								"query": "{{suggestion}}"
							}
						}
					}
				},
				"params": {"field_name" : "title"},
				"prune": true
			},
			"smoothing" : {
				"laplace" : {
					"alpha" : 0.7
				}
			},
			"direct_generator": [{
				"field": "title.trigram",
				"suggest_mode": "always"
			}]
		}
	}`

	phraseSuggesterResult := ps.PhraseSuggester(
		"simple-phrase",
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
	)

	err := phraseSuggesterResult.Err
	suggestQuery := phraseSuggesterResult.Ok
	jsonBody, err := json.Marshal(suggestQuery)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
