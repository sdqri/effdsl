package sum

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type SumBody struct {
	Field     string               `json:"field,omitempty"`
	Missing   any                  `json:"missing,omitempty"`
	Script    *aggregations.Script `json:"script,omitempty"`
	Format    string               `json:"format,omitempty"`
	ValueType string               `json:"value_type,omitempty"`
}

type SumS struct {
	name string
	body SumBody
	aggregations.BaseAggregation
}

func (sumS *SumS) AggregationName() string {
	return sumS.name
}

func (sumS *SumS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("sum", sumS.body, sumS.Extras())
}

type SumOption = aggregations.Option[*SumS]

func Sum(name, field string, opts ...SumOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &SumS{
		name: name,
		body: SumBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: sum aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) SumOption {
	return func(a *SumS) error {
		a.body.Format = format
		return nil
	}
}

func WithMissing(missing any) SumOption {
	return func(a *SumS) error {
		a.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) SumOption {
	return func(a *SumS) error {
		a.body.Script = &script
		return nil
	}
}

func WithValueType(valueType string) SumOption {
	return func(a *SumS) error {
		a.body.ValueType = valueType
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) SumOption {
	return aggregations.WithSubAggregation[*SumS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) SumOption {
	return aggregations.WithNamedSubAggregation[*SumS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) SumOption {
	return aggregations.WithSubAggregationsMap[*SumS](subsMap)
}

func WithMetaField(key string, value any) SumOption {
	return aggregations.WithMetaField[*SumS](key, value)
}

func WithMetaMap(meta map[string]any) SumOption {
	return aggregations.WithMetaMap[*SumS](meta)
}
