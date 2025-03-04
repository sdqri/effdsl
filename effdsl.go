package effdsl

import (
	objs "github.com/sdqri/effdsl/objects"
)

//--------------------------------------------------------------------------------------//
//                                    Type aliasing                                     //
//--------------------------------------------------------------------------------------//

type (
	M                = objs.M
	SearchBody       = objs.SearchBody
	BodyOption       = objs.BodyOption
	QueryResult      = objs.QueryResult
	SortClauseResult = objs.SortClauseResult
	// Other types
	BooleanClause      = objs.BooleanClause
	QueryStringOption  = objs.QueryStringOption
	RangeQueryOption   = objs.RangeQueryOption
	TermQueryOption    = objs.TermQueryOption
	SourceFitlerOption = objs.SourceFitlerOption

	MatchOperator = objs.MatchOperator
	Fuzziness     = objs.Fuzziness

	WildcardQueryFieldParameter = objs.WildcardQueryFieldParameter
	RewriteParameter            = objs.RewriteParameter

	SortOrder = objs.SortOrder

	SuggestSort = objs.SuggestSort
	SuggestMode = objs.SuggestMode
)

//--------------------------------------------------------------------------------------//
//                                        Define                                        //
//--------------------------------------------------------------------------------------//

var (
	// body.go
	WithPIT          = objs.WithPIT
	WithPaginate     = objs.WithPaginate
	WithSearchAfter  = objs.WithSearchAfter
	WithQuery        = objs.WithQuery
	WithSort         = objs.WithSort
	WithCollpse      = objs.WithCollpse
	WithSourceFilter = objs.WithSourceFilter
	Define           = objs.Define
	// q_bool_query.go
	Must               = objs.Must
	Filter             = objs.Filter
	MustNot            = objs.MustNot
	Should             = objs.Should
	MinimumShouldMatch = objs.MinimumShouldMatch
	BoolQuery          = objs.BoolQuery
	// q_exists_query.go
	ExistsQuery = objs.ExistsQuery
	// q_fuzzy_query.go
	WithFuzziness      = objs.WithFuzziness
	WithPrefixLength   = objs.WithPrefixLength
	WithMaxExpansions  = objs.WithMaxExpansions
	WithFQRewrite      = objs.WithFQRewrite
	WithTranspositions = objs.WithTranspositions
	FuzzyQuery         = objs.FuzzyQuery
	// q_match_query.go
	MatchQuery             = objs.MatchQuery
	WithMatchOperator      = objs.WithMatchOperator
	WithFuzzinessParameter = objs.WithFuzzinessParameter
	// q_wildcard_query.go
	WildcardQuery        = objs.WildcardQuery
	WithBoost            = objs.WithBoost
	WithRewriteParameter = objs.WithRewriteParameter
	// q_query_string.go
	WithFields          = objs.WithFields
	WithAnalyzeWildcard = objs.WithAnalyzeWildcard
	QueryString         = objs.QueryString
	// q_range_query.go
	WithGT     = objs.WithGT
	WithGTE    = objs.WithGTE
	WithLT     = objs.WithLT
	WithLTE    = objs.WithLTE
	WithFormat = objs.WithFormat
	RangeQuery = objs.RangeQuery
	// q_regexp_query.go
	WithFlags                 = objs.WithFlags
	WithCaseInsensitive       = objs.WithCaseInsensitive
	WithMaxDeterminizedStates = objs.WithMaxDeterminizedStates
	WithRQRewrite             = objs.WithRQRewrite
	RegexpQuery               = objs.RegexpQuery
	// q_term_query.go
	WithTQBoost = objs.WithTQBoost
	TermQuery   = objs.TermQuery
	// q_term_set_query.go
	WithMinimumShouldMatchField  = objs.WithMinimumShouldMatchField
	WithMinimumShouldMatchScript = objs.WithMinimumShouldMatchScript
	TermsSetQuery                = objs.TermsSetQuery
	// q_terms_query.go
	WithTSQBoost = objs.WithTSQBoost
	TermsQuery   = objs.TermsQuery
	// search_sort.go
	SortClause = objs.SortClause
	// search_source_filtering.go
	WithIncludes = objs.WithIncludes
	WithExcludes = objs.WithExcludes
	SourceFilter = objs.SourceFilter

	// search_source_filtering.go
	Suggest                 = objs.Suggest
	WithSuggest             = objs.WithSuggest
	TermSuggester           = objs.TermSuggester
	WithTermSuggestAnalyzer = objs.WithTermSuggestAnalyzer
	WithTermSuggestSize     = objs.WithTermSuggestSize
	WithTermSuggestSort     = objs.WithTermSuggestSort
	WithTermSuggestMode     = objs.WithTermSuggestMode
)

//--------------------------------------------------------------------------------------//
//                                      constants                                       //
//--------------------------------------------------------------------------------------//

const (
	SORT_DEFAULT objs.SortOrder = objs.SORT_DEFAULT
	SORT_ASC     objs.SortOrder = objs.SORT_ASC
	SORT_DESC    objs.SortOrder = objs.SORT_DESC

	MatchOperatorOR  objs.MatchOperator = objs.MatchOperatorOR
	MatchOperatorAND objs.MatchOperator = objs.MatchOperatorAND

	FuzzinessAUTO objs.Fuzziness = objs.FuzzinessAUTO

	RewriteParameterConstantScoreBlended  objs.RewriteParameter = objs.RewriteParameterConstantScoreBlended
	RewriteParameterConstantScore         objs.RewriteParameter = objs.RewriteParameterConstantScore
	RewriteParameterConstantScoreBoolean  objs.RewriteParameter = objs.RewriteParameterConstantScoreBoolean
	RewriteParameterScoringBoolean        objs.RewriteParameter = objs.RewriteParameterScoringBoolean
	RewriteParameterTopTermsBlendedFreqsN objs.RewriteParameter = objs.RewriteParameterTopTermsBlendedFreqsN
	RewriteParameterTopTermsBoostN        objs.RewriteParameter = objs.RewriteParameterTopTermsBoostN
	RewriteParameterTopTermsN             objs.RewriteParameter = objs.RewriteParameterTopTermsN
)
