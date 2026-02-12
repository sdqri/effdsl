package parent

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type ParentBody struct {
	Type string `json:"type,omitempty"`
}

type ParentS struct {
	name string
	body ParentBody
	aggregations.BaseAggregation
}

func (parent *ParentS) AggregationName() string {
	return parent.name
}

func (parent *ParentS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("parent", parent.body, parent.Extras())
}

type ParentOption = aggregations.Option[*ParentS]

func Parent(name, parentType string, opts ...ParentOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &ParentS{
		name: name,
		body: ParentBody{Type: parentType},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) ParentOption {
	return aggregations.WithSubAggregation[*ParentS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) ParentOption {
	return aggregations.WithNamedSubAggregation[*ParentS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) ParentOption {
	return aggregations.WithSubAggregationsMap[*ParentS](subsMap)
}

func WithMetaField(key string, value any) ParentOption {
	return aggregations.WithMetaField[*ParentS](key, value)
}

func WithMetaMap(meta map[string]any) ParentOption {
	return aggregations.WithMetaMap[*ParentS](meta)
}
