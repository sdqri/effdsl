package rangeagg

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type RangeItem struct {
	Key  string `json:"key,omitempty"`
	From any    `json:"from,omitempty"`
	To   any    `json:"to,omitempty"`
}

type RangeBody struct {
	Field   string      `json:"field,omitempty"`
	Ranges  []RangeItem `json:"ranges,omitempty"`
	Keyed   *bool       `json:"keyed,omitempty"`
	Missing any         `json:"missing,omitempty"`
}

type RangeS struct {
	name string
	body RangeBody
	aggregations.BaseAggregation
}

func (rangeAgg *RangeS) AggregationName() string {
	return rangeAgg.name
}

func (rangeAgg *RangeS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("range", rangeAgg.body, rangeAgg.Extras())
}

type RangeOption = aggregations.Option[*RangeS]

func Range(name, field string, ranges []RangeItem, opts ...RangeOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &RangeS{
		name: name,
		body: RangeBody{
			Field:  field,
			Ranges: ranges,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithKeyed(keyed bool) RangeOption {
	return func(r *RangeS) error {
		r.body.Keyed = &keyed
		return nil
	}
}

func WithMissing(missing any) RangeOption {
	return func(r *RangeS) error {
		r.body.Missing = missing
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) RangeOption {
	return aggregations.WithSubAggregation[*RangeS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) RangeOption {
	return aggregations.WithNamedSubAggregation[*RangeS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) RangeOption {
	return aggregations.WithSubAggregationsMap[*RangeS](subsMap)
}

func WithMetaField(key string, value any) RangeOption {
	return aggregations.WithMetaField[*RangeS](key, value)
}

func WithMetaMap(meta map[string]any) RangeOption {
	return aggregations.WithMetaMap[*RangeS](meta)
}
