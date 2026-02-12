package iprange

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type IPRangeItem struct {
	Key  string `json:"key,omitempty"`
	From any    `json:"from,omitempty"`
	To   any    `json:"to,omitempty"`
	Mask string `json:"mask,omitempty"`
}

type IPRangeBody struct {
	Field   string        `json:"field,omitempty"`
	Ranges  []IPRangeItem `json:"ranges,omitempty"`
	Keyed   *bool         `json:"keyed,omitempty"`
	Missing any           `json:"missing,omitempty"`
}

type IPRangeS struct {
	name string
	body IPRangeBody
	aggregations.BaseAggregation
}

func (r *IPRangeS) AggregationName() string {
	return r.name
}

func (r *IPRangeS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("ip_range", r.body, r.Extras())
}

type IPRangeOption = aggregations.Option[*IPRangeS]

func IPRange(name, field string, ranges []IPRangeItem, opts ...IPRangeOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &IPRangeS{
		name: name,
		body: IPRangeBody{
			Field:  field,
			Ranges: ranges,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithKeyed(keyed bool) IPRangeOption {
	return func(r *IPRangeS) error {
		r.body.Keyed = &keyed
		return nil
	}
}

func WithMissing(missing any) IPRangeOption {
	return func(r *IPRangeS) error {
		r.body.Missing = missing
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) IPRangeOption {
	return aggregations.WithSubAggregation[*IPRangeS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) IPRangeOption {
	return aggregations.WithNamedSubAggregation[*IPRangeS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) IPRangeOption {
	return aggregations.WithSubAggregationsMap[*IPRangeS](subsMap)
}

func WithMetaField(key string, value any) IPRangeOption {
	return aggregations.WithMetaField[*IPRangeS](key, value)
}

func WithMetaMap(meta map[string]any) IPRangeOption {
	return aggregations.WithMetaMap[*IPRangeS](meta)
}
