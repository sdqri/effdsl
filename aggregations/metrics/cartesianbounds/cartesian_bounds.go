package cartesianbounds

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type CartesianBoundsBody struct {
	Field string `json:"field,omitempty"`
}

type CartesianBoundsS struct {
	name string
	body CartesianBoundsBody
	aggregations.BaseAggregation
}

func (bounds *CartesianBoundsS) AggregationName() string {
	return bounds.name
}

func (bounds *CartesianBoundsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("cartesian_bounds", bounds.body, bounds.Extras())
}

type CartesianBoundsOption = aggregations.Option[*CartesianBoundsS]

func CartesianBounds(name, field string, opts ...CartesianBoundsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &CartesianBoundsS{
		name: name,
		body: CartesianBoundsBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) CartesianBoundsOption {
	return aggregations.WithSubAggregation[*CartesianBoundsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) CartesianBoundsOption {
	return aggregations.WithNamedSubAggregation[*CartesianBoundsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) CartesianBoundsOption {
	return aggregations.WithSubAggregationsMap[*CartesianBoundsS](subsMap)
}

func WithMetaField(key string, value any) CartesianBoundsOption {
	return aggregations.WithMetaField[*CartesianBoundsS](key, value)
}

func WithMetaMap(meta map[string]any) CartesianBoundsOption {
	return aggregations.WithMetaMap[*CartesianBoundsS](meta)
}
