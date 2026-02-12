package min

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MinBody struct {
	Field     string               `json:"field,omitempty"`
	Missing   any                  `json:"missing,omitempty"`
	Script    *aggregations.Script `json:"script,omitempty"`
	Format    string               `json:"format,omitempty"`
	ValueType string               `json:"value_type,omitempty"`
}

type MinS struct {
	name string
	body MinBody
	aggregations.BaseAggregation
}

func (minS *MinS) AggregationName() string {
	return minS.name
}

func (minS *MinS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("min", minS.body, minS.Extras())
}

type MinOption = aggregations.Option[*MinS]

func Min(name, field string, opts ...MinOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &MinS{
		name: name,
		body: MinBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: min aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) MinOption {
	return func(m *MinS) error {
		m.body.Format = format
		return nil
	}
}

func WithMissing(missing any) MinOption {
	return func(m *MinS) error {
		m.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) MinOption {
	return func(m *MinS) error {
		m.body.Script = &script
		return nil
	}
}

func WithValueType(valueType string) MinOption {
	return func(m *MinS) error {
		m.body.ValueType = valueType
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MinOption {
	return aggregations.WithSubAggregation[*MinS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MinOption {
	return aggregations.WithNamedSubAggregation[*MinS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MinOption {
	return aggregations.WithSubAggregationsMap[*MinS](subsMap)
}

func WithMetaField(key string, value any) MinOption {
	return aggregations.WithMetaField[*MinS](key, value)
}

func WithMetaMap(meta map[string]any) MinOption {
	return aggregations.WithMetaMap[*MinS](meta)
}
