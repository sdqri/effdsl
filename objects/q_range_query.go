package objects

import (
	"encoding/json"
	"errors"
)

type RangeQueryS struct {
	Field  string `json:"-"`                //(Required, object) Field you wish to search.
	GT     any    `json:"gt,omitempty"`     //(Optional) Greater than.
	GTE    any    `json:"gte,omitempty"`    //(Optional) Greater than or equal to.
	LT     any    `json:"lt,omitempty"`     //(Optional) Less than.
	LTE    any    `json:"lte,omitempty"`    //(Optional) Less than or equal to.
	Format string `json:"format,omitempty"` //(Optional, string) Date format used to convert date values in the query.
}

func (rq RangeQueryS) QueryInfo() string {
	return "Range query"
}

func (rq RangeQueryS) MarshalJSON() ([]byte, error) {
	type RangeQueryAlias RangeQueryS
	return json.Marshal(
		M{
			"range": M{
				rq.Field: (RangeQueryAlias)(rq),
			},
		},
	)
}

type RangeQueryOption func(*RangeQueryS)

func (f DefineType) WithGT(gt any) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.GT = gt
	}
}

func (f DefineType) WithGTE(gte any) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.GTE = gte
	}
}

func (f DefineType) WithLT(lt any) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.LT = lt
	}
}

func (f DefineType) WithLTE(lte any) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.LTE = lte
	}
}

func (f DefineType) WithFormat(format string) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.Format = format
	}
}

// Returns documents that contain terms within a provided range.
// [Range query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html
func (f DefineType) RangeQuery(field string, opts ...RangeQueryOption) QueryResult {
	rangeQuery := RangeQueryS{
		Field: field,
	}
	for _, opt := range opts {
		opt(&rangeQuery)
	}
	if rangeQuery.GT == nil && rangeQuery.GTE == nil && rangeQuery.LT == nil && rangeQuery.LTE == nil {
		return QueryResult{
			Ok:  rangeQuery,
			Err: errors.New("One of WithGT, WithGTE, WithLT, WithLTE should be proveded for range query"),
		}
	}
	return QueryResult{
		Ok:  rangeQuery,
		Err: nil,
	}
}
