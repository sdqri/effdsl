package objects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Search_Suggest_MarshalJSON(t *testing.T) {
	s := Suggest("test",
		TermSuggester("my-suggestion-1", "tring out Elasticsearch", "message"),
		TermSuggester("my-suggestion-2", "tring out Elasticsearch", "message",
			WithTermSuggestMode(SuggestModeAlways),
			WithTermSuggestAnalyzer("test"),
			WithTermSuggestSize(1),
			WithTermSuggestSort(SortScore),
		))

	body, err := s.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"text":"test","my-suggestion-1":{"text":"tring out Elasticsearch","term":{"field":"message"}},"my-suggestion-2":{"text":"tring out Elasticsearch","term":{"field":"message","analyzer":"test","size":1,"sort":"score","suggest_mode":"always"}}}`
	require.Equal(t, expected, string(body))
}
