package avg

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type AvgBody struct {
	Field     string               `json:"field,omitempty"`
	Missing   any                  `json:"missing,omitempty"`
	Script    *aggregations.Script `json:"script,omitempty"`
	Format    string               `json:"format,omitempty"`
	ValueType string               `json:"value_type,omitempty"`
}

type AvgS struct {
	name string
	body AvgBody
	aggregations.BaseAggregation
}

func (avgS *AvgS) AggregationName() string {
	return avgS.name
}

func (avgS *AvgS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("avg", avgS.body, avgS.Extras())
}

type AvgOption = aggregations.Option[*AvgS]

func Avg(name, field string, opts ...AvgOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &AvgS{
		name: name,
		body: AvgBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: avg aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) AvgOption {
	return func(a *AvgS) error {
		a.body.Format = format
		return nil
	}
}

func WithMissing(missing any) AvgOption {
	return func(a *AvgS) error {
		a.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) AvgOption {
	return func(a *AvgS) error {
		a.body.Script = &script
		return nil
	}
}

func WithValueType(valueType string) AvgOption {
	return func(a *AvgS) error {
		a.body.ValueType = valueType
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) AvgOption {
	return aggregations.WithSubAggregation[*AvgS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) AvgOption {
	return aggregations.WithNamedSubAggregation[*AvgS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) AvgOption {
	return aggregations.WithSubAggregationsMap[*AvgS](subsMap)
}

func WithMetaField(key string, value any) AvgOption {
	return aggregations.WithMetaField[*AvgS](key, value)
}

func WithMetaMap(meta map[string]any) AvgOption {
	return aggregations.WithMetaMap[*AvgS](meta)
}
