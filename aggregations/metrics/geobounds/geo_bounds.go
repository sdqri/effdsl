package geobounds

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type GeoBoundsBody struct {
	Field         string `json:"field,omitempty"`
	WrapLongitude *bool  `json:"wrap_longitude,omitempty"`
}

type GeoBoundsS struct {
	name string
	body GeoBoundsBody
	aggregations.BaseAggregation
}

func (geo *GeoBoundsS) AggregationName() string {
	return geo.name
}

func (geo *GeoBoundsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("geo_bounds", geo.body, geo.Extras())
}

type GeoBoundsOption = aggregations.Option[*GeoBoundsS]

func GeoBounds(name, field string, opts ...GeoBoundsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &GeoBoundsS{
		name: name,
		body: GeoBoundsBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithWrapLongitude(wrap bool) GeoBoundsOption {
	return func(g *GeoBoundsS) error {
		g.body.WrapLongitude = &wrap
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) GeoBoundsOption {
	return aggregations.WithSubAggregation[*GeoBoundsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) GeoBoundsOption {
	return aggregations.WithNamedSubAggregation[*GeoBoundsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) GeoBoundsOption {
	return aggregations.WithSubAggregationsMap[*GeoBoundsS](subsMap)
}

func WithMetaField(key string, value any) GeoBoundsOption {
	return aggregations.WithMetaField[*GeoBoundsS](key, value)
}

func WithMetaMap(meta map[string]any) GeoBoundsOption {
	return aggregations.WithMetaMap[*GeoBoundsS](meta)
}
