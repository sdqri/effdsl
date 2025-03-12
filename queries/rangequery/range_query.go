package rangequery

import (
	"encoding/json"
	"errors"

	"github.com/sdqri/effdsl/v2"
)

type RangeQueryS struct {
	Field    string   `json:"-"`                   //(Required, object) Field you wish to search.
	GT       any      `json:"gt,omitempty"`        //(Optional) Greater than.
	GTE      any      `json:"gte,omitempty"`       //(Optional) Greater than or equal to.
	LT       any      `json:"lt,omitempty"`        //(Optional) Less than.
	LTE      any      `json:"lte,omitempty"`       //(Optional) Less than or equal to.
	Format   string   `json:"format,omitempty"`    //(Optional, string) Date format used to convert date values in the query.
	Relation Relation `json:"relation,omitempty"`  //(Optional, string) Indicates how the range query matches values for range fields.
	TimeZone string   `json:"time_zone,omitempty"` //(Optional, string) Coordinated Universal Time (UTC) offset or IANA time zone used to convert date values in the query to UTC.
	Boost    float64  `json:"boost,omitempty"`     //(Optional, float) Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.
}

func (rq RangeQueryS) QueryInfo() string {
	return "Range query"
}

func (rq RangeQueryS) MarshalJSON() ([]byte, error) {
	type RangeQueryBase RangeQueryS
	return json.Marshal(
		effdsl.M{
			"range": effdsl.M{
				rq.Field: (RangeQueryBase)(rq),
			},
		},
	)
}

type RangeQueryOption func(*RangeQueryS)

func WithGT(gt any) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.GT = gt
	}
}

func WithGTE(gte any) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.GTE = gte
	}
}

func WithLT(lt any) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.LT = lt
	}
}

func WithLTE(lte any) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.LTE = lte
	}
}

func WithFormat(format string) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.Format = format
	}
}

type Relation string

const (
	INTERSECTS Relation = "INTERSECTS"
	CONTAINS   Relation = "CONTAINS"
	WITHIN     Relation = "WITHIN"
)

func WithRelation(relation Relation) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.Relation = relation
	}
}

func WithTimeZone(timeZone string) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.TimeZone = timeZone
	}
}

func WithBoost(boost float64) RangeQueryOption {
	return func(rangeQuery *RangeQueryS) {
		rangeQuery.Boost = boost
	}
}

// Returns documents that contain terms within a provided range.
// [Range query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html
func RangeQuery(field string, opts ...RangeQueryOption) effdsl.QueryResult {
	rangeQuery := RangeQueryS{
		Field: field,
	}
	for _, opt := range opts {
		opt(&rangeQuery)
	}
	if rangeQuery.GT == nil && rangeQuery.GTE == nil && rangeQuery.LT == nil && rangeQuery.LTE == nil {
		return effdsl.QueryResult{
			Ok:  rangeQuery,
			Err: errors.New("one of WithGT, WithGTE, WithLT, WithLTE should be provided for range query"),
		}
	}
	return effdsl.QueryResult{
		Ok:  rangeQuery,
		Err: nil,
	}
}
