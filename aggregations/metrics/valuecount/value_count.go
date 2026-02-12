package valuecount

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type ValueCountBody struct {
	Field  string               `json:"field,omitempty"`
	Script *aggregations.Script `json:"script,omitempty"`
}

type ValueCountS struct {
	name string
	body ValueCountBody
	aggregations.BaseAggregation
}

func (count *ValueCountS) AggregationName() string {
	return count.name
}

func (count *ValueCountS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("value_count", count.body, count.Extras())
}

type ValueCountOption = aggregations.Option[*ValueCountS]

func ValueCount(name, field string, opts ...ValueCountOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &ValueCountS{
		name: name,
		body: ValueCountBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: value_count aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithScript(script aggregations.Script) ValueCountOption {
	return func(v *ValueCountS) error {
		v.body.Script = &script
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) ValueCountOption {
	return aggregations.WithSubAggregation[*ValueCountS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) ValueCountOption {
	return aggregations.WithNamedSubAggregation[*ValueCountS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) ValueCountOption {
	return aggregations.WithSubAggregationsMap[*ValueCountS](subsMap)
}

func WithMetaField(key string, value any) ValueCountOption {
	return aggregations.WithMetaField[*ValueCountS](key, value)
}

func WithMetaMap(meta map[string]any) ValueCountOption {
	return aggregations.WithMetaMap[*ValueCountS](meta)
}
