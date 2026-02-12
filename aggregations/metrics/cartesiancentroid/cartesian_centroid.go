package cartesiancentroid

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type CartesianCentroidBody struct {
	Field string `json:"field,omitempty"`
}

type CartesianCentroidS struct {
	name string
	body CartesianCentroidBody
	aggregations.BaseAggregation
}

func (centroid *CartesianCentroidS) AggregationName() string {
	return centroid.name
}

func (centroid *CartesianCentroidS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("cartesian_centroid", centroid.body, centroid.Extras())
}

type CartesianCentroidOption = aggregations.Option[*CartesianCentroidS]

func CartesianCentroid(name, field string, opts ...CartesianCentroidOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &CartesianCentroidS{
		name: name,
		body: CartesianCentroidBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) CartesianCentroidOption {
	return aggregations.WithSubAggregation[*CartesianCentroidS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) CartesianCentroidOption {
	return aggregations.WithNamedSubAggregation[*CartesianCentroidS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) CartesianCentroidOption {
	return aggregations.WithSubAggregationsMap[*CartesianCentroidS](subsMap)
}

func WithMetaField(key string, value any) CartesianCentroidOption {
	return aggregations.WithMetaField[*CartesianCentroidS](key, value)
}

func WithMetaMap(meta map[string]any) CartesianCentroidOption {
	return aggregations.WithMetaMap[*CartesianCentroidS](meta)
}
