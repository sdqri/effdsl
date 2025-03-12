package effdsl_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sdqri/effdsl/v2"
)

func Test_Search_Suggest_MarshalJSON(t *testing.T) {
	s := effdsl.Suggest("test",
		effdsl.TermSuggester("my-suggestion-1", "tring out Elasticsearch", "message"),
		effdsl.TermSuggester("my-suggestion-2", "tring out Elasticsearch", "message",
			effdsl.WithTermSuggesterMode(effdsl.SuggestModeAlways),
			effdsl.WithTermSuggesterAnalyzer("test"),
			effdsl.WithTermSuggesterSize(1),
			effdsl.WithTermSuggesterSort(effdsl.SortScore),
		),
		effdsl.CompletionSuggester("my-suggestion-3", "tri", "message"),
	)

	body, err := s.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"text":"test","my-suggestion-1":{"text":"tring out Elasticsearch","term":{"field":"message"}},"my-suggestion-2":{"text":"tring out Elasticsearch","term":{"field":"message","analyzer":"test","size":1,"sort":"score","suggest_mode":"always"}},"my-suggestion-3":{"prefix":"tri","completion":{"field":"message"}}}`
	require.Equal(t, expected, string(body))

	s2 := effdsl.CompletionSuggester("my-suggestion-4", "tri", "message",
		effdsl.WithCompletionSuggesterSize(10),
		effdsl.WithCompletionSuggesterSkipDuplicates(true),
		effdsl.WithCompletionSuggesterFuzzy(
			effdsl.WithFuzzyFuzziness("3"),
			effdsl.WithFuzzyTranspositions(false),
			effdsl.WithFuzzyMinLength(2),
			effdsl.WithFuzzyPrefixLength(0),
			effdsl.WithFuzzyUnicodeAware(true),
		),
	)
	suggestionS := s2.(effdsl.SuggestionS)
	completionSuggestion, err := suggestionS.MarshalJSON()

	const completionExpected = `{"prefix":"tri","completion":{"field":"message","size":10,"skip_duplicates":true,"fuzzy":{"fuzziness":"3","transpositions":false,"min_length":2,"prefix_length":0,"unicode_aware":true}}}`
	require.Equal(t, completionExpected, string(completionSuggestion))
}
