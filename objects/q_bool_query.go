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

func (f DefineType) MinimumShouldMatch(minimumShouldMatch string) BooleanClause {
	return func(boolQuery *BoolQueryS) error {
		boolQuery.MinimumShouldMatch = minimumShouldMatch
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                     constructor                                      //
//--------------------------------------------------------------------------------------//

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
