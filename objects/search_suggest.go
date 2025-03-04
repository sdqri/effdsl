package objects

import (
	"encoding/json"
)

type SuggestS struct {
	GlobalText string `json:"text,omitempty"` // To avoid repetition of the suggest text, it is possible to define a global text.
	Suggestions
}

func (s SuggestS) MarshalJSON() ([]byte, error) {
	suggestionsJson, err := json.Marshal(s.Suggestions)
	if err != nil {
		return nil, err
	}

	type suggestS struct {
		GlobalText string `json:"text,omitempty"`
	}
	suggestSJson, err := json.Marshal(suggestS{GlobalText: s.GlobalText})
	if err != nil {
		return nil, err
	}
	suggestionsJson[0] = ','
	return append(suggestSJson[:len(suggestSJson)-1], suggestionsJson...), nil
}

func (s SuggestS) SuggestInfo() string {
	return "Suggest"
}

// Suggests similar looking terms based on a provided text by using a suggester.
// [Suggest]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#search-suggesters
func Suggest(globalText string, s ...Suggestion) SuggestResult {
	suggestions := make(Suggestions, len(s))
	for i := range s {
		suggestions[s[i].GetName()] = s[i]
	}
	suggest := SuggestS{
		GlobalText:  globalText,
		Suggestions: suggestions,
	}

	return SuggestResult{
		Ok:  suggest,
		Err: nil,
	}
}

// ----------------------------------------------------

type Suggestions map[string]Suggestion

type Suggestion interface {
	GetName() string
}

// ----------------------------------------------------

type SuggestionS struct {
	Name string `json:"-"`              //
	Text string `json:"text,omitempty"` // The suggest text. The suggest text is a required option that needs to be set globally or per suggestion.
	Suggester
}

func (s SuggestionS) GetName() string {
	return s.Name
}

func (s SuggestionS) MarshalJSON() ([]byte, error) {
	suggesterJson, err := s.Suggester.MarshalJSON()
	if err != nil {
		return nil, err
	}

	type suggestionS struct {
		Text string `json:"text,omitempty"`
	}
	suggestionSJson, err := json.Marshal(suggestionS{Text: s.Text})
	if err != nil {
		return nil, err
	}
	suggesterJson[0] = ','
	return append(suggestionSJson[:len(suggestionSJson)-1], suggesterJson...), nil
}

type Suggester interface {
	json.Marshaler
	SuggesterInfo() string
}

// ----------------------------------------------------

// TermSuggesterS - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#term-suggester
type TermSuggesterS struct {
	Field       string      `json:"field"`              // The field to fetch the candidate suggestions from. This is a required option that either needs to be set globally or per suggestion.
	Analyzer    string      `json:"analyzer,omitempty"` // The analyzer to analyse the suggest text with. Defaults to the search analyzer of the suggest field.
	Size        int         `json:"size,omitempty"`     // The maximum corrections to be returned per suggest text token.
	Sort        SuggestSort `json:"sort,omitempty"`     // Defines how suggestions should be sorted per suggest text term. Two possible values: score, frequency
	SuggestMode SuggestMode `json:"suggest_mode,omitempty"`
}

func (t TermSuggesterS) SuggesterInfo() string {
	return "Term suggester"
}

func (t TermSuggesterS) MarshalJSON() ([]byte, error) {
	type TermSuggesterBase TermSuggesterS
	return json.Marshal(
		M{
			"term": (TermSuggesterBase)(t),
		},
	)
}

// SuggestSort - Defines how suggestions should be sorted per suggest text term.
type SuggestSort string

const (
	// SortScore - Sort by score first, then document frequency and then the term itself.
	SortScore SuggestSort = "score"
	// FrequencyScore - Sort by document frequency first, then similarity score and then the term itself.
	FrequencyScore SuggestSort = "frequency"
)

// SuggestMode - The suggest mode controls what suggestions are included or controls
// for what suggest text terms, suggestions should be suggested.
type SuggestMode string

const (
	// SuggestModeMissing - Only provide suggestions for suggest text terms that are not in the index (default).
	SuggestModeMissing SuggestMode = "score"
	// SuggestModePopular - Only suggest suggestions that occur in more docs than the original suggest text term.
	SuggestModePopular SuggestMode = "popular"
	// Suggest any matching suggestions based on terms in the suggest text.
	SuggestModeAlways SuggestMode = "always"
)

type TermSuggesterOption func(*TermSuggesterS)

func WithTermSuggestAnalyzer(analyzer string) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Analyzer = analyzer
	}
}

func WithTermSuggestSize(size int) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Size = size
	}
}

func WithTermSuggestSort(sort SuggestSort) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Sort = sort
	}
}

func WithTermSuggestMode(mode SuggestMode) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.SuggestMode = mode
	}
}

// The term suggester suggests terms based on edit distance. The provided suggest text is analyzed before terms are suggested. The suggested terms are provided per analyzed suggest text token. The term suggester doesn’t take the query into account that is part of request.
// [TermSuggester]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#term-suggester
func TermSuggester(name, text, field string, opts ...TermSuggesterOption) Suggestion {
	termSuggester := TermSuggesterS{
		Field: field,
	}
	for _, opt := range opts {
		opt(&termSuggester)
	}

	s := SuggestionS{
		Name:      name,
		Text:      text,
		Suggester: termSuggester,
	}
	return s
}
