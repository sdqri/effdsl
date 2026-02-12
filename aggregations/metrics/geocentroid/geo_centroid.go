package geocentroid

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type GeoCentroidBody struct {
	Field string `json:"field,omitempty"`
}

type GeoCentroidS struct {
	name string
	body GeoCentroidBody
	aggregations.BaseAggregation
}

func (geo *GeoCentroidS) AggregationName() string {
	return geo.name
}

func (geo *GeoCentroidS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("geo_centroid", geo.body, geo.Extras())
}

type GeoCentroidOption = aggregations.Option[*GeoCentroidS]

func GeoCentroid(name, field string, opts ...GeoCentroidOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &GeoCentroidS{
		name: name,
		body: GeoCentroidBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) GeoCentroidOption {
	return aggregations.WithSubAggregation[*GeoCentroidS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) GeoCentroidOption {
	return aggregations.WithNamedSubAggregation[*GeoCentroidS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) GeoCentroidOption {
	return aggregations.WithSubAggregationsMap[*GeoCentroidS](subsMap)
}

func WithMetaField(key string, value any) GeoCentroidOption {
	return aggregations.WithMetaField[*GeoCentroidS](key, value)
}

func WithMetaMap(meta map[string]any) GeoCentroidOption {
	return aggregations.WithMetaMap[*GeoCentroidS](meta)
}
