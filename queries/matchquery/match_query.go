package matchquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl"
)

type MatchQueryS struct {
	Field                      string   `json:"-"`                                             // (Required, object) Field you wish to search.
	Query                      string   `json:"query"`                                         // (Required) Text, number, boolean value or date you wish to find in the provided <field>.
	Analyzer                   string   `json:"analyzer,omitempty"`                            // (Optional, string) Analyzer used to convert the text in the query value into tokens.
	AutoGenerateSynonymsPhrase *bool    `json:"auto_generate_synonyms_phrase_query,omitempty"` // (Optional, Boolean) If true, match phrase queries are automatically created for multi-term synonyms.
	Boost                      *float64 `json:"boost,omitempty"`                               // (Optional, float) Floating point number used to decrease or increase the relevance scores of the query.
	Fuzziness                  string   `json:"fuzziness,omitempty"`                           // (Optional, string) Maximum edit distance allowed for matching.
	MaxExpansions              *int     `json:"max_expansions,omitempty"`                      // (Optional, integer) Maximum number of terms to which the query will expand.
	PrefixLength               *int     `json:"prefix_length,omitempty"`                       // (Optional, integer) Number of beginning characters left unchanged for fuzzy matching.
	FuzzyTranspositions        *bool    `json:"fuzzy_transpositions,omitempty"`                // (Optional, Boolean) If true, edits for fuzzy matching include transpositions of two adjacent characters.
	FuzzyRewrite               string   `json:"fuzzy_rewrite,omitempty"`                       // (Optional, string) Method used to rewrite the query.
	Lenient                    *bool    `json:"lenient,omitempty"`                             // (Optional, Boolean) If true, format-based errors are ignored.
	Operator                   string   `json:"operator,omitempty"`                            // (Optional, string) Boolean logic used to interpret text in the query value.
	MinimumShouldMatch         string   `json:"minimum_should_match,omitempty"`                // (Optional, string) Minimum number of clauses that must match for a document to be returned.
	ZeroTermsQuery             string   `json:"zero_terms_query,omitempty"`                    // (Optional, string) Indicates whether no documents are returned if the analyzer removes all tokens, such as when using a stop filter.
}

func (mq MatchQueryS) QueryInfo() string {
	return "Match query"
}

func (mq MatchQueryS) MarshalJSON() ([]byte, error) {
	type MatchQueryBase MatchQueryS
	return json.Marshal(
		effdsl.M{
			"match": effdsl.M{
				mq.Field: (MatchQueryBase)(mq),
			},
		},
	)
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html#match-field-params
type matchQueryFieldParameters struct {
	/*
		(Optional, string) Analyzer used to convert the text in the query value into tokens. Defaults to the index-time analyzer mapped for the <field>. If no analyzer is mapped, the index’s default analyzer is used.
	*/
	Analyzer string
	/*
		(Optional, Boolean) If true, match phrase queries are automatically created for multi-term synonyms. Defaults to true.
	*/
	AutoGenerateSynonymsPhrase *bool
	/*
		(Optional, float) Floating point number used to decrease or increase the relevance scores of the query. Defaults to 1.0.
		Boost values are relative to the default value of 1.0. A boost value between 0 and 1.0 decreases the relevance score. A value greater than 1.0 increases the relevance score.
	*/
	Boost *float64
	/*
		(Optional, string) Maximum edit distance allowed for matching. See [Fuzziness](https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness) for valid values and more information.
	*/
	Fuzziness string
	/*
		(Optional, integer) Maximum number of terms to which the query will expand. Defaults to 50.
	*/
	MaxExpansions *int
	/*
		(Optional, integer) Number of beginning characters left unchanged for fuzzy matching. Defaults to 0.
	*/
	PrefixLength *int
	/*
		(Optional, Boolean) If true, edits for fuzzy matching include transpositions of two adjacent characters (ab → ba). Defaults to true.
	*/
	FuzzyTranspositions *bool
	/*
		(Optional, string) Method used to rewrite the query. See the rewrite parameter for valid values and more information.
		If the fuzziness parameter is not 0, the match query uses a fuzzy_rewrite method of top_terms_blended_freqs_${max_expansions} by default.
	*/
	FuzzyRewrite FuzzyRewrite
	/*
		(Optional, Boolean) If true, format-based errors, such as providing a text query value for a numeric field, are ignored. Defaults to false.
	*/
	Lenient *bool
	/*
		(Optional, string) Boolean logic used to interpret text in the query value. Valid values are:
		* OR (Default)
			* For example, a query value of capital of Hungary is interpreted as capital OR of OR Hungary.
		* AND
			* For example, a query value of capital of Hungary is interpreted as capital AND of AND Hungary.
	*/
	Operator Operator
	/*
		(Optional, string) Minimum number of clauses that must match for a document to be returned. See the minimum_should_match parameter for valid values and more information.
	*/
	MinimumShouldMatch string
	/*
		(Optional, string) Indicates whether no documents are returned if the analyzer removes all tokens, such as when using a stop filter. Valid values are:
		* none (Default)
			* No documents are returned if the analyzer removes all tokens.
		* all
			* Returns all documents, similar to a match_all query.
	*/
	ZeroTermsQuery ZeroTerms
}

type MatchQueryFieldParameter func(params *matchQueryFieldParameters)

func WithAnalyzer(analyzer string) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.Analyzer = analyzer
	}
}

