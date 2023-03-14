package objects

import (
	"encoding/json"
)

type TermSetQueryS struct {
	Field                    string   `json:"-"`                                     //(Required, object) Field you wish to search.
	Terms                    []string `json:"terms"`                                 //(Required, []string) An array of terms you wish to find in the provided <field>. To return a document, at least one of the terms must exactly match the field value, including whitespace and capitalization.
	MinimumShouldMatchField  string   `json:"minimum_should_match_field,omitempty"`  //(Optional, string) The field which holds the minimum number of terms that should match. Only used when `minimum_should_match_script` is not set.
	MinimumShouldMatchScript string   `json:"minimum_should_match_script,omitempty"` //(Optional, string) Script which returns the minimum number of terms that should match.
}

func (tsq TermSetQueryS) QueryInfo() string {
	return "Term set query"
}

func (tsq TermSetQueryS) MarshalJSON() ([]byte, error) {
	type TermSetQueryBase TermSetQueryS
	return json.Marshal(
		M{
			"terms_set": M{
				tsq.Field: (TermSetQueryBase)(tsq),
			},
		},
	)
}

type TermSetQueryOption func(*TermSetQueryS)

func (f DefineType) WithMinimumShouldMatchField(field string) TermSetQueryOption {
	return func(termSetQuery *TermSetQueryS) {
		termSetQuery.MinimumShouldMatchField = field
	}
}

func (f DefineType) WithMinimumShouldMatchScript(script string) TermSetQueryOption {
	return func(termSetQuery *TermSetQueryS) {
		termSetQuery.MinimumShouldMatchScript = script
	}
}

// Returns documents that contain at least one of the specified terms in a provided field.
// [Terms Set query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-set-query.html
func (f DefineType) TermsSetQuery(field string, terms []string, opts ...TermSetQueryOption) QueryResult {
	termSetQuery := TermSetQueryS{
		Field: field,
		Terms: terms,
	}
	for _, opt := range opts {
		opt(&termSetQuery)
	}
	return QueryResult{
		Ok:  termSetQuery,
		Err: nil,
	}
}
