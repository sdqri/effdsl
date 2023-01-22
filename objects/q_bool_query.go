package objects

import (
	"encoding/json"
	"errors"
)

type CompoundQueries string

type BoolQueryS struct {
	Must               []Query `json:"must,omitempty"`
	Filter             []Query `json:"filter,omitempty"`
	MustNot            []Query `json:"must_not,omitempty"`
	Should             []Query `json:"should,omitempty"`
	MinimumShouldMatch string  `json:"minimum_should_match,omitempty"`
}

func (bq BoolQueryS) QueryInfo() string {
	return "Boolean query"
}

func (bq BoolQueryS) MarshalJSON() ([]byte, error) {
	type BoolQueryAlias BoolQueryS
	return json.Marshal(
		map[string]any{
			"bool": (BoolQueryAlias)(bq),
		},
	)
}

type BooleanClause func(*BoolQueryS) error

//--------------------------------------------------------------------------------------//
//                                         must                                         //
//--------------------------------------------------------------------------------------//

// The clause (query) must appear in matching documents and will contribute to the score.
func (f DefineType) Must(queryResults ...QueryResult) BooleanClause {
	clauses := make([]Query, 0)
	for _, qr := range queryResults {
		if qr.Err != nil {
			return func(bqs *BoolQueryS) error {
				return qr.Err
			}
		}
		clauses = append(clauses, qr.Ok)
	}

	return func(boolQuery *BoolQueryS) error {
		if boolQuery.Must == nil {
			boolQuery.Must = clauses
		} else {
			boolQuery.Must = append(boolQuery.Must, clauses...)
		}
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                        filter                                        //
//--------------------------------------------------------------------------------------//

// The clause (query) must appear in matching documents. However unlike must the score of the query will be ignored. Filter clauses are executed in filter context, meaning that scoring is ignored and clauses are considered for caching.
func (f DefineType) Filter(queryResults ...QueryResult) BooleanClause {
	clauses := make([]Query, 0)
	for _, qr := range queryResults {
		if qr.Err != nil {
			return func(bqs *BoolQueryS) error {
				return qr.Err
			}
		}
		clauses = append(clauses, qr.Ok)
	}

	return func(boolQuery *BoolQueryS) error {
		if boolQuery.Filter == nil {
			boolQuery.Filter = clauses
		} else {
			boolQuery.Filter = append(boolQuery.Filter, clauses...)
		}
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                       must_not                                       //
//--------------------------------------------------------------------------------------//

// The clause (query) must not appear in the matching documents. Clauses are executed in filter context meaning that scoring is ignored and clauses are considered for caching. Because scoring is ignored, a score of 0 for all documents is returned.
func (f DefineType) MustNot(queryResults ...QueryResult) BooleanClause {
	clauses := make([]Query, 0)
	for _, qr := range queryResults {
		if qr.Err != nil {
			return func(bqs *BoolQueryS) error {
				return qr.Err
			}
		}
		clauses = append(clauses, qr.Ok)
	}
	return func(boolQuery *BoolQueryS) error {
		if boolQuery.MustNot == nil {
			boolQuery.MustNot = clauses
		} else {
			boolQuery.MustNot = append(boolQuery.MustNot, clauses...)
		}
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                        should                                        //
//--------------------------------------------------------------------------------------//

// The clause (query) should appear in the matching document.
func (f DefineType) Should(queryResults ...QueryResult) BooleanClause {
	clauses := make([]Query, 0)
	for _, qr := range queryResults {
		if qr.Err != nil {
			return func(bqs *BoolQueryS) error {
				return qr.Err
			}
		}
		clauses = append(clauses, qr.Ok)
	}

	return func(boolQuery *BoolQueryS) error {
		if boolQuery.Should == nil {
			boolQuery.Should = clauses
		} else {
			boolQuery.Should = append(boolQuery.Should, clauses...)
		}
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                 minimum_should_match                                 //
//--------------------------------------------------------------------------------------//

// You can use the minimum_should_match parameter to specify the number or percentage of should clauses returned documents must match.
// [minimum_should_match parameter]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-minimum-should-match.html
func (f DefineType) MinimumShouldMatch(minimumShouldMatch string) BooleanClause {
	return func(boolQuery *BoolQueryS) error {
		boolQuery.MinimumShouldMatch = minimumShouldMatch
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                     constructor                                      //
//--------------------------------------------------------------------------------------//

// A query that matches documents matching boolean combinations of other queries. The bool query maps to Lucene BooleanQuery. It is built using one or more boolean clauses, each clause with a typed occurrence.
// [Boolean query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
func (f DefineType) BoolQuery(opts ...BooleanClause) QueryResult {
	boolQuery := new(BoolQueryS)
	if len(opts) == 0 {
		return QueryResult{
			Ok:  nil,
			Err: errors.New("boolean query needs at least one boolean clause(must, filter, ...)"),
		}
	}
	for _, opt := range opts {
		err := opt(boolQuery)
		if err != nil {
			return QueryResult{
				Ok:  nil,
				Err: err,
			}
		}
	}
	return QueryResult{
		Ok:  boolQuery,
		Err: nil,
	}
}
