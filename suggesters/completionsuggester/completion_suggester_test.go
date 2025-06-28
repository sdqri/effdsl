package completionsuggester_test

import (
	"encoding/json"
	"testing"

	cs "github.com/sdqri/effdsl/v2/suggesters/completionsuggester"
	"github.com/stretchr/testify/assert"
)

func TestCompletionSuggester_WithNoOptions(t *testing.T) {
	expectedBody := `{
		"prefix": "nir",
		"completion": {
			"field": "suggest"
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"song-suggest",
		"nir",
		"suggest",
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_WithAllBasicOptions(t *testing.T) {
	expectedBody := `{
		"prefix": "nir",
		"completion": {
			"field": "suggest",
			"size": 10,
			"skip_duplicates": true
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"song-suggest",
		"nir",
		"suggest",
		cs.WithSize(10),
		cs.WithCompletionSuggesterSkipDuplicates(true),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_BasicFuzzy(t *testing.T) {
	expectedBody := `{
		"prefix": "nir",
		"completion": {
			"field": "suggest",
			"fuzzy": {}
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"song-suggest",
		"nir",
		"suggest",
		cs.WithCompletionSuggesterFuzzy(),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_FuzzyWithAllOptions(t *testing.T) {
	expectedBody := `{
		"prefix": "nir",
		"completion": {
			"field": "suggest",
			"fuzzy": {
				"fuzziness": "AUTO",
				"transpositions": true,
				"min_length": 3,
				"prefix_length": 2,
				"unicode_aware": true
			}
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"song-suggest",
		"nir",
		"suggest",
		cs.WithCompletionSuggesterFuzzy(
			cs.WithFuzzyFuzziness("AUTO"),
			cs.WithFuzzyTranspositions(true),
			cs.WithFuzzyMinLength(3),
			cs.WithFuzzyPrefixLength(2),
			cs.WithFuzzyUnicodeAware(true),
		),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggesterRegex_AllRegexOptions(t *testing.T) {
	expectedBody := `{
		"regex": "n.*r",
		"completion": {
			"field": "suggest",
			"flags": "INTERSECTION|COMPLEMENT",
			"max_determinized_states": 12345
		}
	}`

	completionSuggesterResult := cs.CompletionSuggesterRegex(
		"song-suggest",
		"n.*r",
		"suggest",
		cs.WithRegexFlags(cs.RegexFlagIntersection+"|"+cs.RegexFlagComplement),
		cs.WithMaxDeterminizedStates(12345),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_WithMultipleCategoryContext(t *testing.T) {
	expectedBody := `{
		"prefix": "comp",
		"completion": {
			"field": "suggest_field",
			"contexts": {
				"categories": ["electronics", "laptops"]
			}
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"product-suggest",
		"comp",
		"suggest_field",
		cs.WithMultipleCategoryContext("categories", "electronics", "laptops"),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_WithCategoryContext(t *testing.T) {
	expectedBody := `{
		"prefix": "book",
		"completion": {
			"field": "suggest_field",
			"contexts": {
				"genre": {
					"context": "fiction"
				}
			}
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"product-suggest",
		"book",
		"suggest_field",
		cs.WithCategoryContext("genre", "fiction"),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_WithCategoryContextWithOptions(t *testing.T) {
	expectedBody := `{
		"prefix": "book",
		"completion": {
			"field": "suggest_field",
			"contexts": {
				"genre": {
					"context": "science_fiction",
					"boost": 2.5,
					"prefix": true
				}
			}
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"product-suggest",
		"book",
		"suggest_field",
		cs.WithCategoryContext(
			"genre",
			"science_fiction",
			cs.WithCategoryContextBoost(2.5),
			cs.WithCategoryContextPrefix(),
		),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_WithGeoContext(t *testing.T) {
	expectedBody := `{
		"prefix": "cafe",
		"completion": {
			"field": "location_suggest",
			"contexts": {
				"location": {
					"context": {
						"lat": 51.5074,
						"lon": 0.1278
					}
				}
			}
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"place-suggest",
		"cafe",
		"location_suggest",
		cs.WithGeoContext("location", 51.5074, 0.1278),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_WithGeoContextWithOptions(t *testing.T) {
	expectedBody := `{
		"prefix": "restaurant",
		"completion": {
			"field": "location_suggest",
			"contexts": {
				"office_nearby": {
					"context": {
						"lat": 40.7128,
						"lon": -74.0060
					},
					"boost": 3.0,
					"precision": "5km",
					"neighbours": ["brooklyn", "jersey_city"]
				}
			}
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"place-suggest",
		"restaurant",
		"location_suggest",
		cs.WithGeoContext(
			"office_nearby",
			40.7128, -74.0060,
			cs.WithGeoContextBoost(3.0),
			cs.WithGeoContextPrecision("5km"),
			cs.WithGeoContextNeighbours("brooklyn", "jersey_city"),
		),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_WithMixedContexts(t *testing.T) {
	expectedBody := `{
		"prefix": "test",
		"completion": {
			"field": "mixed_field",
			"contexts": {
				"categories": ["electronics", "software"],
				"city": {
					"context": "london"
				},
				"current_location": {
					"context": {
						"lat": 51.5,
						"lon": -0.1
					},
					"boost": 1.2
				}
			}
		}
	}`

	completionSuggesterResult := cs.CompletionSuggester(
		"mixed-suggest",
		"test",
		"mixed_field",
		cs.WithMultipleCategoryContext("categories", "electronics", "software"),
		cs.WithCategoryContext("city", "london"),
		cs.WithGeoContext("current_location", 51.5, -0.1, cs.WithGeoContextBoost(1.2)),
	)
	assert.Nil(t, completionSuggesterResult.Err)
	jsonBody, err := json.Marshal(completionSuggesterResult.Ok)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestCompletionSuggester_WithMultipleCategoryContext_NoContextsError(t *testing.T) {
	completionSuggesterResult := cs.CompletionSuggester(
		"product-suggest",
		"comp",
		"suggest_field",
		cs.WithMultipleCategoryContext("categories"),
	)
	assert.NotNil(t, completionSuggesterResult.Err)
	assert.Contains(t, completionSuggesterResult.Err.Error(), "contexts cannot be empty for MultipleCategoryClause")
}
