package reversenested

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type ReverseNestedBody struct {
	Path string `json:"path,omitempty"`
}

type ReverseNestedS struct {
	name string
	body ReverseNestedBody
	aggregations.BaseAggregation
}

func (nested *ReverseNestedS) AggregationName() string {
	return nested.name
}

func (nested *ReverseNestedS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("reverse_nested", nested.body, nested.Extras())
}

type ReverseNestedOption = aggregations.Option[*ReverseNestedS]

func ReverseNested(name string, opts ...ReverseNestedOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &ReverseNestedS{name: name, body: ReverseNestedBody{}}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithPath(path string) ReverseNestedOption {
	return func(r *ReverseNestedS) error {
		r.body.Path = path
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) ReverseNestedOption {
	return aggregations.WithSubAggregation[*ReverseNestedS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) ReverseNestedOption {
	return aggregations.WithNamedSubAggregation[*ReverseNestedS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) ReverseNestedOption {
	return aggregations.WithSubAggregationsMap[*ReverseNestedS](subsMap)
}

func WithMetaField(key string, value any) ReverseNestedOption {
	return aggregations.WithMetaField[*ReverseNestedS](key, value)
}

func WithMetaMap(meta map[string]any) ReverseNestedOption {
	return aggregations.WithMetaMap[*ReverseNestedS](meta)
}
