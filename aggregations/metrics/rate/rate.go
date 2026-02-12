package rate

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type RateBody struct {
	Field string `json:"field,omitempty"`
	Unit  string `json:"unit,omitempty"`
	Mode  string `json:"mode,omitempty"`
}

type RateS struct {
	name string
	body RateBody
	aggregations.BaseAggregation
}

func (rate *RateS) AggregationName() string {
	return rate.name
}

func (rate *RateS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("rate", rate.body, rate.Extras())
}

type RateOption = aggregations.Option[*RateS]

func Rate(name string, opts ...RateOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &RateS{
		name: name,
		body: RateBody{},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithField(field string) RateOption {
	return func(r *RateS) error {
		r.body.Field = field
		return nil
	}
}

func WithUnit(unit string) RateOption {
	return func(r *RateS) error {
		r.body.Unit = unit
		return nil
	}
}

func WithMode(mode string) RateOption {
	return func(r *RateS) error {
		r.body.Mode = mode
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) RateOption {
	return aggregations.WithSubAggregation[*RateS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) RateOption {
	return aggregations.WithNamedSubAggregation[*RateS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) RateOption {
	return aggregations.WithSubAggregationsMap[*RateS](subsMap)
}

func WithMetaField(key string, value any) RateOption {
	return aggregations.WithMetaField[*RateS](key, value)
}

func WithMetaMap(meta map[string]any) RateOption {
	return aggregations.WithMetaMap[*RateS](meta)
}
