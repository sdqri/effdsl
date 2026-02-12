package children

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type ChildrenBody struct {
	Type string `json:"type,omitempty"`
}

type ChildrenS struct {
	name string
	body ChildrenBody
	aggregations.BaseAggregation
}

func (children *ChildrenS) AggregationName() string {
	return children.name
}

func (children *ChildrenS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("children", children.body, children.Extras())
}

type ChildrenOption = aggregations.Option[*ChildrenS]

func Children(name, childType string, opts ...ChildrenOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &ChildrenS{
		name: name,
		body: ChildrenBody{Type: childType},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) ChildrenOption {
	return aggregations.WithSubAggregation[*ChildrenS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) ChildrenOption {
	return aggregations.WithNamedSubAggregation[*ChildrenS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) ChildrenOption {
	return aggregations.WithSubAggregationsMap[*ChildrenS](subsMap)
}

func WithMetaField(key string, value any) ChildrenOption {
	return aggregations.WithMetaField[*ChildrenS](key, value)
}

func WithMetaMap(meta map[string]any) ChildrenOption {
	return aggregations.WithMetaMap[*ChildrenS](meta)
}
