package simplequerystring

import (
	"encoding/json"

	"github.com/sdqri/effdsl"
)

type Simple struct {
	Query                      string   `json:"query"`                                         // (Required, string) Query string you wish to parse and use for search.
	Fields                     []string `json:"fields,omitempty"`                              // (Optional, array of strings) Array of fields to search. Supports wildcards (*).
	DefaultOperator            Operator `json:"default_operator,omitempty"`                    // (Optional, string) Default boolean logic used to interpret text in the query string.
	AnalyzeWildcard            bool     `json:"analyze_wildcard,omitempty"`                    // (Optional, Boolean) If true, the query attempts to analyze wildcard terms in the query string. Defaults to false.
	Analyzer                   string   `json:"analyzer,omitempty"`                            // (Optional, string) Analyzer used to convert text in the query string into tokens.
	AutoGenerateSynonymsPhrase bool     `json:"auto_generate_synonyms_phrase_query,omitempty"` // (Optional, Boolean) If true, match phrase queries are automatically created for multi-term synonyms. Defaults to true.
	Flags                      string   `json:"flags,omitempty"`                               //(Optional, string) List of enabled operators for the simple query string syntax. Defaults to ALL (all operators). See Limit operators for valid values.
	FuzzyMaxExpansions         int      `json:"fuzzy_max_expansions,omitempty"`                // (Optional, integer) Maximum number of terms for fuzzy matching expansion.
	FuzzyPrefixLength          int      `json:"fuzzy_prefix_length,omitempty"`                 // (Optional, integer) Number of beginning characters left unchanged for fuzzy matching.
	FuzzyTranspositions        bool     `json:"fuzzy_transpositions,omitempty"`                // (Optional, Boolean) If true, edits for fuzzy matching include transpositions of adjacent characters.
	Lenient                    bool     `json:"lenient,omitempty"`                             // (Optional, Boolean) If true, format-based errors are ignored.
	MinimumShouldMatch         string   `json:"minimum_should_match,omitempty"`                // (Optional, string) Minimum number of clauses that must match for a document to be returned.
	QuoteFieldSuffix           string   `json:"quote_field_suffix,omitempty"`                  // (Optional, string) Suffix appended to quoted text in the query string.
}

func (bq Simple) QueryInfo() string {
	return "Simple query string query"
}

func (qs Simple) MarshalJSON() ([]byte, error) {
	type QueryStringBase Simple
	return json.Marshal(
		map[string]any{
			"simple_query_string": (QueryStringBase)(qs),
		},
	)
}

type QueryStringOption func(*Simple)

func WithFields(fields ...string) QueryStringOption {
	return func(queryString *Simple) {
		queryString.Fields = fields
	}
}

type Operator string

const (
	OR  Operator = "OR"
	AND Operator = "AND"
)

func WithDefaultOperator(op Operator) QueryStringOption {
	return func(queryString *Simple) {
		queryString.DefaultOperator = op
	}
}

func WithAnalyzeWildcard() QueryStringOption {
	return func(queryString *Simple) {
		queryString.AnalyzeWildcard = true
	}
}

func WithAnalyzer(analyzer string) QueryStringOption {
	return func(queryString *Simple) {
		queryString.Analyzer = analyzer
	}
}

func WithAutoGenerateSynonymsPhrase(autoGenerate bool) QueryStringOption {
	return func(queryString *Simple) {
		queryString.AutoGenerateSynonymsPhrase = autoGenerate
	}
}

func WithFlags(flags string) QueryStringOption {
	return func(queryString *Simple) {
		queryString.Flags = flags
	}
}

func WithFuzzyMaxExpansions(expansions int) QueryStringOption {
	return func(queryString *Simple) {
		queryString.FuzzyMaxExpansions = expansions
	}
}

func WithFuzzyPrefixLength(length int) QueryStringOption {
	return func(queryString *Simple) {
		queryString.FuzzyPrefixLength = length
	}
}

func WithFuzzyTranspositions(transpositions bool) QueryStringOption {
	return func(queryString *Simple) {
		queryString.FuzzyTranspositions = transpositions
	}
}

func WithLenient(lenient bool) QueryStringOption {
	return func(queryString *Simple) {
		queryString.Lenient = lenient
	}
}

func WithMinimumShouldMatch(minimum string) QueryStringOption {
	return func(queryString *Simple) {
		queryString.MinimumShouldMatch = minimum
	}
}

func WithQuoteFieldSuffix(suffix string) QueryStringOption {
	return func(queryString *Simple) {
		queryString.QuoteFieldSuffix = suffix
	}
}

// Returns documents based on a provided query string, using a parser with a strict syntax
// [Query string query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html
func SimpleQueryString(query string, opts ...QueryStringOption) effdsl.QueryResult {
	queryString := Simple{
		Query: query,
	}
	for _, opt := range opts {
		opt(&queryString)
	}
	return effdsl.QueryResult{
		Ok:  queryString,
		Err: nil,
	}
}
