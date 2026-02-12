package geohexgrid

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type GeoHexGridBody struct {
	Field     string         `json:"field,omitempty"`
	Precision *int           `json:"precision,omitempty"`
	Size      *int           `json:"size,omitempty"`
	ShardSize *int           `json:"shard_size,omitempty"`
	Bounds    map[string]any `json:"bounds,omitempty"`
}

type GeoHexGridS struct {
	name string
	body GeoHexGridBody
	aggregations.BaseAggregation
}

func (grid *GeoHexGridS) AggregationName() string {
	return grid.name
}

func (grid *GeoHexGridS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("geohex_grid", grid.body, grid.Extras())
}

type GeoHexGridOption = aggregations.Option[*GeoHexGridS]

func GeoHexGrid(name, field string, precision int, opts ...GeoHexGridOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &GeoHexGridS{
		name: name,
		body: GeoHexGridBody{
			Field:     field,
			Precision: &precision,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSize(size int) GeoHexGridOption {
	return func(g *GeoHexGridS) error {
		g.body.Size = &size
		return nil
	}
}

func WithShardSize(size int) GeoHexGridOption {
	return func(g *GeoHexGridS) error {
		g.body.ShardSize = &size
		return nil
	}
}

func WithBounds(bounds map[string]any) GeoHexGridOption {
	return func(g *GeoHexGridS) error {
		g.body.Bounds = bounds
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) GeoHexGridOption {
	return aggregations.WithSubAggregation[*GeoHexGridS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) GeoHexGridOption {
	return aggregations.WithNamedSubAggregation[*GeoHexGridS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) GeoHexGridOption {
	return aggregations.WithSubAggregationsMap[*GeoHexGridS](subsMap)
}

func WithMetaField(key string, value any) GeoHexGridOption {
	return aggregations.WithMetaField[*GeoHexGridS](key, value)
}

func WithMetaMap(meta map[string]any) GeoHexGridOption {
	return aggregations.WithMetaMap[*GeoHexGridS](meta)
}
