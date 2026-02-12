package geotilegrid

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type GeoTileGridBody struct {
	Field     string         `json:"field,omitempty"`
	Precision *int           `json:"precision,omitempty"`
	Size      *int           `json:"size,omitempty"`
	ShardSize *int           `json:"shard_size,omitempty"`
	Bounds    map[string]any `json:"bounds,omitempty"`
}

type GeoTileGridS struct {
	name string
	body GeoTileGridBody
	aggregations.BaseAggregation
}

func (grid *GeoTileGridS) AggregationName() string {
	return grid.name
}

func (grid *GeoTileGridS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("geotile_grid", grid.body, grid.Extras())
}

type GeoTileGridOption = aggregations.Option[*GeoTileGridS]

func GeoTileGrid(name, field string, precision int, opts ...GeoTileGridOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &GeoTileGridS{
		name: name,
		body: GeoTileGridBody{
			Field:     field,
			Precision: &precision,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSize(size int) GeoTileGridOption {
	return func(g *GeoTileGridS) error {
		g.body.Size = &size
		return nil
	}
}

func WithShardSize(size int) GeoTileGridOption {
	return func(g *GeoTileGridS) error {
		g.body.ShardSize = &size
		return nil
	}
}

func WithBounds(bounds map[string]any) GeoTileGridOption {
	return func(g *GeoTileGridS) error {
		g.body.Bounds = bounds
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) GeoTileGridOption {
	return aggregations.WithSubAggregation[*GeoTileGridS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) GeoTileGridOption {
	return aggregations.WithNamedSubAggregation[*GeoTileGridS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) GeoTileGridOption {
	return aggregations.WithSubAggregationsMap[*GeoTileGridS](subsMap)
}

func WithMetaField(key string, value any) GeoTileGridOption {
	return aggregations.WithMetaField[*GeoTileGridS](key, value)
}

func WithMetaMap(meta map[string]any) GeoTileGridOption {
	return aggregations.WithMetaMap[*GeoTileGridS](meta)
}
