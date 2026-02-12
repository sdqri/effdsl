package nested

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type NestedBody struct {
	Path string `json:"path,omitempty"`
}

type NestedS struct {
	name string
	body NestedBody
	aggregations.BaseAggregation
}

func (nested *NestedS) AggregationName() string {
	return nested.name
}

func (nested *NestedS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("nested", nested.body, nested.Extras())
}

type NestedOption = aggregations.Option[*NestedS]

func Nested(name, path string, opts ...NestedOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &NestedS{
		name: name,
		body: NestedBody{Path: path},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) NestedOption {
	return aggregations.WithSubAggregation[*NestedS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) NestedOption {
	return aggregations.WithNamedSubAggregation[*NestedS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) NestedOption {
	return aggregations.WithSubAggregationsMap[*NestedS](subsMap)
}

func WithMetaField(key string, value any) NestedOption {
	return aggregations.WithMetaField[*NestedS](key, value)
}

func WithMetaMap(meta map[string]any) NestedOption {
	return aggregations.WithMetaMap[*NestedS](meta)
}
