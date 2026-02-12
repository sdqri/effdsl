package geodistance

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type GeoDistanceRange struct {
	Key  string   `json:"key,omitempty"`
	From *float64 `json:"from,omitempty"`
	To   *float64 `json:"to,omitempty"`
}

type GeoDistanceBody struct {
	Field        string             `json:"field,omitempty"`
	Origin       any                `json:"origin,omitempty"`
	Unit         string             `json:"unit,omitempty"`
	DistanceType string             `json:"distance_type,omitempty"`
	Ranges       []GeoDistanceRange `json:"ranges,omitempty"`
	Keyed        *bool              `json:"keyed,omitempty"`
	Missing      any                `json:"missing,omitempty"`
}

type GeoDistanceS struct {
	name string
	body GeoDistanceBody
	aggregations.BaseAggregation
}

func (geo *GeoDistanceS) AggregationName() string {
	return geo.name
}

func (geo *GeoDistanceS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("geo_distance", geo.body, geo.Extras())
}

type GeoDistanceOption = aggregations.Option[*GeoDistanceS]

func GeoDistance(name, field string, origin any, ranges []GeoDistanceRange, opts ...GeoDistanceOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &GeoDistanceS{
		name: name,
		body: GeoDistanceBody{
			Field:  field,
			Origin: origin,
			Ranges: ranges,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithUnit(unit string) GeoDistanceOption {
	return func(g *GeoDistanceS) error {
		g.body.Unit = unit
		return nil
	}
}

func WithDistanceType(distanceType string) GeoDistanceOption {
	return func(g *GeoDistanceS) error {
		g.body.DistanceType = distanceType
		return nil
	}
}

func WithKeyed(keyed bool) GeoDistanceOption {
	return func(g *GeoDistanceS) error {
		g.body.Keyed = &keyed
		return nil
	}
}

func WithMissing(missing any) GeoDistanceOption {
	return func(g *GeoDistanceS) error {
		g.body.Missing = missing
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) GeoDistanceOption {
	return aggregations.WithSubAggregation[*GeoDistanceS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) GeoDistanceOption {
	return aggregations.WithNamedSubAggregation[*GeoDistanceS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) GeoDistanceOption {
	return aggregations.WithSubAggregationsMap[*GeoDistanceS](subsMap)
}

func WithMetaField(key string, value any) GeoDistanceOption {
	return aggregations.WithMetaField[*GeoDistanceS](key, value)
}

func WithMetaMap(meta map[string]any) GeoDistanceOption {
	return aggregations.WithMetaMap[*GeoDistanceS](meta)
}
