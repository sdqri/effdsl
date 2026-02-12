package filters

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type FiltersBody struct {
	Filters        any    `json:"filters,omitempty"`
	OtherBucket    *bool  `json:"other_bucket,omitempty"`
	OtherBucketKey string `json:"other_bucket_key,omitempty"`
	Keyed          *bool  `json:"keyed,omitempty"`
}

type FiltersS struct {
	name string
	body FiltersBody
	aggregations.BaseAggregation
}

func (filters *FiltersS) AggregationName() string {
	return filters.name
}

func (filters *FiltersS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("filters", filters.body, filters.Extras())
}

type FiltersOption = aggregations.Option[*FiltersS]

func Filters(name string, filters any, opts ...FiltersOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &FiltersS{
		name: name,
		body: FiltersBody{
			Filters: filters,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithOtherBucket(enabled bool) FiltersOption {
	return func(f *FiltersS) error {
		f.body.OtherBucket = &enabled
		return nil
	}
}

func WithOtherBucketKey(key string) FiltersOption {
	return func(f *FiltersS) error {
		f.body.OtherBucketKey = key
		return nil
	}
}

func WithKeyed(keyed bool) FiltersOption {
	return func(f *FiltersS) error {
		f.body.Keyed = &keyed
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) FiltersOption {
	return aggregations.WithSubAggregation[*FiltersS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) FiltersOption {
	return aggregations.WithNamedSubAggregation[*FiltersS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) FiltersOption {
	return aggregations.WithSubAggregationsMap[*FiltersS](subsMap)
}

func WithMetaField(key string, value any) FiltersOption {
	return aggregations.WithMetaField[*FiltersS](key, value)
}

func WithMetaMap(meta map[string]any) FiltersOption {
	return aggregations.WithMetaMap[*FiltersS](meta)
}
