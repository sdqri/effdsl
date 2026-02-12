package geoline

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type GeoLineField struct {
	Field string `json:"field,omitempty"`
}

type GeoLineBody struct {
	Point       GeoLineField  `json:"point,omitempty"`
	Sort        *GeoLineField `json:"sort,omitempty"`
	IncludeSort *bool         `json:"include_sort,omitempty"`
	SortOrder   string        `json:"sort_order,omitempty"`
	Size        *int          `json:"size,omitempty"`
}

type GeoLineS struct {
	name string
	body GeoLineBody
	aggregations.BaseAggregation
}

func (geo *GeoLineS) AggregationName() string {
	return geo.name
}

func (geo *GeoLineS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("geo_line", geo.body, geo.Extras())
}

type GeoLineOption = aggregations.Option[*GeoLineS]

func GeoLine(name, pointField string, opts ...GeoLineOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &GeoLineS{
		name: name,
		body: GeoLineBody{
			Point: GeoLineField{Field: pointField},
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSortField(field string) GeoLineOption {
	return func(g *GeoLineS) error {
		g.body.Sort = &GeoLineField{Field: field}
		return nil
	}
}

func WithIncludeSort(include bool) GeoLineOption {
	return func(g *GeoLineS) error {
		g.body.IncludeSort = &include
		return nil
	}
}

func WithSortOrder(order string) GeoLineOption {
	return func(g *GeoLineS) error {
		g.body.SortOrder = order
		return nil
	}
}

func WithSize(size int) GeoLineOption {
	return func(g *GeoLineS) error {
		g.body.Size = &size
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) GeoLineOption {
	return aggregations.WithSubAggregation[*GeoLineS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) GeoLineOption {
	return aggregations.WithNamedSubAggregation[*GeoLineS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) GeoLineOption {
	return aggregations.WithSubAggregationsMap[*GeoLineS](subsMap)
}

func WithMetaField(key string, value any) GeoLineOption {
	return aggregations.WithMetaField[*GeoLineS](key, value)
}

func WithMetaMap(meta map[string]any) GeoLineOption {
	return aggregations.WithMetaMap[*GeoLineS](meta)
}
