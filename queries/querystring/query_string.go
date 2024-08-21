package querystring

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

type QueryStringS struct {
	Query                      string   `json:"query"`                                         // (Required, string) Query string you wish to parse and use for search.
	DefaultField               string   `json:"default_field,omitempty"`                       // (Optional, string) Default field to search if no field is provided in the query string.
	AllowLeadingWildcard       bool     `json:"allow_leading_wildcard,omitempty"`              // (Optional, Boolean) If true, wildcard characters * and ? are allowed as the first character. Defaults to true.
	AnalyzeWildcard            bool     `json:"analyze_wildcard,omitempty"`                    // (Optional, Boolean) If true, the query attempts to analyze wildcard terms in the query string. Defaults to false.
	Analyzer                   string   `json:"analyzer,omitempty"`                            // (Optional, string) Analyzer used to convert text in the query string into tokens.
	AutoGenerateSynonymsPhrase bool     `json:"auto_generate_synonyms_phrase_query,omitempty"` // (Optional, Boolean) If true, match phrase queries are automatically created for multi-term synonyms. Defaults to true.
	Boost                      float64  `json:"boost,omitempty"`                               // (Optional, float) Floating point number used to adjust relevance scores of the query.
	DefaultOperator            Operator `json:"default_operator,omitempty"`                    // (Optional, string) Default boolean logic used to interpret text in the query string.
	EnablePositionIncrements   bool     `json:"enable_position_increments,omitempty"`          // (Optional, Boolean) If true, enable position increments in queries constructed from a query string search.
	Fields                     []string `json:"fields,omitempty"`                              // (Optional, array of strings) Array of fields to search. Supports wildcards (*).
	Fuzziness                  string   `json:"fuzziness,omitempty"`                           // (Optional, string) Maximum edit distance allowed for fuzzy matching.
	FuzzyMaxExpansions         int      `json:"fuzzy_max_expansions,omitempty"`                // (Optional, integer) Maximum number of terms for fuzzy matching expansion.
	FuzzyPrefixLength          int      `json:"fuzzy_prefix_length,omitempty"`                 // (Optional, integer) Number of beginning characters left unchanged for fuzzy matching.
	FuzzyTranspositions        bool     `json:"fuzzy_transpositions,omitempty"`                // (Optional, Boolean) If true, edits for fuzzy matching include transpositions of adjacent characters.
	Lenient                    bool     `json:"lenient,omitempty"`                             // (Optional, Boolean) If true, format-based errors are ignored.
	MaxDeterminizedStates      int      `json:"max_determinized_states,omitempty"`             // (Optional, integer) Maximum number of automaton states required for the query.
	MinimumShouldMatch         string   `json:"minimum_should_match,omitempty"`                // (Optional, string) Minimum number of clauses that must match for a document to be returned.
	QuoteAnalyzer              string   `json:"quote_analyzer,omitempty"`                      // (Optional, string) Analyzer used to convert quoted text in the query string into tokens.
	PhraseSlop                 int      `json:"phrase_slop,omitempty"`                         // (Optional, integer) Maximum number of positions allowed between matching tokens for phrases.
	QuoteFieldSuffix           string   `json:"quote_field_suffix,omitempty"`                  // (Optional, string) Suffix appended to quoted text in the query string.
	Rewrite                    Rewrite  `json:"rewrite,omitempty"`                             // (Optional, string) Method used to rewrite the query.
	TimeZone                   string   `json:"time_zone,omitempty"`                           // (Optional, string) UTC offset or IANA time zone used to convert date values in the query string to UTC.
}

func (bq QueryStringS) QueryInfo() string {
	return "Query string query"
}

func (qs QueryStringS) MarshalJSON() ([]byte, error) {
	type QueryStringBase QueryStringS
	return json.Marshal(
		map[string]any{
			"query_string": (QueryStringBase)(qs),
		},
	)
}

type QueryStringOption func(*QueryStringS)

func WithDefaultField(defaultField string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.DefaultField = defaultField
	}
}

func WithAllowLeadingWildcard() QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.AllowLeadingWildcard = true
	}
}

func WithAnalyzeWildcard() QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.AnalyzeWildcard = true
	}
}

func WithAnalyzer(analyzer string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.Analyzer = analyzer
	}
}

func WithAutoGenerateSynonymsPhrase(autoGenerate bool) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.AutoGenerateSynonymsPhrase = autoGenerate
	}
}

func WithBoost(boost float64) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.Boost = boost
	}
}

type Operator string

const (
	OR  Operator = "OR"
	AND Operator = "AND"
)

func WithDefaultOperator(op Operator) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.DefaultOperator = op
	}
}

func WithEnablePositionIncrements(enable bool) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.EnablePositionIncrements = enable
	}
}

func WithFields(fields ...string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.Fields = fields
	}
}

func WithFuzziness(fuzziness string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.Fuzziness = fuzziness
	}
}

func WithFuzzyMaxExpansions(expansions int) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.FuzzyMaxExpansions = expansions
	}
}

func WithFuzzyPrefixLength(length int) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.FuzzyPrefixLength = length
	}
}

func WithFuzzyTranspositions(transpositions bool) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.FuzzyTranspositions = transpositions
	}
}

func WithLenient(lenient bool) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.Lenient = lenient
	}
}

func WithMaxDeterminizedStates(states int) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.MaxDeterminizedStates = states
	}
}

func WithMinimumShouldMatch(minimum string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.MinimumShouldMatch = minimum
	}
}

func WithQuoteAnalyzer(quoteAnalyzer string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.QuoteAnalyzer = quoteAnalyzer
	}
}

func WithPhraseSlop(slop int) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.PhraseSlop = slop
	}
}

func WithQuoteFieldSuffix(suffix string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.QuoteFieldSuffix = suffix
	}
}

// Rewrite represents the type of rewrite to use in a fuzzy query.
type Rewrite string

const (
	ConstantScoreBlended  Rewrite = "constant_score_blended"
	ConstantScore         Rewrite = "constant_score"
	ConstantScoreBoolean  Rewrite = "constant_score_boolean"
	ScoringBoolean        Rewrite = "scoring_boolean"
	TopTermsBlendedFreqsN Rewrite = "top_terms_blended_freqs_N"
	TopTermsBoostN        Rewrite = "top_terms_boost_N"
	TopTermsN             Rewrite = "top_terms_N"
)

func WithRewrite(rewrite Rewrite) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.Rewrite = rewrite
	}
}

func WithTimeZone(timeZone string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.TimeZone = timeZone
	}
}

// Returns documents based on a provided query string, using a parser with a strict syntax
// [Query string query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html
func QueryString(query string, opts ...QueryStringOption) effdsl.QueryResult {
	queryString := QueryStringS{
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
