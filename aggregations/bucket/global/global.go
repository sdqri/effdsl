package global

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type GlobalS struct {
	name string
	aggregations.BaseAggregation
}

func (global *GlobalS) AggregationName() string {
	return global.name
}

func (global *GlobalS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("global", map[string]any{}, global.Extras())
}

type GlobalOption = aggregations.Option[*GlobalS]

func Global(name string, opts ...GlobalOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &GlobalS{name: name}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) GlobalOption {
	return aggregations.WithSubAggregation[*GlobalS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) GlobalOption {
	return aggregations.WithNamedSubAggregation[*GlobalS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) GlobalOption {
	return aggregations.WithSubAggregationsMap[*GlobalS](subsMap)
}

func WithMetaField(key string, value any) GlobalOption {
	return aggregations.WithMetaField[*GlobalS](key, value)
}

func WithMetaMap(meta map[string]any) GlobalOption {
	return aggregations.WithMetaMap[*GlobalS](meta)
}