func WithAutoGenerateSynonymsPhrase(auto bool) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.AutoGenerateSynonymsPhrase = &auto
	}
}

func WithBoost(boost float64) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.Boost = &boost
	}
}

const FuzzinessAUTO string = "AUTO"

func WithFuzzinessParameter(f string) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.Fuzziness = f
	}
}

func WithMaxExpansions(maxExpansions int) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.MaxExpansions = &maxExpansions
	}
}

func WithPrefixLength(prefixLength int) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.PrefixLength = &prefixLength
	}
}

func WithFuzzyTranspositions(fuzzyTranspositions bool) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.FuzzyTranspositions = &fuzzyTranspositions
	}
}

// Rewrite represents the type of rewrite to use in a fuzzy query.
type FuzzyRewrite string

const (
	ConstantScore         FuzzyRewrite = "constant_score"            // Query is rewritten to a constant score query.
	ScoringBoolean        FuzzyRewrite = "scoring_boolean"           // Query is rewritten to a scoring boolean query.
	ConstantScoreBoolean  FuzzyRewrite = "constant_score_boolean"    // Query is rewritten to a constant score boolean query.
	TopTermsN             FuzzyRewrite = "top_terms_N"               // Query is rewritten to match the top N scoring terms.
	TopTermsBoostN        FuzzyRewrite = "top_terms_boost_N"         // Query is rewritten to match the top N scoring terms with boosting.
	TopTermsBlendedFreqsN FuzzyRewrite = "top_terms_blended_freqs_N" // Query is rewritten to match the top N scoring terms with blended frequencies.
)

func WithFuzzyRewrite(rewrite FuzzyRewrite) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.FuzzyRewrite = rewrite
	}
}

func WithLenient() MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		lv := true
		params.Lenient = &lv
	}
}

type Operator string

const (
	OR  Operator = "OR"
	AND Operator = "AND"
)

func WithOperator(op Operator) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.Operator = op
	}
}

func MinimumShouldMatch(minimumShouldMatch string) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.MinimumShouldMatch = minimumShouldMatch
	}
}

type ZeroTerms string

const (
	None ZeroTerms = "none"
	All  ZeroTerms = "all"
)

func WithZeroTermsquery(zt ZeroTerms) MatchQueryFieldParameter {
	return func(params *matchQueryFieldParameters) {
		params.ZeroTermsQuery = zt
	}
}

// Returns documents that match a provided text, number, date or boolean value. The provided text is analyzed before matching.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
func MatchQuery(field string, query string, params ...MatchQueryFieldParameter) effdsl.QueryResult {
	matchQuery := MatchQueryS{
		Field: field,
		Query: query,
	}

	var parameters matchQueryFieldParameters
	for _, prm := range params {
		prm(&parameters)
	}

	matchQuery.Operator = string(parameters.Operator)
	matchQuery.Fuzziness = string(parameters.Fuzziness)

	return effdsl.QueryResult{
		Ok:  matchQuery,
		Err: nil,
	}
}
