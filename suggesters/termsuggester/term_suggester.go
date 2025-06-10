package termsuggester

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

// TermSuggesterS - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#term-suggester
type TermSuggesterS struct {
	Name        string             `json:"-"`                  // Required
	Text        string             `json:"-"`                  // Required: Text to generate suggestions for.
	Field       string             `json:"field"`              // Required: The field to fetch the candidate suggestions from. This is a required option that either needs to be set globally or per suggestion.
	Analyzer    string             `json:"analyzer,omitempty"` // The analyzer to analyse the suggest text with. Defaults to the search analyzer of the suggest field.
	Size        *uint64            `json:"size,omitempty"`     // The maximum corrections to be returned per suggest text token.
	Sort        *TermSuggestSort   `json:"sort,omitempty"`     // Defines how suggestions should be sorted per suggest text term. Two possible values: score, frequency
	SuggestMode *TermSuggesterMode `json:"suggest_mode,omitempty"`
}

func (ts TermSuggesterS) SuggestName() string {
	return ts.Name
}

func (ts TermSuggesterS) MarshalJSON() ([]byte, error) {
	type TermSuggesterBase TermSuggesterS
	return json.Marshal(effdsl.M{
		"text": ts.Text,
		"term": (TermSuggesterBase)(ts),
	},
	)
}

// TermSuggestSort - Defines how suggestions should be sorted per suggest text term.
type TermSuggestSort string

const (
	// Sort - Sort by score first, then document frequency and then the term itself.
	ByScore TermSuggestSort = "score"
	// FrequencyScore - Sort by document frequency first, then similarity score and then the term itself.
	ByFrequency TermSuggestSort = "frequency"
)

// TermSuggesterMode - The suggest mode controls what suggestions are included or controls
// for what suggest text terms, suggestions should be suggested.
type TermSuggesterMode string

const (
	// SuggestModeMissing - Only provide suggestions for suggest text terms that are not in the index (default).
	Missing TermSuggesterMode = "score"
	// SuggestModePopular - Only suggest suggestions that occur in more docs than the original suggest text term.
	Popular TermSuggesterMode = "popular"
	// Suggest any matching suggestions based on terms in the suggest text.
	Always TermSuggesterMode = "always"
)

type TermSuggesterOption func(*TermSuggesterS)

func WithAnalyzer(analyzer string) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Analyzer = analyzer
	}
}

func WithSize(size uint64) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Size = &size
	}
}

func WithSort(sort TermSuggestSort) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Sort = &sort
	}
}

func WithMode(mode TermSuggesterMode) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.SuggestMode = &mode
	}
}

// The term suggester suggests terms based on edit distance. The provided suggest text is analyzed before terms are suggested. The suggested terms are provided per analyzed suggest text token. The term suggester doesn’t take the query into account that is part of request.
// [TermSuggester]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#term-suggester
func TermSuggester(suggestName, text, field string, opts ...TermSuggesterOption) effdsl.SuggestResult {
	ts := TermSuggesterS{
		Name:  suggestName,
		Text:  text,
		Field: field,
	}

	for _, opt := range opts {
		opt(&ts)
	}

	return effdsl.SuggestResult{
		Ok:  ts,
		Err: nil,
	}
}
