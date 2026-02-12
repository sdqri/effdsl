package geohashgrid

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type GeoHashGridBody struct {
	Field     string         `json:"field,omitempty"`
	Precision *int           `json:"precision,omitempty"`
	Size      *int           `json:"size,omitempty"`
	ShardSize *int           `json:"shard_size,omitempty"`
	Bounds    map[string]any `json:"bounds,omitempty"`
}

type GeoHashGridS struct {
	name string
	body GeoHashGridBody
	aggregations.BaseAggregation
}

func (grid *GeoHashGridS) AggregationName() string {
	return grid.name
}

func (grid *GeoHashGridS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("geohash_grid", grid.body, grid.Extras())
}

type GeoHashGridOption = aggregations.Option[*GeoHashGridS]

func GeoHashGrid(name, field string, precision int, opts ...GeoHashGridOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &GeoHashGridS{
		name: name,
		body: GeoHashGridBody{
			Field:     field,
			Precision: &precision,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSize(size int) GeoHashGridOption {
	return func(g *GeoHashGridS) error {
		g.body.Size = &size
		return nil
	}
}

func WithShardSize(size int) GeoHashGridOption {
	return func(g *GeoHashGridS) error {
		g.body.ShardSize = &size
		return nil
	}
}

func WithBounds(bounds map[string]any) GeoHashGridOption {
	return func(g *GeoHashGridS) error {
		g.body.Bounds = bounds
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) GeoHashGridOption {
	return aggregations.WithSubAggregation[*GeoHashGridS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) GeoHashGridOption {
	return aggregations.WithNamedSubAggregation[*GeoHashGridS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) GeoHashGridOption {
	return aggregations.WithSubAggregationsMap[*GeoHashGridS](subsMap)
}

func WithMetaField(key string, value any) GeoHashGridOption {
	return aggregations.WithMetaField[*GeoHashGridS](key, value)
}

func WithMetaMap(meta map[string]any) GeoHashGridOption {
	return aggregations.WithMetaMap[*GeoHashGridS](meta)
}
