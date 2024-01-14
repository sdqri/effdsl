package objects

import (
	"encoding/json"
)

type MatchQueryS struct {
	Field     string `json:"-"`                   // (Required, object) Field you wish to search.
	Query     string `json:"query"`               // (Required) Text, number, boolean value or date you wish to find in the provided <field>.
	Operator  string `json:"operator,omitempty"`  // (Optional, string) Boolean logic used to interpret text in the query value.
	Fuzziness string `json:"fuzziness,omitempty"` // (Optional, string) Maximum edit distance allowed for matching.
}

func (mq MatchQueryS) QueryInfo() string {
	return "Match query"
}

func (mq MatchQueryS) MarshalJSON() ([]byte, error) {
	type MatchQueryBase MatchQueryS
	return json.Marshal(
		M{
			"match": M{
				mq.Field: (MatchQueryBase)(mq),
			},
		},
	)
}

type MatchOperator string

const (
	MatchOperatorOR  MatchOperator = "OR"
	MatchOperatorAND MatchOperator = "AND"
)

type Fuzziness string

const (
	FuzzinessAUTO Fuzziness = "AUTO"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html#match-field-params
type matchQueryFieldParameters struct {
	/*
		(Optional, string) Boolean logic used to interpret text in the query value. Valid values are:

		* OR (Default)
			* For example, a query value of capital of Hungary is interpreted as capital OR of OR Hungary.
		* AND
			* For example, a query value of capital of Hungary is interpreted as capital AND of AND Hungary.
	*/
	Operator MatchOperator

	/*
		(Optional, string) Maximum edit distance allowed for matching. See [Fuzziness](https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness) for valid values and more information.
	*/
	Fuzziness Fuzziness
}

type MatchQueryFieldParameter func(prms *matchQueryFieldParameters)

// WithMatchOperator ...
func WithMatchOperator(op MatchOperator) MatchQueryFieldParameter {
	return func(prms *matchQueryFieldParameters) {
		prms.Operator = op
	}
}

// WithFuzzinessParameter ...
func WithFuzzinessParameter(f Fuzziness) MatchQueryFieldParameter {
	return func(prms *matchQueryFieldParameters) {
		prms.Fuzziness = f
	}
}

// Returns documents that match a provided text, number, date or boolean value. The provided text is analyzed before matching.
// [Match query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
func MatchQuery(field string, query string, prms ...MatchQueryFieldParameter) QueryResult {
	matchQuery := MatchQueryS{
		Field: field,
		Query: query,
	}

	var parameters matchQueryFieldParameters
	for _, prm := range prms {
		prm(&parameters)
	}

	matchQuery.Operator = string(parameters.Operator)
	matchQuery.Fuzziness = string(parameters.Fuzziness)

	return QueryResult{
		Ok:  matchQuery,
		Err: nil,
	}
}
