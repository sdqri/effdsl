package max

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MaxBody struct {
	Field     string               `json:"field,omitempty"`
	Missing   any                  `json:"missing,omitempty"`
	Script    *aggregations.Script `json:"script,omitempty"`
	Format    string               `json:"format,omitempty"`
	ValueType string               `json:"value_type,omitempty"`
}

type MaxS struct {
	name string
	body MaxBody
	aggregations.BaseAggregation
}

func (maxS *MaxS) AggregationName() string {
	return maxS.name
}

func (maxS *MaxS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("max", maxS.body, maxS.Extras())
}

type MaxOption = aggregations.Option[*MaxS]

func Max(name, field string, opts ...MaxOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &MaxS{
		name: name,
		body: MaxBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: max aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) MaxOption {
	return func(m *MaxS) error {
		m.body.Format = format
		return nil
	}
}

func WithMissing(missing any) MaxOption {
	return func(m *MaxS) error {
		m.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) MaxOption {
	return func(m *MaxS) error {
		m.body.Script = &script
		return nil
	}
}

func WithValueType(valueType string) MaxOption {
	return func(m *MaxS) error {
		m.body.ValueType = valueType
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MaxOption {
	return aggregations.WithSubAggregation[*MaxS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MaxOption {
	return aggregations.WithNamedSubAggregation[*MaxS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MaxOption {
	return aggregations.WithSubAggregationsMap[*MaxS](subsMap)
}

func WithMetaField(key string, value any) MaxOption {
	return aggregations.WithMetaField[*MaxS](key, value)
}

func WithMetaMap(meta map[string]any) MaxOption {
	return aggregations.WithMetaMap[*MaxS](meta)
}
