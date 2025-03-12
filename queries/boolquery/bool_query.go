package booleanquery

import (
	"encoding/json"
	"errors"

	"github.com/sdqri/effdsl/v2"
)

type CompoundQueries string

type BoolQueryS struct {
	Must               []effdsl.Query `json:"must,omitempty"`
	Filter             []effdsl.Query `json:"filter,omitempty"`
	Should             []effdsl.Query `json:"should,omitempty"`
	MustNot            []effdsl.Query `json:"must_not,omitempty"`
	MinimumShouldMatch string         `json:"minimum_should_match,omitempty"`
}

func (bq BoolQueryS) QueryInfo() string {
	return "Boolean query"
}

func (bq BoolQueryS) MarshalJSON() ([]byte, error) {
	type BoolQueryBase BoolQueryS
	return json.Marshal(
		map[string]any{
			"bool": (BoolQueryBase)(bq),
		},
	)
}

type BooleanClause func(*BoolQueryS) error

// The clause (query) must appear in matching documents and will contribute to the score.
func Must(queryResults ...effdsl.QueryResult) BooleanClause {
	clauses := make([]effdsl.Query, 0)
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

// The clause (query) must appear in matching documents. However unlike must the score of the query will be ignored. Filter clauses are executed in filter context, meaning that scoring is ignored and clauses are considered for caching.
func Filter(queryResults ...effdsl.QueryResult) BooleanClause {
	clauses := make([]effdsl.Query, 0)
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

// The clause (query) should appear in the matching document.
func Should(queryResults ...effdsl.QueryResult) BooleanClause {
	clauses := make([]effdsl.Query, 0)
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

// The clause (query) must not appear in the matching documents. Clauses are executed in filter context meaning that scoring is ignored and clauses are considered for caching. Because scoring is ignored, a score of 0 for all documents is returned.
func MustNot(queryResults ...effdsl.QueryResult) BooleanClause {
	clauses := make([]effdsl.Query, 0)
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

// You can use the minimum_should_match parameter to specify the number or percentage of should clauses returned documents must match.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-minimum-should-match.html
func WithMinimumShouldMatch(minimumShouldMatch string) BooleanClause {
	return func(boolQuery *BoolQueryS) error {
		boolQuery.MinimumShouldMatch = minimumShouldMatch
		return nil
	}
}

// A query that matches documents matching boolean combinations of other queries. The bool query maps to Lucene BooleanQuery. It is built using one or more boolean clauses, each clause with a typed occurrence.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
func BoolQuery(opts ...BooleanClause) effdsl.QueryResult {
	boolQuery := new(BoolQueryS)
	if len(opts) == 0 {
		return effdsl.QueryResult{
			Ok:  nil,
			Err: errors.New("boolean query needs at least one boolean clause(must, filter, ...)"),
		}
	}

	for _, opt := range opts {
		err := opt(boolQuery)
		if err != nil {
			return effdsl.QueryResult{
				Ok:  nil,
				Err: err,
			}
		}
	}
	return effdsl.QueryResult{
		Ok:  boolQuery,
		Err: nil,
	}
}
