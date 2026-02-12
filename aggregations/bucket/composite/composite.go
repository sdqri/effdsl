package composite

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type CompositeBody struct {
	Sources []any          `json:"sources,omitempty"`
	Size    *int           `json:"size,omitempty"`
	After   map[string]any `json:"after,omitempty"`
}

type CompositeS struct {
	name string
	body CompositeBody
	aggregations.BaseAggregation
}

func (comp *CompositeS) AggregationName() string {
	return comp.name
}

func (comp *CompositeS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("composite", comp.body, comp.Extras())
}

type CompositeOption = aggregations.Option[*CompositeS]

func Composite(name string, sources []any, opts ...CompositeOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &CompositeS{
		name: name,
		body: CompositeBody{
			Sources: sources,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSize(size int) CompositeOption {
	return func(c *CompositeS) error {
		c.body.Size = &size
		return nil
	}
}

func WithAfter(after map[string]any) CompositeOption {
	return func(c *CompositeS) error {
		c.body.After = after
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) CompositeOption {
	return aggregations.WithSubAggregation[*CompositeS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) CompositeOption {
	return aggregations.WithNamedSubAggregation[*CompositeS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) CompositeOption {
	return aggregations.WithSubAggregationsMap[*CompositeS](subsMap)
}

func WithMetaField(key string, value any) CompositeOption {
	return aggregations.WithMetaField[*CompositeS](key, value)
}

func WithMetaMap(meta map[string]any) CompositeOption {
	return aggregations.WithMetaMap[*CompositeS](meta)
}
