package missing

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MissingBody struct {
	Field string `json:"field,omitempty"`
}

type MissingS struct {
	name string
	body MissingBody
	aggregations.BaseAggregation
}

func (missing *MissingS) AggregationName() string {
	return missing.name
}

func (missing *MissingS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("missing", missing.body, missing.Extras())
}

type MissingOption = aggregations.Option[*MissingS]

func Missing(name, field string, opts ...MissingOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &MissingS{
		name: name,
		body: MissingBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) MissingOption {
	return aggregations.WithSubAggregation[*MissingS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MissingOption {
	return aggregations.WithNamedSubAggregation[*MissingS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MissingOption {
	return aggregations.WithSubAggregationsMap[*MissingS](subsMap)
}

func WithMetaField(key string, value any) MissingOption {
	return aggregations.WithMetaField[*MissingS](key, value)
}

func WithMetaMap(meta map[string]any) MissingOption {
	return aggregations.WithMetaMap[*MissingS](meta)
}
