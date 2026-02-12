package weightedavg

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type WeightedAvgField struct {
	Field   string               `json:"field,omitempty"`
	Missing any                  `json:"missing,omitempty"`
	Script  *aggregations.Script `json:"script,omitempty"`
}

type WeightedAvgBody struct {
	Value  WeightedAvgField `json:"value,omitempty"`
	Weight WeightedAvgField `json:"weight,omitempty"`
	Format string           `json:"format,omitempty"`
}

type WeightedAvgS struct {
	name string
	body WeightedAvgBody
	aggregations.BaseAggregation
}

func (avg *WeightedAvgS) AggregationName() string {
	return avg.name
}

func (avg *WeightedAvgS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("weighted_avg", avg.body, avg.Extras())
}

type WeightedAvgOption = aggregations.Option[*WeightedAvgS]

func WeightedAvg(name, valueField, weightField string, opts ...WeightedAvgOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &WeightedAvgS{
		name: name,
		body: WeightedAvgBody{
			Value:  WeightedAvgField{Field: valueField},
			Weight: WeightedAvgField{Field: weightField},
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Value.Script != nil && agg.body.Value.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: weighted_avg value cannot have both field and script set")}
	}

	if agg.body.Weight.Script != nil && agg.body.Weight.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: weighted_avg weight cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithValueField(field string) WeightedAvgOption {
	return func(w *WeightedAvgS) error {
		w.body.Value.Field = field
		return nil
	}
}

func WithValueMissing(missing any) WeightedAvgOption {
	return func(w *WeightedAvgS) error {
		w.body.Value.Missing = missing
		return nil
	}
}

func WithValueScript(script aggregations.Script) WeightedAvgOption {
	return func(w *WeightedAvgS) error {
		w.body.Value.Script = &script
		return nil
	}
}

func WithWeightField(field string) WeightedAvgOption {
	return func(w *WeightedAvgS) error {
		w.body.Weight.Field = field
		return nil
	}
}

func WithWeightMissing(missing any) WeightedAvgOption {
	return func(w *WeightedAvgS) error {
		w.body.Weight.Missing = missing
		return nil
	}
}

func WithWeightScript(script aggregations.Script) WeightedAvgOption {
	return func(w *WeightedAvgS) error {
		w.body.Weight.Script = &script
		return nil
	}
}

func WithFormat(format string) WeightedAvgOption {
	return func(w *WeightedAvgS) error {
		w.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) WeightedAvgOption {
	return aggregations.WithSubAggregation[*WeightedAvgS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) WeightedAvgOption {
	return aggregations.WithNamedSubAggregation[*WeightedAvgS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) WeightedAvgOption {
	return aggregations.WithSubAggregationsMap[*WeightedAvgS](subsMap)
}

func WithMetaField(key string, value any) WeightedAvgOption {
	return aggregations.WithMetaField[*WeightedAvgS](key, value)
}

func WithMetaMap(meta map[string]any) WeightedAvgOption {
	return aggregations.WithMetaMap[*WeightedAvgS](meta)
}
