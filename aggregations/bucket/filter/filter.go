package filter

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type FilterBody struct {
	Filter any `json:"filter,omitempty"`
}

type FilterS struct {
	name string
	body FilterBody
	aggregations.BaseAggregation
}

func (filter *FilterS) AggregationName() string {
	return filter.name
}

func (filter *FilterS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("filter", filter.body, filter.Extras())
}

type FilterOption = aggregations.Option[*FilterS]

func Filter(name string, filterQuery any, opts ...FilterOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &FilterS{
		name: name,
		body: FilterBody{
			Filter: filterQuery,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) FilterOption {
	return aggregations.WithSubAggregation[*FilterS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) FilterOption {
	return aggregations.WithNamedSubAggregation[*FilterS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) FilterOption {
	return aggregations.WithSubAggregationsMap[*FilterS](subsMap)
}

func WithMetaField(key string, value any) FilterOption {
	return aggregations.WithMetaField[*FilterS](key, value)
}

func WithMetaMap(meta map[string]any) FilterOption {
	return aggregations.WithMetaMap[*FilterS](meta)
}
