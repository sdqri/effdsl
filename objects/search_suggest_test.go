package objects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Search_Suggest_MarshalJSON(t *testing.T) {
	s := Suggest("test",
		TermSuggester("my-suggestion-1", "tring out Elasticsearch", "message"),
		TermSuggester("my-suggestion-2", "tring out Elasticsearch", "message",
			WithTermSuggesterMode(SuggestModeAlways),
			WithTermSuggesterAnalyzer("test"),
			WithTermSuggesterSize(1),
			WithTermSuggesterSort(SortScore),
		),
		CompletionSuggester("my-suggestion-3", "tri", "message"),
	)

	body, err := s.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"text":"test","my-suggestion-1":{"text":"tring out Elasticsearch","term":{"field":"message"}},"my-suggestion-2":{"text":"tring out Elasticsearch","term":{"field":"message","analyzer":"test","size":1,"sort":"score","suggest_mode":"always"}},"my-suggestion-3":{"prefix":"tri","completion":{"field":"message"}}}`
	require.Equal(t, expected, string(body))

	s2 := CompletionSuggester("my-suggestion-4", "tri", "message",
		WithCompletionSuggesterSize(10),
		WithCompletionSuggesterSkipDuplicates(true),
		WithCompletionSuggesterFuzzy(
			WithFuzzyFuzziness("3"),
			WithFuzzyTranspositions(false),
			WithFuzzyMinLength(2),
			WithFuzzyPrefixLength(0),
			WithFuzzyUnicodeAware(true),
		),
	)
	suggestionS := s2.(SuggestionS)
	completionSuggestion, err := suggestionS.MarshalJSON()

	const completionExpected = `{"prefix":"tri","completion":{"field":"message","size":10,"skip_duplicates":true,"fuzzy":{"fuzziness":"3","transpositions":false,"min_length":2,"prefix_length":0,"unicode_aware":true}}}`
	require.Equal(t, completionExpected, string(completionSuggestion))
}
