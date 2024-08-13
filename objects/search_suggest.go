package objects

import "encoding/json"

type SuggestS struct {
	GlobalText string `json:"text,omitempty"` // To avoid repetition of the suggest text, it is possible to define a global text.
	Suggester
}

func (s SuggestS) MarshalJSON() ([]byte, error) {
	type SuggesterBase SuggestS
	return json.Marshal(
		M{
			"suggest": (SuggesterBase)(s),
		},
	)
}

func (s SuggestS) SuggestInfo() string {
	return "Suggest"
}

// Suggests similar looking terms based on a provided text by using a suggester.
// [Suggesters]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#search-suggesters
func Suggesters(globalText string, s Suggester) SuggestResult {
	sugest := SuggestS{
		GlobalText: globalText,
		Suggester:  s,
	}

	return SuggestResult{
		Ok:  sugest,
		Err: nil,
	}
}

type Suggester interface {
	json.Marshaler
	_type() string
}

// TermSuggester - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#term-suggester
type TermsSuggest []TermSuggest

func (ts TermsSuggest) MarshalJSON() ([]byte, error) {
	type TermSuggesterBase TermSuggest

	terms := make(M, len(ts))
	for i := range ts {
		terms[ts[i].GetName()] = (TermSuggesterBase)(ts[i])
	}

	return json.Marshal(terms)
}

func (ts TermsSuggest) _type() string {
	return "term suggester"
}

type TermSuggesterS struct {
	Name string `json:"-"`              //
	Text string `json:"text,omitempty"` // The suggest text. The suggest text is a required option that needs to be set globally or per suggestion.
	Term TermS  `json:"term"`           // (Required) The term suggester suggests terms based on edit distance. The provided suggest text is analyzed before terms are suggested. The suggested terms are provided per analyzed suggest text token. The term suggester doesn’t take the query into account that is part of request.
}

func (t TermSuggesterS) GetName() string {
	return t.Name
}

type TermS struct {
	Field       string      `json:"field"`              // The field to fetch the candidate suggestions from. This is a required option that either needs to be set globally or per suggestion.
	Analyzer    string      `json:"analyzer,omitempty"` // The analyzer to analyse the suggest text with. Defaults to the search analyzer of the suggest field.
	Size        int         `json:"size,omitempty"`     // The maximum corrections to be returned per suggest text token.
	Sort        SuggestSort `json:"sort,omitempty"`     // Defines how suggestions should be sorted per suggest text term. Two possible values: score, frequency
	SuggestMode SuggestMode `json:"suggest_mode,omitempty"`
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

type TermSuggestOption func(*TermS)

func WithTermSuggestAnalyzer(analyzer string) TermSuggestOption {
	return func(termSuggest *TermS) {
		termSuggest.Analyzer = analyzer
	}
}

func WithTermSuggestSize(size int) TermSuggestOption {
	return func(termSuggest *TermS) {
		termSuggest.Size = size
	}
}

func WithTermSuggestSort(sort SuggestSort) TermSuggestOption {
	return func(termSuggest *TermS) {
		termSuggest.Sort = sort
	}
}

func WithTermSuggestMode(mode SuggestMode) TermSuggestOption {
	return func(termSuggest *TermS) {
		termSuggest.SuggestMode = mode
	}
}

func Term(name, text, field string, opts ...TermSuggestOption) TermSuggest {
	s := TermSuggesterS{
		Name: name,
		Text: text,
		Term: TermS{
			Field: field,
		},
	}
	for _, opt := range opts {
		opt(&s.Term)
	}
	return s
}

type TermSuggest interface {
	GetName() string
}

// TermSuggester - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#term-suggester
func TermSuggester(terms ...TermSuggest) Suggester {
	v := make(TermsSuggest, 0, len(terms))
	for i := range terms {
		v = append(v, terms[i])
	}
	return v
}
