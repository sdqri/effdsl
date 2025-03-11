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
	Name   string `json:"-"`              //
	Text   string `json:"text,omitempty"` // The suggest text. The suggest text is a required option that needs to be set globally or per suggestion.
	Prefix string `json:"prefix,omitempty"`
	Regex  string `json:"regex,omitempty"` // The completion suggester also supports regex queries meaning you can express a prefix as a regular expression
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
		Text   string `json:"text,omitempty"`
		Prefix string `json:"prefix,omitempty"`
		Regex  string `json:"regex,omitempty"`
	}
	suggestionSJson, err := json.Marshal(suggestionS{
		Text:   s.Text,
		Prefix: s.Prefix,
		Regex:  s.Regex,
	})
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
	Size        *uint64     `json:"size,omitempty"`     // The maximum corrections to be returned per suggest text token.
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

func WithTermSuggesterAnalyzer(analyzer string) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Analyzer = analyzer
	}
}

func WithTermSuggesterSize(size uint64) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Size = &size
	}
}

func WithTermSuggesterSort(sort SuggestSort) TermSuggesterOption {
	return func(termSuggest *TermSuggesterS) {
		termSuggest.Sort = sort
	}
}

func WithTermSuggesterMode(mode SuggestMode) TermSuggesterOption {
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

// ----------------------------------------------------

// CompletionSuggesterS - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
type CompletionSuggesterS struct {
	Field          string  `json:"field"`                     // The name of the field on which to run the query (required).
	Size           *uint64 `json:"size,omitempty"`            // The number of suggestions to return (defaults to 5).
	SkipDuplicates bool    `json:"skip_duplicates,omitempty"` // Whether duplicate suggestions should be filtered out (defaults to false).
	Fuzzy          *FuzzyS `json:"fuzzy,omitempty"`           // Fuzzy options for the completion suggester
	// TODO: add "contexts" field
}

func (c CompletionSuggesterS) MarshalJSON() ([]byte, error) {
	type CompletionSuggesterSBase CompletionSuggesterS
	return json.Marshal(
		M{
			"completion": (CompletionSuggesterSBase)(c),
		},
	)
}

func (c CompletionSuggesterS) SuggesterInfo() string {
	return "Completion suggester"
}

type CompletionSuggesterOption func(*CompletionSuggesterS)

func WithCompletionSuggesterSize(size uint64) CompletionSuggesterOption {
	return func(c *CompletionSuggesterS) {
		c.Size = &size
	}
}

func WithCompletionSuggesterSkipDuplicates(skipDuplicates bool) CompletionSuggesterOption {
	return func(c *CompletionSuggesterS) {
		c.SkipDuplicates = skipDuplicates
	}
}

func WithCompletionSuggesterFuzzy(opt ...FuzzyOption) CompletionSuggesterOption {
	fuzzy := completionSuggesterFuzzy(opt...)
	return func(c *CompletionSuggesterS) {
		c.Fuzzy = &fuzzy
	}
}

// The completion suggester provides auto-complete/search-as-you-type functionality. This is a navigational feature to guide users to relevant results as they are typing, improving search precision. It is not meant for spell correction or did-you-mean functionality like the term or phrase suggesters.
// [CompletionSuggester]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
func CompletionSuggester(name, prefix, field string, opts ...CompletionSuggesterOption) Suggestion {
	completionSuggester := CompletionSuggesterS{
		Field: field,
	}
	for _, opt := range opts {
		opt(&completionSuggester)
	}

	s := SuggestionS{
		Name:      name,
		Prefix:    prefix,
		Suggester: completionSuggester,
	}
	return s
}

// The completion suggester also supports regex queries meaning you can express a prefix as a regular expression.
// [CompletionSuggesterRegex]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#regex
func CompletionSuggesterRegex(name, regex, field string, opts ...CompletionSuggesterOption) Suggestion {
	completionSuggester := CompletionSuggesterS{
		Field: field,
	}
	for _, opt := range opts {
		opt(&completionSuggester)
	}

	s := SuggestionS{
		Name:      name,
		Regex:     regex,
		Suggester: completionSuggester,
	}
	return s
}

// ----------------------------------------------------

type FuzzyS struct {
	Fuzziness      Fuzziness `json:"fuzziness,omitempty"`      // The fuzziness factor, defaults to AUTO.
	Transpositions *bool     `json:"transpositions,omitempty"` // If set to true, transpositions are counted as one change instead of two, defaults to true.
	MinLength      *uint64   `json:"min_length,omitempty"`     // Minimum length of the input before fuzzy suggestions are returned, defaults 3.
	PrefixLength   *uint64   `json:"prefix_length,omitempty"`  // Minimum length of the input, which is not checked for fuzzy alternatives, defaults to 1.
	UnicodeAware   bool      `json:"unicode_aware,omitempty"`  // If true, all measurements (like fuzzy edit distance, transpositions, and lengths) are measured in Unicode code points instead of in bytes. This is slightly slower than raw bytes, so it is set to false by default.
}

type FuzzyOption func(*FuzzyS)

func WithFuzzyFuzziness(fuzziness Fuzziness) FuzzyOption {
	return func(f *FuzzyS) {
		f.Fuzziness = fuzziness
	}
}

func WithFuzzyTranspositions(transpositions bool) FuzzyOption {
	return func(f *FuzzyS) {
		f.Transpositions = &transpositions
	}
}

func WithFuzzyMinLength(minLength uint64) FuzzyOption {
	return func(f *FuzzyS) {
		f.MinLength = &minLength
	}
}

func WithFuzzyPrefixLength(prefixLength uint64) FuzzyOption {
	return func(f *FuzzyS) {
		f.PrefixLength = &prefixLength
	}
}

func WithFuzzyUnicodeAware(unicodeAware bool) FuzzyOption {
	return func(f *FuzzyS) {
		f.UnicodeAware = unicodeAware
	}
}

func completionSuggesterFuzzy(opt ...FuzzyOption) FuzzyS {
	fuzzy := FuzzyS{}
	for _, o := range opt {
		o(&fuzzy)
	}
	return fuzzy
}
