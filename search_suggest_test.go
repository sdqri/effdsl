package effdsl_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	effdsl "github.com/sdqri/effdsl"
)

func Test_Search_Suggest_MarshalJSON(t *testing.T) {
	s := effdsl.Suggesters("test", effdsl.TermSuggester(
		effdsl.Term("my-suggestion-1", "tring out Elasticsearch", "message"),
		effdsl.Term("my-suggestion-2", "tring out Elasticsearch", "message",
			effdsl.WithTermSuggestMode(effdsl.SuggestModeAlways),
			effdsl.WithTermSuggestAnalyzer("test"),
			effdsl.WithTermSuggestSize(1),
			effdsl.WithTermSuggestSort(effdsl.SortScore),
		),
	))

	body, err := s.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"suggest":{"my-suggestion-1":{"text":"tring out Elasticsearch","term":{"field":"message"}},"my-suggestion-2":{"text":"tring out Elasticsearch","term":{"field":"message","analyzer":"test","size":1,"sort":"score","suggest_mode":"always"}}}}`
	require.Equal(t, expected, string(body))
}
